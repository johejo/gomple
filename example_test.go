package gomple_test

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/johejo/gomple"
)

func handle(w http.ResponseWriter, r *http.Request) error {
	// do something
	return nil // could return error
}

func Example() {
	// with http.NewServeMux
	g := gomple.New()
	mux := http.NewServeMux()
	mux.HandleFunc("/", g.WrapFunc(handle))
}

func errHandler(w http.ResponseWriter, r *http.Request, err error) {
	if err != nil {
		// your own error handling
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ExampleMux() {
	g := gomple.New(gomple.WithErrorHandler(errHandler))
	c := &controller{gomple: g}
	mux := gomple.NewMuxWithGomple(g)
	mux.Post("/users", c.createUser)
}

type controller struct {
	gomple *gomple.Gomple
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (c *controller) createUser(w http.ResponseWriter, r *http.Request) error {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		return err
	}
	log.Println(u)
	w.WriteHeader(http.StatusCreated)
	return nil
}
