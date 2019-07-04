package main

import (
	"net/http"
	"os"
	"log"
	"fmt"
)

func main() {
	port := GetPort()
	log.Println("[-] Listening on...", port)
	http.HandleFunc("/", func (res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "hello, world")
	})

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4747"
		log.Println("[-] No PORT environment variable detected. Setting to ", port)
	}
	return ":" + port
}