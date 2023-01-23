## multi-cluster-k8s-manager-demo
## k8s多集群后端接口展示demo
### 项目思路：
对多个集群采用informer的方式管理k8s资源，server中使用key为集群名的方式区分不同集群。
![](https://github.com/googs1025/multi-cluster-k8s-manager-demo/blob/main/image/%E6%B5%81%E7%A8%8B%E5%9B%BE.jpg?raw=true)


### 启动步骤：
1. 目录下创建一个resources文件，把集群的.kube/config文件复制一份放入(记得cluster server需要改成"公网ip")。
2. 可以放置多个.kube/config配置文件，支持多集群list查询。
3. go run main.go启动服务
```
➜  mutli-cluster-k8s-manager git:(main) ✗ go run main.go
cluster https://xxxxxx:6443 start informer!!
cluster https://xxxxxx:6443 start informer!!
cluster https://xxxxxx start informer!!
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

2023/01/23 16:05:39 open /Users/zhenyu.jiang/go/src/golanglearning/new_project/mutli-cluster-k8s-manager/application.yaml: no such file or directory
[GIN-debug] GET    /deployments              --> github.com/shenyisyn/goft-gin/goft.JsonResponder.RespondTo.func1 (3 handlers)
[GIN-debug] GET    /pods                     --> github.com/shenyisyn/goft-gin/goft.JsonResponder.RespondTo.func1 (3 handlers)
[GIN-debug] GET    /nodes                    --> github.com/shenyisyn/goft-gin/goft.JsonResponder.RespondTo.func1 (3 handlers)
[GIN-debug] GET    /clusters                 --> github.com/shenyisyn/goft-gin/goft.JsonResponder.RespondTo.func1 (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8080

```
4. 接口调用测试：
注意：调用接口时，如果不写query，cluster默认取第一个集群的结果，namespace默认取default
```
http://localhost:8080/pods
http://localhost:8080/deployments
http://localhost:8080/clusters
http://localhost:8080/nodes?cluster=cluster1
http://localhost:8080/pods?cluster=cluster0
http://localhost:8080/pods?cluster=cluster0&namespace=default
http://localhost:8080/deployments?cluster=cluster1&namespace=default
```

### RoadMap
1. 本项目基于多集群实现，目前预计支持多个资源的展示与查询。
2. 未来会加入前端，方便可视化。
