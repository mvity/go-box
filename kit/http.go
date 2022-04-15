// Copyright © 2021 - 2022 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package k

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// HttpGet 执行Get请求
func HttpGet(rUrl string) (bool, string, int) {
	req, err := http.NewRequest("GET", rUrl, nil)
	if err != nil {
		return false, fmt.Sprintf("[%s] %s", "k.HttpGet", "Network Create Error："+err.Error()), 0
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Sprintf("[%s] %s", "k.HttpGet", "Network Request Error："+err.Error()), 0
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Sprintf("[%s] %s", "k.HttpGet", "Network Response Error："+err.Error()), resp.StatusCode
	}
	return true, string(body), resp.StatusCode

}
