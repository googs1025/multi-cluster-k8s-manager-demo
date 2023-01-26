package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
	"mutli-cluster-k8s-manager/pkg/init_multi_cluster"
	"mutli-cluster-k8s-manager/pkg/services"
)

// pod控制器
type PodCtl struct {
	K8sClient kubernetes.Interface  `inject:"-"`
	PodService *services.PodService `inject:"-"`

}

func NewPodCtl() *PodCtl {
	return &PodCtl{
		PodService: services.NewPodService(),
	}
}

// Name 实现pod controller 框架规范
func (*PodCtl) Name() string {
	return "PodCtl"
}

// Build 实现pod controller 路由 框架规范
func (p *PodCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/pods", p.List)
}

// List 获取dep列表
func (p *PodCtl) List(c *gin.Context) goft.Json {
	namespace := c.DefaultQuery("namespace", "default")  // 请求：GET /pods?namespace=xxx
	clusterName := c.DefaultQuery("cluster", "cluster0") // 请求：GET /pods?namespace=xxx&cluster=xxx
	// 配合前端

	clusterName1 := init_multi_cluster.GetClusterName(clusterName)
	klog.Info("pod list!")
	return gin.H{
		"code": 20000,
		"data": p.PodService.ListByNamespace(namespace, clusterName1),
	}

}
