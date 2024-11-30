package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HandlerError struct {
	Message string `json:"message"`
	Cause   string `json:"cause"`
}

func NewHandlerError(message string, cause error) *HandlerError {
	he := HandlerError{
		Message: message,
		Cause:   "",
	}
	if cause != nil {
		he.Cause = cause.Error()
	}
	return &he
}

func (e *HandlerError) Send(w http.ResponseWriter, statusCode int) error {
	b, err := json.Marshal(*e)
	if err != nil {
		return err
	}

	w.WriteHeader(statusCode)
	if byteCount, err := w.Write(b); err != nil {
		return err
	} else if byteCount != len(b) {
		return fmt.Errorf("could not write entire response body: %d / %d bytes", byteCount, len(b))
	}

	return nil
}
