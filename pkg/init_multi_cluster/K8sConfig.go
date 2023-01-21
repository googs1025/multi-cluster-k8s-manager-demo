package init_multi_cluster

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"mutli-cluster-k8s-manager/pkg/models"
	"mutli-cluster-k8s-manager/pkg/services"
)

type K8sConfig struct {
	DepHandler []*services.DeploymentHandler `inject:"-"`
	//RsHandler []*services.RsHandler `inject:"-"`
	//PodHandler []*services.PodHandler        `inject:"-"`
	//NsHandler  []*services.NamespaceHandler  `inject:"-"`
	//EventHandler []*services.EventHandler `inject:"-"`
	//NodeHandler []*services.NodeHandler `inject:"-"`
}

func NewK8sConfig() *K8sConfig {
	return &K8sConfig{}
}

// 初始化 系统 配置
func(*K8sConfig) InitSysConfig() *models.SysConfig{
	b, err := ioutil.ReadFile("app.yaml")
	if err != nil {
		log.Fatal(err)
	}
	config := &models.SysConfig{}
	err = yaml.Unmarshal(b,config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}

func K8sRestConfig(path string) *rest.Config{
	config, err := clientcmd.BuildConfigFromFlags("", path)
	config.Insecure = true
	if err != nil {
		log.Fatal(err)
	}
	return config
}


func InitClient(config *rest.Config) kubernetes.Interface {


	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	
	return client
}

func InitInformer(initClient kubernetes.Interface, clusterName string) informers.SharedInformerFactory {

	fact := informers.NewSharedInformerFactory(initClient, 0)
	fmt.Printf("cluster %s start informer!! \n", clusterName)
	deploymentInformer := fact.Apps().V1().Deployments()
	deploymentInformer.Informer().AddEventHandler(services.MultiClusterResourceHandler.DeploymentHandlerList[clusterName])

	//rsInformer := fact.Apps().V1().ReplicaSets()
	//rsInformer.Informer().AddEventHandler(k.RsHandler)
	//
	podInformer := fact.Core().V1().Pods() //监听pod
	podInformer.Informer().AddEventHandler(services.MultiClusterResourceHandler.PodHandlerList[clusterName])
	//
	//nsInformer := fact.Core().V1().Namespaces()
	//nsInformer.Informer().AddEventHandler(k.NsHandler)
	//
	//eventInformer := fact.Core().V1().Events()
	//eventInformer.Informer().AddEventHandler(k.EventHandler)
	//
	//
	//NodeInformer := fact.Core().V1().Nodes()
	//NodeInformer.Informer().AddEventHandler(k.NodeHandler)

	fact.Start(wait.NeverStop)


	return fact

}