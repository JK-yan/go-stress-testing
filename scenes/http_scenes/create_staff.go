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
		URL:    "http://qa-inner-pre.rcrai.com/iam/v1/staff", // 请求url
		Form:   "http",                                       // 请求方式 示例参数:http/webSocket/tcp
		Method: "POST",                                       // 请求方法 示例参数:GET/POST/PUT
		Headers: map[string]string{
			"Rcrai-Bid": "6226ca21cbc95900014de733",
		}, // headers 头信息
		Body:    "{\"departmentExternalId\":\"root\",\"email\":\"1660566073770@qq.com\",\"externalId\":\"$externalId\",\"managerExternalId\":\"523de8a3-32d0-460a-a47f-fe98a0282c8f\",\"name\":\"White\",\"roleIds\":[]}",
		Verify:  "statusCode",     // 验证的方法 示例参数:statusCode、json
		Timeout: 30 * time.Second, // 是否开启Debug模式
		Debug:   false,            // 是否开启Debug模式
		Code:    200,
		Variables: map[string]interface{}{
			"externalId": "$get_uuid()",
		},
		//Extract: map[string]string{
		//	"accessToken": "token.accessToken",
		//},
	})
	ReqListScenes["create_staff"] = clients
}
