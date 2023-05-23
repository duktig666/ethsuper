//  Copyright © 2022-2023. duktig666 Limited.
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package timetool

import (
	"math/rand"
	"time"
)

// RandomTime Between min Duration and max Duration, random values are taken
func RandomTime(minDuration, maxDuration time.Duration) time.Duration {
	duration := time.Duration(rand.Int63n(int64(maxDuration-minDuration+1)) + int64(minDuration))
	return duration
}

// East8Time 返回东八区时间
func East8Time() time.Time {
	// 中国没有夏令时，使用一个固定的8小时的UTC时差。
	// 对于很多其他国家需要考虑夏令时。
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	// FixedZone 返回始终使用给定区域名称和偏移量(UTC 以东秒)的 Location。
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

	// 与上述北京时间等同
	//shanghai, err := time.LoadLocation("Asia/Shanghai")

	return time.Now().In(beijing)
}

// East8TimeDefaultFormat 东八区时间，默认格式化：YYYY-MM-DD HH:MM:SS
func East8TimeDefaultFormat() string {
	return East8Time().Format("2006-01-02 15:04:05")
}
