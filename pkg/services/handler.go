package services

import (
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"log"
)

// DeploymentHandler 使用informer后 回调的方法
type DeploymentHandler struct {
	DeploymentMap *DeploymentMap         `inject:"-"`
	//DeploymentService *DeploymentService `inject:"-"`
}

func NewDeploymentHandler() *DeploymentHandler {
	deploymentMap := NewDeploymentMap()
	//deploymentService := NewDeploymentService()
	return &DeploymentHandler{
		DeploymentMap: deploymentMap,
		//DeploymentService: deploymentService,
	}
}

func (d *DeploymentHandler) OnAdd(obj interface{}) {
	if dep, ok := obj.(*v1.Deployment); ok {
		d.DeploymentMap.Add(dep)
	}
	//ns := obj.(*v1.Deployment).Namespace
	//wscore.ClientMap.SendAll(
	//	gin.H{
	//		"type":"deployments",
	//		"result":gin.H{"ns": ns,"data": d.DeploymentService.ListAll(obj.(*v1.Deployment).Namespace)},
	//	},
	//)
}

func (d *DeploymentHandler) OnDelete(obj interface{}) {
	if dep, ok := obj.(*v1.Deployment); ok {
		d.DeploymentMap.Delete(dep)
	}
	//ns := obj.(*v1.Deployment).Namespace
	//wscore.ClientMap.SendAll(
	//	gin.H{
	//		"type":"deployments",
	//		"result":gin.H{"ns": ns,"data": d.DeploymentService.ListAll(obj.(*v1.Deployment).Namespace)},
	//	},
	//)
}

func (d *DeploymentHandler) OnUpdate(oldObj, newObj interface{}) {
	err := d.DeploymentMap.Update(newObj.(*v1.Deployment))
	if err != nil {
		log.Println(err)
	}
	//} else {
	//	ns := newObj.(*v1.Deployment).Namespace
	//	wscore.ClientMap.SendAll(
	//		gin.H{
	//			"type":"deployments",
	//			"result":gin.H{"ns": ns,"data": d.DeploymentService.ListAll(newObj.(*v1.Deployment).Namespace)},
	//		},
	//	)
	//}
}


// pod相关的回调handler
type PodHandler struct {
	PodMap *PodMap `inject:"-"`
	PodService *PodService `inject:"-"`
}

func NewPodHandler() *PodHandler {
	podMap := NewPodMap()
	//deploymentService := NewDeploymentService()
	return &PodHandler{
		PodMap: podMap,
		//DeploymentService: deploymentService,
	}
}

func(p *PodHandler) OnAdd(obj interface{}){
	p.PodMap.Add(obj.(*corev1.Pod))
	//ns := obj.(*corev1.Pod).Namespace
	//wscore.ClientMap.SendAll(
	//	gin.H{
	//		"type":"pods",
	//		"result":gin.H{
	//			"ns": ns,
	//			"data":p.PodService.ListByNamespace(obj.(*corev1.Pod).Namespace),
	//		},
	//	},
	//)
}

func(p *PodHandler) OnUpdate(oldObj, newObj interface{}){
	err := p.PodMap.Update(newObj.(*corev1.Pod))
	if err != nil {
		log.Println(err)
	}
	//ns := newObj.(*corev1.Pod).Namespace
	//wscore.ClientMap.SendAll(
	//	gin.H{
	//		"type":"pods",
	//		"result":gin.H{
	//			"ns": ns,
	//			"data":p.PodService.ListByNamespace(newObj.(*corev1.Pod).Namespace),
	//		},
	//	},
	//)
}

func(p *PodHandler)	OnDelete(obj interface{}){
	if d, ok := obj.(*corev1.Pod); ok {
		p.PodMap.Delete(d)
	}
	//ns := obj.(*corev1.Pod).Namespace
	//wscore.ClientMap.SendAll(
	//	gin.H{
	//		"type":"pods",
	//		"result":gin.H{
	//			"ns": ns,
	//			"data":p.PodService.ListByNamespace(obj.(*corev1.Pod).Namespace),
	//		},
	//	},
	//)
}