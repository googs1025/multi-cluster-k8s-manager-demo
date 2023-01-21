package main

import (
	"mutli-cluster-k8s-manager/pkg/controllers"
)

var (
	depController *controllers.DeploymentCtl
)

func initController() {
	depController = controllers.NewDeploymentCtl()
}

//func register(router *gin.Engine) {
//	initController()
//	api := router.Group("/api")
//
//
//	{
//		manager := api.Group("/manager")
//		{
//			deployments := manager.Group("/deployments")
//
//			{
//				deployments.GET("", depController.List)
//			}
//		}
//	}
//}
