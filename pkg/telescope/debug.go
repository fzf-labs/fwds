package telescope

type Debug struct {
	Index       int         `json:"index"`        // 索引ID
	Timestamp   string      `json:"timestamp"`    // 时间，格式：2006-01-02 15:04:05
	Key         string      `json:"key"`          // 标示
	Value       interface{} `json:"value"`        // 值
	CostSeconds float64     `json:"cost_seconds"` // 执行时间(单位秒)
}
