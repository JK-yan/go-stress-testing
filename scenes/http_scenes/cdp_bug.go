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
		URL:    "http://172.16.4.206/v1/conversations/text", // 请求url
		Form:   "http",                                      // 请求方式 示例参数:http/webSocket/tcp
		Method: "POST",                                      // 请求方法 示例参数:GET/POST/PUT
		Headers: map[string]string{
			"Rcrai-Bid": "61b04d95cbc959000174b38f",
		}, // headers 头信息
		Body:    "{\"start_time\":1618390937,\"end_time\":1618391937,\"messages\":[{\"message_type\":\"text\",\"speaker_type\":\"c\",\"id\":\"202007070620142128\",\"message\":\"你好\",\"timestamp\":1617086604},{\"message_type\":\"text\",\"speaker_type\":\"s\",\"id\":\"202007070620142130\",\"message\":\"你好\",\"timestamp\":1617086604},{\"message_type\":\"text\",\"speaker_type\":\"s\",\"id\":\"202007070620142131\",\"message\":\" 你好 我这边是循环智能 请问你是吴越吗？ \",\"timestamp\":1617086604},{\"message_type\":\"text\",\"speaker_type\":\"c\",\"id\":\"202007070620142145\",\"message\":\"哎,你好.请问你有什么事情？我这边没有太多时间。\",\"timestamp\":1617086604},{\"message_type\":\"text\",\"speaker_type\":\"s\",\"id\":\"202007070620142146\",\"message\":\"嗯，那我长话短说。我这边目前有一款保险比较适合您，不知道您这边想不想了解一下。\",\"timestamp\":1617086604},{\"message_type\":\"text\",\"speaker_type\":\"c\",\"id\":\"202007070620142149\",\"message\":\"多少钱啊？贵吗？\",\"timestamp\":1617086604},{\"message_type\":\"text\",\"speaker_type\":\"s\",\"id\":\"202007070620142151\",\"message\":\"这个是一个套餐，可以进行进行不同的组合，不一样的组合价格不一样，你这边可以根据自己的情况进行选择。\",\"timestamp\":1617086604},{\"message_type\":\"text\",\"speaker_type\":\"c\",\"id\":\"202007070620142154\",\"message\":\"有价格在5000到10000之间的套餐吗？\",\"timestamp\":1617086604},{\"message_type\":\"text\",\"speaker_type\":\"s\",\"id\":\"202007070620142158\",\"message\":\"有的，目前有一个每年8000块钱，适合年龄30-35之间男性的保险套餐，包括：意外险、大病医疗\",\"timestamp\":1617086604},{\"message_type\":\"text\",\"speaker_type\":\"c\",\"id\":\"202007070620142200\",\"message\":\"嗯，你这边有相关的资料吗？\",\"timestamp\":1617086604},{\"message_type\":\"text\",\"speaker_type\":\"s\",\"id\":\"202007070620142300\",\"message\":\"有的，有的。要不我加一下你的微信，我给你发一下。\",\"timestamp\":1617086604},{\"message_type\":\"text\",\"speaker_type\":\"c\",\"id\":\"202007070620142320\",\"message\":\"好的，我的微信号15120069282\",\"timestamp\":1617086604},{\"message_type\":\"text\",\"speaker_type\":\"s\",\"id\":\"202007070620142360\",\"message\":\"已经添加了，您通过一下。\",\"timestamp\":1617086604},{\"message_type\":\"text\",\"speaker_type\":\"c\",\"id\":\"202007070620142380\",\"message\":\"好的，好的。\",\"timestamp\":1617086604},{\"message_type\":\"text\",\"speaker_type\":\"s\",\"id\":\"202007070620142400\",\"message\":\"嗯，好的，那再见。\",\"timestamp\":1617086604},{\"message_type\":\"text\",\"speaker_type\":\"c\",\"id\":\"202007070620142420\",\"message\":\"嗯，再见。你大爷的，哈哈\",\"timestamp\":1617086604},{\"speaker_type\":\"s\",\"message_type\":\"text\",\"id\":\"9802900426815511738_1619093660_external\",\"message\":\"[亲亲][得意][嘴唇][发呆]\",\"timestamp\":1617086604},{\"speaker_type\":\"s\",\"message_type\":\"我是一个其他类型\",\"id\":\"9802900426815511738_1619093690_external\",\"message\":\"我是一个其他类型\",\"timestamp\":1617086604},{\"speaker_type\":\"c\",\"message_type\":\"others\",\"id\":\"9802900426815511738_1619093672_external\",\"message\":\"类型为others\",\"timestamp\":1617086604},{\"speaker_type\":\"s\",\"message_type\":\" \",\"id\":\"9802900426815511738_1619093673_external\",\"message\":\"类型为空格\",\"timestamp\":1617086604},{\"speaker_type\":\"s\",\"message_type\":\"!@#$%^&*()\",\"id\":\"9802900426815511738_1619093673_external\",\"message\":\"类型为!@#$%^&*()\",\"timestamp\":1617086604}],\"staff\":{\"name\":\"1645519416220\",\"id\":\"33ae9fa1-a719-4307-be3c-62c7f62cf7f9\"},\"source_id\":\"$source_id\",\"category\":\"默认分类\",\"customer\":{\"phone\":\"13850162551\",\"level\":\"1\",\"is_high_sea\":false,\"name\":\"yace123\",\"id\":\"444a311a-1111-45d9-a913-d149c779997c\",\"source\":\"source\"}}",
		Verify:  "statusCode",     // 验证的方法 示例参数:statusCode、json
		Timeout: 30 * time.Second, // 是否开启Debug模式
		Debug:   true,             // 是否开启Debug模式
		Code:    200,
		Variables: map[string]interface{}{
			"source_id": "$get_uuid()",
		},
	})

	ReqListScenes["cdp_bug"] = clients
}
