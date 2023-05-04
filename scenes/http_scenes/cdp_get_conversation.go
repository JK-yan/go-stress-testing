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
			"Rcrai-Bid": "631aec34cbc9590001780b62",
		}, // headers 头信息
		Body:    "{\"callback\":\"http://qa-inner-pre.rcrai.com/cdp/api/automation\",\"category\":\"默认分类\",\"customer\":{\"id\":\"$customer_id\",\"is_high_sea\":false,\"level\":\"1\",\"name\":\"成星泽\",\"phone\":\"13382061619\",\"source\":\"source\"},\"is_initial\":false,\"roles\":[\"c\",\"s\"],\"source_id\":\"$source_id\",\"staff_id\":\"f60b23e8-485e-4e01-87db-f0df6c2573e2\",\"timestamp\":1658409425636,\"url\":\"https://fileserver.rcrai.com/QA-file/test-data/call/call_mp3_5facd382cbc959000112c912.mp3\",\"user_defined_values\":{},\"messages\":[{\"id\":\"202007070620142128\",\"message\":\"你好\",\"message_type\":\"text\",\"speaker_type\":\"c\",\"start_time\":1617086604,\"end_time\":1617086604},{\"id\":\"202007070620142130\",\"message\":\"你好\",\"message_type\":\"text\",\"speaker_type\":\"s\",\"start_time\":1617086604,\"end_time\":1617086604}]}",
		Verify:  "statusCode",     // 验证的方法 示例参数:statusCode、json
		Timeout: 30 * time.Second, // 是否开启Debug模式
		Debug:   false,            // 是否开启Debug模式
		Code:    200,
		Variables: map[string]interface{}{
			"oid": "$get_oid_one()",
		},
	})

	ReqListScenes["get_conversation"] = clients
}
