package fc

// Config defines fc config
type Config struct {
	Endpoint        string // fc地址
	APIVersion      string // API版本
	AccountID       string // Account ID
	AccessKeyID     string // accessKeyID
	AccessKeySecret string // accessKeySecret
	SecurityToken   string // STS securityToken
	UserAgent       string // SDK名称/版本/系统信息
	IsDebug         bool   // 是否开启调试模式，默认false
	Timeout         uint   // 超时时间，默认60秒
	host            string // Set host from endpoint
}

// NewConfig get default config
func NewConfig() *Config {
	return &Config{
		Endpoint:		 "",
		APIVersion:      "2016-08-15",
		AccountID:       "",
		AccessKeyID:     "",
		AccessKeySecret: "",
		SecurityToken:   "",
		UserAgent:       "go-sdk-0.1",
		IsDebug:         false,
		Timeout:         60,
		host:            "",
	}
}
