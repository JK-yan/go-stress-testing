package tools

import (
	"math/rand"
	"time"
)

var Functions = map[string]interface{}{
	"get_timestamp": getTimestamp, // call without arguments
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
func getTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// MergeMap 以master 为主合并second。
func MergeMap(master map[string]interface{}, second map[string]interface{}) map[string]interface{} {
	for k, v := range second {
		if _, ok := master[k]; !ok {
			master[k] = v
		}
	}
	return master
}
