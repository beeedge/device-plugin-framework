// Package shared contains shared data between the host and plugins.
package shared

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/beeedge/device-plugin-framework/proto"
	"github.com/hashicorp/go-plugin"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "DEVICE_PLUGIN",
	MagicCookieValue: "BeeThings",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"converter": &ConverterPlugin{},
}

// Converter is the interface that we're exposing as a plugin.
type Converter interface {
	// ConvertIssueMessage2Device converts issue request to protocol that device understands, which has four return parameters:
	// 1. inputMessages: device issue protocols for each of command input param.
	// 2. outputMessages: device data report protocols for each of command output param.
	// 3. issueTopic: device issue MQTT topic for input params.
	// 4. issueResponseTopic: device issue response MQ topic for output params.
	ConvertIssueMessage2Device(deviceId, modelId, featureId string, values map[string]string, convertedDeviceFeatureMap string) ([]string, []string, string, string, error)
	// ConvertDeviceMessages2MQFormat receives device command issue responses and converts it to RabbitMQ normative format.
	ConvertDeviceMessages2MQFormat(messages []string, convertedDeviceFeatureMap string) (string, []byte, error)
}

// This is the implementation of plugin.Plugin so we can serve/consume this.
// We also implement GRPCPlugin so that this plugin can be served over
// gRPC.
type ConverterPlugin struct {
	plugin.NetRPCUnsupportedPlugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl Converter
}

func (p *ConverterPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterConverterServer(s, &GRPCServer{
		Impl:   p.Impl,
		broker: broker,
	})
	return nil
}

func (p *ConverterPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{
		client: proto.NewConverterClient(c),
		broker: broker,
	}, nil
}

var _ plugin.GRPCPlugin = &ConverterPlugin{}
