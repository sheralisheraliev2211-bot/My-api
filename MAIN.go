package main

import (
 "encoding/json"
 "net/http"
)

type Task struct {
 ID    int    json:"id"
 Title string json:"title"
 Done  bool   json:"done"
}

var tasks = []Task{
 {ID: 1, Title: "Learn Go", Done: false},
 {ID: 2, Title: "Build API", Done: true},
 {ID: 3, Title: "Learn Docker", Done: false},
}

func home(w http.ResponseWriter, r *http.Request) {
 w.Header().Set("Content-Type", "application/json")

 json.NewEncoder(w).Encode(map[string]string{
  "message": "Tasks API is running",
 })
}

func getTasks(w http.ResponseWriter, r *http.Request) {
 w.Header().Set("Content-Type", "application/json")

 json.NewEncoder(w).Encode(tasks)
}

func main() {
 http.HandleFunc("/", home)
 http.HandleFunc("/tasks", getTasks)

 http.ListenAndServe(":8080", nil)
}
