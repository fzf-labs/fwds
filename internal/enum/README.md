# 枚举

枚举值定义与自动生成

使用 golang.org/x/tools/cmd/stringer

# 操作:

1. 使用type第一枚举值,然后写入注释
    ```go
    //go:generate stringer -type=StatusType
    type StatusType int
    
    const (
        Normal  StatusType = 1  //正常
        Disable StatusType = -1 //禁用
        Del     StatusType = -2 //删除
    )
    ```

1. 在根目录下使用 `go generate ./...` 自动生成