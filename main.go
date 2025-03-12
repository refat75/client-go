package main

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/home/rabbani/.kube/config", "Location of kube config")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		//handle error
		fmt.Printf("Error building kubeconfig from flag: %s", err.Error())
		config, err = rest.InClusterConfig()
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		//handle error
		fmt.Printf("Error building clientset: %s", err.Error())
	}

	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})

	if err != nil {
		// handle error
		fmt.Printf("Error listing pods: %s", err.Error())
	}
	fmt.Printf("Pods Name:\n")
	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}

	deployments, err := clientset.AppsV1().Deployments("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		//handle error
		fmt.Printf("Error listing deployments: %s", err.Error())
	}

	fmt.Printf("Deoloyments Name:\n")
	for _, deployment := range deployments.Items {
		fmt.Println(deployment.Name)
	}
}
