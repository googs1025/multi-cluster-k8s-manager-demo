package init_multi_cluster

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"mutli-cluster-k8s-manager/pkg/common"
	"mutli-cluster-k8s-manager/pkg/services"
	"strconv"
)

func (m *MultiClusterClient) ReadMultiClusterConfig() {

	p := common.GetWd()
	files, _ := ioutil.ReadDir(p + "/resources")
	path := "/Users/zhenyu.jiang/go/src/golanglearning/new_project/mutli-cluster-k8s-manager/resources"
	m.parseConfigs(files, path)
}

func (m *MultiClusterClient) parseConfigs(files []fs.FileInfo, path string) {

	for i, f := range files {

		res := K8sRestConfig(path + "/" + f.Name())

		client := InitClient(res)
		MultiClusterController.clients[res.Host] = client

		depHandler := services.NewDeploymentHandler()
		services.MultiClusterResourceHandler.DeploymentHandlerList[res.Host] = depHandler

		podHandler := services.NewPodHandler()
		services.MultiClusterResourceHandler.PodHandlerList[res.Host] = podHandler

		informerFactory := InitInformer(client, res.Host)
		MultiClusterController.facts[res.Host] = informerFactory

		clusterName := fmt.Sprintf("cluster%s", strconv.Itoa(i))
		MultiClusterController.setClusterName(res.Host, clusterName)



	}

}

func (m *MultiClusterClient) setClusterName(hostUrl string, clusterName string) {
	m.clusters[clusterName] = hostUrl
}
