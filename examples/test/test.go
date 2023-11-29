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
	"time"
)

func main() {

	//now := time.Date(2023, 8, 31, 0, 0, 0, 0, time.Now().Location())
	//
	//fmt.Println(x.GetWeekdayOfDay(now, time.Monday))
	//fmt.Println(x.GetWeekdayOfDay(now, time.Tuesday))
	//fmt.Println(x.GetWeekdayOfDay(now, time.Wednesday))
	//fmt.Println(x.GetWeekdayOfDay(now, time.Thursday))
	//fmt.Println(x.GetWeekdayOfDay(now, time.Friday))
	//fmt.Println(x.GetWeekdayOfDay(now, time.Saturday))
	//fmt.Println(x.GetWeekdayOfDay(now, time.Sunday))

	fmt.Println(time.Now().Format("200601021505") + x.RandomString(5, false, true))
}
