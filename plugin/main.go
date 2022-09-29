package main

import (
	"github.com/beeedge/device-plugin-framework/shared"
	"github.com/hashicorp/go-plugin"
)

// Here is a real implementation of device-plugin.
type Converter struct {
}

// ConvertReportMessage2Devices converts data report request to protocol that device understands for each device of this device model,
func (c *Converter) ConvertReportMessage2Devices(modelId, featureId string) ([]string, error) {
	// TODO: concrete implement
	return []string{"Have a good try!!!"}, nil
}

// ConvertIssueMessage2Device converts issue request to protocol that device understands, which has four return parameters:
// 1. inputMessages: device issue protocols for each of command input param.
// 2. outputMessages: device data report protocols for each of command output param.
// 3. issueTopic: device issue MQTT topic for input params.
// 4. issueResponseTopic: device issue response MQ topic for output params.
func (c *Converter) ConvertIssueMessage2Device(deviceId, modelId, featureId string, values map[string]string) ([]string, []string, string, string, error) {
	// TODO: concrete implement
	return nil, nil, "", "", nil
}

// ConvertDeviceMessages2MQFormat receives device command issue responses and converts it to RabbitMQ normative format.
func (c *Converter) ConvertDeviceMessages2MQFormat(messages []string) (string, []byte, error) {
	// TODO: concrete implement
	return "", nil, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins: map[string]plugin.Plugin{
			"converter": &shared.ConverterPlugin{Impl: &Converter{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
