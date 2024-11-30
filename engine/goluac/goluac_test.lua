local actor = require("actor")
local json = require("json")
local errors = require("errors")
local goluac_test_mod = require("goluac_test_mod")

goluac_test_mod.check()

actor.on_receive(function(ctx)
    local packet = ctx.message()
    router[packet.name](ctx, packet.data)
end)

router = {
    ["test"] = function(ctx, message)
        print("receive test message: " .. json.encode(message))
        ctx.reply("reply", message)
    end,
}
