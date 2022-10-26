package main

//how to use r.BasicAuth() ?

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/hello.text", helloText)
	http.HandleFunc("/hello.html", helloHtml)
	http.HandleFunc("/name-get", nameGet)
	http.HandleFunc("/name-post", namePost)
	http.HandleFunc("/request-name", nameReqGet)
	http.HandleFunc("/request-name-post", nameReqPost)
	err := http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
	log.Fatal("ListenAndServeTLS: ", err)
}

func helloText(w http.ResponseWriter, r *http.Request) {
	if r.Method != "HEAD" && r.Method != "GET" {
		http.Error(w, "method no allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Authorization") == "" {
		w.Header().Set("www-authenticate", "basic realm=\"tp1\"")
		http.Error(w, "Haha!", http.StatusUnauthorized)

		return
	}
	//fmt.Fprintln(w, r.Header.Get("Authorization"))
	user, pass, ok := r.BasicAuth()
	fmt.Fprintln(w, user)
	fmt.Fprintln(w, pass)
	fmt.Fprintln(w, ok)
	if !ok || user != "curie" || pass != "pol" {
		w.Header().Set("www-authenticate", "basic realm=\"tp1\"")
		http.Error(w, "Nope!", http.StatusUnauthorized)
		return
	}
	http.Error(w, "Yes!", 200)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "Bonjour!")
} //attention aux attaques au dictionnaire (pleins de mots de passe qui sont testés)

func helloHtml(w http.ResponseWriter, r *http.Request) {
	const data = `<!DOCTYPE html>
<html>
<head></head>
<body>
<p>Bonjouuuuuuuuuuur Sarah!</p>
</body>
</html>`
	if r.Method != "HEAD" && r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, data)
}

func nameGet(w http.ResponseWriter, r *http.Request) {
	const data = `<!DOCTYPE html>
<html>
<head></head>
<body>
<form action="/request-name" method="get">
Votre nom: <input type="text" name="name"/> <input type="submit"/></form>
</body>
</html>`
	if r.Method != "HEAD" && r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, data)
}

func namePost(w http.ResponseWriter, r *http.Request) {
	const data = `<!DOCTYPE html>
<html>
<head></head>
<body>
<form action="/request-name-post" method="post">
Votre nom: <input type="text" name="name"/> <input type="submit"/></form>
</body>
</html>`
	if r.Method != "HEAD" && r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, data)
}

func nameReqGet(w http.ResponseWriter, r *http.Request) {

	if r.Method != "HEAD" && r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	r.ParseForm()
	fmt.Fprint(w, "Votre nom est : ", r.Form.Get("name"))
}

func nameReqPost(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	io.Copy(os.Stdout, r.Body)
	r.ParseForm()
	fmt.Fprintln(w, "Type du corps de la requête : ", r.Header.Get("Content-Type"))
	//fmt.Fprintln(w, "Votre nom est : ", r.Form.Get("name"))

}
