package fetch

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
	"gitlab.com/piorun102/lg"
	"net/http"
)

var (
	v       = validator.New(validator.WithRequiredStructEnabled())
	decoder = schema.NewDecoder()
)

const (
	Admin Role = "admin"
	User  Role = "user"
)

type Role string

func Post[Req any, Resp any](
	pattern string,
	r chi.Router,
	auth func(header http.Header) (*Role, *Error),
	cb func(ctx lg.CtxLogger, req *Req, role *Role) (*Resp, *Error),
) {
	r.Post(pattern, func(w http.ResponseWriter, r *http.Request) {
		var body *Req
		var (
			role *Role
			err  *Error
			ctx  = lg.Ctx(r.Context(), r.Header)
		)
		defer ctx.Send()
		defer r.Body.Close()

		ctx.SpanLog("REQ", fmt.Sprintf("%+v", r))
		ctx.Tracef("REQ %v\n%+v", r.URL.Path, r)

		if role, err = auth(r.Header); err != nil {
			ctx.Errorf("auth: %+v", err)
			writeResponse(w, nil, err)
			return
		}

		if errD := DecodeAndValidate(r, &body); err != nil {
			ctx.Errorf("DecodeAndValidate failed: %v", errD)
			writeResponse(w, nil, &Error{http.StatusBadRequest, errD.Error()})
			return
		}

		resp, err := cb(ctx, body, role)
		if err != nil {
			ctx.Errorf(err.Message)
			writeResponse(w, nil, err)
			return
		}

		ctx.SpanLog("RESP", fmt.Sprintf("%+v, %+v", resp, err))
		ctx.Tracef("RESP %v\n%+v", r.URL.Path, fmt.Sprintf("%+v, %+v", resp, err))
		writeResponse(w, resp, nil)
		return
	})
}

func Patch[Req any, Resp any](
	pattern string,
	r chi.Router,
	auth func(header http.Header) (*Role, *Error),
	cb func(ctx lg.CtxLogger, req *Req, role *Role) (*Resp, *Error),
) {
	r.Patch(pattern, func(w http.ResponseWriter, r *http.Request) {
		var body *Req
		var (
			role *Role
			err  *Error
			ctx  = lg.Ctx(r.Context(), r.Header)
		)
		defer ctx.Send()
		defer r.Body.Close()

		ctx.SpanLog("REQ", fmt.Sprintf("%+v", r))
		ctx.Tracef("REQ %v\n%+v", r.URL.Path, r)

		if role, err = auth(r.Header); err != nil {
			ctx.Errorf("auth: %+v", err)
			writeResponse(w, nil, err)
			return
		}

		if errD := DecodeAndValidate(r, &body); err != nil {
			ctx.Errorf("DecodeAndValidate failed: %v", err)
			writeResponse(w, nil, &Error{http.StatusBadRequest, errD.Error()})
			return
		}

		resp, err := cb(ctx, body, role)
		if err != nil {
			ctx.Errorf(err.Message)
			writeResponse(w, nil, err)
			return
		}

		ctx.SpanLog("RESP", fmt.Sprintf("%+v, %+v", resp, err))
		ctx.Tracef("RESP %v\n%+v", r.URL.Path, fmt.Sprintf("%+v, %+v", resp, err))
		writeResponse(w, resp, nil)
		return
	})
}

func DecodeAndValidate[Req any](r *http.Request, req Req) error {
	return json.NewDecoder(r.Body).Decode(&req)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func Get[reqType any, respType any](
	pattern string,
	r chi.Router,
	auth func(header http.Header) (*Role, *Error),
	cb func(ctx lg.CtxLogger, req *reqType, role *Role) (*respType, *Error),
) {
	r.Get(pattern, func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		var (
			role *Role
			req  reqType
			err  *Error
			ctx  = lg.Ctx(r.Context(), r.Header)
		)
		defer ctx.Send()

		if role, err = auth(r.Header); err != nil {
			ctx.Errorf("auth: %+v", err)
			writeResponse(w, nil, err)
			return
		}

		if errD := DecodeQueryAndValidate(r, &req); errD != nil {
			ctx.Errorf("DecodeQueryAndValidate failed: %v", errD)
			writeResponse(w, nil, &Error{http.StatusBadRequest, errD.Error()})
			return
		}

		ctx.SpanLog("REQ", fmt.Sprintf("%+v", req))
		ctx.Tracef("REQ %v\n%+v", r.URL.Path, req)
		resp, err := cb(ctx, &req, role)
		if err != nil {
			ctx.Errorf(err.Message)
			writeResponse(w, nil, err)
			return
		}
		ctx.SpanLog("RESP", fmt.Sprintf("%+v, %+v", resp, err))
		ctx.Tracef("RESP %v\n%+v", r.URL.Path, fmt.Sprintf("%+v, %+v", resp, err))
		writeResponse(w, resp, nil)
		return
	})
}

func DecodeQueryAndValidate[reqType any](r *http.Request, req reqType) (err error) {
	if err = decoder.Decode(req, r.URL.Query()); err != nil {
		return
	}
	return v.Struct(req)
}
