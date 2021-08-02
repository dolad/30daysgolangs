package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

type server struct{}

func post(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

func put(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called"}`))
}

func delete(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "delete called"}`))
}

func notFound(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

func get(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

// func params(w http.ResponseWriter, r*http.Request){
// 	pathParam := mux.Vars(r)
// 	w.Header().Set("Content-Type", "application/json")
// 	userID := -1
// 	var err error

// 	if val, ok := pathParams["userID"]; ok {
//         userID, err = strconv.Atoi(val)
//         if err != nil {
//             w.WriteHeader(http.StatusInternalServerError)
//             w.Write([]byte(`{"message": "need a number"}`))
//             return
//         }
//     }
// }


func main() {
    r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
    api.HandleFunc("/", get).Methods(http.MethodGet)
	api.HandleFunc("/", put).Methods(http.MethodPut)
	api.HandleFunc("/", post).Methods(http.MethodPut)
	api.HandleFunc("/", delete).Methods(http.MethodDelete)
	api.HandleFunc("/", notFound)
    log.Fatal(http.ListenAndServe(":8080", r))
}