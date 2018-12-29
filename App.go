package main

import (
	"fmt"
	"net/http"
	"crypto/sha1"
)

var urlDB map[string]string

func processURL(inputURL string) (tinyURL string){
	h := sha1.New()
	h.Write([]byte(inputURL))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs[:4])
}

func saveURL(inputURL string, tinyURL string, urlDB map[string]string) (err error){
	urlDB["/" + tinyURL] = inputURL
	fmt.Println("saving new post request, db=", urlDB)
	return nil
}

func urlLookup(urlDB map[string]string, tinyUrl string) (orURL string) {
	fmt.Println("urlDB[tinyUrl] is ", urlDB[tinyUrl])
	fmt.Println("urlDB", urlDB)
	return urlDB[tinyUrl]
}

func routeTo(w http.ResponseWriter, r *http.Request){

	switch r.Method {
	case "GET":
		fmt.Fprint(w, "Welcome to TinyUrl, this is a server, why dont you try to make a request?")
		if r.URL.String() != "/" {
			ogURL := urlLookup(urlDB, r.URL.String())
			fmt.Fprint(w, "\n now redirect to ... ", ogURL)
		}

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

		fmt.Fprintf(w, "saved tinyURL is %s\n", tinyURL)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func bootUpServer() {
	urlDB = map[string]string{}
	http.HandleFunc("/", routeTo)
	http.ListenAndServe(":8000", nil)
}

func main() {
	bootUpServer()
}

