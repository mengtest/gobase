/**************************************************************************************
Code Description    : 服务定义
Code Vesion         :
					|------------------------------------------------------------|
						  Version    					Editor            Time
							1.0        					yuansudong        2016.4.12
					|------------------------------------------------------------|
Version Description	:
                    |------------------------------------------------------------|
						  Version
							1.0
								 ....
					|------------------------------------------------------------|
***************************************************************************************/

package websocket

import (
	"encoding/json"
	"gateway-websocket-/global"
	"gobase/redis"
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/websocket"
)

var (
	servInstance *Server
)

func init() {
	servInstance = newServer()
}

type (
	// Server  用于描述一个Server
	Server struct {
		mutex                     sync.Mutex
		gatewayData               *message.HotelGatewayData
		navigate                  *router
		echoConnectMaxIdleTimeout time.Duration
		echoConnectReadTimeout    time.Duration
		service                   *http.Server
	}
)

// SetMaxConnectNumber 用于设置最大连接数量,当达到了这个数量,该服务器将不会再对外多余的连接服务
func (s *Server) SetMaxConnectNumber(maxConnectNumber int32) {
	s.gatewayData.MaxConnectCount = maxConnectNumber
}

// IncConnectCount 用于增加连接,并向redis中进行更新
func (s *Server) IncConnectCount() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.gatewayData.ConnectCount >= s.gatewayData.MaxConnectCount {
		return false
	}
	s.gatewayData.ConnectCount++
	data, _ := json.Marshal(s.gatewayData)
	redis.HSet(message.GetRedisGatewaySetKey(), global.GetInstance().GatewayName, string(data))
	return true
}

// DecConnectCount 用于减少连接,并向redis中进行更新
func (s *Server) DecConnectCount() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.gatewayData.ConnectCount == 0 {
		return
	}
	s.gatewayData.ConnectCount--
	data, _ := json.Marshal(s.gatewayData)
	redis.HSet(message.GetRedisGatewaySetKey(), global.GetInstance().GatewayName, string(data))
}

// SetEchoConnectMaxIdleTimeout 用于设置每个连接的最大驻留时间,在不活跃的状态下
func (s *Server) SetEchoConnectMaxIdleTimeout(t time.Duration) {
	s.echoConnectMaxIdleTimeout = t
}

// GetEchoConnectMaxIdleTimeout 用于获取每个连接的最大驻留时间,当连接在不活跃的状态下才可以
func (s *Server) GetEchoConnectMaxIdleTimeout() time.Duration {
	return s.echoConnectMaxIdleTimeout
}

// SetEchoConnectReadTimeout 用于设置每个连接的最大读取超时时间
func (s *Server) SetEchoConnectReadTimeout(t time.Duration) {
	s.echoConnectReadTimeout = t
}

// GetEchoConnectReadTimeout 用于获取每个连接的读取最大超时时间
func (s *Server) GetEchoConnectReadTimeout() time.Duration {
	return s.echoConnectReadTimeout
}

// GetServer 用于获取不知道的事情
func GetServer() *Server {
	return servInstance
}

// NewServer 创建一个新的server实例
func newServer() *Server {
	pServer := &Server{
		navigate:                  newRouter(),
		echoConnectMaxIdleTimeout: time.Millisecond * 1000 * 30,
		echoConnectReadTimeout:    time.Second * 10,
		gatewayData: &message.HotelGatewayData{
			MaxConnectCount: 1000,
			ConnectCount:    0,
			Addr:            global.GetInstance().WebsocketServerAddr,
			RPCAddr:         global.GetInstance().RPCServerAddr,
		},
	}
	return pServer
}

// SetMaxConnectArrivedCallBack 当达到了最大连接的时候,会回调该函数,并且向客户端发送提示信息
func (s *Server) SetMaxConnectArrivedCallBack(maxConnectArrivedCallBack FuncInnerCallBack) {
	s.navigate.maxConnectArrivedNavigate = maxConnectArrivedCallBack
}

// SetUnknownCodeCallBack 用于设置当发生未知行为码的时候要做的事情
func (s *Server) SetUnknownCodeCallBack(unknownRouterCallBack FuncRouterCallBack) {
	s.navigate.unknownNavigate = unknownRouterCallBack
}

// SetPingCallBack 用于设置Ping的回调
func (s *Server) SetPingCallBack(pingCallBack FuncInnerCallBack) {
	s.navigate.pingNavigate = pingCallBack
}

// SetOpenConnectCallBack 用于设置有新的连接到来的时候,需要调用的函数
func (s *Server) SetOpenConnectCallBack(openConnectCallBack FuncInnerCallBack) {
	s.navigate.openConnectNavigate = openConnectCallBack
}

// SetCloseConnectCallBack 用于在连接关闭的时候,需要调用的函数,一般用于在管理者中删除该会话
func (s *Server) SetCloseConnectCallBack(closeConnectCallBack FuncInnerCallBack) {
	s.navigate.closeConnectNavigate = closeConnectCallBack
}

// AddActionFunc 用于添加路由
func (s *Server) AddActionFunc(actionCode int, actionCallback FuncRouterCallBack) {
	s.navigate.AddFunction(actionCode, actionCallback)
}

//Run 用于创建一个Server 的实例
func (s *Server) Run(url string) {
	data, _ := json.Marshal(s.gatewayData)
	gatewayName := global.GetInstance().GatewayName
	redis.HSet(message.GetRedisGatewaySetKey(), gatewayName, string(data))
	http.Handle("/", websocket.Handler(runSession))
	s.service = &http.Server{
		Addr: url,
	}
	s.service.ListenAndServe()
	redis.Lock(message.GetRedisSessionSetLock(gatewayName))
	redis.DEL(message.GetRedisSessionSetKey(gatewayName))
	redis.UnLock(message.GetRedisSessionSetLock(gatewayName))
	redis.HDel(message.GetRedisGatewaySetKey(), gatewayName)
}

// Stop 用于停止Websocket 服务
func (s *Server) Stop() {
	defer func() {
		if err := recover(); err != nil {

		}
	}()
	s.service.Shutdown(nil)
}
