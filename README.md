## Project Server

### i18n

### 接口签名

### 延时队列

### 网关

### 注册中心

### 中间键

### 文件上传和处理

### redis锁
### zayn

###运行配置：
```
--registry=etcdv3
```

###网关启动：
```
.\micro.exe --registry=etcdv3 api --handler=proxy --namespace=tb.server --address=:10081
```
