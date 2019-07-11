package main

import (
	"deploysimple/driver"
	_ "deploysimple/driver"
	"deploysimple/routers"
	"log"
	"os"
)

func main() {
	defer driver.GetConn()
	port := GetPort()
	log.Println("[-] Listening on...", port)
	routers.Run(port)
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4747"
		log.Println("[-] No PORT environment variable detected. Setting to", port)
	}
	return ":" + port
}
