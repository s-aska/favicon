package web

import (
	"net/http"
	"strings"

	"context"

	"google.golang.org/appengine"
)

// Web Web
type Web struct {
	Context context.Context
	Request *http.Request
	Writer  http.ResponseWriter
	Data    map[string]interface{}
}

// NewContext NewContext
func NewContext(w http.ResponseWriter, r *http.Request) *Web {
	ctx := appengine.NewContext(r)
	web := &Web{
		Context: ctx,
		Writer:  w,
		Data:    map[string]interface{}{},
	}
	web.Request = contextSet(r, "Web", web)
	return web
}

// Context Context
func Context(r *http.Request) *Web {
	v := contextGet(r, "Web")
	if v == nil {
		return nil
	}
	c, ok := v.(*Web)
	if !ok {
		return nil
	}
	c.Request = r
	return c
}

// IsDev Local Machine
func (c *Web) IsDev() bool {
	return appengine.IsDevAppServer()
}

// IsBeta Beta
func (c *Web) IsBeta() bool {
	return strings.HasSuffix(appengine.AppID(c.Context), "-beta")
}

// Scheme Scheme
func (c *Web) Scheme() string {
	if appengine.IsDevAppServer() {
		return "http"
	}
	return "https"
}

// RequestURI RequestURI
func (c *Web) RequestURI() string {
	return c.Scheme() + "://" + c.Request.Host + c.Request.RequestURI
}
