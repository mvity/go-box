// Copyright © 2021 - 2022 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package k

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
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
