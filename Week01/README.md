# 第一周学习笔记

## 微服务
起因，单体架构不容易快速扩容去承载突发流量，将功能分拆，方便按需扩充。

## API Gateway
￼将入口流量按照设定配置，进行分切，做到服务高可用、后端流量负载均衡。

## gRPC
• 服务而非对象、消息而非引用：促进微服务的系统间粗粒度消息交互设计理念。
• 负载无关的：不同的服务需要使用不同的消息类型和编码，例如protocol buffers、JSON、XML和Thrift。
• 流: Streaming API。
• 阻塞式和非阻塞式:支持异步和同步处理在客户端和服务端间交互的消息序列。
• 元数据交换:常见的横切关注点，如认证或跟踪，依赖数据交换。
• 标准化状态码：客户端通常以有限的方式响应API调用返回的错误。

## 多集群
两地三中心，防止服务中断服务。

## 多租户
虚拟多个线上环境，互相隔离，可以将测试环境的流量单独流入指定环境。



## 链接

[书籍、文章记录][1]

[1]:	[https://shimo.im/docs/GYvDrQT8qW8RgkGY/read]