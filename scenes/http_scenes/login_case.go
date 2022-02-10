package http_scenes

import (
	"github.com/link1st/go-stress-testing/model"
	"time"
)

func init() {
	// 需要压测的接口参数
	clients := make([]*model.Request, 0)
	// 压测第一步
	clients = append(clients, &model.Request{
		URL:    "https://iam-02.rcrai.com/v1/login", // 请求url
		Form:   "http",                              // 请求方式 示例参数:http/webSocket/tcp
		Method: "POST",                              // 请求方法 示例参数:GET/POST/PUT
		Headers: map[string]string{
			"Rcrai-Bid": "60f697f4cbc9590001290b99",
		}, // headers 头信息
		Body:    "{\"password\":\"test1234!\",\"email\":\"admin@zsk.com\"}",
		Verify:  "statusCode",     // 验证的方法 示例参数:statusCode、json
		Timeout: 30 * time.Second, // 是否开启Debug模式
		Debug:   false,            // 是否开启Debug模式
		Code:    200,
		Extract: map[string]string{
			"accessToken": "token.accessToken",
		},
	})

	// 压测第二步
	clients = append(clients, &model.Request{
		URL:    "https://iam-02.rcrai.com/v1/token/$accessToken", // 请求url
		Form:   "http",                                           // 请求方式 示例参数:http/webSocket/tcp
		Method: "GET",                                            // 请求方法 示例参数:GET/POST/PUT
		Headers: map[string]string{
			"Rcrai-Bid": "60f697f4cbc9590001290b99",
		}, // headers 头信息
		Verify:  "statusCode",     // 验证的方法 示例参数:statusCode、json
		Timeout: 30 * time.Second, // 是否开启Debug模式
		Debug:   false,            // 是否开启Debug模式
		Variables: map[string]interface{}{
			"accessToken": "$accessToken",
		},
	})
	ReqListScenes["login"] = clients
}