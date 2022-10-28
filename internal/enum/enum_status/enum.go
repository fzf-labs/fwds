package enum_status

//go:generate stringer -type=StatusType
type StatusType int

const (
	Normal  StatusType = 1  //正常
	Disable StatusType = -1 //禁用
	Del     StatusType = -2 //删除
)
