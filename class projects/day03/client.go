package main

import (
	"bytes"
	"encoding/json"

	"log"
	"net/http"
)

type Student struct {
	Name    string
	Age     int
	Address string
	Courses []string
}

func main() {
	w:=Student{
		"walid",
		20,
		"map 010",
		[]string{"Go","Algorithms"},
	}
	jsonData,err:=json.Marshal(w)
	if err!=nil{
		log.Fatal(err)
	}else{
		bodyReader:=bytes.NewReader(jsonData)
		http.Post("http://localhost:8080/students","",bodyReader)
	}



}
