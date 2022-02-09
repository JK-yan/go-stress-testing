package tools

import (
	"fmt"
	"github.com/link1st/go-stress-testing/model"
	"strings"
)

func replaceStringVariables(str string, m map[string]interface{}) string {
	if strings.Contains(str, "$") {
		for k, v := range m {
			str = strings.Replace(str, fmt.Sprint("$", k), v.(string), -1)
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

func RequestFormat(res *model.Request, m map[string]interface{}) *model.Request {
	request := &model.Request{
		URL:       replaceStringVariables(res.URL, m),
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
