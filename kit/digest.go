// Copyright © 2021 - 2022 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package k

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

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
