package server

import (
	"ne_cache/client_server/common"
	"neko_server_go/utils"
	"net"
	"strings"
)

type RequestParam struct {
	Content []byte
	Length  int
}

func (r *RequestParam) RequestParamEnd() bool {
	return r.Length == len(r.Content)
}

type Request struct {
	Command     common.RedisCommand
	Params      []RequestParam
	ParamsCount int
}

func (r *Request) LastParamEnd() bool {
	if len(r.Params) > 0 {
		return r.Params[len(r.Params)-1].RequestParamEnd()
	}
	return true
}

func (r *Request) LastParamLength() int {
	if len(r.Params) > 0 {
		return r.Params[len(r.Params)-1].Length
	}
	return 0
}

type RequestHandler struct {
	Conn              net.Conn
	UnhandleBuffer    []byte
	CurrentRequest    *Request
	WaitHandleRequest []*Request
	EndConn           bool
}

func (r *RequestHandler) HandleRecv() {
	if r.CurrentRequest == nil {
		r.CurrentRequest = &Request{}
	}
	// 处理Param
	// 如果存在最后一个Param，则读的长度为最后一个Param的长度
	if len(r.CurrentRequest.Params) > 0 && !r.CurrentRequest.LastParamEnd() && len(r.UnhandleBuffer) == r.CurrentRequest.LastParamLength() {
		r.CurrentRequest.Params[len(r.CurrentRequest.Params)-1].Content = r.UnhandleBuffer
		// 清空UnhandleBuffer
		r.UnhandleBuffer = make([]byte, 0)

		// 判断这个request是否是完整可用的
		if len(r.CurrentRequest.Params) == r.CurrentRequest.ParamsCount {
			c := strings.ToUpper(string(r.CurrentRequest.Params[0].Content))
			r.CurrentRequest.Command = common.RedisCommand(c)
			r.WaitHandleRequest = append(r.WaitHandleRequest, r.CurrentRequest)
			r.CurrentRequest = nil
		}
	} else if len(r.UnhandleBuffer) > 2 && string(r.UnhandleBuffer[len(r.UnhandleBuffer)-2:]) == "\r\n" {
		// 可能出现第一个是空的情况
		index := 1
		if r.UnhandleBuffer[0] == 0 {
			index = 2
		}
		if string(r.UnhandleBuffer[index-1]) == "*"{
			// *表示参数数量
			r.CurrentRequest.ParamsCount = common.BytesStringToInt(r.UnhandleBuffer[index:len(r.UnhandleBuffer)-2])
			// 清空UnhandleBuffer
			r.UnhandleBuffer = make([]byte, 0)
		} else if string(r.UnhandleBuffer[index-1]) == "$" {
			p := RequestParam{
				Length: common.BytesStringToInt(r.UnhandleBuffer[index:len(r.UnhandleBuffer)-2]),
			}
			r.CurrentRequest.Params = append(r.CurrentRequest.Params, p)
			// 清空UnhandleBuffer
			r.UnhandleBuffer = make([]byte, 0)
		} else {
			utils.LogError("出现未知模式, UnhandleBuffer: ", r.UnhandleBuffer)
		}
	} else if len(r.UnhandleBuffer) == 2 && string(r.UnhandleBuffer) == "\r\n" {
		// 只有\r\n是遗留的情况，直接清理
		// 清空UnhandleBuffer
		r.UnhandleBuffer = make([]byte, 0)
	}
}

func (r *RequestHandler) Parse(buffer []byte, length int) {
	for i := 0; i <= length; i++ {
		r.UnhandleBuffer = append(r.UnhandleBuffer, buffer[i])
		r.HandleRecv()
	}
}

func (r *RequestHandler) OneStep(settings common.SettingsBase) {
	buf := make([]byte, settings.SettingsBufferSize)
	n, err := r.Conn.Read(buf)
	if err != nil {
		utils.LogError("conn read error: ", err)
		// 关闭连接
		r.EndConn = true
	}
	if n > 0 {
		r.Parse(buf, n)
	}
}

func (r *RequestHandler) Process(settings common.SettingsBase) {
	for {
		// 关闭连接以及没有可以处理的请求内容，退出循环
		if r.EndConn == true && len(r.WaitHandleRequest) <= 0 {
			break
		}

		// 先处理没有处理的请求
		if len(r.WaitHandleRequest) > 0 {
			if handler, ok := Router[r.WaitHandleRequest[0].Command]; ok {
				// TODO 支持预校验request
				handler(settings, r.WaitHandleRequest[0], r.Conn)
			} else {
				// 404
				resp := Response{
					Conn: r.Conn,
				}
				resp.UnknownCommand(r.WaitHandleRequest[0].Command)
			}
			// 处理完拿出这个WaitHandleRequest
			r.WaitHandleRequest = r.WaitHandleRequest[1:]
		} else {
			r.OneStep(settings)
		}
	}
}
