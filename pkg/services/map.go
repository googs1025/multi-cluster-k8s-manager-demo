package services

import (
	"fmt"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"mutli-cluster-k8s-manager/pkg/models"
	"reflect"
	"sort"
)

// DeploymentMap 使用informer监听资源变化后，事件变化加入map中
type DeploymentMap struct {
	data map[string]interface{} // [key string] []*v1.Deployment    key=>namespace
}

func NewDeploymentMap() *DeploymentMap {
	data := make(map[string]interface{}, 0)
	return &DeploymentMap{data: data}
}

func (d *DeploymentMap) Add(deployment *v1.Deployment) {

	if deploymentList, ok := d.data[deployment.Namespace]; ok {
		deploymentList = append(deploymentList.([]*v1.Deployment), deployment)
		d.data[deployment.Namespace] = deploymentList
	} else {
		newDeploymentList := make([]*v1.Deployment, 0)
		newDeploymentList = append(newDeploymentList, deployment)
		d.data[deployment.Namespace] = newDeploymentList
	}

}

func (d *DeploymentMap) Delete(deployment *v1.Deployment) {

	if deploymentList, ok := d.data[deployment.Namespace]; ok {
		list := deploymentList.([]*v1.Deployment)
		for k, needDeleteDeployment := range list {
			if deployment.Name == needDeleteDeployment.Name {
				newList := append(list[:k], list[k+1:]...)
				d.data[deployment.Namespace] = newList
				break
			}
		}
	}


}

func (d *DeploymentMap) Update(deployment *v1.Deployment) error {

	if deploymentList, ok := d.data[deployment.Namespace]; ok {
		list := deploymentList.([]*v1.Deployment)
		for k, needUpdateDeployment := range list {
			if deployment.Name == needUpdateDeployment.Name {
				list[k] = deployment
			}
		}
		return nil

	}

	return fmt.Errorf("deployment-%s update error", deployment.Name)

}

// ListDeploymentByNamespace 内存中读取deploymentList
func (d *DeploymentMap) ListDeploymentByNamespace(namespace string) ([]*v1.Deployment, error) {
	if deploymentList, ok := d.data[namespace]; ok {
		return deploymentList.([]*v1.Deployment), nil
	}

	return []*v1.Deployment{}, nil
}

// GetDeployment 内存中读取deployment
func (d *DeploymentMap) GetDeployment(namespace string, deploymentName string) (*v1.Deployment, error) {
	if deploymentList, ok := d.data[namespace]; ok {
		list := deploymentList.([]*v1.Deployment)
		for _, dep := range list {
			if dep.Name == deploymentName {
				return dep, nil
			}
		}
	}

	return nil, fmt.Errorf("get deployment error, not found")
}

// 保存Pod集合
type PodMap struct {
	data map[string]interface{}  // [key string] []*v1.Pod    key=>namespace
}

func NewPodMap() *PodMap {
	data := make(map[string]interface{}, 0)
	return &PodMap{data: data}
}

func(p *PodMap) ListByNamespace(namespace string) []*corev1.Pod {
	if podList, ok := p.data[namespace]; ok {
		return podList.([]*corev1.Pod)
	}
	return nil
}

func(p *PodMap) GetPod(namespace string, podName string) *corev1.Pod{
	if podList, ok := p.data[namespace]; ok {
		list := podList.([]*corev1.Pod)
		for _, pod := range list {
			if pod.Name == podName {
				return pod
			}
		}
	}
	return nil
}

func (p *PodMap) Add(pod *corev1.Pod) {
	if podList, ok := p.data[pod.Namespace]; ok {
		podList = append(podList.([]*corev1.Pod), pod)
		p.data[pod.Namespace] = podList
	} else {
		newPodList := make([]*corev1.Pod, 0)
		newPodList = append(newPodList, pod)
		p.data[pod.Namespace] = newPodList

	}
}

func (p *PodMap) Update(pod *corev1.Pod) error {
	if podList, ok := p.data[pod.Namespace]; ok {
		list := podList.([]*corev1.Pod)
		for i, needUpdatePod := range list {
			if needUpdatePod.Name == pod.Name {
				list[i] = pod
			}
		}
		return nil
	}
	return fmt.Errorf("pod-%s update error",pod.Name)
}

