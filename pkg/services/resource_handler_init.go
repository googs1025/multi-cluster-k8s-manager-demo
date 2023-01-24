package services


// ResourcesHandler 这个结构体存放所有informer到的资源，依照集群名称为key，value为该集群该资源对应的map[string]interface{}结构。
// ex: MultiClusterResourceHandler.EventHandlerList[clusterName].EventMap
type ResourcesHandler struct {
	DeploymentHandlerList map[string]*DeploymentHandler
	PodHandlerList        map[string]*PodHandler
	EventHandlerList      map[string]*EventHandler
	NodeHandlerList      map[string]*NodeHandler
	NamespaceHandlerList map[string]*NamespaceHandler


}

func NewResourcesHandler() *ResourcesHandler {
	return &ResourcesHandler{
		DeploymentHandlerList: map[string]*DeploymentHandler{},
		PodHandlerList: map[string]*PodHandler{},
		EventHandlerList: map[string]*EventHandler{},
		NodeHandlerList: map[string]*NodeHandler{},
		NamespaceHandlerList: map[string]*NamespaceHandler{},
	}
}

var MultiClusterResourceHandler *ResourcesHandler


func init() {

	MultiClusterResourceHandler = NewResourcesHandler()
}

