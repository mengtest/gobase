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
