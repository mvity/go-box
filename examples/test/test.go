/*
 * Copyright © 2021 - 2025 vity <vityme@icloud.com>.
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

	start := x.GetZeroTime(x.ParseDateTime("2024-11-03 09:10:30"))
	finish := x.GetZeroTime(x.ParseDateTime("2024-11-05 08:30:40"))

	dur := finish.Sub(start)

	fmt.Println(dur.Hours())
	//
	//step := 10
	//
	//fmt.Println(float64(9)/float64(step), math.Ceil(float64(9)/float64(step)))
	//fmt.Println(float64(19)/float64(step), math.Ceil(float64(19)/float64(step)))
	//fmt.Println(float64(20)/float64(step), math.Ceil(float64(20)/float64(step)))
	//fmt.Println(float64(21)/float64(step), math.Ceil(float64(21)/float64(step)))

	//minute := 15
	//
	//t := time.Now().Add(time.Minute * time.Duration(-minute))

	//fmt.Println(t)

	//addr := "https://xjl-ops.wh.lepincx.cn"
	//
	//if uri, err := url.Parse(addr); err == nil {
	//	code := strings.Split(uri.Host, "-")[0]
	//
	//	fmt.Println(code)
	//}

	//now := time.Date(2023, 8, 31, 0, 0, 0, 0, time.Now().Location())
	//
	//fmt.Println(x.GetWeekdayOfDay(now, time.Monday))
	//fmt.Println(x.GetWeekdayOfDay(now, time.Tuesday))
	//fmt.Println(x.GetWeekdayOfDay(now, time.Wednesday))
	//fmt.Println(x.GetWeekdayOfDay(now, time.Thursday))
	//fmt.Println(x.GetWeekdayOfDay(now, time.Friday))
	//fmt.Println(x.GetWeekdayOfDay(now, time.Saturday))
	//fmt.Println(x.GetWeekdayOfDay(now, time.Sunday))

	//key := "udik876ehjde32dU61edsxsf"

	//fmt.Println("rst : ", x.DESTripleEncrypt(key, "7wsVZ9EX"))

	//{
	//	str := x.RandomString(8, true, true)
	//
	//	fmt.Println("str : ", str)
	//
	//	key := "udik876ehjde32dU61edsxsf"
	//
	//	fmt.Println("rst : ", x.DESTripleEncrypt(key, str))
	//
	//	//fmt.Println("bak : ", x.DESTripleDecrypt(key, x.DESTripleEncrypt(key, str)))
	//}
	//
	//{
	//	str := x.RandomString(32, true, true)
	//
	//	fmt.Println("str : ", str)
	//
	//	key := "udik876ehjde32dU61edsxsf"
	//
	//	fmt.Println("rst : ", x.DESTripleEncrypt(key, str))
	//
	//	//fmt.Println("bak : ", x.DESTripleDecrypt(key, x.DESTripleEncrypt(key, str)))
	//
	//}
	//
	//{
	//	str := x.RandomString(64, true, true)
	//
	//	fmt.Println("str : ", str)
	//
	//	key := "udik876ehjde32dU61edsxsf"
	//
	//	fmt.Println("rst : ", x.DESTripleEncrypt(key, str))
	//
	//	//fmt.Println("bak : ", x.DESTripleDecrypt(key, x.DESTripleEncrypt(key, str)))
	//
	//}
	//
	//{
	//	str := x.RandomString(128, true, true)
	//
	//	fmt.Println("str : ", str)
	//
	//	key := "udik876ehjde32dU61edsxsf"
	//
	//	fmt.Println("rst : ", x.DESTripleEncrypt(key, str))
	//
	//	//fmt.Println("bak : ", x.DESTripleDecrypt(key, x.DESTripleEncrypt(key, str)))
	//
	//}
}
