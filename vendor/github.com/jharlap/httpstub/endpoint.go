package httpstub

import (
	"net/http"
	"time"
)

// An Endpoint is added to a stub server with server.Path(), and implements a handler that will be matched against a path prefix.
type Endpoint struct {
	path        string
	status      int
	contentType string
	body        []byte
	delay       time.Duration

	forMethod map[string]*Endpoint
}

// WithMethod creates a method-specific endpoint within a path.
func (e *Endpoint) WithMethod(m string) *Endpoint {
	me := &Endpoint{
		status:      e.status,
		contentType: e.contentType,
		body:        e.body,
		delay:       e.delay,
	}

	if e.forMethod == nil {
		e.forMethod = make(map[string]*Endpoint)
	}
	e.forMethod[m] = me
	return me
}

// WithStatus sets the response status code for the endpoint.
func (e *Endpoint) WithStatus(s int) *Endpoint {
	e.status = s
	return e
}

// WithContentType overrides the server's default content type for the endpoint.
func (e *Endpoint) WithContentType(t string) *Endpoint {
	e.contentType = t
	return e
}

// WithBody sets the body the endpoint should return.
func (e *Endpoint) WithBody(b string) *Endpoint {
	e.body = []byte(b)
	return e
}

// WithDelay sets a delay before the endpoint responds
func (e *Endpoint) WithDelay(d time.Duration) *Endpoint {
	e.delay = d
	return e
}

// ServeHTTP lets Endpoint implement http.Handler.
func (e *Endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// use the method-specific handler if available
	if me, ok := e.forMethod[r.Method]; ok {
		me.ServeHTTP(w, r)
		return
	}

	if e.delay > 0 {
		time.Sleep(e.delay)
	}

	if len(e.contentType) > 0 {
		w.Header().Set("Content-Type", e.contentType)
	}

	if e.status > 0 {
		w.WriteHeader(e.status)
	}

	if len(e.body) > 0 {
		w.Write(e.body)
	}
}
