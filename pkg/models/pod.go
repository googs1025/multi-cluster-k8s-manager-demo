package models


type Pod struct {
	Name 	   string
	Namespace  string
	Images 	   string
	NodeName   string
	Ip   	   []string
	Phase      string
	IsReady    bool
	Message    string
	CreateTime string
}

type ContainerModel struct {
	Name string
}