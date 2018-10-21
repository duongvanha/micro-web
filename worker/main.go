package main

import (
	"context"
	model "github.com/duongvanha/micro-web/proto"
	"github.com/duongvanha/micro-web/worker/movie"
	"google.golang.org/grpc"
	"log"
	"net"
)

var movieRepository = movie.Repository{}

type MovieController struct{}

func (t *MovieController) GetByPage(ctx context.Context, page *model.Page) (listMovies *model.ListMovie, err error) {

	movies, err := movieRepository.GetByPage(page.Page)
	if err != nil {
		return
	}

	return &model.ListMovie{Movies: movies}, err

}

func main() {

	//port := os.Getenv("PORT")
	port := "50050"

	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}

	log.Println("Starting HTTP service at " + port)

	server := grpc.NewServer()
	model.RegisterMovieRepositoryServer(server, new(MovieController))
	server.Serve(listen)

}
