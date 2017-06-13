package main

import (
	"fmt"
	"log"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalln(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	}
	ings, err := clientset.ExtensionsV1beta1().Ingresses("").List(v1.ListOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	for _, ing := range ings.Items {
		fmt.Println(ing.ObjectMeta.Namespace, ing.ObjectMeta.Name)
	}
}
