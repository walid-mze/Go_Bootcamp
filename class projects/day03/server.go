package main

import (
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

var students = []Student{
	{
		Name:    "John Doe",
		Age:     20,
		Address: "101 Maple Street",
		Courses: []string{"Go", "Algorithms"},
	},
	{
		Name:    "Jane Smith",
		Age:     22,
		Address: "202 Oak Avenue",
		Courses: []string{"Python", "Data Structures"},
	},
}

func handleGet(w http.ResponseWriter, _ *http.Request){

		jsonResp, err := json.Marshal(students)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonResp)
}




func studentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		handleGet(w,r)

	case http.MethodPost:
		var newStudent Student
		err := json.NewDecoder(r.Body).Decode(&newStudent)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		if newStudent.Name == "" {
			http.Error(w, "Name is required", http.StatusBadRequest)
			return
		}
		students=append(students,newStudent)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newStudent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/students", studentsHandler)

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
