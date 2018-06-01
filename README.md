# gobase
这个是一个为自己将来的服务器生涯独挡一面所准备的库
# buffer
   这个是一个ringbuffer,主要是在io数据接受那块
# distribute
   这个是自己设计的一个分布式的消息ID,在所有节点上都是唯一
# idle
   这个是自己设计的一个驻留式池,可用于连接,以及驻留式协程池
# logger
   这个是自己准备的一个日志库,目前还不支持分布式日志   
# network
   这个是自己准备的网络库,目前支持tcp,websocket
# packet
   这个是自己准备的一个数据交换的包,目前支持json,protobuff
# redis
   这个是自己封装的一些redis中的命令, 并且支持连接池
# rpc
   这是自己封装的rpc包
      目前支持服务端,和单连接, 后面准备加入驻留式连接池
# service
   这是一个构建rpc 服务的包
# util
   一些常用的工具函数
