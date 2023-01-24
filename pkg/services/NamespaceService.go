package services

import (
	"mutli-cluster-k8s-manager/pkg/models"
)

// @service
type NamespaceService struct {
	NamespaceMap *NamespaceMap `inject:"-"`
}

func NewNamespaceService() *NamespaceService {
	return &NamespaceService{}
}

// ListAllNamespaces 显示所有namespace
func (ns *NamespaceService) ListAllNamespaces(clusterName string) []*models.NamespaceModel{

	list := MultiClusterResourceHandler.NamespaceHandlerList[clusterName].NamespaceMap.ListAll()

	return list
}

