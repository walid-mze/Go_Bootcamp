package main

import (
	"fmt"
	"net/http"
	"io"
	"log"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func main() {
	resp,err:=http.Get("http://localhost:3000/")
	if err!=nil{
		log.Fatal(err)
	}
	message,err:=io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Printf("%s",message)
	http.HandleFunc("/", hello)
	http.ListenAndServe(":3000", nil)
}
