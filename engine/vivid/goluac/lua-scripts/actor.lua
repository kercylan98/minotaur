local actor = {}

---on_receive 传入一个函数，用来接收并处理消息
---@param handler function 消息处理函数，它接收一个参数，即消息内容
function actor.on_receive(handler)
    assert(type(handler) == "function", "handler must be a function, got: " .. type(handler))
    actor.on_receive = handler
end

return actor