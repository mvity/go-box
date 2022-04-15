// Copyright © 2021 - 2022 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import k "github.com/mvity/go-box/kit"

func main() {

	idw := k.SnowfIDWorker(1, 1650032075572)
	println(idw.SnowfID())
}
