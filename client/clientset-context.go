package main

import (
	"context"
	"fmt"
	"k8s-cheat/conf"
	"k8s-cheat/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"time"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", conf.KubeConfigPath)
	utils.HandlerCheck(err)
	// kubernetes.NewForConfig用于生成kubernetes.Clients
	clientset, err := kubernetes.NewForConfig(config)
	utils.HandlerCheck(err)
	NamespacesList , err:= clientset.CoreV1().Namespaces().List(context.TODO(),  metav1.ListOptions{})
	utils.HandlerCheck(err)
	/// all namespaces list
	for _, namespaces := range NamespacesList.Items {
		//fmt.Println(namespaces.Name)
		fmt.Printf("\n==== namespace: %s ==== \n", namespaces.Name)
		podClient := clientset.CoreV1().Pods(namespaces.Name)
		// 内部实际使用的也是RESTClient发起请求

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		list, err := podClient.List(ctx, metav1.ListOptions{Limit: 500})
		if err != nil {
			panic(err)
		}

		for _, d := range list.Items {
			fmt.Printf("namespace: %v\tname: %v\tstatus: %+v\tnode: %+v\tnode: %+v\t\n",
				d.Namespace,
				d.Name,
				d.Status.Phase,
				d.Spec.NodeName, // node
				time.Now().Sub(d.CreationTimestamp.Time).Seconds(),
			)
		}
	}

}