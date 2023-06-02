/*
 * Copyright © 2021 - 2023 vity <vityme@icloud.com>.
 *
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

package main

import (
	"fmt"
	"github.com/mvity/go-box/x"
)

func main() {

	p := x.NewGeoPoint("113.685214,34.789401")
	ps := x.NewGeoPoints("113.736638,34.842372:113.615533,34.843031:113.501244,34.810442:113.503647,34.737158:113.705633,34.611918:113.849389,34.772623")
	fmt.Println(x.GeoCheckInAreas(p, ps))
	//
	//runReqs()
}

func runReqs() {
	//var res *http.Response
	//var body []byte
	//var err error
	//res = doApi1()
	//defer func(Body io.ReadCloser) {
	//	_ = Body.Close()
	//}(res.Body)
	//body, err = io.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(body))
	//
	//res = doApi2()
	//defer func(Body io.ReadCloser) {
	//	_ = Body.Close()
	//}(res.Body)
	//body, err = io.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(body))
	//
	//res = doApi3()
	//defer func(Body io.ReadCloser) {
	//	_ = Body.Close()
	//}(res.Body)
	//body, err = io.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(body))

}

//func doPost(uri string, headers map[string]string, body string) *http.Response {
//	proxy, _ := url.Parse("http://127.0.0.1:8888")
//	client := http.Client{
//		Transport: &http.Transport{
//			Proxy:       http.ProxyURL(proxy),
//			DialContext: (&net.Dialer{KeepAlive: 30 * time.Second}).DialContext,
//		},
//	}
//	t1 := time.Now()
//	req, err := http.NewRequest("POST", uri, bytes.NewBuffer([]byte(body)))
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
//	fmt.Println("Req ", uri, "time : ", t2)
//	return res
//}
//
//func doApi1() *http.Response {
//	uri := ""
//	headers := map[string]string{
//		"User-Agent":      "UnityPlayer/2020.3.34f1 (UnityWebRequest/1.0, libcurl/7.80.0-DEV)",
//		"Accept":          "*/*",
//		"Accept-Encoding": "deflate, gzip",
//		"Context-Type":    "application/octet-stream",
//		"x_acts":          "",
//		"X-Unity-Version": "2020.3.34f1",
//	}
//	body := ""
//	return doPost(uri, headers, body)
//}
//
//func doApi2() *http.Response {
//	uri := ""
//	headers := map[string]string{
//		"User-Agent":      "UnityPlayer/2020.3.34f1 (UnityWebRequest/1.0, libcurl/7.80.0-DEV)",
//		"Accept":          "*/*",
//		"Accept-Encoding": "deflate, gzip",
//		"Context-Type":    "application/octet-stream",
//		"x_acts":          "",
//		"X-Unity-Version": "2020.3.34f1",
//	}
//	body := ""
//	return doPost(uri, headers, body)
//}
//
//func doApi3() *http.Response {
//	uri := ""
//	headers := map[string]string{
//		"User-Agent":      "UnityPlayer/2020.3.34f1 (UnityWebRequest/1.0, libcurl/7.80.0-DEV)",
//		"Accept":          "*/*",
//		"Accept-Encoding": "deflate, gzip",
//		"Context-Type":    "application/octet-stream",
//		"x_acts":          "",
//		"X-Unity-Version": "2020.3.34f1",
//	}
//	body := ""
//	return doPost(uri, headers, body)
//}
