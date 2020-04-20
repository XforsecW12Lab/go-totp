/* This program is free software. It comes without any warranty, to
 * the extent permitted by applicable law. You can redistribute it
 * and/or modify it under the terms of the Do What The Fuck You Want
 * To Public License, Version 2, as published by Sam Hocevar. See
 * http://www.wtfpl.net/ for more details. */

package totp

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"log"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const DefaultT0 = 0
const DefaultTI = 30
const DefaultDigital = 6

type Auth interface {
	GenerateSecret() string
	GenerateCode(key string) string
	VerifyCode(key, value string) bool
}

type Totp struct {
	t0  int
	ti  int
	dig int
}

func NewTotp(t0, ti, dig int) *Totp {
	return &Totp{
		t0:  t0,
		ti:  ti,
		dig: dig,
	}
}

func (t *Totp) GenerateSecret() string {
	b := make([]byte, 20)
	rand.Seed(time.Now().UnixNano())
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return strings.ToUpper(base32.StdEncoding.EncodeToString(b))
}

func (t *Totp) VerifyCode(key, value string) bool {
	return t.GenerateCode(key) == value
}

func (t *Totp) hmacCrypto(key, c []byte) []byte {
	client := hmac.New(sha1.New, key)
	client.Write(c)
	return client.Sum(nil)
}

func toBytes(value int64) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytesBuffer, binary.BigEndian, value)
	return bytesBuffer.Bytes()
}

func (t *Totp) GenerateCode(key string) string {
	sk, err := base32.StdEncoding.DecodeString(key)
	if err != nil {
		log.Println(err)
	}
	now := (time.Now().Unix() - int64(t.t0)) / int64(t.ti)
	steps := toBytes(now)
	hmacSha1 := t.hmacCrypto(sk, steps)
	length := len(hmacSha1)
	offset := hmacSha1[length-1] & 0x0F
	codeBinary :=
		(uint32(hmacSha1[offset]&0x7F) << 24) +
			(uint32(hmacSha1[offset+1]) << 16) +
			(uint32(hmacSha1[offset+2]) << 8) +
			(uint32(hmacSha1[offset+3]))
	otp := int(codeBinary) % int(math.Pow10(t.dig))
	result := strconv.Itoa(otp)
	for len(result) < t.dig {
		result = "0" + result
	}
	return result
}

func GetGoogle2FAAuth() Auth {
	return Auth(NewTotp(0, 30, 6))
}
