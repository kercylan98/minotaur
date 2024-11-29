local errors = {}

function errors.new(message)
    return {
        type = "error",
        message = message
    }
end

function errors.is_error(value)
    return type(value) == "table" and value.type == "error"
end

function errors.to_string(value)
    return value.message
end

function errors.throw(message)
    error(errors.new(message))
end

return errors