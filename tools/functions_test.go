package tools

import (
	"testing"
	"time"
)

func TestMergeMap(t *testing.T) {
	//b := make(map[string]interface{})
	a := MergeMap(nil, nil)
	c := len(a)
	t.Log(a == nil)
	t.Log(c)
	O := ReadCsv("/Users/yanjiankai/Documents/work/xunhuan/base/go-stress-testing/tools/aa.csv")
	SetOid(O)
	for i := 0; i < 1000; i++ {
		go t.Log(getOidOneByOne())
	}

	time.Sleep(10 * time.Second)

}
