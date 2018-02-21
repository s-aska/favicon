package web

import (
	"encoding/json"
	"fmt"
	"html/template"
	"lib/app/web/functions"
	"net/http"
	"strings"

	"github.com/justinas/nosurf"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// Error ...
func (c *Web) Error(e error) {
	log.Errorf(c.Context, "error:%v", e)
	http.Error(c.Writer, e.Error(), http.StatusInternalServerError)
}

// Forbidden ...
func (c *Web) Forbidden(e error) {
	http.Error(c.Writer, e.Error(), http.StatusForbidden)
}

// NotFound ...
func (c *Web) NotFound() {
	http.Error(c.Writer, "NotFound", http.StatusNotFound)
}

// Redirect ...
func (c *Web) Redirect(path string) {
	http.Redirect(c.Writer, c.Request, path, http.StatusFound)
}

// Success ...
func (c *Web) Success() {
	c.Writer.Header().Set("Content-Type", "application/json")
	fmt.Fprint(c.Writer, "{\"success\":true}")
}

// JSON ...
func (c *Web) JSON(i interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(c.Writer).Encode(i)
}

// ForURI ...
func (c *Web) ForURI(key string, value string) string {
	params := c.Request.URL.Query()
	params.Set(key, value)
	if key != "page" {
		params.Del("page")
	}
	return "?" + params.Encode()
}

// HTML ...
func (c *Web) HTML(v string) {
	funcMap := template.FuncMap{
		"Raw":         functions.Raw,
		"RelTime":     functions.RelTime,
		"RelDuration": functions.RelDuration,
		"YMDHMS":      functions.YMDHMS,
		"TimeParse":   functions.TimeParse,
		"Atoi":        functions.Atoi,
		"Commify":     functions.Commify,
		"Add":         functions.Add,
		"lnToBr":      functions.LnToBr,
		"Has":         functions.Has,
		"Join":        strings.Join,
		"HasPrefix":   strings.HasPrefix,
		"Contains":    strings.Contains,
		"ForURI":      c.ForURI,
		"ToDateTime":  functions.ToDateTime,
		"Set": func(k string, v interface{}) string {
			c.Data[k] = v
			return ""
		},
	}
	c.Data["req"] = c.Request
	c.Data["c"] = c
	c.Data["VersionID"] = appengine.VersionID(c.Context)
	csrfToken := nosurf.Token(c.Request)
	c.Data["csrfField"] = template.HTML(fmt.Sprintf(`<input type="hidden" name="csrf_token" value="%s">`, csrfToken))
	c.Data["csrfToken"] = csrfToken
	tmpl := template.Must(template.New("").Funcs(funcMap).ParseFiles("views/layout.html", v))
	err := tmpl.ExecuteTemplate(c.Writer, "layout", c.Data)
	if err != nil {
		c.Error(err)
	}
}
