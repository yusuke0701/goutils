package chi

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/render"
)

// ErrResponse renderer type for handling all error types.
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

// Render renders a single payload and responds to the client request.
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// ErrInvalidRequest is the http error for Status 400 - Bad Request.
// ref: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400
func ErrInvalidRequest(ctx context.Context, err error) render.Renderer {
	log.Print(err)
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

// ErrInvalidTaskName is the http error for the No X-Appengine-Taskname request header error.
func ErrInvalidTaskName(ctx context.Context) render.Renderer {
	err := errors.New("invalid request: No X-Appengine-Taskname request header found")
	log.Print(err)
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

// ErrForbidden is the http error for Status 403 - Forbidden.
// ref: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/403
func ErrForbidden(ctx context.Context, err error) render.Renderer {
	log.Print(err)
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusForbidden,
		StatusText:     "Forbidden",
		ErrorText:      err.Error(),
	}
}

// ErrConflict is the http error for Status 409 - Conflict.
// ref: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/409
func ErrConflict(ctx context.Context, err error, ac int64) render.Renderer {
	log.Print(err)
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusConflict,
		StatusText:     "Conflict",
		AppCode:        ac,
		ErrorText:      err.Error(),
	}
}

// ErrNotFound is the http error for Status 404 - Not Found.
// ref: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404
func ErrNotFound(ctx context.Context, err error, ac int64) render.Renderer {
	log.Print(err)
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusNotFound,
		StatusText:     "Not found.",
		AppCode:        ac,
		ErrorText:      err.Error(),
	}
}

// ErrInternalServerError is the http error for Status 500 - Internal Server Error.
// ref: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500
func ErrInternalServerError(ctx context.Context, err error) render.Renderer {
	log.Print(err)
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		StatusText:     "Internal server error",
		ErrorText:      err.Error(),
	}
}

// ErrRender renders http error responses
func ErrRender(ctx context.Context, err error) render.Renderer {
	log.Print(err)
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusUnprocessableEntity,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}
