package gomple

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

type ErrorHandler func(w http.ResponseWriter, r *http.Request, err error)

func DefaultErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

type config struct {
	errHandler ErrorHandler
}

type Option func(cfg *config)

func WithErrorHandler(eh ErrorHandler) Option {
	if eh == nil {
		eh = DefaultErrorHandler
	}
	return func(cfg *config) {
		cfg.errHandler = eh
	}
}

func defaults() []Option {
	return []Option{
		WithErrorHandler(nil),
	}
}

func New(opts ...Option) *Gomple {
	cfg := new(config)
	for _, opt := range append(defaults(), opts...) {
		opt(cfg)
	}
	return &Gomple{
		errHandler: cfg.errHandler,
	}
}

type Gomple struct {
	errHandler ErrorHandler
}

func (g *Gomple) WrapFunc(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			g.errHandler(w, r, err)
			return
		}
	}
}

func (g *Gomple) JSON(w http.ResponseWriter, r *http.Request, j interface{}) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(j); err != nil {
		g.errHandler(w, r, err)
		return
	}
	_, _ = io.Copy(w, &buf)
}
