/**************************************************************************************
Code Description    : 服务
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

package service

import (
	"encoding/json"
	"gobase/util"
)

// Service 用于对外的服务
type Service struct {
	slots []map[string]*Single
}

// NewService 用于创建一个新的服务
func NewService() *Service {
	pService := &Service{
		slots: make([]map[string]*Single, slotMask),
	}
	for i := 0; i < slotMask; i++ {
		pService.slots[i] = make(map[string]*Single)
	}
	return pService
}

// Register 用于注册一个服务,传入的是一个指针类型的
func (s *Service) Register(service interface{}) {
	name, single := NewSingleWith(service)
	code := util.StringToInt(name)
	s.slots[code%slotMask][name] = single
}

// Execute 用于执行相应的数据体
func (s *Service) Execute(request []byte) []byte {
	p := GetPacket()
	var err error
	var single *Single
	var ok bool
	err = json.Unmarshal(request, p)
	if err != nil {
		p.ServiceName = exceptionModel
		p.ServiceMethod = errMethod
		p.Data = []byte(err.Error())
		goto end
	}
	if single, ok = s.slots[util.StringToInt(p.ServiceName)%slotMask][p.ServiceName]; ok {
		p.Data = single.Execute(p.ServiceMethod, p.Data)
	} else {
		p.ServiceMethod = exceptionModel
		p.ServiceMethod = unknownMethod
		p.Data = unknownServiceDesc
	}
end:
	data, _ := json.Marshal(p)
	PutPacket(p)
	return data
}
