package tools

import (
	"fmt"
	"github.com/link1st/go-stress-testing/model"
	"reflect"
	"strings"
)

func ReplaceStringVariables(str string, m map[string]interface{}) string {
	if strings.Contains(str, "$") {
		for k, v := range m {
			strValue := v.(string)
			if strings.HasPrefix(strValue, "$") && strings.HasSuffix(strValue, "()") {
				funcName := strValue[1:strings.Index(strValue, "(")]
				if f, ok := Functions[funcName]; ok {
					value := CallBackFunc(f)
					strValue = value
				}
				continue
			}
			str = strings.Replace(str, fmt.Sprint("$", k), strValue, -1)
		}
	}
	return str
}
func replaceMapVariables(va map[string]string, m map[string]interface{}) map[string]string {
	for k, v := range va {
		str := strings.Replace(v, "$", "", 1)
		if _, ok := m[str]; ok {
			va[k] = m[str].(string)
		}
	}
	return va
}
func CallBackFunc(i interface{}) string {
	var args []reflect.Value
	fn := reflect.ValueOf(i)
	return fmt.Sprintf("%v", fn.Call(args)[0])
}

func RequestFormat(res *model.Request, m map[string]interface{}) *model.Request {
	request := &model.Request{
		URL:       ReplaceStringVariables(res.URL, m),
		Form:      res.Form,
		Method:    res.Method,
		Headers:   replaceMapVariables(res.Headers, m),
		Body:      res.Body,
		Verify:    res.Verify,
		Timeout:   res.Timeout,
		Debug:     res.Debug,
		MaxCon:    res.MaxCon,
		HTTP2:     res.HTTP2,
		Keepalive: res.Keepalive,
		Code:      res.Code,
	}
	return request
}
