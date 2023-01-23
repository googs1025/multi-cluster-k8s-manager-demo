package init_multi_cluster

func ListClusterName() map[string]string {
	clusterNameMap := make(map[string]string, 0)

	for k, v := range MultiClusterController.clusters {
		clusterNameMap[k] = v
	}

	return clusterNameMap
}

func GetClusterName(name string) string {
	if value, ok := MultiClusterController.clusters[name]; ok {
		return value
	}

	return ""
}
