package main

import (
	"fmt"

	appsv1beta1 "k8s.io/api/apps/v1beta1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var (
	clientSet *kubernetes.Clientset
)

func initializeClients() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientSet, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Initialized clients!")
}

// GetPods returns the pods from the cluster
func GetPods() string {

	pods, err := clientSet.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	return fmt.Sprintf("There are %d pods in the cluster\n", len(pods.Items))
}

// GetDeployments returns the deployments from the cluster
func GetDeployments() string {

	return "Trying to get Kubernetes deployments!"
}

// GetNamespaces returns all namespaces in the cluster
func GetNamespaces() (string, error) {
	namespaces, err := clientSet.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		return "Cannot get namespaces", err
	}

	var responseString = "Namespaces in cluster: "

	for _, namespace := range namespaces.Items {
		responseString += namespace.Name + ", "
	}

	return responseString, nil
}

// GetClusterInformation return generic information about the cluster
func GetClusterInformation() string {

	var responseString = "Hi, there, here is some information about your cluster: "

	namespaces, err := clientSet.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		return "Cannot get namespaces"
	}
	responseString += fmt.Sprintf("You have %d namespaces ", len(namespaces.Items))

	pods, err := getKubernetesPods()
	if err != nil {
		return "Cannot get pods"
	}
	responseString += fmt.Sprintf("with %d containers, ", len(pods.Items))

	services, err := getKubernetesServices()
	if err != nil {
		return "Cannot get services"
	}
	responseString += fmt.Sprintf("There are %d public services", len(services.Items)) + " and all systems are up and running!"

	return responseString
}

func getKubernetesPods() (*v1.PodList, error) {
	pods, err := clientSet.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return pods, nil
}

func getKubernetesServices() (*v1.ServiceList, error) {
	services, err := clientSet.CoreV1().Services("").List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return services, nil
}

func CreateDeployment() string {

	deploymentsClient := clientSet.AppsV1beta1().Deployments(v1.NamespaceDefault)

	deployment := &appsv1beta1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "azure-summerschool",
		},
		Spec: appsv1beta1.DeploymentSpec{
			Replicas: int32Ptr(2),
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  "azure-summerschool",
							Image: "radumatei/azure-summerschool",
							Ports: []v1.ContainerPort{
								{
									Name:          "http",
									Protocol:      v1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	result, err := deploymentsClient.Create(deployment)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())

	return "Created deployment! Check the public IP!"
}

func int32Ptr(i int32) *int32 { return &i }
