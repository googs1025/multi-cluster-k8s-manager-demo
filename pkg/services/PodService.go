package services

import (
	"mutli-cluster-k8s-manager/pkg/models"
)

//@Service
type PodService struct {
	PodMap *PodMap `inject:"-"`
	Common *CommonService `inject:"-"`
}

func NewPodService() *PodService {
	return &PodService{}
}

// GetPodContainer 获取Pod容器列表
func(p *PodService) GetPodContainer(ns , podName string) []*models.ContainerModel{
	ret := make([]*models.ContainerModel, 0)
	pod := p.PodMap.GetPod(ns, podName)
	if pod != nil{
		for _, c := range pod.Spec.Containers {
			ret = append(ret, &models.ContainerModel{
				Name: c.Name,
			})
		}
	}
	return ret
}

func(p *PodService) ListByNamespace(namespace string, clusterName string) []*models.Pod {
	res := make([]*models.Pod, 0)

	podList := MultiClusterResourceHandler.PodHandlerList[clusterName].PodMap.ListByNamespace(namespace)
	//goft.Error(err)

	for _, pod := range podList {
		res = append(res, &models.Pod{
			Name: pod.Name,
			Namespace: pod.Namespace,
			Images: p.Common.GetImagesByPod(pod.Spec.Containers), // 查看pod镜像
			NodeName: pod.Spec.NodeName,
			Phase: string(pod.Status.Phase),
			Ip: []string{pod.Status.PodIP, pod.Status.HostIP},
			IsReady: p.Common.PodIsReady(pod), // 查看pod是否ready
			Message: MultiClusterResourceHandler.EventHandlerList[clusterName].EventMap.GetMessage(pod.Namespace,"Pod", pod.Name),
			CreateTime: pod.CreationTimestamp.Format("2006-01-02 15:04:05"),

		})
	}

	return res

}
