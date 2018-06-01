/**************************************************************************************
Code Description    : session中的附加数据
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

// SessionAttachData 用于描述 session的附加数据
type SessionAttachData struct {
	UID     int    `json:"uid"`
	HotelID int    `json:"hotel_id"`
	IP      string `json:"ip"`
	Token   string `json:"token"`
}

// CreateSessionAttachData 用于创建附加数据
func CreateSessionAttachData() *SessionAttachData {
	return &SessionAttachData{}
}
