package conf

import "time"

// Conf
// @Description: 全局配置变量
var Conf = &Config{}

// Config
// @Description: 全局配置
type Config struct {
	App  AppConfig  `json:"app"`
	Http HttpConfig `json:"http"`
	Grpc GrpcConfig `json:"grpc"`

	Log      LogConfig      `json:"log"`
	Jwt      JwtConfig      `json:"jwt"`
	HashIds  HashIdsConfig  `json:"hash_ids"`
	Mysql    MysqlConfigs   `json:"mysql,omitempty"`
	Redis    RedisConfig    `json:"redis"`
	Email    EmailConfig    `json:"email"`
	Trace    TraceConfig    `json:"trace"`
	Jaeger   JaegerConfig   `json:"jaeger"`
	Oss      OssConfig      `json:"oss"`
	Mq       MqConfig       `json:"mq"`
	Mongo    MongoDBConfigs `json:"mongo"`
	Notify   NotifyConfig   `json:"notify"`
	DingTalk DingTalkConfig `json:"ding_talk"`
	Sms      SmsConfig      `json:"sms"`
}

// AppConfig
// @Description: 框架配置
type AppConfig struct {
	Name      string `json:"name,omitempty"`
	Env       string `json:"env,omitempty"`
	Version   string `json:"version,omitempty"`
	Mode      string `json:"mode,omitempty"`
	PprofPort string `json:"pprof_port,omitempty"`
	Host      string `json:"host,omitempty"`
	SSL       bool   `json:"ssl,omitempty"`
	CSRF      bool   `json:"csrf,omitempty"`
	Debug     bool   `json:"debug,omitempty"`
}

// HttpConfig
// @Description: http服务配置
type HttpConfig struct {
	Addr         string        `json:"addr,omitempty"`
	ReadTimeout  time.Duration `json:"read_timeout,omitempty"`
	WriteTimeout time.Duration `json:"write_timeout,omitempty"`
}

// GrpcConfig
// @Description: grpc服务配置
type GrpcConfig struct {
	Addr         string        `json:"addr,omitempty"`
	ReadTimeout  time.Duration `json:"read_timeout,omitempty"`
	WriteTimeout time.Duration `json:"write_timeout,omitempty"`
}

// LogConfig
// @Description: 日志配置文件
type LogConfig struct {
	Name              string `json:"name,omitempty"`
	Development       bool   `json:"development,omitempty"`
	DisableCaller     bool   `json:"disable_caller,omitempty"`
	DisableStacktrace bool   `json:"disable_stacktrace,omitempty"`
	Encoding          string `json:"encoding,omitempty"`
	Level             string `json:"level,omitempty"`
	Writers           string `json:"writers,omitempty"`
	FormatText        bool   `json:"format_text,omitempty"`
	RollingPolicy     string `json:"rolling_policy,omitempty"`
	RotateDate        int    `json:"rotate_date,omitempty"`
	RotateSize        int    `json:"rotate_size,omitempty"`
	BackupCount       uint   `json:"backup_count,omitempty"`
	File              string `json:"file,omitempty"`
	WarnFile          string `json:"warn_file,omitempty"`
	ErrorFile         string `json:"error_file,omitempty"`
}

type JwtConfig struct {
	JwtSecret   string        `json:"jwt_secret,omitempty"`
	JwtDuration time.Duration `json:"jwt_duration,omitempty"`
}

type HashIdsConfig struct {
	Secret string `json:"secret,omitempty"`
	Length int    `json:"length,omitempty"`
}

type MysqlConfigs map[string]MysqlConfig

// MysqlConfig
// @Description: mysql配置
type MysqlConfig struct {
	DSN             string        `json:"dsn,omitempty"`
	ShowLog         bool          `json:"show_log,omitempty"`
	MaxIdleConn     int           `json:"max_idle_conn,omitempty"`
	MaxOpenConn     int           `json:"max_open_conn,omitempty"`
	ConnMaxLifeTime time.Duration `json:"conn_max_life_time,omitempty"`
}

// RedisConfig
// @Description: redis 配置
type RedisConfig struct {
	Addr         string        `json:"addr,omitempty"`
	Password     string        `json:"password,omitempty"`
	DB           int           `json:"db,omitempty"`
	MinIdleConn  int           `json:"min_idle_conn,omitempty"`
	DialTimeout  time.Duration `json:"dial_timeout,omitempty"`
	ReadTimeout  time.Duration `json:"read_timeout,omitempty"`
	WriteTimeout time.Duration `json:"write_timeout,omitempty"`
	PoolSize     int           `json:"pool_size,omitempty"`
	PoolTimeout  time.Duration `json:"pool_timeout,omitempty"`
}

