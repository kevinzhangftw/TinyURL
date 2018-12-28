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
	return string(bs[:8])
}

func saveURL(tinyURL string){
	fmt.Println("Hello, 世界")
}

func routeTo(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case "GET":
		fmt.Fprint(w, "Welcome to TinyUrl, this is a server, why dont you try to make a request?")

	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		inputURL := r.FormValue("inputURL")

		tinyURL := processURL(inputURL)
		//if err := saveURL(tinyURL); err != nil {
		//	fmt.Fprintf(w, "saveURL err: %v", err)
		//	return
		//}

		fmt.Fprintf(w, "tinyURL = %s\n", tinyURL)

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

