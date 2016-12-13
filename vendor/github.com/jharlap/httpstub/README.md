Easily configure http test servers for stubbing external dependencies, for testing in Go.

API doc: http://godoc.org/github.com/jharlap/httpstub

```go
ts := httpstub.New().WithDefaultContentType(ctJSON)
defer ts.Close()

// the default status for name requests will be 204 no content, this will match PUT, DELETE, ...
nameEndpoint := ts.Path("/user/*/name").WithStatus(http.StatusNoContent)

// GET overrides the status and body
nameEndpoint.WithMethod("GET").WithBody(`{"id":"a1","name":"Alice"}`).WithStatus(http.StatusOK)

// endpoint-specific content type
ts.Path("/user/*/xml").WithContentType(ctXML).WithBody(`<user id="a1"><name>Alice</name></user>`)

// note that paths are matched first to last, so the longest paths must appear first
ts.Path("/user").WithBody(`{"id":"a1","name":"Alice","gender":"f"}`)

client := mine{a3rdPartyServerURL: ts.URL}
client.DoSomething() // that makes HTTP requests to the 3rd party server
```

Note that the `With...` methods are designed for chaining, and mutate the object they are invoked on.