// EmailConfig
// @Description: 邮件配置
type EmailConfig struct {
	Host      string `json:"host,omitempty"`
	Port      int    `json:"port,omitempty"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	Name      string `json:"name,omitempty"`
	Address   string `json:"address,omitempty"`
	ReplyTo   string `json:"reply_to,omitempty"`
	KeepAlive int    `json:"keep_alive,omitempty"`
}

// TraceConfig
// @Description: 链路追踪配置
type TraceConfig struct {
	ServiceName string `json:"service_name,omitempty"`
	Open        int    `json:"open,omitempty"`
	TraceAgent  string `json:"trace_agent,omitempty"`
}

// JaegerConfig
// @Description: 链路追踪Jaeger配置
type JaegerConfig struct {
	SamplingServerURL      string  `json:"sampling_server_url,omitempty"`       // Set the sampling server url
	SamplingType           string  `json:"sampling_type,omitempty"`             // Set the sampling type
	SamplingParam          float64 `json:"sampling_param,omitempty"`            // Set the sampling parameter
	LocalAgentHostPort     string  `json:"local_agent_host_port,omitempty"`     // Set jaeger-agent's host:port that the reporter will used
	Gen128Bit              bool    `json:"gen_128_bit,omitempty"`               // Generate 128 bit span IDs
	Propagation            string  `json:"propagation,omitempty"`               // Which propagation format to use (jaeger/b3)
	TraceContextHeaderName string  `json:"trace_context_header_name,omitempty"` // Set the header to use for the trace-id
	CollectorEndpoint      string  `json:"collector_endpoint,omitempty"`        // Instructs reporter to send spans to jaeger-collector at this URL
	CollectorUser          string  `json:"collector_user,omitempty"`            // CollectorUser for basic http authentication when sending spans to jaeger-collector
	CollectorPassword      string  `json:"collector_password,omitempty"`        // CollectorPassword for basic http authentication when sending spans to jaeger
}

// OssConfig
// @Description: 对象存储
type OssConfig struct {
	AliYun AliYunOssConfig `json:"ali_yun"`
}

type AliYunOssConfig struct {
	AccessKey string `json:"access_key,omitempty"`
	SecretKey string `json:"secret_key,omitempty"`
	Bucket    string `json:"bucket,omitempty"`
	Endpoint  string `json:"endpoint,omitempty"`
}

// MqConfig
// @Description: 消息队列
type MqConfig struct {
	Switch      bool
	Use         string
	RocketMqAli RocketMqAliConfig
	RocketMq    RocketMqConfig
	Nsq         NsqConfig
	Business    Business
}

type RocketMqAliConfig struct {
	AccessKey  string `json:"access_key"`
	SecretKey  string `json:"secret_key"`
	Endpoint   string `json:"endpoint"` //设置HTTP协议客户端接入点，进入消息队列RocketMQ版控制台实例详情页面的接入点区域查看。
	InstanceId string `json:"instance_id"`
}

type RocketMqConfig struct {
	Endpoint string `json:"endpoint,omitempty"`
}

type NsqConfig struct {
	Lookupds []string `json:"lookupds,omitempty"`
}

type Business map[string]BusinessConfig

// BusinessConfig 不同业务的配置
type BusinessConfig struct {
	// Topic所属的实例ID，在消息队列RocketMQ版控制台创建。
	// 若实例有命名空间，则实例ID必须传入；若实例无命名空间，则实例ID传入null空值或字符串空值。实例的命名空间可以在消息队列RocketMQ版控制台的实例详情页面查看。
	Name string `json:"name,omitempty"`
	// 消息所属的Topic，在消息队列RocketMQ版控制台创建。
	//不同消息类型的Topic不能混用，例如普通消息的Topic只能用于收发普通消息，不能用于收发其他类型的消息。
	Topic string `json:"topic,omitempty"`
	//标签
	Tag string `json:"tag,omitempty"`
	// 您在控制台创建的Group ID。
	GroupId string `json:"group_id,omitempty"`
}
type MongoDBConfigs map[string]MongoDB

type MongoDB struct {
	Uri      string `json:"uri"`
	Database string `json:"database"`
	Coll     string `json:"coll"`
}

type NotifyConfig map[string]NotifyBusiness

type NotifyBusiness struct {
	Email    string `json:"email"`
	Sms      string `json:"sms"`
	DingTalk string `json:"ding_talk"`
}

type DingTalkConfig map[string]DingTalkBusiness

type DingTalkBusiness struct {
	Url    string `json:"url"`
	Secret string `json:"secret"`
}

type SmsConfig struct {
	Use     string           `json:"use"`
	AliYun  SmsAliYunConfig  `json:"aliyun"`
	Tencent SmsTencentConfig `json:"tencent"`
}

type SmsAliYunConfig struct {
	AccessKey string                  `json:"access_key"`
	SecretKey string                  `json:"secret_key"`
	RegionId  string                  `json:"region_id"`
	Business  SmsAliYunBusinessConfig `json:"business"`
}

type SmsAliYunBusinessConfig map[string]SmsAliYunBusiness

type SmsAliYunBusiness struct {
	SignName string `json:"sign_name"`
	Template string `json:"template"`
}

type SmsTencentConfig struct {
	SecretId  string                   `json:"secret_id"`
	SecretKey string                   `json:"secret_key"`
	SdkAppId  string                   `json:"sdk_app_id"`
	Business  SmsTencentBusinessConfig `json:"business"`
}

type SmsTencentBusinessConfig map[string]SmsTencentBusiness

type SmsTencentBusiness struct {
	SignName   string `json:"sign_name"`
	TemplateId string `json:"template_id"`
}
