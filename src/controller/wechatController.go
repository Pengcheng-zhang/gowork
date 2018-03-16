package controller

import (
	"encoding/json"
	"compress/gzip"
	"io"
	"io/ioutil"
	"encoding/hex"
	"encoding/base64"
	"fmt"
	"crypto/sha1"
	"strings"
	"sort"
	"net/http"
	"github.com/martini-contrib/render"
	"wechat"
	"biz"

)

type WechatController struct {
	wechatBiz biz.WechatBiz
}

func (this *WechatController) Index(r render.Render,req *http.Request) {
	code := req.FormValue("code")
	fmt.Println("code =", code)
	if code == "" {
		requestUrl := this.wechatBiz.GetAuthCodeUrl("http://4fa13010.ngrok.io/wechat")
		//r.JSON(200, map[string]interface{}{"code": 10000, "message":"微信链接成功", "result": ""})
		fmt.Println("request_url = ", requestUrl)
		r.Redirect(requestUrl)
		return
	}
	result, err := this.wechatBiz.GetAccessTokenByCode(code)
	if err != nil {
		return 
	}
	var accessToken wechat.AccessToken
	err = json.Unmarshal(result, &accessToken)
	if err != nil {
		fmt.Printf("json decode error: %v", err)
		var accessError wechat.WechatError
		err = json.Unmarshal(result, &accessError)
		if err != nil {
			fmt.Printf("json decode error: %v", err)
			return 
		}
		fmt.Printf("access token request error: %s", accessError.Errmsg)
		return
	}
	fmt.Printf("access token is: %s", accessToken.AccessToken)
	return
}

func (this *WechatController) Login(r render.Render,req *http.Request) {
	code := req.FormValue("code")
	fmt.Println("wechat access code ",code)
}

//二维码扫描登录
func (this *WechatController) QrScanLogin (r render.Render,req *http.Request) {
	var redirect_uri string ="http://7227a312.ngrok.io/wechat/login"
	redirect_uri = base64.URLEncoding.EncodeToString([]byte(redirect_uri))  //该回调需要url编码
	var scope string ="snsapi_login" //写死，微信暂时只支持这个值
	//准备向微信发请求
	requestParams := []string{"https://open.weixin.qq.com/connect/qrconnect?appid=", wechat.WeConfig.AppID, "&redirect_uri=", redirect_uri, "&response_type=code&scope=",scope, "&state=STATE#wechat_redirect"}
	var requestUrl = strings.Join(requestParams,"")
	fmt.Println(requestUrl)
	//请求返回的结果(实际上是个html的字符串)
	response,err := http.Get(requestUrl)
	if err != nil{
		r.Redirect("/404")
		return
	}
	defer response.Body.Close()
	var reader io.ReadCloser
	if response.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(response.Body)
		if err != nil {
			return
		}
	}else{
		reader = response.Body
	}
	body,berr := ioutil.ReadAll(reader)
	if berr != nil {
		r.Redirect("/404")
		return
	}

	//替换图片的src才能显示二维码
	result := strings.Replace(string(body), "/connect/qrcode/", "https://open.weixin.qq.com/connect/qrcode/", -1)
	r.Header().Set("Content-Type", response.Header.Get("Content-Type"))
	r.Text(200, result)
}
func check(r render.Render,req *http.Request) {
	signature := req.FormValue("signature")
	timestamp := req.FormValue("timestamp")
	nonce := req.FormValue("nonce")
	echostr := req.FormValue("echostr")

	token := wechat.WeConfig.AppToken
	checkSort := []string{nonce, timestamp, token}
	sort.Strings(checkSort)
	sortString := strings.Join(checkSort, "")
	fmt.Printf("sort=%s\n", sortString)
	h := sha1.New()
	h.Write([]byte(sortString))
	encryptString := hex.EncodeToString(h.Sum(nil))
	fmt.Printf("ens=%s\n", encryptString)
	if encryptString == signature {
		r.Data(200, []byte(echostr))
		return
	}
}