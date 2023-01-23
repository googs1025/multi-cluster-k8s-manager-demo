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
	Cpu  float64
	Memory float64
	Pods int
}

func NewNodeCapacity(cpu int64, memory int64, pods int64) *NodeCapacity {
	return &NodeCapacity{Cpu: cpu, Memory: memory, Pods: pods}
}

func NewNodeUsage(pods int, cpu float64, memory float64) *NodeUsage {
	return &NodeUsage{Pods: pods, Cpu: cpu, Memory: memory}
}
