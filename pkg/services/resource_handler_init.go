package services

type ResourcesHandler struct {
	DeploymentHandlerList map[string]*DeploymentHandler
	PodHandlerList        map[string]*PodHandler
}

func NewResourcesHandler() *ResourcesHandler {
	return &ResourcesHandler{
		DeploymentHandlerList: map[string]*DeploymentHandler{},
		PodHandlerList: map[string]*PodHandler{},
	}
}

var MultiClusterResourceHandler *ResourcesHandler


func init() {

	MultiClusterResourceHandler = NewResourcesHandler()
}

