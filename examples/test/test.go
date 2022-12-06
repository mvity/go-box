// Copyright © 2021 - 2022 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	k "github.com/mvity/go-box/kit"
)

func main() {
	//
	//jsonA := `{
	//  "default": true,
	//  "year": "2020",
	//  "color": null,
	//  "door": "4",
	//  "type": "三厢",
	//  "pic": "https://xxxx.com/xxx.webp"
	//}`
	//node := k.ParseJSONString(jsonA, false)
	//
	//fmt.Println(node.Name("color").IsEmpty())
	//fmt.Println(node.Name("color").Size())
	//fmt.Println(node.Name("color").Int64())
	//fmt.Println(node.Name("color").String())
	//fmt.Println(node.Name("color").Boolean())
	//fmt.Println(node.Name("color").Name("0"))
	//
	//digitsRegexp := regexp.MustCompile(`\d+`)
	//
	//fmt.Println(digitsRegexp.FindAllString("5人", -1))

	//println(base64.StdEncoding.EncodeToString([]byte("abcdef律框架1123abcdef律框架1123abcdef律框架1123&=023")))

	// 				count := int64(math.Max(math.Ceil(float64(kit.HourToDayValue(hours)/15)), 1))
	//println(int64(math.Max(math.Ceil(float64(16)/15), 1)))

	fmt.Println("Hello, 世界")
	//var val uint64 = 730
	//var ratio uint64 = 50000
	//res := val * ratio / 1e6
	//fmt.Println("res:", res)
	//
	//var val2 float64 = 10
	//var ratio2 float64 = 100
	//res2 := val2 * (ratio2 / 1000000.00)
	//fmt.Println("res2:", res2)

	var total int64 = 150
	var count int64 = 50

	//percent := (float64(count) * float64(total)) / float64(total)
	percent := float64(count) / float64(total) * 100.0

	k.CastToBool(11)
	println(fmt.Sprintf("%.0f", percent))

}
