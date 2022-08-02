package deployment

import (
	"context"
	"fmt"
	"github.com/why19970628/k8s-cheatsheet/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"time"
)

func ListPods(client kubernetes.Interface) {
	NamespacesList, err := client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	utils.HandlerCheck(err)
	/// all namespaces list
	for _, namespaces := range NamespacesList.Items {
		fmt.Printf("\n==== namespace: %s ==== \n", namespaces.Name)
		podClient := client.CoreV1().Pods(namespaces.Name)
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
