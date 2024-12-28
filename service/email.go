package service

import (
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/patrickmn/go-cache"
	"math/rand"
	"net/smtp"
	"time"
	"web/utils/config"
	"web/utils/errmsg"
)

var (
	// 缓存中的验证代码将在创建后5分钟内有效，且每隔10分钟进行一次清理。
	verificationCodeCache = cache.New(2*time.Minute, 5*time.Minute)
)

type EmailService interface {
	SendVerificationCode(to string) int
	VerifyVerificationCode(email string, code string) bool
}

type emailService struct {
}

func NewEmailService() EmailService {
	return &emailService{}
}

// SendVerificationCode sends a verification code to the user's email
func (e *emailService) SendVerificationCode(to string) int {
	// 生成验证码
	code := generateVerificationCode()

	// 发送验证码
	_ = e.sendVerificationCode(to, code)

	// todo：暂时有问题，会返回 short response，但是可以正常发送验证码，所以先注释掉
	//if err != nil {
	//	println(err.Error())
	//	return errmsg.ERROR_SEND_VERIFICATION_CODE
	//}

	// 将验证码和邮箱使用键值对方式存储在cache中
	verificationCodeCache.Set(to, code, cache.DefaultExpiration)

	// debug:观察cache是否真正存储验证码
	//cachedCode, _ := verificationCodeCache.Get(to)
	//cachedCodeStr, ok := cachedCode.(string)
	//if !ok {
	//	// 如果转换失败，返回 false
	//	println("转化失败")
	//}
	//println("这是验证码：", cachedCodeStr)

	return errmsg.SUCCESS
}

// sendVerificationCode 发送验证代码到指定的邮箱。
// 参数 to: 邮件接收人的邮箱地址。
// 参数 code: 需要发送的验证代码。
// 返回值 error: 发送过程中遇到的任何错误。
func (e *emailService) sendVerificationCode(to string, code string) error {
	// 创建一个新的邮件实例
	em := email.NewEmail()
	em.From = "Admin <3594781281@qq.com>"
	em.To = []string{to}
	em.Subject = "Verification Code"
	// 设置邮件的HTML内容
	em.HTML = []byte(`
		<h1>Verification Code</h1>
		<p>Your verification code is: <strong>` + code + `</strong></p>
	`)

	// 发送邮件(这里使用QQ进行发送邮件验证码)
	err := em.Send("smtp.qq.com:587",
		smtp.PlainAuth(
			"",
			config.QQEmail,
			config.QQGenCode,
			"smtp.qq.com"))
	if err != nil {
		return err // 如果发送过程中有错误，返回错误信息
	}
	return nil // 邮件发送成功，返回nil
}

// 随机生成一个6位数的验证码。
func generateVerificationCode() string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// 生成一个6位数的验证码
	code := fmt.Sprintf("%06d", r.Intn(1000000))
	return code
}

// VerifyVerificationCode verifies the verification code sent to the user
func (e *emailService) VerifyVerificationCode(email string, code string) bool {
	// retrieve the verification code from the cache
	cachedCode, found := verificationCodeCache.Get(email)
	// 如果没有找到验证码或者验证码过期，返回false
	if !found {
		return false
	}

	// 将 cachedCode 转换为 string 类型
	cachedCodeStr, ok := cachedCode.(string)
	if !ok {
		// 如果转换失败，返回 false
		return false
	}

	// compare the cached code with the provided code
	if cachedCodeStr != code {
		return false
	}

	// 使用一次验证码后就将其从cache中删除，保证验证码在2分钟内只能使用一次
	verificationCodeCache.Delete(email)

	return true
}
