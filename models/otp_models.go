package models

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"strings"
	"time"
	"wisdom-portal/wisdom-portal/logger"
)

type GoogleAuth struct {
	QrCode string `json:"qr_code"`
	Sercet string `json:"sercet"`
}

func NewGoogleAuth() *GoogleAuth {
	return &GoogleAuth{}
}

func (g *GoogleAuth) un() int64 {
	return time.Now().UnixNano() / 1000 / 30
}

func (g *GoogleAuth) hmacSha1(key, data []byte) []byte {
	h := hmac.New(sha1.New, key)
	if total := len(data); total > 0 {
		h.Write(data)
	}
	return h.Sum(nil)
}

func (g *GoogleAuth) base32encode(src []byte) string {
	return base32.StdEncoding.EncodeToString(src)
}

func (g *GoogleAuth) base32decode(s string) ([]byte, error) {
	return base32.StdEncoding.DecodeString(s)
}

func (g *GoogleAuth) toBytes(value int64) []byte {
	var result []byte
	mask := int64(0xFF)
	shifts := [8]uint16{56, 48, 40, 32, 24, 16, 8, 0}
	for _, shift := range shifts {
		result = append(result, byte((value>>shift)&mask))
	}
	return result
}

func (g *GoogleAuth) toUint32(bts []byte) uint32 {
	return (uint32(bts[0]) << 24) + (uint32(bts[1]) << 16) +
		(uint32(bts[2]) << 8) + uint32(bts[3])
}

func (g *GoogleAuth) oneTimePassword(key []byte, data []byte) uint32 {
	hash := g.hmacSha1(key, data)
	offset := hash[len(hash)-1] & 0x0F
	hashParts := hash[offset : offset+4]
	hashParts[0] = hashParts[0] & 0x7F
	number := g.toUint32(hashParts)
	return number % 1000000
}

// 获取秘钥
func (g *GoogleAuth) GetSecret() string {
	var buf bytes.Buffer
	_ = binary.Write(&buf, binary.BigEndian, g.un())
	g.Sercet = strings.ToUpper(g.base32encode(g.hmacSha1(buf.Bytes(), nil)))
	return g.Sercet
}

// 获取动态码
func (g *GoogleAuth) GetCode(secret string) (string, error) {
	secretUpper := strings.ToUpper(secret)
	secretKey, err := g.base32decode(secretUpper)
	if err != nil {
		return "", err
	}
	number := g.oneTimePassword(secretKey, g.toBytes(time.Now().Unix()/30))
	return fmt.Sprintf("%06d", number), nil
}

// 获取动态码二维码内容
// 代码调用
// g.GetQrCode(user, secret)
func (g *GoogleAuth) GetQrCode(user, secret string) string {
	g.QrCode = fmt.Sprintf("otpauth://totp/%s?secret=%s", user, secret)
	return g.QrCode
}

// 验证动态码
func (g *GoogleAuth) VerifyCode(secret, code string) (bool, error) {
	_code, err := g.GetCode(secret)
	logger.Debug(fmt.Sprintf("%s %s %v", _code, code, err))
	if err != nil {
		return false, err
	}
	return _code == code, nil
}