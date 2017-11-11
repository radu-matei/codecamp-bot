package main

import (
	"io"
	"net/http"
)

func main() {
	initializeClients()

	http.HandleFunc("/get/pods", getPods)
	http.HandleFunc("/get/deployments", getDeployments)
	http.HandleFunc("/get/namespaces", getNamespaces)
	http.HandleFunc("/get/cluster", getClusterInformation)

	http.HandleFunc("/create", createDeployment)

	http.ListenAndServe(":80", nil)
}

func getPods(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, GetPods())
}

func getDeployments(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, GetDeployments())
}

func getNamespaces(writer http.ResponseWriter, request *http.Request) {
	ns, _ := GetNamespaces()
	io.WriteString(writer, ns)
}

func getClusterInformation(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, GetClusterInformation())
}

func createDeployment(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, CreateDeployment())
}
