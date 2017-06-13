package main

import (
	"flag"
	"fmt"
	"log"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var flagNamespace = flag.String("ns", "", "namespace")

func main() {
	flag.Parse()
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalln(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	}
	services, err := clientset.CoreV1().Services(*flagNamespace).List(v1.ListOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	for _, ing := range services.Items {
		fmt.Println(ing.ObjectMeta.Namespace, ing.ObjectMeta.Name)
	}
}
