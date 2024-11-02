package core

// Unit 数据单位
type Unit uint64

const (
	Byte Unit = 1
	// KB KB单位
	KB = Byte * 1024
	// MB MB单位
	MB = KB * 1024
	// GB GB单位
	GB = MB * 1024
	// TB TB单位
	TB = GB * 1024
)

// Byte 返回字节
func (u Unit) Byte() uint64 {
	return uint64(u)
}

// KB 返回KB
func (u Unit) KB() float64 {
	return float64(u) / float64(KB)
}

// MB 返回MB
func (u Unit) MB() float64 {
	return float64(u) / float64(MB)
}

// GB 返回GB
func (u Unit) GB() float64 {
	return float64(u) / float64(GB)
}

// TB 返回TB
func (u Unit) TB() float64 {
	return float64(u) / float64(TB)
}
