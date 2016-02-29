package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/gorilla/mux"
)


func main() {
	myRouter := mux.NewRouter()
    myRouter.HandleFunc("/", Home)

	http.Handle("/", myRouter)
	port := os.Getenv("PORT")

	
	fmt.Println("listening on port "+port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

func Home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "hello, world!")
}
