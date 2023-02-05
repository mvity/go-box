/*
 * Copyright © 2021 - 2023 vity <vityme@icloud.com>.
 *
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

package x

import "regexp"

const (
	regexpNumeric    = "^(\\-|\\+)?\\d+(\\.\\d+)?$"
	regxpChinaID     = "^(\\d{6})(19|20)(\\d{2})(1[0-2]|0[1-9])(0[1-9]|[1-2][0-9]|3[0-1])(\\d{3})(\\d|X|x)?$"
	regxpChinaMobile = "^1[3|4|5|6|7|8|9][0-9]\\d{8}$"
	regxpChinaCarNO  = "^(([京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领][A-Z](([0-9]{5}[ABCDEFGHJK])|([ABCDEFGHJK]([A-HJ-NP-Z0-9])[0-9]{4})))|([京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领]\\d{3}\\d{1,3}[领])|([京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领][A-Z][A-HJ-NP-Z0-9]{4}[A-HJ-NP-Z0-9挂学警港澳使领]))$"
)

var (
	RegexpNumeric     = regexp.MustCompile(regexpNumeric)    // 数字格式
	RegexpChinaID     = regexp.MustCompile(regxpChinaID)     // 中国身份证号
	RegexpChinaMobile = regexp.MustCompile(regxpChinaMobile) // 中国11位手机号
	RegexpChinaCarNO  = regexp.MustCompile(regxpChinaCarNO)  // 中国车牌号
)
