package telescope

type Redis struct {
	Index     int    `json:"index"`     // 索引ID
	Timestamp string `json:"timestamp"` // 时间，格式：2006-01-02 15:04:05
	Handle    string `json:"handle"`    // 操作，SET/GET 等
	Cmd       string `json:"cmd"`       // 命令
	Err       error  `json:"err"`       // 错误
}
