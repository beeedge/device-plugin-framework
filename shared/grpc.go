package shared

import (
	"github.com/beeedge/device-plugin-framework/proto"
	plugin "github.com/hashicorp/go-plugin"
	"golang.org/x/net/context"
)

// GRPCClient is an implementation of Converter that talks over RPC.
type GRPCClient struct {
	broker *plugin.GRPCBroker
	client proto.ConverterClient
}

func (m *GRPCClient) ConvertIssueMessage2Device(deviceId, modelId, featureId string, values map[string]string, convertedDeviceFeatureMap string) ([]string, []string, string, string, error) {
	resp, err := m.client.ConvertIssueMessage2Device(context.Background(), &proto.GetDeviceIssueRequest{
		DeviceId:                  deviceId,
		ModelId:                   modelId,
		FeatureId:                 featureId,
		Values:                    values,
		ConvertedDeviceFeatureMap: convertedDeviceFeatureMap,
	})
	if err != nil {
		return nil, nil, "", "", err
	}

	return resp.InputMessages, resp.OutputMessages, resp.IssueTopic, resp.IssueResponseTopic, nil
}

func (m *GRPCClient) ConvertDeviceMessages2MQFormat(messages []string, convertedDeviceFeatureMap string) (string, []byte, error) {
	resp, err := m.client.ConvertDeviceMessages2MQFormat(context.Background(), &proto.GetMQFormatRequest{
		Messages:                  messages,
		ConvertedDeviceFeatureMap: convertedDeviceFeatureMap,
	})
	if err != nil {
		return "", nil, err
	}

	return resp.RoutingKey, resp.RabbitMQMsgBody, nil
}

// Here is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl Converter

	broker *plugin.GRPCBroker
}

func (m *GRPCServer) ConvertIssueMessage2Device(ctx context.Context, req *proto.GetDeviceIssueRequest) (*proto.GetDeviceIssueResponse, error) {
	inputMessages, outputMessages, issueTopic, issueResponseTopic, err := m.Impl.ConvertIssueMessage2Device(req.DeviceId, req.ModelId, req.FeatureId, req.Values, req.ConvertedDeviceFeatureMap)
	return &proto.GetDeviceIssueResponse{
		InputMessages:      inputMessages,
		OutputMessages:     outputMessages,
		IssueTopic:         issueTopic,
		IssueResponseTopic: issueResponseTopic,
	}, err
}

func (m *GRPCServer) ConvertDeviceMessages2MQFormat(ctx context.Context, req *proto.GetMQFormatRequest) (*proto.GetMQFormatResponse, error) {
	routingKey, rabbitMQMsgBody, err := m.Impl.ConvertDeviceMessages2MQFormat(req.Messages, req.ConvertedDeviceFeatureMap)
	return &proto.GetMQFormatResponse{
		RoutingKey:      routingKey,
		RabbitMQMsgBody: rabbitMQMsgBody,
	}, err
}
