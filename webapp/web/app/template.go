package app

import (
	"html/template"
	"io"
	"os"
	"path"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/s-aska/favicon/webapp/web/appengine"
	"github.com/s-aska/favicon/webapp/web/functions"
)

func scheme() string {
	if appengine.IsDevAppServer() {
		return "http"
	}
	return "https"
}

// Renderer ...
type Renderer struct {
	dir string
}

// Render ...
func (r *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	d, ok := data.(map[string]interface{})
	if ok {
		req := c.Request()
		d["c"] = c
		d["req"] = req
		d["VersionID"] = appengine.VersionID
		d["IsDev"] = appengine.IsDevAppServer()
		d["IsProd"] = appengine.IsProd()
	}
	t := template.New("")
	funcMap := functions.NewFuncMap()
	funcMap["root"] = func() interface{} {
		return data
	}
	funcMap["Set"] = func(key string, value interface{}) interface{} {
		d[key] = value
		return nil
	}
	t = t.Funcs(funcMap)
	t = template.Must(t.ParseFiles(path.Join(r.dir, "layout.html"), path.Join(r.dir, name)))
	t = template.Must(t.ParseGlob(path.Join(r.dir, "/includes/*.html")))
	err := t.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Errorf("[Render] data:%#+v error:%v", data, err)
	}
	return err
}

func registerTemplates(e *echo.Echo) {
	sv := os.Getenv("GAE_SERVICE")
	if sv == "default" {
		sv = "app"
	}
	views := sv + "/views/"
	r := &Renderer{
		dir: views,
	}
	e.Renderer = r
}
