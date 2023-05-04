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
		URL:    "http://qa-inner-pre.rcrai.com/cdp/v1/conversations/$oid", // 请求url
		Form:   "http",                                                    // 请求方式 示例参数:http/webSocket/tcp
		Method: "GET",                                                     // 请求方法 示例参数:GET/POST/PUT
		Headers: map[string]string{
			"Rcrai-Bid": "6226ca21cbc95900014de733",
		}, // headers 头信息
		Body:    "",
		Verify:  "statusCode",     // 验证的方法 示例参数:statusCode、json
		Timeout: 30 * time.Second, // 是否开启Debug模式
		Debug:   false,            // 是否开启Debug模式
		Code:    200,
		Variables: map[string]interface{}{
			"oid": "$get_oid()",
		},
	})

	ReqListScenes["cdp_key"] = clients
}
