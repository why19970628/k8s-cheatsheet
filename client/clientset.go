package main

import (
	"k8s-cheat/conf"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)


var Clientset *kubernetes.Clientset



func kubeInitOutCluster()  {
	config, err := clientcmd.BuildConfigFromFlags("", conf.KubeConfigPath)
	if err != nil {
		panic(err)
	}
	// kubernetes.NewForConfig用于生成kubernetes.Clients
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	Clientset = clientset
}