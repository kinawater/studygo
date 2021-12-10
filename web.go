package main

import (
	"fmt"
	"log"
	"net/http"
)

func handle(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w,"hello,I love %s!",r.URL.Path[1:])
}
func main() {
	http.HandleFunc("/",handle)
	log.Fatal(http.ListenAndServe(":8080",nil))
}