func (p *PodMap) Delete(pod *corev1.Pod) {
	if podList, ok := p.data[pod.Namespace]; ok {
		list := podList.([]*corev1.Pod)
		for i, needDeletePod := range list {
			if needDeletePod.Name == pod.Name {
				newList := append(list[:i], list[i+1:]...)
				p.data[pod.Namespace] = newList
				break
			}
		}
	}
}

// ListByLabels 根据标签获取 POD列表
func(p *PodMap) ListByLabels(namespace string,labels []map[string]string) ([]*corev1.Pod,error){
	ret := make([]*corev1.Pod, 0)
	if podList, ok := p.data[namespace]; ok {
		list := podList.([]*corev1.Pod)
		for _, pod := range list {
			for _, label := range labels {
				if reflect.DeepEqual(pod.Labels, label) {  //标签完全匹配
					ret = append(ret, pod)
				}
			}
		}
		return ret,nil
	}
	return nil, fmt.Errorf("pods not found ")
}


// GetNum 根据节点名称 获取pods数量
func(p *PodMap) GetNum(nodeName string) (num int){

	for _, v := range p.data {
		podList := v.([]*corev1.Pod)
		for _, pod := range podList {
			if pod.Spec.NodeName == nodeName {
				num++
			}
		}
	}


	return
}

// event 事件map 相关
// EventSet 集合 用来保存事件, 只保存最新的一条
type EventMap struct {
	data map[string]interface{} // [key string] *v1.Event
	// key=>namespace+"_"+kind+"_"+name 这里的name 不一定是pod ,这样确保唯一
}

func NewEventMap() *EventMap {
	data := make(map[string]interface{}, 0)
	return &EventMap{data: data}
}

func(e *EventMap) GetMessage(ns string, kind string, name string) string {
	key := fmt.Sprintf("%s_%s_%s", ns, kind, name)
	if v, ok := e.data[key]; ok {
		return v.(*corev1.Event).Message
	}

	return ""
}


// node map
type NodeMap struct {
	data map[string]interface{} // [key string] v1.Node
}

func NewNodeMap() *NodeMap {
	data := make(map[string]interface{}, 0)
	return &NodeMap{data: data}
}

func(n *NodeMap) Get(name string) *corev1.Node{
	if node, ok := n.data[name]; ok{
		return node.(*corev1.Node)
	}
	return nil
}

func(n *NodeMap) Add(item *corev1.Node){

	n.data[item.Name] = item
}

func(n *NodeMap) Update(item *corev1.Node) bool {
	n.data[item.Name] = item
	return true
}

func(n *NodeMap) Delete(node *corev1.Node){

	delete(n.data, node.Name)
}

func(n *NodeMap) ListAll() []*corev1.Node {

	ret := make([]*corev1.Node, 0)

	for _, v := range n.data {
		node := v.(*corev1.Node)
		ret = append(ret, node)
	}

	return ret
}

// namespace map
type NamespaceMap struct {
	data map[string]interface{} // [key string] *corev1.Namespace
}

func NewNamespaceMap() *NamespaceMap {
	data := make(map[string]interface{}, 0)
	return &NamespaceMap{data: data}
}

func (n *NamespaceMap) Add(ns *corev1.Namespace) {
	n.data[ns.Name] = ns
}

func (n *NamespaceMap) Update(ns *corev1.Namespace) {
	n.data[ns.Name] = ns
}

func (n *NamespaceMap) Delete(ns *corev1.Namespace) {
	delete(n.data, ns.Name)
}

func (n *NamespaceMap) Get(namespace string) *corev1.Namespace {
	if item, ok := n.data[namespace]; ok {
		ns := item.(*corev1.Namespace)
		return ns
	}
	return nil
}

func (n *NamespaceMap) ListAll() []*models.NamespaceModel {

	items := convertToMapItems(n.data)
	sort.Sort(items)

	res := make([]*models.NamespaceModel, 0)

	for _, value := range items {
		nsName := &models.NamespaceModel{Name: value.key}
		res = append(res, nsName)
	}


	return res
}


type MapItems []*MapItem
type MapItem struct {
	key string
	value interface{}
}
//把sync.map  转为 自定义切片
func convertToMapItems(m map[string]interface{}) MapItems{
	items := make(MapItems, 0)

	for key, value := range m {
		items = append(items, &MapItem{key:key,value:value})
	}

	return items
}

func(m MapItems) Len() int{
	return len(m)
}
func(m MapItems) Less(i, j int) bool{
	return m[i].key < m[j].key
}
func(m MapItems) Swap(i, j int){
	m[i], m[j] = m[j], m[i]
}
