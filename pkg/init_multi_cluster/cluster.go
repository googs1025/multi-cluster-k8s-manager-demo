package init_multi_cluster


func ListClusterName() []string {
	clusterNameList := make([]string, 0)

	for _, v := range MultiClusterController.clusters {
		clusterNameList = append(clusterNameList, v)
	}

	return clusterNameList
}

func GetClusterName(name string) string {
	if value, ok := MultiClusterController.clusters[name]; ok {
		return value
	}

	return ""
}
