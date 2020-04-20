package main

import (
	"log"
	"w12lab.com/go-totp"
)

func main() {
	// RFC 6232
	log.Println("RFC 6232")
	// Default: T0:0  TI: 30 Digital: 6
	client := totp.NewTotp(totp.DefaultT0, totp.DefaultTI, totp.DefaultDigital)

	secret := client.GenerateSecret()
	log.Println("Secret/秘钥:", secret)

	code := client.GenerateCode(secret)
	log.Println("VerifyCode/验证码:", code)

	res := client.VerifyCode(secret, code)
	log.Println("Verify/验证:", res)

	log.Println("--------------------------------------")

	// Google Authenticator
	log.Println("Google Authenticator")
	googleAuth := totp.GetGoogle2FAAuth()

	secret = googleAuth.GenerateSecret()
	log.Println("Secret/秘钥:", secret)

	code = googleAuth.GenerateCode(secret)
	log.Println("VerifyCode/验证码:", code)

	res = googleAuth.VerifyCode(secret, code)
	log.Println("Verify/验证:", res)

}
