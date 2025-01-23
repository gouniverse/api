# API 
<a href="https://gitpod.io/#https://github.com/gouniverse/api" style="float:right:"><img src="https://gitpod.io/button/open-in-gitpod.svg" alt="Open in Gitpod" loading="lazy"></a>

[![Open in Codeanywhere](https://codeanywhere.com/img/open-in-codeanywhere-btn.svg)](https://app.codeanywhere.com/#https://github.com/gouniverse/api)

[![Tests Status](https://github.com/gouniverse/api/workflows/tests/badge.svg)](https://github.com/gouniverse/api/workflows/tests/badge.svg)
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
api.RespondWithStatusCode(w, r, api.Error("endpoint not found"), http.StatusNotFound)

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

## What is the Difference Between Unauthenticated, Unauthorized and Forbidden?

1. Unauthenticated: This is the most fundamental level of access control. If a user or system cannot be verified as a legitimate entity, it will be denied access to any resource.
2. Unauthorized: Once a user or system has been authenticated, it must also be authorized to access a particular resource. If it lacks the necessary permissions, it will be denied access.
3. Forbidden: This is the highest level of access control. Even if a user or system is both authenticated and authorized, it may still be denied access to a resource if there are specific conditions or rules in place that prohibit access.

In summary:

Unauthenticated -> Unauthorized -> Forbidden

This order reflects the increasing levels of verification and permission required to access a resource.

However in REST context, these are represented with ambiguous status codes:

- Unauthenticated is represented by 401 Unauthorized: This indicates that the client needs to authenticate itself with the server. Typically, this involves providing credentials like a username and password.

- Unauthorized & Forbidden are both represented by 403 Forbidden: This means the server understood the request but refuses to fulfill it. This could be due to a lack of necessary permissions or a policy restriction.
