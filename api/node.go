package api

import (
	"context"
	"fmt"
	"github.com/why19970628/k8s-cheatsheet/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	//corev1 "k8s.io/api/core/v1"
	//"k8s.io/client-go/rest"
)

func ListNodes(ctx context.Context, c kubernetes.Interface) error {
	modeList, err := c.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return err
	}
	utils.HandlerCheck(err)
	for _, node := range modeList.Items {
		fmt.Printf("Name: %v\nPhase: %v\nAddresses: %v\nNodeInfo: %v\n",
			node.Name,
			node.Status.Phase,
			node.Status.Addresses,
			node.Status.NodeInfo,
		)
	}
	return nil
}
