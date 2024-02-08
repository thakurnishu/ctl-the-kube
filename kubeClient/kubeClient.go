package kubeclient

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func kubeconfigHome() string {

	// To get kubeconfig file location
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("ERROR getting UserHome dir \n%v\n", err.Error())
	}

	kubeconfigPath := filepath.Join(homeDir, ".kube", "config")

	return kubeconfigPath
}

func getTypedClientSet(kubeconfig *string) *kubernetes.Clientset {

	var config *rest.Config

	if _, err := os.Stat(*kubeconfig); err == nil {

		config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			log.Fatalf("[ERROR Building Config from Flag] \n%v\n", err.Error())
		}
	} else {

		config, err = rest.InClusterConfig()
		if err != nil {
			log.Fatalf("ERROR getting Config from K8S Cluster \n%v\n", err.Error())
		}
	}

	// Typed ClientSet
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("ERROR Creating Typed ClientSet from Config \n%v\n", err.Error())
	}

	return clientSet
}

func GetClient() *kubernetes.Clientset {
	kubeconfigHome := kubeconfigHome()
	// File location from user ( --kubeconfig ) or homeDir/.kube/config
	kubeconfig := flag.String("kubeconfig", kubeconfigHome, "Location of KubeConfig file")

	// Parse flags once
	flag.Parse()

	clientset := getTypedClientSet(kubeconfig)

    return clientset
}
