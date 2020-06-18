package common

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"errors"
)

var pbKey = `MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCsxKFbUOc7QOvpOK2XdfXC7yAq
lSmALUK1LtOSb76vhEAkIJh8W5p8YMxknSeD3fF44PDqnSl42AnVr1qT/yBE9AYX
uXXnZPUOCyBnejQLMVQ3vheQuCrcp6a54EylEWe0sthlPh1xn5DjdHn3smM+m29z
FC8umQHtaZo/4Gti4QIDAQAB`

//解密前端密码
var pKey = `MIICXgIBAAKBgQCsxKFbUOc7QOvpOK2XdfXC7yAqlSmALUK1LtOSb76vhEAkIJh8
W5p8YMxknSeD3fF44PDqnSl42AnVr1qT/yBE9AYXuXXnZPUOCyBnejQLMVQ3vheQ
uCrcp6a54EylEWe0sthlPh1xn5DjdHn3smM+m29zFC8umQHtaZo/4Gti4QIDAQAB
AoGAH9DyOiPTAXl6OG/koADsKmLpFI51nxI2t7EQ62XCwwXi3gRWsIgaEg+tdFXw
ofssbetW0o3wxj1aykxJrPmN0nZP7NMHlc/gFsRSPDGsfuqbsvUi+FaPcBlRchK0
/uGJPLBmTA7SObi1vN8PcuI5EVx3Y0dOH6NSYwNI9UwdYFUCQQCwvq5Y4rFKE2D8
mJr5STU31r51KqawAxEotWu4/xo8pzLa9/nySLYqXHRMsH7k5nWYEgLRWvkvWvqS
fLrwUV8vAkEA+j1we+hmSZdRCHw03Y3PvHkieGB2scQsHmtY4hwmV27GZ/c7nE0I
GIs4ApjAJLsXgU1dn23i3qy7cXMoj5Na7wJBAK6Bplymfot39MOR3TmSuZPHWEcQ
9IFGlfOOpKyEW6BVKRYfzs4UUqAipsVtYeyZSrP53IRegTfraQmnU1+hyAsCQQDg
EMd5i3ybhGxAhsbnYyWRg33D/8wIHLnMex5ZSA9k1oG9cfjkWW1YXMBIQvI5cXT4
iIe+S6yK6mvkh9LXzH/tAkEAi6Bz8JMXdT5vUh4wO+DY2wlhmZBZXWz2Utt1Uwh+
yLMNqHe9cawzNhbcLZg21dz6TDNQDTzIHfdbcv/0Qq3S1A==`

func RsaPassword(pk string) (password string, err error) {
	encryptedDecodeBytes, err := base64.StdEncoding.DecodeString(pk)
	if err != nil {
		return pk, err
	}
	key, _ := base64.StdEncoding.DecodeString(pKey)
	prvKey, _ := x509.ParsePKCS1PrivateKey(key)
	originalData, err := rsa.DecryptPKCS1v15(rand.Reader, prvKey, encryptedDecodeBytes)
	if err != nil {
		return pk, err
	}
	//w := md5.New()
	//io.WriteString(w, string(originalData))  //将str写入到w中
	//password = fmt.Sprintf("%x", w.Sum(nil)) //w.Sum(nil)将w的hash转成[]byte格式
	return string(originalData), nil
}

func GetRsaPassword(pk []byte) (string, error) {
	if pk == nil {
		return "", errors.New("输入的空加密数据")
	}
	key, _ := base64.StdEncoding.DecodeString(pbKey)
	prvKey, err := x509.ParsePKIXPublicKey(key)
	if err != nil {
		return "", err
	}
	pub := prvKey.(*rsa.PublicKey)
	encrypateData, err := rsa.EncryptPKCS1v15(rand.Reader, pub, pk)
	return base64.StdEncoding.EncodeToString(encrypateData), err
}
