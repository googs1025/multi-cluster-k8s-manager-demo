package services

import (
	"fmt"
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
	return &PodHandler{
		PodMap: podMap,

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

// event 事件相关的handler
type EventHandler struct {
	EventMap *EventMap  `inject:"-"`
}

func NewEventHandler() *EventHandler {
	eventMap := NewEventMap()

	return &EventHandler{
		EventMap: eventMap,
	}
}

func(e *EventHandler) storeData(obj interface{}, isdelete bool){
	if event, ok := obj.(*corev1.Event); ok {
		key := fmt.Sprintf("%s_%s_%s", event.Namespace, event.InvolvedObject.Kind, event.InvolvedObject.Name)
		if !isdelete {

			e.EventMap.data[key] = event
		} else {
			delete(e.EventMap.data,key)

		}
	}
}

func(e *EventHandler) OnAdd(obj interface{}){
	e.storeData(obj,false)
}
func(e *EventHandler) OnUpdate(oldObj, newObj interface{}){
	e.storeData(newObj,false)
}
func(e *EventHandler) OnDelete(obj interface{}){
	e.storeData(obj,true)
}

//Node相关的handler
type NodeHandler struct {
	NodeMap *NodeMap  `inject:"-"`
	NodeService *NodeService `inject:"-"`
}

func NewNodeHandler() *NodeHandler {
	nodeMap := NewNodeMap()

	return &NodeHandler{
		NodeMap: nodeMap,
	}
}

func(nm *NodeHandler) OnAdd(obj interface{}){
	nm.NodeMap.Add(obj.(*corev1.Node))

	//wscore.ClientMap.SendAll(
	//	gin.H{
	//		"type":"node",
	//		"result":gin.H{"ns":"node",
	//			"data":nm.NodeService.ListAllNodes()},
	//	},
	//)
}
func(nm *NodeHandler) OnUpdate(oldObj, newObj interface{}){
	//重点： 只要update返回true 才会发送 。否则不发送
	if nm.NodeMap.Update(newObj.(*corev1.Node)){
		//wscore.ClientMap.SendAll(
		//	gin.H{
		//		"type":"node",
		//		"result":gin.H{"ns":"node",
		//			"data":nm.NodeService.ListAllNodes()},
		//	},
		//)
	}
}
func(nm *NodeHandler) OnDelete(obj interface{}){
	nm.NodeMap.Delete(obj.(*corev1.Node))

	//wscore.ClientMap.SendAll(
	//	gin.H{
	//		"type":"node",
	//		"result":gin.H{"ns":"node",
	//			"data":nm.NodeService.ListAllNodes()},
	//	},
	//)
}