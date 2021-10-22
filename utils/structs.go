package utils

import (
	"net/http"

	"github.com/go-chi/render"
)

//--
// Error response payloads & renderers
//--

// ErrResponse - Response with error
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int    `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

// Render - Make HTTP status code equal to status code in struct
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	e.AppCode = e.HTTPStatusCode
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// ErrRender - Return reder error with code
func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response",
		ErrorText:      err.Error(),
	}
}

// ErrBadRequest - Если нет параметров нахуй
func ErrBadRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     err.Error(),
	}
}

// ErrNotFound - Возвращает ошибку 404 со статусом
func ErrNotFound(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 404,
		StatusText:     err.Error(),
	}
}

// ErrForbidden - Возвращает ошибку 403 со статусом
func ErrForbidden(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 403,
		StatusText:     err.Error(),
	}
}

// ErrMethodNotAllowed - Возвращает ошибку 404 со статусом
func ErrMethodNotAllowed() render.Renderer {
	return &ErrResponse{
		HTTPStatusCode: 405,
		StatusText:     "Method Not Allowed",
	}
}

// ErrUnauthorized - Хахаха
func ErrUnauthorized() render.Renderer {
	return &ErrResponse{
		HTTPStatusCode: 401,
		StatusText:     "Unauthorized",
	}
}

//--
// Success response payloads & renderers
//--

// SuccessResponse - Success response
type SuccessResponse struct {
	HTTPStatusCode int `json:"-"` // http response status code

	StatusText string      `json:"status"`            // user-level status message
	AppCode    int         `json:"code,omitempty"`    // application-specific error code
	Payload    interface{} `json:"payload,omitempty"` // user-level payload
}

// Render - Make HTTP status code equal to status code in struct
func (s *SuccessResponse) Render(w http.ResponseWriter, r *http.Request) error {
	s.AppCode = s.HTTPStatusCode
	render.Status(r, s.HTTPStatusCode)
	return nil
}

// SuccessOK - Ну ок же все
func SuccessOK(message string) render.Renderer {
	return &SuccessResponse{
		HTTPStatusCode: 200,
		StatusText:     message,
	}
}
