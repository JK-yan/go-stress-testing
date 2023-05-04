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
		URL:    "http://172.16.0.108/v1/staff?page.count=10000&page.startIndex=0", // 请求url
		Form:   "http",                                                            // 请求方式 示例参数:http/webSocket/tcp
		Method: "GET",                                                             // 请求方法 示例参数:GET/POST/PUT
		Headers: map[string]string{
			"Rcrai-Bid": "6226ca21cbc95900014de733",
		}, // headers 头信息
		//Body:    "{\"password\":\"1234qwerASDF!@#$\",\"email\":\"jichufuwu@rcrai.com\"}",
		Verify:  "statusCode",     // 验证的方法 示例参数:statusCode、json
		Timeout: 30 * time.Second, // 是否开启Debug模式
		Debug:   false,            // 是否开启Debug模式
		Code:    200,
		//Extract: map[string]string{
		//	"accessToken": "token.accessToken",
		//},
	})
	// 压测第二步
	//clients = append(clients, &model.Request{
	//	URL:    "http://debug.rcrai-staging.rcrai.com/zeus-iam/v1/token/$accessToken", // 请求url
	//	Form:   "http",                                                                // 请求方式 示例参数:http/webSocket/tcp
	//	Method: "GET",                                                                 // 请求方法 示例参数:GET/POST/PUT
	//	Headers: map[string]string{
	//		"Rcrai-Bid": "60f697f4cbc9590001290b99",
	//	}, // headers 头信息
	//	Verify:    "statusCode",     // 验证的方法 示例参数:statusCode、json
	//	Timeout:   30 * time.Second, // 是否开启Debug模式
	//	Debug:     true,             // 是否开启Debug模式
	//	Code:      200,
	//	Variables: map[string]interface{}{},
	//})
	ReqListScenes["list_staff"] = clients
}
