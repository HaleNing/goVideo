package cache

import (
	"fmt"
	"strconv"
)

const (
	DailyRankKey = "rank:daily"
)

// view:video:1--->10000点击数
// view:video:2--->10点击数

// 创建某视频的点击数的key
func VideoViewKey(id uint) string {
	return fmt.Sprintf("view:video:%s", strconv.Itoa(int(id)))
}
