package http_scenes

import (
	"encoding/json"
	"github.com/jmespath/go-jmespath"
	"github.com/link1st/go-stress-testing/model"
	"github.com/link1st/go-stress-testing/server/client"
	"github.com/link1st/go-stress-testing/tools"
	"io/ioutil"
	"net/http"
)

// ReqListScenes 接口分步压测
var ReqListScenes = make(map[string][]*model.Request)

// StepByStepRequest 分别压测数据初始化
type StepByStepRequest struct {
	localVariables map[string]interface{}
}

func GetReqListScenes(key string) ([]*model.Request, bool) {
	value, ok := ReqListScenes[key]
	return value, ok
}

// SendList  多个接口分步压测
func (f StepByStepRequest) SendList(requestList []*model.Request) (isSucceed bool, errCode int, requestTime uint64, contentLength int64) {
	errCode = model.HTTPOk
	f.localVariables = make(map[string]interface{})
	for _, request := range requestList {
		succeed, code, u, length := f.send(request)
		isSucceed = succeed
		errCode = code
		requestTime = requestTime + u
		contentLength = contentLength + length
		if succeed == false {
			break
		}
	}
	return
}

// send 发送一次请求
func (f StepByStepRequest) send(request *model.Request) (bool, int, uint64, int64) {
	var (
		// startTime = time.Now()
		isSucceed     = false
		errCode       = model.HTTPOk
		contentLength = int64(0)
		err           error
		resp          *http.Response
		requestTime   uint64
		respStep      []byte
	)
	//newRequest := getRequest(request)
	if request.Variables != nil {
		//以request 的参数为准合并extract 的数据
		request.Variables = tools.MergeMap(request.Variables, f.localVariables)
		request = tools.RequestFormat(request, request.Variables)
	}

	resp, requestTime, err = client.HTTPRequest(request)
	//提取接口返回结果的参数
	if request.Extract != nil {
		respStep, err = ioutil.ReadAll(ioutil.NopCloser(resp.Body))
		for k, v := range request.Extract {
			va, _ := extractValue(respStep, v)
			f.localVariables[k] = va
		}
	}
	if err != nil {
		errCode = model.RequestErr // 请求错误
	} else {
		contentLength = resp.ContentLength
		// 验证请求是否成功
		errCode, isSucceed = request.GetVerifyHTTP()(request, resp)
	}

	return isSucceed, errCode, requestTime, contentLength
}
func extractValue(r []byte, path string) (interface{}, error) {
	var jsonData interface{}
	err := json.Unmarshal(r, &jsonData)
	if err != nil {
		return nil, err
	}

	res, err := jmespath.Search(path, jsonData)
	return res, err
}
