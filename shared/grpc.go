package shared

import (
	"github.com/device-plugin-framework/examples/bidirectional/proto"
	hclog "github.com/hashicorp/go-hclog"
	plugin "github.com/hashicorp/go-plugin"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// GRPCClient is an implementation of KV that talks over RPC.
type GRPCClient struct {
	broker *plugin.GRPCBroker
	client proto.ConverterClient
}

func (m *GRPCClient) Put(key string, value int64, a AddHelper) error {
	addHelperServer := &GRPCAddHelperServer{Impl: a}

	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		s = grpc.NewServer(opts...)
		proto.RegisterAddHelperServer(s, addHelperServer)

		return s
	}

	brokerID := m.broker.NextId()
	go m.broker.AcceptAndServe(brokerID, serverFunc)

	_, err := m.client.Put(context.Background(), &proto.PutRequest{
		AddServer: brokerID,
		Key:       key,
		Value:     value,
	})

	s.Stop()
	return err
}

func (m *GRPCClient) Get(key string) (int64, error) {
	resp, err := m.client.Get(context.Background(), &proto.GetRequest{
		Key: key,
	})
	if err != nil {
		return 0, err
	}

	return resp.Value, nil
}

func (m *GRPCClient) ConvertReportMessage(deviceId, modelId, featureId, msgId string) ([]string, error) {
	resp, err := m.client.ConvertReportMessage(context.Background(), &proto.GetReportRequest{
		DeviceId:  deviceId,
		ModelId:   modelId,
		FeatureId: featureId,
		MsgId:     msgId,
	})
	if err != nil {
		return nil, err
	}

	return resp.Messages, nil
}

func (m *GRPCClient) ConvertIssueMessage(deviceId, modelId, featureId, msgId string, values map[string]string) ([]string, []string, error) {
	resp, err := m.client.ConvertIssueMessage(context.Background(), &proto.GetIssueRequest{
		DeviceId:  deviceId,
		ModelId:   modelId,
		FeatureId: featureId,
		MsgId:     msgId,
		Values:    values,
	})
	if err != nil {
		return nil, nil, err
	}

	return resp.ReportMessages, resp.IssueMessages, nil
}

func (m *GRPCClient) RevConvertMessages(messages []string) (string, []string, error) {
	resp, err := m.client.RevConvertMessages(context.Background(), &proto.GetRevRequest{
		Messages: messages,
	})
	if err != nil {
		return "", nil, err
	}

	return resp.RoutingKey, resp.Messages, nil
}

// Here is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl Converter

	broker *plugin.GRPCBroker
}

func (m *GRPCServer) Put(ctx context.Context, req *proto.PutRequest) (*proto.Empty, error) {
	conn, err := m.broker.Dial(req.AddServer)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	a := &GRPCAddHelperClient{proto.NewAddHelperClient(conn)}
	return &proto.Empty{}, m.Impl.Put(req.Key, req.Value, a)
}

func (m *GRPCServer) Get(ctx context.Context, req *proto.GetRequest) (*proto.GetResponse, error) {
	v, err := m.Impl.Get(req.Key)
	return &proto.GetResponse{Value: v}, err
}

func (m *GRPCServer) ConvertReportMessage(ctx context.Context, req *proto.GetReportRequest) (*proto.GetReportResponse, error) {
	v, err := m.Impl.ConvertReportMessage(req.DeviceId, req.ModelId, req.FeatureId, req.MsgId)
	return &proto.GetReportResponse{Messages: v}, err
}

func (m *GRPCServer) ConvertIssueMessage(ctx context.Context, req *proto.GetIssueRequest) (*proto.GetIssueResponse, error) {
	v1, v2, err := m.Impl.ConvertIssueMessage(req.DeviceId, req.ModelId, req.FeatureId, req.MsgId, req.Values)
	return &proto.GetIssueResponse{ReportMessages: v1, IssueMessages: v2}, err
}

func (m *GRPCServer) RevConvertMessages(ctx context.Context, req *proto.GetRevRequest) (*proto.GetRevResponse, error) {
	v1, v2, err := m.Impl.RevConvertMessages(req.Messages)
	return &proto.GetRevResponse{RoutingKey: v1, Messages: v2}, err
}

// GRPCClient is an implementation of KV that talks over RPC.
type GRPCAddHelperClient struct{ client proto.AddHelperClient }

func (m *GRPCAddHelperClient) Sum(a, b int64) (int64, error) {
	resp, err := m.client.Sum(context.Background(), &proto.SumRequest{
		A: a,
		B: b,
	})
	if err != nil {
		hclog.Default().Info("add.Sum", "client", "start", "err", err)
		return 0, err
	}
	return resp.R, err
}

// Here is the gRPC server that GRPCClient talks to.
type GRPCAddHelperServer struct {
	// This is the real implementation
	Impl AddHelper
}

func (m *GRPCAddHelperServer) Sum(ctx context.Context, req *proto.SumRequest) (resp *proto.SumResponse, err error) {
	r, err := m.Impl.Sum(req.A, req.B)
	if err != nil {
		return nil, err
	}
	return &proto.SumResponse{R: r}, err
}
