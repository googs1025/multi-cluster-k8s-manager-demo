package init_multi_cluster

import (
	"mutli-cluster-k8s-manager/pkg/models"
)

func ListClusterName() map[string]string {
	clusterNameMap := make(map[string]string, 0)

	for k, v := range MultiClusterController.clusters {
		clusterNameMap[k] = v
	}

	return clusterNameMap
}

func ListClusterNameSlice() []string {
	clusterNameSlice := make([]string, 0)

	for _, v := range MultiClusterController.clusters {

		clusterNameSlice = append(clusterNameSlice, v)
	}

	return clusterNameSlice
}

func ListCluster() []*models.ClusterModel {


	res := make([]*models.ClusterModel, 0)

	for _, value := range MultiClusterController.clusters {
		clusterName := &models.ClusterModel{Name: value}
		res = append(res, clusterName)
	}


	return res
}

func GetClusterName(name string) string {
	if value, ok := MultiClusterController.clusters[name]; ok {
		return value
	}

	return ""
}
