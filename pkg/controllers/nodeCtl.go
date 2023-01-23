package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s.io/client-go/kubernetes"
	"mutli-cluster-k8s-manager/pkg/init_multi_cluster"
	"mutli-cluster-k8s-manager/pkg/services"
)

//@controller
type NodeCtl struct {
	NodeService *services.NodeService `inject:"-"`
	Client *kubernetes.Interface  `inject:"-"`
}

func NewNodeCtl() *NodeCtl {
	return &NodeCtl{
		NodeService: services.NewNodeService(),
	}
}

func(n *NodeCtl) ListAll(c *gin.Context) goft.Json{

	clusterName := c.DefaultQuery("cluster", "cluster0")
	// 配合前端

	clusterName1 := init_multi_cluster.GetClusterName(clusterName)

	return gin.H{
		"code": 20000,
		"data": n.NodeService.ListAllNodes(clusterName1),
	}

}

func(n *NodeCtl)  Build(goft *goft.Goft){
	goft.Handle("GET","/nodes", n.ListAll)
}

func(*NodeCtl) Name() string{
	return "NodeCtl"
}
