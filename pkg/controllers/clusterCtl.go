package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
	"mutli-cluster-k8s-manager/pkg/init_multi_cluster"
	"mutli-cluster-k8s-manager/pkg/services"
)

// deployment控制器
type ClusterCtl struct {
	K8sClient kubernetes.Interface  `inject:"-"`
	PodService *services.PodService `inject:"-"`

}

func NewClusterCtl() *ClusterCtl {
	return &ClusterCtl{
		PodService: services.NewPodService(),
	}
}

// Name 实现deployment controller 框架规范
func (*ClusterCtl) Name() string {
	return "ClusterCtl"
}

// Build 实现deployment controller 路由 框架规范
func (cl *ClusterCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/clusters", cl.List)
}

// List 获取dep列表
func (cl *ClusterCtl) List(c *gin.Context) goft.Json {

	// 配合前端

	clusterList := init_multi_cluster.ListCluster()
	klog.Info("cluster list!")
	return gin.H{
		"code": 20000,
		"data": clusterList,
	}

}
