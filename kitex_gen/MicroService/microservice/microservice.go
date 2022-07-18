// Code generated by Kitex v0.3.2. DO NOT EDIT.

package microservice

import (
	"context"
	"fmt"
	"github.com/caoxl-c/kitex-proto/kitex_gen/MicroService"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
	"google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return microServiceServiceInfo
}

var microServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "MicroService"
	handlerType := (*MicroService.MicroService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Ping": kitex.NewMethodInfo(pingHandler, newPingArgs, newPingResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "proto",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.3.2",
		Extra:           extra,
	}
	return svcInfo
}

func pingHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(MicroService.PingRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(MicroService.MicroService).Ping(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *PingArgs:
		success, err := handler.(MicroService.MicroService).Ping(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PingResult)
		realResult.Success = success
	}
	return nil
}
func newPingArgs() interface{} {
	return &PingArgs{}
}

func newPingResult() interface{} {
	return &PingResult{}
}

type PingArgs struct {
	Req *MicroService.PingRequest
}

func (p *PingArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PingArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *PingArgs) Unmarshal(in []byte) error {
	msg := new(MicroService.PingRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PingArgs_Req_DEFAULT *MicroService.PingRequest

func (p *PingArgs) GetReq() *MicroService.PingRequest {
	if !p.IsSetReq() {
		return PingArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PingArgs) IsSetReq() bool {
	return p.Req != nil
}

type PingResult struct {
	Success *MicroService.PingResponse
}

var PingResult_Success_DEFAULT *MicroService.PingResponse

func (p *PingResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PingResult")
	}
	return proto.Marshal(p.Success)
}

func (p *PingResult) Unmarshal(in []byte) error {
	msg := new(MicroService.PingResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PingResult) GetSuccess() *MicroService.PingResponse {
	if !p.IsSetSuccess() {
		return PingResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PingResult) SetSuccess(x interface{}) {
	p.Success = x.(*MicroService.PingResponse)
}

func (p *PingResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Ping(ctx context.Context, Req *MicroService.PingRequest) (r *MicroService.PingResponse, err error) {
	var _args PingArgs
	_args.Req = Req
	var _result PingResult
	if err = p.c.Call(ctx, "Ping", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
