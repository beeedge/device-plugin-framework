package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/device-plugin-framework/examples/bidirectional/shared"
	"github.com/hashicorp/go-plugin"
)

// Here is a real implementation of KV that writes to a local file with
// the key name and the contents are the value of the key.
type Converter struct {
}

type data struct {
	Value int64
}

func (k *Converter) Put(key string, value int64, a shared.AddHelper) error {
	v, _ := k.Get(key)

	r, err := a.Sum(v, value)
	if err != nil {
		return err
	}

	buf, err := json.Marshal(&data{r})
	if err != nil {
		return err
	}

	return ioutil.WriteFile("kv_"+key, buf, 0644)
}

func (k *Converter) Get(key string) (int64, error) {
	dataRaw, err := ioutil.ReadFile("kv_" + key)
	if err != nil {
		return 0, err
	}

	data := &data{}
	err = json.Unmarshal(dataRaw, data)
	if err != nil {
		return 0, err
	}

	return data.Value, nil
}

func (k *Converter) ConvertReportMessage(deviceId, modelId, featureId, msgId string) ([]string, error) {
	// TODO
	return nil, nil
}

func (k *Converter) ConvertIssueMessage(deviceId, modelId, featureId, msgId string, values map[string]string) ([]string, []string, error) {
	// TODO
	return nil, nil, nil
}

func (k *Converter) RevConvertMessages(messages []string) (string, []string, error) {
	// TODO
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
