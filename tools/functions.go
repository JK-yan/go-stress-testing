package tools

import (
	"math/rand"
	"time"
)

var Functions = map[string]interface{}{}

func init() {
	rand.Seed(time.Now().Unix())
}
func MergeMap(master map[string]interface{}, second map[string]interface{}) map[string]interface{} {
	for k, v := range second {
		if _, ok := master[k]; !ok {
			master[k] = v
		}
	}
	return master
}
