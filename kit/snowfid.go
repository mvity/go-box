// Copyright © 2021 - 2022 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package k

import (
	"sync"
	"time"
)

const (
	snowfIDnumberBits  uint8 = 12                                    // 表示每个集群下的每个节点，1毫秒内可生成的id序号的二进制位
	snowfIDworkerBits  uint8 = 10                                    // 每台机器(节点)的ID位数 10位最大可以有2^10=1024个节点数 即每毫秒可生成 2^12-1=4096个唯一ID
	snowfIDworkerMax   int64 = -1 ^ (-1 << snowfIDworkerBits)        // 节点ID的最大值，用于防止溢出
	snowfIDnumberMax   int64 = -1 ^ (-1 << snowfIDnumberBits)        // 同上，用来表示生成id序号的最大值
	snowfIDtimeShift         = snowfIDworkerBits + snowfIDnumberBits // 时间戳向左的偏移量
	snowfIDworkerShift       = snowfIDnumberBits                     // 节点ID向左的偏移量
)

type snowfIDWorker struct {
	mu        sync.Mutex
	timestamp int64
	workerId  int64
	number    int64
	startTime int64 // 固定时间戳，毫秒值，使用后不可更改
}

// SnowfIDWorker 创建雪花ID实例
func SnowfIDWorker(workerId, startTime int64) *snowfIDWorker {
	if workerId < 0 || workerId > snowfIDworkerMax {
		workerId = 0
	}
	// 生成一个新节点
	return &snowfIDWorker{
		timestamp: 0,
		workerId:  workerId,
		number:    0,
		startTime: startTime,
	}
}

func (w *snowfIDWorker) GetID() uint64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	nowTime := time.Now()
	now := nowTime.UnixNano() / 1e6
	if w.timestamp == now {
		w.number++
		if w.number > snowfIDnumberMax {
			for now <= w.timestamp {
				now = nowTime.UnixNano() / 1e6
			}
		}
	} else {
		w.number = 0
		w.timestamp = now
	}
	id := (now-w.startTime)<<snowfIDtimeShift | (w.workerId << snowfIDworkerShift) | (w.number)
	return uint64(id)
}
