package wechat

type wechatConfig struct {
	AppID string
	AppSecret string
	AppToken string
	AppHost string
	ApiHost string
}
type AccessToken struct {
	AccessToken string	`json:"access_token"`
	ExpiresIn int `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid string `json:"openid"`
	Scope string `json:"scope"`
	
}

//api请求错误
type WechatError struct {
	Errcode string `json:"errcode"`
	Errmsg string `json:"errmsg"`
}

var WeConfig wechatConfig
func init() {
	WeConfig.AppID = "wx9db02d7fe62f044f"
	WeConfig.AppSecret = "1f3f16f3c45743c190092ca4647e5cca"
	WeConfig.AppToken = "zhangpch666"
	WeConfig.AppHost = "https://open.weixin.qq.com"
	WeConfig.ApiHost = "https://api.weixin.qq.com"
}