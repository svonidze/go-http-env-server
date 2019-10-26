package main

import (
	"fmt"
	"net/http"
	"os"
)

func dealwithErr(err error) {
	if err != nil {
		fmt.Println(err)
		//os.Exit(-1)
	}
}

func SayName(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	dealwithErr(err)

	w.Write([]byte("Hello, I'm a machine and my name is " + name))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", SayName)

	http.ListenAndServe(":8080", mux)

}
