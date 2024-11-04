package vivid

type Topic = string

const (
	// TransportTopic 来自 Shared 传输异常的订阅主题，在 Actor 订阅后可通过 *OnTransportError 类型断言获取到该消息
	//  - 该订阅仅在本地生效
	TransportTopic Topic = "_shared_transport"
)
