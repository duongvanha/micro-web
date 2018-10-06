package main

import (
	"encoding/json"
	"github.com/duongvanha/micro-web/worker/movie"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {

	r := mux.NewRouter()

	movieRepository := movie.Repository{}

	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		data, err := movieRepository.GetByPage(1)

		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		}

		movies, err := json.Marshal(data)

		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		}

		writer.Write(movies)
	})

	http.Handle("/", r)

	port := os.Getenv("PORT")

	log.Println("Starting HTTP service at " + port)

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}

}
