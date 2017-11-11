package main

import (
	"io"
	"net/http"
)

func main() {
	initializeClients()

	http.HandleFunc("/get/cluster", getClusterInformation)

	http.HandleFunc("/update", updateDeployment)
	http.HandleFunc("/create", createDeployment)

	http.ListenAndServe(":80", nil)
}

func getClusterInformation(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, GetClusterInformation())
}

func createDeployment(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, CreateDeployment())
}

func updateDeployment(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, CreateDeployment())
}
