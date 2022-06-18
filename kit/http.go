// Copyright © 2021 - 2022 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package k

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
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

// HttpPostForm 执行Post Form 请求
func HttpPostForm(requestUrl string, data url.Values) (bool, string, int) {
	form := ioutil.NopCloser(strings.NewReader(data.Encode()))
	req, err := http.NewRequest("POST", requestUrl, form)
	if err != nil {
		return false, fmt.Sprintf("[%s] %s", "k.HttpPostJson", "Network Create Error："+err.Error()), 0
	}
	client := http.Client{}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Sprintf("[%s] %s", "k.HttpPostJson", "Network Request Error："+err.Error()), 0
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Sprintf("[%s] %s", "k.HttpPostJson", "Network Response Error："+err.Error()), 0
	}
	return true, string(body), resp.StatusCode

	//client := http.Client{}
	//resp, err := client.PostForm(requestUrl, data)
	//if err != nil {
	//	return false, fmt.Sprintf("[%s] %s", "k.HttpPostForm", "Network Request Error："+err.Error()), 0
	//}
	//defer func(Body io.ReadCloser) {
	//	_ = Body.Close()
	//}(resp.Body)
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return false, fmt.Sprintf("[%s] %s", "k.HttpPostForm", "Network Response Error："+err.Error()), 0
	//}
	//return true, string(body), resp.StatusCode
}

// HttpPostJson 执行Post JSON请求
func HttpPostJson(requestUrl string, json string) (bool, string, int) {
	req, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer([]byte(json)))
	if err != nil {
		return false, fmt.Sprintf("[%s] %s", "k.HttpPostJson", "Network Create Error："+err.Error()), 0
	}
	client := http.Client{}
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Sprintf("[%s] %s", "k.HttpPostJson", "Network Request Error："+err.Error()), 0
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Sprintf("[%s] %s", "k.HttpPostJson", "Network Response Error："+err.Error()), 0
	}
	return true, string(body), resp.StatusCode
}

// HttpPostJsonWithHeader 执行Post JSON请求
func HttpPostJsonWithHeader(requestUrl string, json string, headers map[string]string) (bool, string, int) {
	req, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer([]byte(json)))
	if err != nil {
		return false, fmt.Sprintf("[%s] %s", "k.HttpPostJsonWithHeader", "Network Create Error："+err.Error()), 0
	}
	client := http.Client{}
	req.Header.Set("Content-Type", "application/json")
	for key, val := range headers {
		req.Header.Set(key, val)
	}
	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Sprintf("[%s] %s", "k.HttpPostJsonWithHeader", "Network Request Error："+err.Error()), 0
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Sprintf("[%s] %s", "k.HttpPostJsonWithHeader", "Network Response Error："+err.Error()), 0
	}
	return true, string(body), resp.StatusCode
}

// HttpPostXmlSecure 执行Post XML 证书请求
func HttpPostXmlSecure(requestUrl string, xml string, certFile string, keyFile string, rootCaFile string) (bool, string, int) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return false, fmt.Sprintf("[%s] %s", "k.HttpPostXmlSecure", "Init Cert & Key error:"+err.Error()), 0
	}
	root, err := ioutil.ReadFile(rootCaFile)
	if err != nil {
		return false, fmt.Sprintf("[%s] %s", "k.HttpPostXmlSecure", "Init Rootca Cert error:"+err.Error()), 0
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(root)
	conf := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	c := &http.Client{Transport: &http.Transport{TLSClientConfig: conf}}

	req, err := http.NewRequest(http.MethodPost, requestUrl, bytes.NewBuffer([]byte(xml)))
	if err != nil {
		return false, fmt.Sprintf("[%s] %s", "k.HttpPostXmlSecure", "Network Create Error："+err.Error()), 0
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Do(req)
	if err != nil {
		return false, fmt.Sprintf("[%s] %s", "k.HttpPostXmlSecure", "Network Request Error："+err.Error()), 0
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Sprintf("[%s] %s", "k.HttpPostXmlSecure", "Network Response Error："+err.Error()), 0
	}
	return true, string(body), resp.StatusCode
}
