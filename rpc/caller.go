package rpc

type (
	// NonBlockingCaller 非阻塞 RPC 调用器，用于在不关心调用结果的情况下进行异步调用
	//   - 该调用器不一定能够成功执行
	NonBlockingCaller func(param any) error

	// BlockingCaller 阻塞 RPC 调用器，用于在不关心调用结果的情况下进行同步调用，除了与 NonBlockingCaller 的区别外，该函数可以保证调用成功
	BlockingCaller func(param any) error

	// NonBlockingRequestCaller 非阻塞 RPC 请求调用器，该调用器会在调用成功后将结果通过 Reader 返回到回调函数中
	NonBlockingRequestCaller func(param any, callback func(reader Reader)) error

	// BlockingRequestCaller 阻塞 RPC 请求调用器，该调用器会在调用成功后将结果通过 Reader 返回到调用方
	BlockingRequestCaller func(param any) (Reader, error)
)
