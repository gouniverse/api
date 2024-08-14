# API <a href="https://gitpod.io/#https://github.com/gouniverse/api" style="float:right:"><img src="https://gitpod.io/button/open-in-gitpod.svg" alt="Open in Gitpod" loading="lazy"></a>

[![Tests Status](https://github.com/gouniverse/api/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/gouniverse/api/actions/workflows/tests.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/gouniverse/api)](https://goreportcard.com/report/github.com/gouniverse/api)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/gouniverse/api)](https://pkg.go.dev/github.com/gouniverse/api)

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

// return error response with HTTP status sode
api.RespondWithStatus(w, r, api.Error("endpoint not found"), http.StatusNotFound)

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
