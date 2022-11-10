device-plugin-framework
=======================

`device-plugin-framework` provides IoT device protocol parse framework for [beethings](https://github.com/beeedge/beethings).

## Framework

`device-plugin-framework` bases on [go-plugin](https://github.com/hashicorp/go-plugin/tree/master/examples/bidirectional) and we only care about one interfaces with three methods as below:

```go
// Converter is the interface that we're exposing as a plugin.
type Converter interface {
    // ConvertIssueMessage2Device converts issue request to protocol that device understands, which has four return parameters:
    // 1. inputMessages: device issue protocols for each of command input param.
    // 2. outputMessages: device data report protocols for each of command output param.
    // 3. issueTopic: device issue MQTT topic for input params.
    // 4. issueResponseTopic: device issue response MQ topic for output params.
    ConvertIssueMessage2Device(deviceId, modelId, featureId string, values map[string]string) ([]string, []string, string, string, error)
    // ConvertDeviceMessages2MQFormat receives device command issue responses and converts it to RabbitMQ normative format.
    ConvertDeviceMessages2MQFormat(messages []string) (string, []byte, error)
}
```

We can add on one device plugin by implementing above three methods in `plugin/main.go`. 

## Usage

```sh
# This builds the main CLI(main.go)
$ go build -o converter

# This builds the plugin written in Go(plugin/main.go)
$ go build -o converter-go-grpc

# This tells the Converter binary to use the "converter-go-grpc" binary
$ export DEVICE_PLUGIN="./converter-go-grpc"

# gRPC calls
$ ./converter
[Have a good try!!!]
```

### Updating the Protocol

If you update the protocol buffers file, you can regenerate the file using the following command from this directory. You do not need to run this if you're just trying the example.

For Go:

```bash
$ protoc -I proto/ proto/converter.proto --go_out=plugins=grpc:proto --go-grpc_out=require_unimplemented_servers=false:proto
```

## Refs

* [bidirectional go-plugin](https://github.com/hashicorp/go-plugin/tree/master/examples/bidirectional)
* [Alibaba Paho-MQTT Go](https://help.aliyun.com/document_detail/146503.html#section-lk1-zyq-byp)
* [IoT Device SDK](https://support.huaweicloud.com/sdkreference-iothub/iot_02_0178.html)