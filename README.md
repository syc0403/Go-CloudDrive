# Go-CloudDrive

一个基于 Go 语言实现的分布式文件上传服务（仿百度网盘），实现个人文件存储云盘。

### 安装库

如下：
```shell

```

### 进度情况

* [x] 简单的文件上传服务
* [ ] mysql存储文件元数据
* [ ] 账号系统, 注册/登录/查询用户或文件数据
* [ ] 基于帐号的文件操作接口
* [ ] 文件秒传功能
* [ ] 文件分块上传/断点续传功能
* [ ] 搭建及使用Ceph对象存储集群
* [ ] 使用阿里云OSS对象存储服务
* [ ] 使用RabbitMQ实现异步任务队列
* [ ] 微服务化(API网关, 服务注册, RPC通讯)
* [ ] CI/CD(持续集成)

### 参考项目

- yddeng：[fileCloud](https://github.com/yddeng/filecloud)
- blankjee：[file-storage-system](https://github.com/blankjee/file-storage-system)
### 参考资料

- Go入门: [语言之旅](https://tour.go-zh.org/welcome/1)
- MySQL: [偶然翻到的一位大牛翻译的使用手册](https://chhy2009.github.io/document/mysql-reference-manual.pdf)
- Redis: [命令手册](http://redisdoc.com/)
- Ceph: [中文社区](http://ceph.org.cn/) [中文文档](http://docs.ceph.org.cn/)
- RabbitMQ: [英文官方](http://www.rabbitmq.com/getstarted.html) [一个中文版文档](http://rabbitmq.mr-ping.com/)
- 阿里云OSS: [文档首页](https://help.aliyun.com/product/31815.html?spm=a2c4g.750001.3.1.47287b13LQI3Ah)
- gRPC: [官方文档中文版](http://doc.oschina.net/grpc?t=56831)
- k8s: [中文社区](https://www.kubernetes.org.cn/docs)
