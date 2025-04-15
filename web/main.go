package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//go:embed templates/home-page
var home_page embed.FS

func main(){
	router := httprouter.New()

	directory,err := fs.Sub(home_page,"templates/home-page")
	if err != nil{
		panic(err)
	}
	router.ServeFiles("/cv/*filepath",http.FS(directory))

	server := http.Server{
		Addr: "localhost:4000",
		Handler: router,
	}

	err = server.ListenAndServe()
	if err != nil{
		panic(err)
	}
}