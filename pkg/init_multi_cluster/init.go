package init_multi_cluster

import (
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"mutli-cluster-k8s-manager/pkg/helpers"
)

// MultiClusterController
var MultiClusterController *MultiClusterClient

// MultiClusterClient 全局存放多集群的 client informer cluster名的数据结构。
type MultiClusterClient struct {
	clients 		map[string]kubernetes.Interface
	facts   		map[string]informers.SharedInformerFactory
	clusters		map[string]string
}

func NewMultiClusterClient() *MultiClusterClient {
	return &MultiClusterClient{
		clients: make(map[string]kubernetes.Interface, 0),
		facts: make(map[string]informers.SharedInformerFactory, 0),
		clusters: make(map[string]string, 0),
	}
}



func init() {
	helpers.MultiVersionClusterController = helpers.NewMultiClusterVersionClient()
	MultiClusterController = NewMultiClusterClient()
	// 初始化多集群资源。
	MultiClusterController.ReadMultiClusterConfig()
}