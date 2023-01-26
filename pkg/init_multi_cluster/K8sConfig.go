package init_multi_cluster

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	"k8s.io/metrics/pkg/client/clientset/versioned"
	"log"
	"mutli-cluster-k8s-manager/pkg/models"
	"mutli-cluster-k8s-manager/pkg/services"
)


// 初始化 系统 配置
func InitSysConfig() *models.SysConfig{
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

// metric客户端
func InitMetricClient(config *rest.Config) *versioned.Clientset {

	c, err := versioned.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func InitInformer(initClient kubernetes.Interface, clusterName string) informers.SharedInformerFactory {

	fact := informers.NewSharedInformerFactory(initClient, 0)

	klog.Infof("cluster %s start informer!! \n", clusterName)
	deploymentInformer := fact.Apps().V1().Deployments()
	deploymentInformer.Informer().AddEventHandler(services.MultiClusterResourceHandler.DeploymentHandlerList[clusterName])

	//rsInformer := fact.Apps().V1().ReplicaSets()
	//rsInformer.Informer().AddEventHandler(k.RsHandler)
	//
	podInformer := fact.Core().V1().Pods() //监听pod
	podInformer.Informer().AddEventHandler(services.MultiClusterResourceHandler.PodHandlerList[clusterName])
	//
	nsInformer := fact.Core().V1().Namespaces()
	nsInformer.Informer().AddEventHandler(services.MultiClusterResourceHandler.NamespaceHandlerList[clusterName])
	//
	eventInformer := fact.Core().V1().Events()
	eventInformer.Informer().AddEventHandler(services.MultiClusterResourceHandler.EventHandlerList[clusterName])
	//
	//
	NodeInformer := fact.Core().V1().Nodes()
	NodeInformer.Informer().AddEventHandler(services.MultiClusterResourceHandler.NodeHandlerList[clusterName])

	fact.Start(wait.NeverStop)


	return fact

}