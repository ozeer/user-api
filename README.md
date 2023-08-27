## user-api go http server

## consul使用
1. 创建consul数据卷
```bash
docker volume create consul-data
```
2. 单机部署运行
```bash
docker run -id --name=consul -p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8500:8500 -p 8600:8600 -v consul-data:/consul/data consul:1.15.4 agent -server -ui -node=n1 -bootstrap-expect=1 -client=0.0.0.0
```
