package utils

import (
	"encoding/json"
	"net/http"

	"github.com/neghi14/starter"
)

type httpResponse struct {
	Status  string      `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Code    string      `json:"error_code,omitempty"`
	Limit   int         `json:"limit,omitempty"`
	Page    int         `json:"page,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type HttpResponseBuilder struct {
	w          http.ResponseWriter
	statusCode int
	res        *httpResponse
}

func JSON(w http.ResponseWriter) *HttpResponseBuilder {
	return &HttpResponseBuilder{
		w:   w,
		res: &httpResponse{},
	}
}

func (h *HttpResponseBuilder) SetStatusCode(code int) *HttpResponseBuilder {
	h.statusCode = code
	return h
}
func (h *HttpResponseBuilder) SetStatus(status starter.ResponseStatus) *HttpResponseBuilder {
	h.res.Status = status.String()
	return h
}
func (h *HttpResponseBuilder) SetMessage(message string) *HttpResponseBuilder {
	h.res.Message = message
	return h
}
func (h *HttpResponseBuilder) SetData(data interface{}) *HttpResponseBuilder {
	h.res.Data = data
	return h
}
func (h *HttpResponseBuilder) SetErrorCode(code string) *HttpResponseBuilder {
	h.res.Code = code
	return h
}
func (h *HttpResponseBuilder) SetLimit(limit int) *HttpResponseBuilder {
	h.res.Limit = limit
	return h
}
func (h *HttpResponseBuilder) SetPage(page int) *HttpResponseBuilder {
	h.res.Page = page
	return h
}
func (h *HttpResponseBuilder) Send() {
	h.w.Header().Add("Content-Type", "application/json")
	h.w.WriteHeader(h.statusCode)
	json.NewEncoder(h.w).Encode(&h.res)
}
