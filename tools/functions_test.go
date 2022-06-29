package tools

import "testing"

func TestMergeMap(t *testing.T) {
	//b := make(map[string]interface{})
	a := MergeMap(nil, nil)
	c := len(a)
	t.Log(a == nil)
	t.Log(c)
}
