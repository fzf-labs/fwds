# 签名校验

## 如何给调用方开通 KEY 和 SECRET？

1. 新增调用方，输入调用方标识、调用方对接人、备注等信息；

2. 授权调用方可调用的接口；

3. 查看详情，将调用方的 KEY、SECRET 发给调用方；

## 调用方如何传递 Token？

基于 HTTP Header 中的两个参数 Authorization、Authorization-Date 存储签名信息。

1. Authorization 存储签名信息，格式：调用方 KEY + 空格分隔符 + 摘要(加密串)，例如： Authorization:blog MjJjMDE1MWFkZjMwOWFmYjFlNzViNDFjYjYwMWFlMmM=
2. Authorization-Date 存储时间信息，格式：0000-00-00 00:00:00，使用 Asia/Shanghai 时区，例如；

Authorization-Date:2021-04-03 21:12:36

## 代码示例

```go
func New(key, secret string, ttl time.Duration) Signature {
return &signature{
key:    key,
secret: secret,
ttl:    ttl,
}
}

// Generate
// path 请求的路径 (不附带 querystring)
func (s *signature) Generate(path string, method string, params url.Values) (authorization, date string, err error) {
if path == "" {
err = errors.New("path required")
return
}

if method == "" {
err = errors.New("method required")
return
}

methodName := strings.ToUpper(method)
if !methods[methodName] {
err = errors.New("method param error")
return
}

// Date
date = time_parse.CSTLayoutString()

// Encode() 方法中自带 sorted by key
sortParamsEncode, err := url.QueryUnescape(params.Encode())
if err != nil {
err = errors.Errorf("url QueryUnescape %v", err)
return
}

// 加密字符串规则
buffer := bytes.NewBuffer(nil)
buffer.WriteString(path)
buffer.WriteString(delimiter)
buffer.WriteString(methodName)
buffer.WriteString(delimiter)
buffer.WriteString(sortParamsEncode)
buffer.WriteString(delimiter)
buffer.WriteString(date)

// 对数据进行 sha256 加密，并进行 base64 encode
hash := hmac.New(sha256.New, []byte(s.secret))
hash.Write(buffer.Bytes())
digest := base64.StdEncoding.EncodeToString(hash.Sum(nil))

authorization = fmt.Sprintf("%s %s", s.key, digest)
return
}

// 模拟数据
const (
key = "blog"
secret = "i1ydX9RtHyuJTrw7frcu"
ttl = time.Minute * 10
)

func TestSignature_Generate(t *testing.T) {
path := "/echo"
method := "POST"

params := url.Values{}
params.Add("a", "a1")
params.Add("d", "d1")
params.Add("c", "c1 c2*")

authorization, date, err := New(key, secret, ttl).Generate(path, method, params)
t.Log("authorization:", authorization)
t.Log("authorization-date:", date)
t.Log("err:", err)
}
```

```php
// 模拟数据
$key    = "blog";
$secret = "i1ydX9RtHyuJTrw7frcu";

$path = "/echo";
$method = "POST";

$params['a'] = "a1";
$params['d'] = "d1";
$params['c'] = "c1 c2*";

// 对 params key 进行排序
ksort($params);

// 对 sortParams 进行操作
$sortParamsEncode = rawurldecode(http_build_query($params, "", "&", PHP_QUERY_RFC3986));

// 时间 使用 Asia/Shanghai 时区
$date = date("Y-m-d H:i:s", time());

// 加密字符串规则
$encryptStr = $path."|".strtoupper($method)."|".$sortParamsEncode."|".$date;

// 对数据进行 sha256 加密，并进行 base64 encode
$digest = base64_encode(hash_hmac("sha256", $encryptStr, $secret, true));

$authorization = $key." ".$digest;

echo "authorization:{$authorization}";
echo "---";
echo "authorization-date:{$date}";
```

```js
let key = "blog";
let secret = "i1ydX9RtHyuJTrw7frcu";

let date = new Date();
let datetime = date.getFullYear() + "-" // "年"
    + ((date.getMonth() + 1) > 10 ? (date.getMonth() + 1) : "0" + (date.getMonth() + 1)) + "-" // "月"
    + (date.getDate() < 10 ? "0" + date.getDate() : date.getDate()) + " " // "日"
    + (date.getHours() < 10 ? "0" + date.getHours() : date.getHours()) + ":" // "小时"
    + (date.getMinutes() < 10 ? "0" + date.getMinutes() : date.getMinutes()) + ":" // "分钟"
    + (date.getSeconds() < 10 ? "0" + date.getSeconds() : date.getSeconds()); // "秒"

let path = "/echo";
let method = "POST";
let params = {a: 'a1', d: 'd1', c: 'c1 c2*'};

let sortParamsEncode = decodeURIComponent(jQuery.param(ksort(params)));
let encryptStr = path + "|" + method.toUpperCase() + "|" + sortParamsEncode + "|" + datetime;
let digest = CryptoJS.enc.Base64.stringify(CryptoJS.HmacSHA256(encryptStr, secret));
console.log({authorization: key + " " + digest, date: datetime});
```