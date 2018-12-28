package main

import (
	"fmt"
	"net/http"
	"crypto/sha1"
)

func processURL(inputURL string) (tinyURL string){
	h := sha1.New()
	h.Write([]byte(inputURL))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs[:4])
}

func saveURL(inputURL string, tinyURL string, urlDB map[string]string) (err error){
	urlDB[inputURL] = tinyURL
	return nil
}

func routeTo(w http.ResponseWriter, r *http.Request){
	urlDB := make(map[string]string)

	switch r.Method {
	case "GET":
		fmt.Fprint(w, "Welcome to TinyUrl, this is a server, why dont you try to make a request?")
		fmt.Fprint(w, "request URL is ", r.URL)

	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		inputURL := r.FormValue("inputURL")

		tinyURL := processURL(inputURL)
		if err := saveURL(inputURL, tinyURL, urlDB); err != nil {
			fmt.Fprintf(w, "saveURL err: %v", err)
			return
		}

		fmt.Fprintf(w, "saved tinyURL is %s\n", urlDB[inputURL])

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func bootUpServer() {
	http.HandleFunc("/", routeTo)
	http.ListenAndServe(":8000", nil)
}

func main() {
	bootUpServer()
}

