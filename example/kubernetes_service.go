package main

import (
	"flag"
	"fmt"
	apiV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	coreV1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
)

type srv struct {
	service coreV1.ServiceInterface
}

func main() {
	var s srv
	kubernetesConfig := flag.String("kubeconfig", "/Users/heyang/.kube/config", "upload kubeconfig direct")
	config, err := clientcmd.BuildConfigFromFlags("", *kubernetesConfig)
	if err != nil {
		fmt.Println("11")
		panic(err.Error())
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	s.service = clientSet.CoreV1().Services(apiV1.NamespaceDefault)
	s.list()
}

func (s *srv) list() {
	serviceList, err := s.service.List(metaV1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, i := range serviceList.Items {
		fmt.Printf("%s \n", i.Name)
	}
}
