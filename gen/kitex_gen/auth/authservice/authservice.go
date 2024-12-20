// Code generated by Kitex v0.11.3. DO NOT EDIT.

package authservice

import (
	"context"
	"errors"
	auth "github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/auth"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"GetToken": kitex.NewMethodInfo(
		getTokenHandler,
		newGetTokenArgs,
		newGetTokenResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"ParseAccessToken": kitex.NewMethodInfo(
		parseAccessTokenHandler,
		newParseAccessTokenArgs,
		newParseAccessTokenResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"ParseRefreshToken": kitex.NewMethodInfo(
		parseRefreshTokenHandler,
		newParseRefreshTokenArgs,
		newParseRefreshTokenResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"VerifyToken": kitex.NewMethodInfo(
		verifyTokenHandler,
		newVerifyTokenArgs,
		newVerifyTokenResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"ExecRefreshToken": kitex.NewMethodInfo(
		execRefreshTokenHandler,
		newExecRefreshTokenArgs,
		newExecRefreshTokenResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	authServiceServiceInfo                = NewServiceInfo()
	authServiceServiceInfoForClient       = NewServiceInfoForClient()
	authServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return authServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return authServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return authServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "AuthService"
	handlerType := (*auth.AuthService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "auth",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.11.3",
		Extra:           extra,
	}
	return svcInfo
}

func getTokenHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(auth.UserId)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(auth.AuthService).GetToken(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetTokenArgs:
		success, err := handler.(auth.AuthService).GetToken(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetTokenResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetTokenArgs() interface{} {
	return &GetTokenArgs{}
}

func newGetTokenResult() interface{} {
	return &GetTokenResult{}
}

type GetTokenArgs struct {
	Req *auth.UserId
}

func (p *GetTokenArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(auth.UserId)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetTokenArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetTokenArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetTokenArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetTokenArgs) Unmarshal(in []byte) error {
	msg := new(auth.UserId)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetTokenArgs_Req_DEFAULT *auth.UserId

func (p *GetTokenArgs) GetReq() *auth.UserId {
	if !p.IsSetReq() {
		return GetTokenArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetTokenArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetTokenArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetTokenResult struct {
	Success *auth.TwoToken
}

var GetTokenResult_Success_DEFAULT *auth.TwoToken

func (p *GetTokenResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(auth.TwoToken)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetTokenResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetTokenResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetTokenResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetTokenResult) Unmarshal(in []byte) error {
	msg := new(auth.TwoToken)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetTokenResult) GetSuccess() *auth.TwoToken {
	if !p.IsSetSuccess() {
		return GetTokenResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetTokenResult) SetSuccess(x interface{}) {
	p.Success = x.(*auth.TwoToken)
}

func (p *GetTokenResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetTokenResult) GetResult() interface{} {
	return p.Success
}

func parseAccessTokenHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(auth.AccessToken)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(auth.AuthService).ParseAccessToken(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ParseAccessTokenArgs:
		success, err := handler.(auth.AuthService).ParseAccessToken(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ParseAccessTokenResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newParseAccessTokenArgs() interface{} {
	return &ParseAccessTokenArgs{}
}

func newParseAccessTokenResult() interface{} {
	return &ParseAccessTokenResult{}
}

type ParseAccessTokenArgs struct {
	Req *auth.AccessToken
}

func (p *ParseAccessTokenArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(auth.AccessToken)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ParseAccessTokenArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ParseAccessTokenArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ParseAccessTokenArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ParseAccessTokenArgs) Unmarshal(in []byte) error {
	msg := new(auth.AccessToken)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ParseAccessTokenArgs_Req_DEFAULT *auth.AccessToken

func (p *ParseAccessTokenArgs) GetReq() *auth.AccessToken {
	if !p.IsSetReq() {
		return ParseAccessTokenArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ParseAccessTokenArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ParseAccessTokenArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ParseAccessTokenResult struct {
	Success *auth.UserId
}

var ParseAccessTokenResult_Success_DEFAULT *auth.UserId

func (p *ParseAccessTokenResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(auth.UserId)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ParseAccessTokenResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ParseAccessTokenResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ParseAccessTokenResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ParseAccessTokenResult) Unmarshal(in []byte) error {
	msg := new(auth.UserId)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ParseAccessTokenResult) GetSuccess() *auth.UserId {
	if !p.IsSetSuccess() {
		return ParseAccessTokenResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ParseAccessTokenResult) SetSuccess(x interface{}) {
	p.Success = x.(*auth.UserId)
}

func (p *ParseAccessTokenResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ParseAccessTokenResult) GetResult() interface{} {
	return p.Success
}

func parseRefreshTokenHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(auth.RefreshToken)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(auth.AuthService).ParseRefreshToken(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ParseRefreshTokenArgs:
		success, err := handler.(auth.AuthService).ParseRefreshToken(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ParseRefreshTokenResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newParseRefreshTokenArgs() interface{} {
	return &ParseRefreshTokenArgs{}
}

func newParseRefreshTokenResult() interface{} {
	return &ParseRefreshTokenResult{}
}

type ParseRefreshTokenArgs struct {
	Req *auth.RefreshToken
}

func (p *ParseRefreshTokenArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(auth.RefreshToken)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ParseRefreshTokenArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ParseRefreshTokenArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ParseRefreshTokenArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ParseRefreshTokenArgs) Unmarshal(in []byte) error {
	msg := new(auth.RefreshToken)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ParseRefreshTokenArgs_Req_DEFAULT *auth.RefreshToken

func (p *ParseRefreshTokenArgs) GetReq() *auth.RefreshToken {
	if !p.IsSetReq() {
		return ParseRefreshTokenArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ParseRefreshTokenArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ParseRefreshTokenArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ParseRefreshTokenResult struct {
	Success *auth.UserId
}

var ParseRefreshTokenResult_Success_DEFAULT *auth.UserId

func (p *ParseRefreshTokenResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(auth.UserId)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ParseRefreshTokenResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ParseRefreshTokenResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ParseRefreshTokenResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ParseRefreshTokenResult) Unmarshal(in []byte) error {
	msg := new(auth.UserId)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ParseRefreshTokenResult) GetSuccess() *auth.UserId {
	if !p.IsSetSuccess() {
		return ParseRefreshTokenResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ParseRefreshTokenResult) SetSuccess(x interface{}) {
	p.Success = x.(*auth.UserId)
}

func (p *ParseRefreshTokenResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ParseRefreshTokenResult) GetResult() interface{} {
	return p.Success
}

func verifyTokenHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(auth.AccessToken)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(auth.AuthService).VerifyToken(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *VerifyTokenArgs:
		success, err := handler.(auth.AuthService).VerifyToken(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*VerifyTokenResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newVerifyTokenArgs() interface{} {
	return &VerifyTokenArgs{}
}

func newVerifyTokenResult() interface{} {
	return &VerifyTokenResult{}
}

type VerifyTokenArgs struct {
	Req *auth.AccessToken
}

func (p *VerifyTokenArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(auth.AccessToken)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *VerifyTokenArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *VerifyTokenArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *VerifyTokenArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *VerifyTokenArgs) Unmarshal(in []byte) error {
	msg := new(auth.AccessToken)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var VerifyTokenArgs_Req_DEFAULT *auth.AccessToken

func (p *VerifyTokenArgs) GetReq() *auth.AccessToken {
	if !p.IsSetReq() {
		return VerifyTokenArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *VerifyTokenArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *VerifyTokenArgs) GetFirstArgument() interface{} {
	return p.Req
}

type VerifyTokenResult struct {
	Success *auth.UserId
}

var VerifyTokenResult_Success_DEFAULT *auth.UserId

func (p *VerifyTokenResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(auth.UserId)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *VerifyTokenResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *VerifyTokenResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *VerifyTokenResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *VerifyTokenResult) Unmarshal(in []byte) error {
	msg := new(auth.UserId)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *VerifyTokenResult) GetSuccess() *auth.UserId {
	if !p.IsSetSuccess() {
		return VerifyTokenResult_Success_DEFAULT
	}
	return p.Success
}

func (p *VerifyTokenResult) SetSuccess(x interface{}) {
	p.Success = x.(*auth.UserId)
}

func (p *VerifyTokenResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *VerifyTokenResult) GetResult() interface{} {
	return p.Success
}

func execRefreshTokenHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(auth.RefreshToken)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(auth.AuthService).ExecRefreshToken(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ExecRefreshTokenArgs:
		success, err := handler.(auth.AuthService).ExecRefreshToken(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ExecRefreshTokenResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newExecRefreshTokenArgs() interface{} {
	return &ExecRefreshTokenArgs{}
}

func newExecRefreshTokenResult() interface{} {
	return &ExecRefreshTokenResult{}
}

type ExecRefreshTokenArgs struct {
	Req *auth.RefreshToken
}

func (p *ExecRefreshTokenArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(auth.RefreshToken)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ExecRefreshTokenArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ExecRefreshTokenArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ExecRefreshTokenArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ExecRefreshTokenArgs) Unmarshal(in []byte) error {
	msg := new(auth.RefreshToken)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ExecRefreshTokenArgs_Req_DEFAULT *auth.RefreshToken

func (p *ExecRefreshTokenArgs) GetReq() *auth.RefreshToken {
	if !p.IsSetReq() {
		return ExecRefreshTokenArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ExecRefreshTokenArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ExecRefreshTokenArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ExecRefreshTokenResult struct {
	Success *auth.TwoToken
}

var ExecRefreshTokenResult_Success_DEFAULT *auth.TwoToken

func (p *ExecRefreshTokenResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(auth.TwoToken)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ExecRefreshTokenResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ExecRefreshTokenResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ExecRefreshTokenResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ExecRefreshTokenResult) Unmarshal(in []byte) error {
	msg := new(auth.TwoToken)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ExecRefreshTokenResult) GetSuccess() *auth.TwoToken {
	if !p.IsSetSuccess() {
		return ExecRefreshTokenResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ExecRefreshTokenResult) SetSuccess(x interface{}) {
	p.Success = x.(*auth.TwoToken)
}

func (p *ExecRefreshTokenResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ExecRefreshTokenResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetToken(ctx context.Context, Req *auth.UserId) (r *auth.TwoToken, err error) {
	var _args GetTokenArgs
	_args.Req = Req
	var _result GetTokenResult
	if err = p.c.Call(ctx, "GetToken", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ParseAccessToken(ctx context.Context, Req *auth.AccessToken) (r *auth.UserId, err error) {
	var _args ParseAccessTokenArgs
	_args.Req = Req
	var _result ParseAccessTokenResult
	if err = p.c.Call(ctx, "ParseAccessToken", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ParseRefreshToken(ctx context.Context, Req *auth.RefreshToken) (r *auth.UserId, err error) {
	var _args ParseRefreshTokenArgs
	_args.Req = Req
	var _result ParseRefreshTokenResult
	if err = p.c.Call(ctx, "ParseRefreshToken", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) VerifyToken(ctx context.Context, Req *auth.AccessToken) (r *auth.UserId, err error) {
	var _args VerifyTokenArgs
	_args.Req = Req
	var _result VerifyTokenResult
	if err = p.c.Call(ctx, "VerifyToken", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ExecRefreshToken(ctx context.Context, Req *auth.RefreshToken) (r *auth.TwoToken, err error) {
	var _args ExecRefreshTokenArgs
	_args.Req = Req
	var _result ExecRefreshTokenResult
	if err = p.c.Call(ctx, "ExecRefreshToken", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
