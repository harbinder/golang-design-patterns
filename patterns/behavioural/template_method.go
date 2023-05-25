package behavioural

import (
	"fmt"
)

/*
Template Method is a beharioural design pattern.
If we have a set of steps or algorithm, which can be implemented by any entity in similar way, then we can use this pattern.

For instance, OTP generation/verification related implementation via SMS, Email, Whatsapp or Push Notification can use this.

*/

type iOTP interface {
	generateOTP() int
	saveOTPForVerification(int) error
	createOTPContent(int) string
	sendOTP(string) error
}

func generateAndSendOTP(otp iOTP) (err error) {
	otpVal := otp.generateOTP()
	if err = otp.saveOTPForVerification(otpVal); err != nil {
		return
	}
	content := otp.createOTPContent(otpVal)
	otp.sendOTP(content)
	return
}

type SMS struct {
	mobile string
}
type Email struct {
	emailID string
}

func (s *SMS) generateOTP() (otp int) {
	otp = 2369
	fmt.Println("SMS: Generate random OTP : ", otp)
	return
}

func (s *SMS) saveOTPForVerification(otp int) (err error) {
	fmt.Println("SMS: Save OTP in cache/DB")
	return
}

func (s *SMS) createOTPContent(otp int) (content string) {
	fmt.Println("SMS: create content to be sent")
	content = "SMS OTP content"
	return
}
func (s *SMS) sendOTP(content string) (err error) {
	fmt.Println("SMS: OTP sent successfully to ", s.mobile)
	return
}

func (s *Email) generateOTP() (otp int) {
	otp = 2369
	fmt.Println("Email: Generate random OTP : ", otp)
	return
}

func (s *Email) saveOTPForVerification(otp int) (err error) {
	fmt.Println("Email: Save OTP in cache/DB")
	return
}

func (s *Email) createOTPContent(otp int) (content string) {
	fmt.Println("Email: create content to be sent")
	content = "Email OTP content"
	return
}
func (s *Email) sendOTP(content string) (err error) {
	fmt.Println("Email: OTP sent successfully to ", s.emailID)
	return
}

func ExecuteTemplateMethod() {
	sms := &SMS{mobile: "9910825975"}
	email := &Email{emailID: "harry@gmail.com"}

	generateAndSendOTP(sms)
	generateAndSendOTP(email)
}
