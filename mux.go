package gomple

import (
	"net/http"

	"github.com/go-chi/chi"
)

func NewMux(opts ...Option) *Mux {
	return &Mux{
		g:   New(opts...),
		mux: chi.NewMux(),
	}
}

func NewMuxWithGomple(g *Gomple) *Mux {
	return &Mux{
		g:   g,
		mux: chi.NewMux(),
	}
}

type Mux struct {
	g   *Gomple
	mux *chi.Mux
}

var _ http.Handler = (*Mux)(nil)

func (m *Mux) Gomple() *Gomple {
	return m.g
}

func (m *Mux) Raw() *chi.Mux {
	return m.mux
}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.mux.ServeHTTP(w, r)
}

func (m *Mux) Get(pattern string, h HandlerFunc) {
	m.mux.Get(pattern, m.g.WrapFunc(h))
}

func (m *Mux) Put(pattern string, h HandlerFunc) {
	m.mux.Put(pattern, m.g.WrapFunc(h))
}

func (m *Mux) Post(pattern string, h HandlerFunc) {
	m.mux.Post(pattern, m.g.WrapFunc(h))
}

func (m *Mux) Options(pattern string, h HandlerFunc) {
	m.mux.Options(pattern, m.g.WrapFunc(h))
}

func (m *Mux) Head(pattern string, h HandlerFunc) {
	m.mux.Head(pattern, m.g.WrapFunc(h))
}

func (m *Mux) Delete(pattern string, h HandlerFunc) {
	m.mux.Delete(pattern, m.g.WrapFunc(h))
}

func (m *Mux) Connect(pattern string, h HandlerFunc) {
	m.mux.Connect(pattern, m.g.WrapFunc(h))
}

func (m *Mux) Trace(pattern string, h HandlerFunc) {
	m.mux.Trace(pattern, m.g.WrapFunc(h))
}

func (m *Mux) HandleFunc(pattern string, h HandlerFunc) {
	m.mux.HandleFunc(pattern, m.g.WrapFunc(h))
}
