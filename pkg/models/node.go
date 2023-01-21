package models

type NodeModel struct {
	Name string
	IP   string
	HostName string
	OriginLabels  map[string]string
	OriginTaints  map[string]string
	Labels   []string
	Taints   []string
	Capacity *NodeCapacity
	Usage    *NodeUsage
	CreateTime string
}


// NodeCapacity 最大容量
type NodeCapacity struct {
	Cpu  int64
	Memory int64
	Pods int64

}


// NodeUsage 使用情况
type NodeUsage struct {
	Cpu  int64
	Memory int64
	Pods int64
}
