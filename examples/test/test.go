// Copyright © 2021 - 2022 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"regexp"
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
	//node := k.ParseJSONStringForce(jsonA, false)
	//
	//fmt.Println(node.Name("color").IsEmpty())
	//fmt.Println(node.Name("color").Size())
	//fmt.Println(node.Name("color").Int64())
	//fmt.Println(node.Name("color").String())
	//fmt.Println(node.Name("color").Boolean())
	//fmt.Println(node.Name("color").Name("0"))

	digitsRegexp := regexp.MustCompile(`\d+`)

	fmt.Println(digitsRegexp.FindAllString("5人", -1))
}
