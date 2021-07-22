package client_server

import (
	"neko_server_go/utils"
	"net"
)

type RequestParam struct {
	Content []byte
	Length int
}

func (r *RequestParam)RequestParamEnd() bool {
	return r.Length == len(r.Content)
}

type Request struct {
	Command     RedisCommand
	Params      []RequestParam
	ParamsCount int
}

func (r *Request)LastParamEnd() bool {
	if len(r.Params) > 0 {
		return r.Params[len(r.Params)-1].RequestParamEnd()
	}
	return true
}

func (r *Request)LastParamLength() int {
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


func (r *RequestHandler)HandleRecv() {
	// 处理Param
	// 如果存在最后一个Param，则读的长度为最后一个Param的长度
	if r.CurrentRequest != nil && len(r.CurrentRequest.Params) > 0 && !r.CurrentRequest.LastParamEnd() && len(r.UnhandleBuffer) == r.CurrentRequest.LastParamLength() {
		r.CurrentRequest.Params[len(r.CurrentRequest.Params)-1].Content = r.UnhandleBuffer
		// 清空UnhandleBuffer
		r.UnhandleBuffer = make([]byte, 0)
	} else if string(r.UnhandleBuffer[:len(r.UnhandleBuffer)-2]) == "\\r\\n" {
		if r.CurrentRequest == nil {
			r.CurrentRequest = &Request{}
		}
		// 判断这个request是否是完整可用的
		if r.CurrentRequest.ParamsCount != 0 && len(r.CurrentRequest.Params) + 1 == r.CurrentRequest.ParamsCount {
			r.WaitHandleRequest = append(r.WaitHandleRequest, r.CurrentRequest)
			r.CurrentRequest = nil
			// 清空UnhandleBuffer
			r.UnhandleBuffer = make([]byte, 0)
		}
		// 判断这个参数应该去哪里
		if string(r.UnhandleBuffer[0]) == "*" {
			// *表示参数数量
			r.CurrentRequest.ParamsCount = BytesToInt(r.UnhandleBuffer[1:])
			// 清空UnhandleBuffer
			r.UnhandleBuffer = make([]byte, 0)
		} else if string(r.UnhandleBuffer[0]) == "$" {
			p := RequestParam{
				Length: BytesToInt(r.UnhandleBuffer[1:]),
			}
			r.CurrentRequest.Params = append(r.CurrentRequest.Params, p)
			// 清空UnhandleBuffer
			r.UnhandleBuffer = make([]byte, 0)

		}
		// TODO
	}
}


func (r *RequestHandler)Parse(buffer []byte) {
	for _, b := range buffer {
		r.UnhandleBuffer = append(r.UnhandleBuffer, b)
		r.HandleRecv()
	}
}

func (r *RequestHandler)OneStep() {
	buf := make([]byte, SettingsBufferSize)
	n, err := r.Conn.Read(buf)
	if err != nil {
		utils.LogError("conn read error: ", err)
		// 关闭连接
		r.EndConn = true
	}
	if n > 0 {
		r.Parse(buf)
		r.UnhandleBuffer = append(r.UnhandleBuffer, buf...)
	}
}

func (r *RequestHandler) Process() {
	for {
		// 关闭连接以及没有可以处理的请求内容，退出循环
		if r.EndConn == true && len(r.WaitHandleRequest) <= 0 {
			break
		}

		// 先处理没有处理的请求
		if len(r.WaitHandleRequest) > 0 {

		} else {
			r.OneStep()
		}
	}
}
