package buffer

/*************************************************************************

Descritption: Ring 用于描述一个缓冲区

Version 	:
			No       coder
			1.0  ==> 袁苏东
			        func newRingWith() 用于创建一个环形buffer
			        func TryRead(Container) 用于尝试读满整个容器
					func TryWrite(Container) 用于将指定的字节数写到缓冲区
					func Capacity() 用于返回该环形buffer的容量
					func



*************************************************************************/

// Ring 用于描述一个环形缓冲区
type Ring struct {
	start, end, cap, len int
	buffer               []byte
}

// NewRingWith 用于创建一个环形的buffer
func NewRingWith(size int) *Ring {
	return &Ring{
		start:  0,
		end:    0,
		cap:    size,
		len:    0,
		buffer: make([]byte, size),
	}
}

// TryRead 用于在Ring buffer 中尝试读取指定的字节数
func (r *Ring) TryRead(Container []byte) bool {
	isHave := true
	size := len(Container)

	if r.len == 0 {
		isHave = false
		goto end
	}
	if size >= r.len {
		isHave = false
		goto end
	}

	for i := 0; i < size; i++ {
		Container[i] = r.buffer[r.start]
		r.start++
		r.len--
		if r.start == r.cap {
			r.start = 0
		}
	}

end:
	return isHave

}

// TryWrite 用于尝试写数据到Ring buffer 中
func (r *Ring) TryWrite(data []byte) bool {
	isSuccess := true
	return isSuccess
}

// Capacity 用于获取剩余的容量
func (r *Ring) Capacity() int {
	return r.cap
}

// Length 用于获取缓冲区中的数据长度
func (r *Ring) Length() int {
	return r.len
}
