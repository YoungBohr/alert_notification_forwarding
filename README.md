# Aliyun__Alert_Notification_Forwarding

## Install

### Git clone

```shell
mkdir -p ${GOPATH}/src/aliyun/alert_notification_forwarding
```

```shell
cd ${GOPATH}/src/alert_notification_forwarding
git clone https://gitlab.taidu8.com/devops/scripts.git
```

### Build

```shell
go test
go build -o main cmd/main/main.go
chmod +x main
docker build -t registry.cn-zhangjiakou.aliyuncs.com/data100/alert:latest .
```

## Quick Start

```shell
docker run -d -p 19099:19099 registry.cn-zhangjiakou.aliyuncs.com/data100/alert:latest
```
