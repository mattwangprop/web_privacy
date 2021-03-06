package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	// "encoding/json"
	"github.com/gorilla/mux"
)

func loadIndex(w http.ResponseWriter) {
	pwd, _ := os.Getwd()
	filepath := pwd + "/index.html"
	index, _ := ioutil.ReadFile(filepath)

	io.WriteString(w, string(index))
}

func serveFile(w http.ResponseWriter, filePath string) {
	pwd, _ := os.Getwd()
	filePath = pwd + "/" + filePath
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		loadIndex(w)
	} else {
		io.WriteString(w, string(file))
	}
}

func handlerURL(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		loadIndex(w)
	} else {
		serveFile(w, r.URL.Path)
	}
}

func addRoutes() {

	http.HandleFunc("/", handlerURL)

	router := mux.NewRouter()

	router.HandleFunc("/api/words", viewWords).Methods("GET")
	router.HandleFunc("/api/word/detail/{id}", viewFindWordByID).Methods("GET")
	router.HandleFunc("/api/word/new", addNewWord).Methods("POST")
	router.HandleFunc("/api/word/update/{id}", updateWord).Methods("POST")
	router.HandleFunc("/api/word/remove/{id}", removeWord).Methods("POST")

	http.Handle("/api/", router)

}


