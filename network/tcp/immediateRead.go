package tcp

import "base/util"

var (
	invalidContentLength = -1
)

type immediateRead struct {
	headStart, headEnd       int
	contentStart, contentEnd int
	contentData              []byte
	headData                 []byte
	input                    *reader
}

func newImmediateRead() *immediateRead {
	return &immediateRead{
		headStart:    0,
		headEnd:      packetHeadSize,
		contentStart: 0,
		contentEnd:   0,
		headData:     make([]byte, packetHeadSize),
		contentData:  nil,
		input:        newReaderWith(readBufferSize),
	}
}

// read 用于读取数据
func (ir *immediateRead) read() ([]byte, bool) {
	var retData []byte
	isHave := false
	if ir.contentData == nil {
		ir.headStart += ir.input.read(ir.headStart, ir.headEnd, ir.headData)
		if ir.headStart == ir.headEnd {
			ir.headStart = 0
			ir.contentEnd = util.BytesToInt(ir.headData)
			ir.contentData = make([]byte, ir.contentEnd)
		}
	} else {
		ir.contentStart += ir.input.read(ir.contentStart, ir.contentEnd, ir.contentData)
		if ir.contentStart == ir.contentEnd {
			retData = ir.contentData
			isHave = true
			ir.contentEnd = 0
			ir.contentStart = 0
			ir.contentData = nil
		}
	}
	return retData, isHave
}

// update 用于更新缓冲区 中的数据
func (ir *immediateRead) update(session *Session) {
	ir.input.tryUpdate(session)
}
