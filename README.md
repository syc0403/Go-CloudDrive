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

### 一些问题
```shell
sudo docker container run -p 3307:3306 --name mysql-master -v /home/mysql/master.cnf:/etc/mysql/conf.d/master.cnf -e MYSQL_ROOT_PASSWORD=root -d mysql:5.7

docker exec -it mysql-slave bash
```
[docker安装MySQL并实现主从复制](https://blog.csdn.net/wyg1973017714/article/details/112601802?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522167367945516800186532263%2522%252C%2522scm%2522%253A%252220140713.130102334..%2522%257D&request_id=167367945516800186532263&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~sobaiduend~default-2-112601802-null-null.142^v71^one_line,201^v4^add_ask&utm_term=dockermysql%E4%B8%BB%E4%BB%8E%E5%A4%8D%E5%88%B6&spm=1018.2226.3001.4187)

[MySQL主从同步show master status; Empty set (0.01 sec)主库无master状态问题解决](https://blog.csdn.net/lucky_ykcul/article/details/102809957)