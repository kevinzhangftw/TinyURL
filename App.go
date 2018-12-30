package main

import (
	"crypto/sha1"
	"fmt"
	"net/http"
)

type ServerImpl struct {
	urlDB map[string]string
}

var s ServerImpl

func (s ServerImpl) processURL(inputURL string) (tinyURL string) {
	h := sha1.New()
	h.Write([]byte(inputURL))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs[:4])
}

func (s *ServerImpl) saveURL(inputURL string, tinyURL string) (err error) {
	s.urlDB["/"+tinyURL] = inputURL
	fmt.Println("saving new post request, db=", s.urlDB)
	return nil
}

func (s ServerImpl) urlLookup(tinyUrl string) (orURL string) {
	fmt.Println("urlDB[tinyUrl] is ", s.urlDB[tinyUrl])
	fmt.Println("urlDB", s.urlDB)
	return s.urlDB[tinyUrl]
}

func routeTo(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		fmt.Fprint(w, "Welcome to TinyUrl, this is a server, why dont you try to make a request?")
		if r.URL.String() != "/" {
			ogURL := s.urlLookup( r.URL.String())
			fmt.Fprint(w, "\n now redirect to ... ", ogURL)
		}

	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		inputURL := r.FormValue("inputURL")

		tinyURL := s.processURL(inputURL)
		if err := s.saveURL(inputURL, tinyURL); err != nil {
			fmt.Fprintf(w, "saveURL err: %v", err)
			return
		}

		fmt.Fprintf(w, "saved tinyURL is %s\n", tinyURL)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func bootUpServer() {
	s = ServerImpl{map[string]string{}}
	http.HandleFunc("/", routeTo)
	http.ListenAndServe(":8000", nil)
}

func main() {
	bootUpServer()
}
