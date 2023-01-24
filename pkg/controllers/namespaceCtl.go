package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"mutli-cluster-k8s-manager/pkg/init_multi_cluster"
	"mutli-cluster-k8s-manager/pkg/services"
)

type NamespaceCtl struct {
	NamespaceService *services.NamespaceService `inject:"-"`
}

func NewNamespaceCtl() *NamespaceCtl {
	return &NamespaceCtl{
		NamespaceService: services.NewNamespaceService(),
	}
}

func (n *NamespaceCtl) ListAll(c *gin.Context) goft.Json {

	clusterName := c.DefaultQuery("cluster", "cluster0")
	// 配合前端

	clusterName1 := init_multi_cluster.GetClusterName(clusterName)

	return gin.H{
		"code": 20000,
		"data": n.NamespaceService.ListAllNamespaces(clusterName1),
	}
}

func (n *NamespaceCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/namespaces", n.ListAll)
}

func (*NamespaceCtl) Name() string {
	return "NamespaceCtl"
}

