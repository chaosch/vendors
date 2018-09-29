// example of HTTP server that uses the captcha package.
package main

import (
	"encoding/json"
	"fmt"
	"github.com/mojocn/base64Captcha"
	"log"
	"net/http"
	"strconv"
)

// base64Captcha verify http handler
func captchaVerifyHandle(w http.ResponseWriter, r *http.Request) {
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
	}
	json.NewEncoder(w).Encode(body)
}

// base64Captcha create http handler
func generateCaptchaHandler(w http.ResponseWriter, r *http.Request) {
	//parse request parameters
	//接收客户端发送来的请求参数

	r.ParseForm()
	formData := r.Form
	captchaId := formData.Get("captchaId")
	DotCount, _ := strconv.Atoi(formData.Get("DotCount"))
	MaxSkew, _ := strconv.ParseFloat(formData.Get("MaxSkew"), 64)
	PngWidth, _ := strconv.Atoi(formData.Get("PngWidth"))
	PngHeight, _ := strconv.Atoi(formData.Get("PngHeight"))
	DefaultLen, _ := strconv.Atoi(formData.Get("DefaultLen"))

	//create base64 encoding captcha
	//创建base64图像验证码
	base64Png := base64Captcha.GenerateCaptchaPngBase64String(captchaId, PngWidth, PngHeight, DotCount, DefaultLen, MaxSkew)
	//or you can do this
	//你也可以是用默认参数 生成图像验证码
	//base64Png := captcha.GenerateCaptchaPngBase64StringDefault(captchaId)

	//set json response
	//设置json响应

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	body := map[string]interface{}{"code": 1, "data": base64Png, "msg": "success", "debug": formData}
	json.NewEncoder(w).Encode(body)
}

//start a net/http server
//启动golang net/http 服务器
func main() {
	//serve Vuejs+ElementUI+Axios Web Application
	http.Handle("/", http.FileServer(http.Dir("./static")))
	//api for create captcha
	http.HandleFunc("/api/getCaptcha", generateCaptchaHandler)
	//api for verify captcha
	http.HandleFunc("/api/verifyCaptcha", captchaVerifyHandle)

	fmt.Println("Server is at localhost:777")
	if err := http.ListenAndServe("localhost:777", nil); err != nil {
		log.Fatal(err)
	}
}
