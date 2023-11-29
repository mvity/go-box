/*
 * Copyright © 2021 - 2023 vity <vityme@icloud.com>.
 *
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

package x

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
)

// AESEncrypt AES 加密内容 AES-128/CBC/PKCS5Padding ，key：32，iv：16
func AESEncrypt(key string, iv string, data string) string {
	block, _ := aes.NewCipher([]byte(key))
	cbc := cipher.NewCBCEncrypter(block, []byte(iv))
	content := []byte(data)
	padding := block.BlockSize() - len(content)%block.BlockSize()
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	content = append(content, padtext...)
	crypted := make([]byte, len(content))
	cbc.CryptBlocks(crypted, content)
	return base64.StdEncoding.EncodeToString(crypted)
}

// AESDecrypt AES 解密内容 AES-128/CBC/PKCS5Padding ，key：32，iv：16
func AESDecrypt(key string, iv string, data string) string {
	crypt, _ := base64.StdEncoding.DecodeString(data)
	block, _ := aes.NewCipher([]byte(key))
	cbc := cipher.NewCBCDecrypter(block, []byte(iv))
	decrypted := make([]byte, len(crypt))
	cbc.CryptBlocks(decrypted, crypt)
	padding := decrypted[len(decrypted)-1]
	return string(decrypted[:len(decrypted)-int(padding)])
}

// DESEncrypt DES 加密内容 CBC/PKCS5Padding key：8，iv：8
func DESEncrypt(key string, iv string, data string) string {
	block, _ := des.NewCipher([]byte(key))
	cbc := cipher.NewCBCEncrypter(block, []byte(iv))
	content := []byte(data)
	padding := block.BlockSize() - len(content)%block.BlockSize()
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	content = append(content, padtext...)
	crypted := make([]byte, len(content))
	cbc.CryptBlocks(crypted, content)
	return base64.StdEncoding.EncodeToString(crypted)
}

// DESDecrypt DES 解密内容 CBC/PKCS5Padding ，key：8，iv：8
func DESDecrypt(key string, iv string, data string) string {
	crypt, _ := base64.StdEncoding.DecodeString(data)
	block, _ := des.NewCipher([]byte(key))
	cbc := cipher.NewCBCDecrypter(block, []byte(iv))
	decrypted := make([]byte, len(crypt))
	cbc.CryptBlocks(decrypted, crypt)
	padding := decrypted[len(decrypted)-1]
	return string(decrypted[:len(decrypted)-int(padding)])
}

// DESTripleEncrypt 3DES 加密内容 CBC/PKCS5Padding
func DESTripleEncrypt(key string, data string) string {
	block, _ := des.NewTripleDESCipher([]byte(key))
	cbc := cipher.NewCBCEncrypter(block, []byte(key)[:block.BlockSize()])
	content := []byte(data)
	crypted := make([]byte, len(content))
	cbc.CryptBlocks(crypted, content)
	return base64.StdEncoding.EncodeToString(crypted)
}

// DESTripleDecrypt DES 解密内容 CBC/PKCS5Padding
func DESTripleDecrypt(key string, data string) string {
	crypt, _ := base64.StdEncoding.DecodeString(data)
	block, _ := des.NewTripleDESCipher([]byte(key))
	cbc := cipher.NewCBCEncrypter(block, []byte(key)[:block.BlockSize()])
	decrypted := make([]byte, len(crypt))
	cbc.CryptBlocks(decrypted, crypt)
	padding := decrypted[len(decrypted)-1]
	return string(decrypted[:len(decrypted)-int(padding)])
}

// MD5String 对字符串MD5处理
func MD5String(plain string) string {
	hMD5 := md5.New()
	hMD5.Write([]byte(plain))
	return hex.EncodeToString(hMD5.Sum(nil))
}

// SHA1String  对字符串SHA1处理
func SHA1String(plain string) string {
	hSha1 := sha1.New()
	hSha1.Write([]byte(plain))
	return hex.EncodeToString(hSha1.Sum(nil))
}

// SHA256String  对字符串SHA256处理
func SHA256String(plain string) string {
	hSha256 := sha256.New()
	hSha256.Write([]byte(plain))
	return hex.EncodeToString(hSha256.Sum(nil))
}

// SHA512String  对字符串SHA512处理
func SHA512String(plain string) string {
	hSha512 := sha512.New()
	hSha512.Write([]byte(plain))
	return hex.EncodeToString(hSha512.Sum(nil))
}
