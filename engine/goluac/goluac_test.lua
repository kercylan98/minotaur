local actor = require("actor")
local json = require("json")
local errors = require("errors")

actor.on_receive(function(message)
    print("lua local receive message: " .. message)
end)

actor.on_receive("test receive callback")

local future = actor.future_ask("monitor://localhost/user/1", { 
    name = "John", 
    age = 30
})

local result = future.result()
if errors.is_error(result) then
    print("lua receive future error: " .. errors.to_string(result))
else
    print("lua receive future result: " .. result)
    print("lua parse to json got name: " ..json.decode(result).name)
end 