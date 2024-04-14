package fetch

import (
	"encoding/json"
	"gitlab.com/piorun102/lg"
	"net/http"
)

const (
	SuccessCode    = 1
	NotCreatedCode = 2
	ErrorCode      = 3
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func writeResponse(w http.ResponseWriter, output any, error *Error) {
	var resp RespTemplate
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	if error != nil {
		//resp.Message = fmt.Sprint(err)
		resp.Data = nil
		w.WriteHeader(error.Code)
		_, err := w.Write([]byte(error.Message))
		if err != nil {
			lg.Error(err)
		}
		return
	}
	resp.Data = output

	r, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		lg.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(r)
	lg.Error(err)
}

type RespTemplate struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data"`
}
