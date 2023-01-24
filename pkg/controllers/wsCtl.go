package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
	"mutli-cluster-k8s-manager/pkg/wscore"
)

//@Controller
type WsCtl struct {
	Client *kubernetes.Clientset  `inject:"-"`
	Config *rest.Config  `inject:"-"`
}


func NewWsCtl() *WsCtl {
	return &WsCtl{}
}

func(w *WsCtl) Connect(c *gin.Context) string  {
	client, err := wscore.Upgrader.Upgrade(c.Writer,c.Request,nil)  //升级
	if err != nil {
		log.Println(err)
		return err.Error()
	} else {
		wscore.ClientMap.Store(client)
		return "success"
	}

}


func(w *WsCtl)  Build(goft *goft.Goft){
	goft.Handle("GET","/ws", w.Connect)
}

func(w *WsCtl) Name() string{
	return "WsCtl"
}