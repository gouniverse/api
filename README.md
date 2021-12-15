# API

Concise library for generating API responses in JSON format


## Shortcut methods

- Error(msg string)
- ErrorWithData(msg string, data []interface)
- Forbidden(msg string)
- Success(msg string)
- SuccessWithData(msg string, data []interface)
- Unauthenticated(msg string)
- Unauthorized(msg string)


# Example

```
// return error response
api.Respond(w, r, api.Error("api key is required"))

// return success response with data payload
api.Respond(w, r, api.SuccessWithData("success", map[string]interface{}{
  "key1": "value1",
  "key2": "value1",
  "key3": "value1",
}))
```

## Custom responses

```
api.Respond(w, r, api.Response{
  Status:  "custom_status",
  Message: "message",
  Data: map[string]interface{}{
    "key1": "value1",
    "key2": "value1",
    "key3": "value1",
  },
})
```
