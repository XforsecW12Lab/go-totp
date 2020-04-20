package totp

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestTotp_GenerateSecret(t *testing.T) {
	client := GetGoogle2FAAuth()
	key := client.GenerateSecret()
	log.Println(key)
	log.Println(time.Now().Unix())
	fmt.Println(key)
	fmt.Println(client.GenerateCode(key))
}
