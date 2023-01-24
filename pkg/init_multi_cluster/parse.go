package init_multi_cluster

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"mutli-cluster-k8s-manager/pkg/common"
	"mutli-cluster-k8s-manager/pkg/helpers"
	"mutli-cluster-k8s-manager/pkg/services"
	"strconv"
)

func (m *MultiClusterClient) ReadMultiClusterConfig() {

	p := common.GetWd()
	files, _ := ioutil.ReadDir(p + "/resources")
	path := p + "/resources"
	m.parseConfigs(files, path)
}

func (m *MultiClusterClient) parseConfigs(files []fs.FileInfo, path string) {

	for i, f := range files {

		res := K8sRestConfig(path + "/" + f.Name())

		client := InitClient(res)
		MultiClusterController.clients[res.Host] = client

		versionClient := InitMetricClient(res)
		helpers.MultiVersionClusterController.VersionClients[res.Host] = versionClient

		depHandler := services.NewDeploymentHandler()
		services.MultiClusterResourceHandler.DeploymentHandlerList[res.Host] = depHandler

		podHandler := services.NewPodHandler()
		services.MultiClusterResourceHandler.PodHandlerList[res.Host] = podHandler

		eventHandler := services.NewEventHandler()
		services.MultiClusterResourceHandler.EventHandlerList[res.Host] = eventHandler

		nodeHandler := services.NewNodeHandler()
		services.MultiClusterResourceHandler.NodeHandlerList[res.Host] = nodeHandler

		namespaceHandler := services.NewNamespaceHandler()
		services.MultiClusterResourceHandler.NamespaceHandlerList[res.Host] = namespaceHandler

		informerFactory := InitInformer(client, res.Host)
		MultiClusterController.facts[res.Host] = informerFactory

		clusterName := fmt.Sprintf("cluster%s", strconv.Itoa(i))
		MultiClusterController.setClusterName(res.Host, clusterName)



	}

}

func (m *MultiClusterClient) setClusterName(hostUrl string, clusterName string) {
	m.clusters[clusterName] = hostUrl
}
