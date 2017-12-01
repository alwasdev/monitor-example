package main

import (
	"fmt"
	"monitor-server"
	"monitor-client"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello World")

	monitor_server.MonitorServer();
	monitor_client.MonitorClient("http://localhost:8080", "HelloApplication")

	log.Fatal(http.ListenAndServe(":8080", nil))

}