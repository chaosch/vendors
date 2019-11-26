package captcha

import (
	. "egov/common"
	"egov/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/mojocn/base64Captcha"
	"net/http"
)

//获取验证码
// base64Captcha create http handler
func GetCaptcha(w http.ResponseWriter, r *http.Request, prams httprouter.Params) {
	//parse request parameters
	//接收客户端发送来的请求参数
	r.ParseForm()
	formData := r.Form
	captchaId := formData.Get("captchaId")
	//create base64 encoding captcha
	//创建base64图像验证码
	base64Png := base64Captcha.GenerateCaptchaPngBase64String(captchaId, 160, 80, 20, 5, 0.5)
	//or you can do this
	//你也可以是用默认参数 生成图像验证码
	//base64Png := captcha.GenerateCaptchaPngBase64StringDefault(captchaId)

	//set json response
	//设置json响应

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	body := map[string]interface{}{"code": 1, "data": base64Png, "msg": "success"}
	//json.NewEncoder(w).Encode(body)
	bod, err := json.Marshal(RetOk(body))
	if err != nil {
		fmt.Fprintf(w, "验证码生成错误")
	}
	fmt.Fprint(w, string(bod))
}

//校验验证码
// base64Captcha verify http handler
func VerifyCaptcha(w http.ResponseWriter, r *http.Request, prams httprouter.Params) *ResultTemplate {
	//parse request parameters
	//接收客户端发送来的请求参数
	r.ParseForm()
	formData := r.Form
	captchaId := formData.Get("captchaId")
	captchaDigits := formData.Get("captchaDigits")

	//verify the captcha
	//比较图像验证码
	verifyResult := base64Captcha.VerifyCaptcha(captchaId, captchaDigits)

	//set json response
	//设置json响应
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	body := map[string]interface{}{"code": "error", "data": "验证失败", "msg": "captcha failed", "debug": formData}
	if verifyResult {
		body = map[string]interface{}{"code": "success", "data": "验证通过", "msg": "captcha verified", "debug": formData}
		return RetOk(body)
	} else {
		return RetErr(NewErrorCode(1007))
	}
	//json.NewEncoder(w).Encode(body)

}
