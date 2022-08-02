package test

import (
	"fmt"
	"github.com/why19970628/k8s-cheatsheet/api"
	client "github.com/why19970628/k8s-cheatsheet/client"
	"testing"
)

func TestDeploymentSize(t *testing.T) {
	c := client.NewKubeClient("")
	size := api.CountDeploymentSize(c, "aplum-service", "test-ad-report-srv")
	fmt.Println(*size)
}
