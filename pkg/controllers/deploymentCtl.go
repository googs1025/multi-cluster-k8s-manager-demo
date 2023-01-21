package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s.io/client-go/kubernetes"
	"mutli-cluster-k8s-manager/pkg/init_multi_cluster"
	"mutli-cluster-k8s-manager/pkg/services"
)

// deployment控制器
type DeploymentCtl struct {
	K8sClient kubernetes.Interface                `inject:"-"`
	DeploymentService *services.DeploymentService `inject:"-"`
}

func NewDeploymentCtl() *DeploymentCtl {
	return &DeploymentCtl{
		DeploymentService: services.NewDeploymentService(),
	}
}

// Name 实现deployment controller 框架规范
func (*DeploymentCtl) Name() string {
	return "DeploymentCtl"
}

// Build 实现deployment controller 路由 框架规范
func (d *DeploymentCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/deployments", d.List)
}

// List 获取dep列表
func (d *DeploymentCtl) List(c *gin.Context) goft.Json {
	namespace := c.DefaultQuery("namespace", "default") // 请求： GET /deployments?namespace=xxxxxxx
	clusterName := c.DefaultQuery("cluster", "cluster0")
	// 配合前端

	clusterName1 := init_multi_cluster.GetClusterName(clusterName)

	return gin.H{
		"code": 20000,
		"data": d.DeploymentService.ListAll(namespace, clusterName1),
	}

}

//func (d *DeploymentCtl) List(c *gin.Context) {
//	namespace := c.DefaultQuery("namespace", "default") // 请求： GET /deployments?namespace=xxxxxxx
//	clusterName := c.DefaultQuery("cluster", "cluster0")
//	// 配合前端
//	//return gin.H{
//	//	"code": 20000,
//	//	"data": d.DeploymentService.ListAll(namespace, clusterName),
//	//}
//
//	clusterName1 := init_multi_cluster.GetClusterName(clusterName)
//
//	deploymentList := d.DeploymentService.ListAll(namespace, clusterName1)
//	c.JSON(400, gin.H{"data": deploymentList})
//
//}





