package main

import (
	"context"
	"fmt"
	"k8s.io/client-go/tools/clientcmd"
	"k8s-cheat/utils"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	//corev1 "k8s.io/api/core/v1"
	//"k8s.io/client-go/rest"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "config/kube.conf")
	if err != nil {
		panic(err)
	}

	// kubernetes.NewForConfig用于生成kubernetes.Clientset
	clientset, err := kubernetes.NewForConfig(config)
	utils.HandlerCheck(err)
	//
	modeList, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	utils.HandlerCheck(err)


	//podClient := clientset.CoreV1().Pods(corev1.NamespaceDefault)
	//// 内部实际使用的也是RESTClient发起请求
	//list, err := podClient.List(context.TODO(), metav1.ListOptions{Limit: 500})
	//if err != nil {
	//	panic(err)
	//}
	//
	for _, node := range modeList.Items {
		fmt.Printf("Name: %v\nPhase: %v\nAddresses: %v\nNodeInfo: %v\n",
			node.Name,
			node.Status.Phase,
			node.Status.Addresses,
			node.Status.NodeInfo,
		)
	}
}

//var clientset *kubernetes.Clientset
//
//func kubeInitOutCluster() error {
//	kubeconfig := "C:\\config"
//	// build config from kubeconfig file
//	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
//	if err != nil {
//		return err
//	}
//	// create the clientset
//	clientset, err = kubernetes.NewForConfig(restConfig)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func kubeInitInCluster() error {
//	restConfig, err := rest.InClusterConfig()
//	if err != nil {
//		return err
//	}
//	// create the clientset
//	clientset, err = kubernetes.NewForConfig(restConfig)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func kubeInit() error {
//	kubeconfig := "C:\\config"
//	// build config from kubeconfig file
//	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
//	if err != nil {
//		restConfig, err = rest.InClusterConfig()
//		if err != nil {
//			return err
//		}
//	}
//	// create the clientset
//	clientset, err = kubernetes.NewForConfig(restConfig)
//	if err != nil {
//		return err
//	}
//	return nil
//}

