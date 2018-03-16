package biz

import (
	"net/url"
	"io/ioutil"
	"fmt"
	"wechat"
	"net/http"
	"strings"
)
type WechatBiz struct {

}

//网页授权
func (this *WechatBiz)GetAuthCodeUrl(redirectUri string) string{
	redirectUri = url.QueryEscape(redirectUri)
	requestParams := []string{wechat.WeConfig.AppHost, "/connect/oauth2/authorize?appid=", wechat.WeConfig.AppID, "&redirect_uri=", redirectUri, "&response_type=code&scope=snsapi_userinfo&state=0012#wechat_redirect"}
	requestUrl := strings.Join(requestParams, "")
	return requestUrl
}

func GetAccessTokenByCodeUrl(code string) (url string){
	urlParams := []string{wechat.WeConfig.ApiHost, "/sns/oauth2/access_token?appid=", wechat.WeConfig.AppID, "&secret=", wechat.WeConfig.AppSecret, "&code=", code, "&grant_type=authorization_code"}
	requestUrl := strings.Join(urlParams, "")
	fmt.Printf("access token url %s", requestUrl)
	return requestUrl
}

func (this *WechatBiz)GetAccessTokenByCode (code string) ([]byte, error){
	requestUrl := GetAccessTokenByCodeUrl(code)
	result,err := requestForWechatApiURL(requestUrl)
	fmt.Printf("access code result :%s", result)
	return result, err
}
//http请求
func requestForWechatApiURL(requestUrl string) ([]byte,error){
	var body []byte
	client := &http.Client{}
	fmt.Println("request begin url=", requestUrl)
	request, err := http.NewRequest("GET", requestUrl, nil)
	request.Header.Set("Content-Type","application/json;charset=utf-8")
	if err != nil {
		fmt.Println("http request fail, url=", requestUrl)
		return body, err
	}
	resp, reserr := client.Do(request)
	if reserr != nil {
		fmt.Printf("response error: %v", reserr)
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("wechat request get body err: %v", err)
	}
	fmt.Println("body=", string(body))
	return body, err
}