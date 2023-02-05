/*
 * Copyright © 2021 - 2023 vity <vityme@icloud.com>.
 *
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */
package main

//
//import (
//	"fmt"
//	"io"
//	"net"
//	"net/http"
//	"net/url"
//	"time"
//)
//
//func main() {
//
//	headers := map[string]string{
//		"User-Agent":      "UnityPlayer/2020.3.34f1 (UnityWebRequest/1.0, libcurl/7.80.0-DEV)",
//		"Accept":          "*/*",
//		"Accept-Encoding": "deflate, gzip",
//		"Context-Type":    "application/octet-stream",
//		"x_acts":          "Single.gate_entry",
//		"X-Unity-Version": "2020.3.34f1",
//	}
//
//	proxy, _ := url.Parse("http://127.0.0.1:8888")
//	client := http.Client{
//		Transport: &http.Transport{
//			Proxy:       http.ProxyURL(proxy),
//			DialContext: (&net.Dialer{KeepAlive: 30 * time.Second}).DialContext,
//		},
//	}
//	t1 := time.Now()
//	req, err := http.NewRequest("POST", "http://127.0.0.1:32145/dev/test/api111", nil)
//	if err != nil {
//		println(err)
//	}
//	for key, val := range headers {
//		req.Header.Set(key, val)
//	}
//
//	res, err := client.Do(req)
//	t2 := time.Now().Sub(t1)
//	if err != nil {
//		println(err)
//	}
//	defer func(Body io.ReadCloser) {
//		_ = Body.Close()
//	}(res.Body)
//	body, err := io.ReadAll(res.Body)
//	fmt.Println(string(body))
//	fmt.Println("time : ", t2)
//}
