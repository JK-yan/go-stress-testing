package tools

import (
	"fmt"
	"testing"
)

func TestReplaceStringVariables(t *testing.T) {
	m := map[string]interface{}{
		"aa": "$get_timestamp()",
	}
	a := ReplaceStringVariables("ss$aa", m)
	fmt.Println(a)

}
func TestCallBackFunc(t *testing.T) {
	a := Functions["get_timestamp"]
	b := CallBackFunc(a)
	fmt.Println(b)

}
