package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:jacobfullstack@proton.me\">jacobfullstack@proton.me</a>.</p>")
}

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
		case "/":
			homeHandler(w, r)
		case "/contact":
			contactHandler(w, r)
		default:
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "<h1>404 - Page Not Found</h1><p>The page you are looking for does not exist.</p>")
			// http.Error(w, "404 - Page Not Found", http.StatusNotFound)
	}
}

func main() {
	var router Router
	fmt.Println("Starting server on :3000...")
	http.ListenAndServe(":3000", router)
}