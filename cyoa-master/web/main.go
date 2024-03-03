package main

import (
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLogger *log.Logger
	infoLogger  *log.Logger
	mapp        map[string]Compononet
}

func main() {

	app := &application{
		errorLogger: log.New(os.Stderr, "Error", log.Ldate|log.Ltime),
		infoLogger:  log.New(os.Stdout, "Info", log.Ldate|log.Ltime),
	}

	story := app.buildMap()
	app.mapp = story
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.handleMain)

	app.infoLogger.Println("starting application at port 4000")

	http.ListenAndServe(":4000", mux)
}
