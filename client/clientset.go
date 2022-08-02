package client

import (
	"github.com/why19970628/k8s-cheatsheet/conf"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func NewKubeClient(configPath string) *kubernetes.Clientset {
	// config/kube.conf
	if len(configPath) == 0 {
		configPath = conf.KubeDefaultConfigPath
	}
	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		panic(err)
	}
	// kubernetes.NewForConfig用于生成kubernetes.Clients
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return clientset
}
