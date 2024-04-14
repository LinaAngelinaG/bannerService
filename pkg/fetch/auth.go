package fetch

import (
	"github.com/puzpuzpuz/xsync/v3"
	"net/http"
	"time"
)

var access = &Access{
	tokens: xsync.NewMapOf[string, *AccessInfo](),
}

type AccessInfo struct {
	ExpireTime time.Time
	Role       Role
}

type AccessToken struct {
	Token      string    `db:"token"`
	Role       Role      `db:"role"`
	ExpireTime time.Time `db:"expire_time"`
}

type Access struct {
	tokens *xsync.MapOf[string, *AccessInfo]
	//userTokens  *xsync.MapOf[string, time.Time]
	//adminTokens *xsync.MapOf[string, time.Time]
}

func (a *Access) HasAdminAccess(token string) bool {
	//_, ok := a.adminTokens.Load(token)
	//return ok
	info, ok := a.tokens.Load(token)
	return ok && info.Role == Admin
}

func (a *Access) HasUserAccess(token string) bool {
	//_, ok := a.userTokens.Load(token)
	//return ok
	info, ok := a.tokens.Load(token)
	return ok && (info.Role == User || info.Role == Admin)
}

func InitAccessTokens(accessTokens []AccessToken) {
	for _, token := range accessTokens {
		access.tokens.Store(token.Token, &AccessInfo{Role: token.Role, ExpireTime: token.ExpireTime})
	}
}

func (a *Access) addAccessTokens(accessTokens *xsync.MapOf[string, *AccessInfo]) {
	accessTokens.Range(func(key string, value *AccessInfo) bool {
		a.tokens.Store(key, value)
		return true
	})
}

func (a *Access) removeExpiredTokens() {
	keysToDelete := make([]string, 0)
	a.tokens.Range(func(key string, value *AccessInfo) bool {
		if value.ExpireTime.Before(time.Now()) {
			keysToDelete = append(keysToDelete, key)
		}
		return true
	})
	for _, key := range keysToDelete {
		a.tokens.Delete(key)
	}
}

func AdminAccess(header http.Header) (*Role, *Error) {
	token := header.Get("token")
	a := Admin

	if token == "" {
		return &a, &Error{
			Code:    http.StatusUnauthorized,
			Message: "Требуется заголовок token",
		}
	}

	if !access.HasAdminAccess(token) {
		return &a, &Error{
			Code:    http.StatusUnauthorized,
			Message: "Недействительный токен",
		}
	}
	return &a, nil
}

func UserAccess(header http.Header) (*Role, *Error) {
	token := header.Get("token")
	a := User
	if token == "" {
		return &a, &Error{
			Code:    http.StatusUnauthorized,
			Message: "Требуется заголовок token",
		}
	}

	if !access.HasUserAccess(token) {
		return &a, &Error{
			Code:    http.StatusUnauthorized,
			Message: "Недействительный токен",
		}
	}
	return &a, nil
}
