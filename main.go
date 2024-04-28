package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/TomShep998/blog/models"
	"github.com/gorilla/mux"
)

func main() {
	setup()

}

func setup() {

	r := mux.NewRouter()

	r.HandleFunc("/", index)
	r.HandleFunc("/posts", getPosts).Methods("GET")
	r.HandleFunc("/posts/{id}", getPost).Methods("GET")
	r.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")
	r.HandleFunc("/createpost", createPost).Methods("POST")
	r.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	http.ListenAndServe(":3000", r)

}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./layouts/index.html"))
	tmpl.Execute(w, nil)

}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.GetPosts())

}
func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	convId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(models.GetPost(convId))
}

func updatePost(w http.ResponseWriter, r *http.Request) {

	var p models.Post
	/*vars := mux.Vars(r)
	id := vars["id"]

	convid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	*/

	json.NewDecoder(r.Body).Decode(&p)

	models.UpdatePost(p)

}
func deletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	convId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	models.DeletePost(convId)

}

func createPost(w http.ResponseWriter, r *http.Request) {

	var p models.Post
	json.NewDecoder(r.Body).Decode(&p)
	models.CreatePost(p)

}
