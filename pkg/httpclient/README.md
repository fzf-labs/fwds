# httpclient

## 模块介绍

httpclient 是基于 `net/http` 封装的 Go HTTP 客户端请求包，支持常用的请求方式、常用设置，比如：

- 支持设置 Mock 信息
- 支持设置失败时告警

- 支持设置失败时重试
- 支持设置项目内部的 Trace

- 支持设置超时时间、Header 等

## 请求说明

| 方法名                 | 描述                  |
| ---------------------- | --------------------- |
| httpclient.Get()       | GET 请求              |
| httpclient.PostForm()  | POST 请求，form 形式  |
| httpclient.PostJSON()  | POST 请求，json 形式  |
| httpclient.PutForm()   | PUT 请求，form 形式   |
| httpclient.PutJSON()   | PUT 请求，json 形式   |
| httpclient.PatchForm() | PATCH 请求，form 形式 |
| httpclient.PatchJSON() | PATCH 请求，json 形式 |
| httpclient.Delete()    | DELETE 请求           |

## 配置说明

| 配置项                        | 配置方法                                                     |
| ----------------------------- | ------------------------------------------------------------ |
| 设置 TTL 本次请求最大超时时间 | httpclient.WithTTL(ttl time.Duration)                        |
| 设置 Header 信息              | httpclient.WithHeader(key, value string)                     |
| 设置 Logger 信息              | httpclient.WithLogger(logger *zap.Logger)                    |
| 设置 Trace 信息               | httpclient.WithTrace(t trace.T)                              |
| 设置 Mock 信息                | httpclient.WithMock(m Mock)                                  |
| 设置失败时告警                | httpclient.WithOnFailedAlarm(alarmTitle string, alarmObject AlarmObject, alarmVerify AlarmVerify) |
| 设置失败时重试                | httpclient.WithOnFailedRetry(retryTimes int, retryDelay time.Duration, retryVerify RetryVerify) |

### 设置 TTL

```go
// 设置本次请求最大超时时间为 5s
httpclient.WithTTL(time.Second*5),
```

### 设置 Header 信息

可以调用多次进行设置多对 key-value 信息。

```go
// 设置多对 key-value 信息，比如这样：
httpclient.WithHeader("Authorization", "xxxx"),
httpclient.WithHeader("Date", "xxxx"),
```

### 设置 Logger 信息

传递的 logger 便于 httpclient 打印日志。

```go
// 使用上下文中的 logger，比如这样：
httpclient.WithLogger(ctx.Logger()),
```

### 设置 Trace 信息

传递的 trace 便于记录使用 httpclient 调用第三方接口的链路日志。

```go
// 使用上下文中的 trace，比如这样：
httpclient.WithTrace(ctx.Trace()),
```

### 设置 Mock 信息

```go
// Mock 类型
type Mock func() (body []byte)

// 需实现 Mock 方法，比如这样：
func MockDemoPost() (body []byte) {
	res := new(demoPostResponse)
	res.Code = 1
	res.Msg = "ok"
	res.Data.Name = "mock_Name"
	res.Data.Job = "mock_Job"

	body, _ = json.Marshal(res)
	return body
}

// 使用时：
httpclient.WithMock(MockDemoPost),
```

传递的 Mock 方式便于设置调用第三方接口的 Mock 数据。只要约定了接口文档，即使对方接口未开发时，也不影响数据联调。

### 设置失败时告警

```go
// alarmTitle 设置失败告警标题 String

// AlarmObject 告警通知对象，可以是邮件、短信或微信
type AlarmObject interface {
	Send(subject, body string) error
}

// 需要去实现 AlarmObject 接口，比如这样：
var _ httpclient.AlarmObject = (*AlarmEmail)(nil)

type AlarmEmail struct{}

func (a *AlarmEmail) Send(subject, body string) error {
	options := &mail.Options{
		MailHost: "smtp.163.com",
		MailPort: 465,
		MailUser: "xx@163.com",
		MailPass: "",
		MailTo:   "",
		Subject:  subject,
		Body:     body,
	}
	return mail.Send(options)
}

// AlarmVerify 定义符合告警的验证规则
type AlarmVerify func(body []byte) (shouldAlarm bool)

// 需要去实现 AlarmVerify 方法，比如这样：
func alarmVerify(body []byte) (shouldalarm bool) {
	if len(body) == 0 {
		return true
	}

	type Response struct {
		Code int `json:"code"`
	}
	resp := new(Response)
	if err := json.Unmarshal(body, resp); err != nil {
		return true
	}

    // 当第三方接口返回的 code 不等于约定的成功值（1）时，就要进行告警
	return resp.Code != 1
}

// 使用时：
httpclient.WithOnFailedAlarm("接口告警", new(third_party_request.AlarmEmail), alarmVerify),
```

### 设置失败时重试

```go
// retryTimes 设置重试次数 Int，默认：3

// retryDelay 设置重试前延迟等待时间 time.Duration，默认：time.Millisecond * 100

// RetryVerify 定义符合重试的验证规则
type RetryVerify func(body []byte) (shouldRetry bool)

// 需要去实现 RetryVerify 方法，比如这样：
func retryVerify(body []byte) (shouldRetry bool) {
	if len(body) == 0 {
		return true
	}

	type Response struct {
		Code int `json:"code"`
	}
	resp := new(Response)
	if err := json.Unmarshal(body, resp); err != nil {
		return true
	}

    // 当第三方接口返回的 code 等于约定值（10010）时，就要进行重试
	return resp.Code = 10010
}

// RetryVerify 也可以为 nil , 当为 nil 时，默认重试规则为 http_code 为如下情况：
// http.StatusRequestTimeout, 408
// http.StatusLocked, 423
// http.StatusTooEarly, 425
// http.StatusTooManyRequests, 429
// http.StatusServiceUnavailable, 503
// http.StatusGatewayTimeout, 504

// 使用时：
httpclient.WithOnFailedRetry(3, time.Second*1, retryVerify),
```

## 示例

```go
// 以 httpclient.PostForm 为例

api := "http://127.0.0.1:9999/demo/post"
params := url.Values{}
params.Set("name", name)
body, err := httpclient.PostForm(api, params,
	httpclient.WithTTL(time.Second*5),
	httpclient.WithTrace(ctx.Trace()),
	httpclient.WithLogger(ctx.Logger()),
	httpclient.WithHeader("Authorization", "xxxx"),
	httpclient.WithMock(MockDemoPost),
    httpclient.WithOnFailedRetry(3, time.Second*1, retryVerify),
    httpclient.WithOnFailedAlarm("接口告警", new(third_party_request.AlarmEmail), alarmVerify),                             
)

if err != nil {
    return nil, err
}

res = new(demoPostResponse)
err = json.Unmarshal(body, res)
if err != nil {
    return nil, errors.Wrap(err, "DemoPost json unmarshal error")
}

if res.Code != 1 {
    return nil, errors.New(fmt.Sprintf("code err: %d-%s", res.Code, res.Msg))
}

return res, nil
```

