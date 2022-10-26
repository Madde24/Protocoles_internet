package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/number-get", numberGet)
	http.HandleFunc("/request-number-get", numberReqGet)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal("ListenAndServe: ", err)
}

func numberGet(w http.ResponseWriter, r *http.Request) {
	const data = `<!DOCTYPE html>
<html>
<head></head>
<body>
<form action="/request-number-get" method="get">
Votre numero: <input type="number" name="value"/> <input type="submit"/></form>
</body>
</html>`
	if r.Method != "HEAD" && r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, data)
}

func numberPrime(n int) []int {
	a := make([]int, n)
	for i := 2; i < n; i++ {
		for j := 2 * i; j < n; j += i {
			if j%i == 0 {
				a[j] = -1
			}
		}
	}

	return a

}

func numberReqGet(w http.ResponseWriter, r *http.Request) {

	if r.Method != "HEAD" && r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	r.ParseForm()
	i, err := strconv.ParseInt(r.FormValue("value"), 10, 64)
	if err != nil {
		panic(err)
	}
	/**
	if err != nil {
		panic(err)
		return
	}**/
	fmt.Fprintln(w, "Votre numéro est : ", r.FormValue("value"))
	fmt.Fprintln(w,"Les nombres premiers de 1 à ", i, "sont :")
	a :=numberPrime(int(i))
	
	for j := 1; j < int(i); j++ {
		if a[j] != -1 {
			fmt.Fprintln(w,j)
		}
	}
}
