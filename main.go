package main

import "net/http"

func main() {
	http.HandleFunc("/helloworld", func (w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello world"))
	})
	http.ListenAndServe(":8080", nil)
}

