package init_multi_cluster

import (
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
)

var MultiClusterController *MultiClusterClient

type MultiClusterClient struct {
	clients 	map[string]kubernetes.Interface
	facts   	map[string]informers.SharedInformerFactory
	clusters	map[string]string
}

func NewMultiClusterClient() *MultiClusterClient {
	return &MultiClusterClient{
		clients: make(map[string]kubernetes.Interface, 0),
		facts: make(map[string]informers.SharedInformerFactory, 0),
		clusters: make(map[string]string, 0),
	}
}



func init() {

	MultiClusterController = NewMultiClusterClient()
	MultiClusterController.ReadMultiClusterConfig()
}