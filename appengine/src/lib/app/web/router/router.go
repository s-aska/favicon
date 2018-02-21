package router

import (
	"lib/app/web"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/nosurf"
	"google.golang.org/appengine"
)

// Router App router
type Router struct {
	router     *mux.Router
	authorizer func(r *http.Request) bool
}

// Get GETメソッド限定
func (i *Router) Get(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route {
	return i.HandleFunc("GET", path, f)
}

// Post POSTメソッド限定
func (i *Router) Post(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route {
	return i.HandleFunc("POST", path, f)
}

// Put PUTメソッド限定
func (i *Router) Put(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route {
	return i.HandleFunc("PUT", path, f)
}

// Delete DELETEメソッド限定
func (i *Router) Delete(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route {
	return i.HandleFunc("DELETE", path, f)
}

// HandleFunc ルーティング定義
func (i *Router) HandleFunc(method string, path string, f func(http.ResponseWriter, *http.Request)) *mux.Route {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if i.authorizer(r) {
			f(w, r)
		}
	}
	return i.router.HandleFunc(path, fn).Methods(method)
}

// InAuth 認証あり
func InAuth(router *mux.Router, authorizer func(r *http.Request) bool, f func(i *Router)) {
	f(&Router{
		router:     router,
		authorizer: authorizer,
	})
}

// InPublic 認証なし
func InPublic(router *mux.Router, f func(i *Router)) {
	f(&Router{
		router: router,
		authorizer: func(r *http.Request) bool {
			return true
		},
	})
}

// InWorkerOrCron 認証なし
func InWorkerOrCron(router *mux.Router, f func(i *Router)) {
	f(&Router{
		router: router,
		authorizer: func(r *http.Request) bool {
			return (r.Header.Get("X-AppEngine-Cron") == "true" && r.RemoteAddr == "0.1.0.1") || r.RemoteAddr == "0.1.0.2"
		},
	})
}

// WithCSRF CSRFチェック掛けるやつ
func WithCSRF(h http.Handler) http.Handler {
	n := nosurf.New(h)
	c := http.Cookie{}
	c.HttpOnly = true
	c.Path = "/"
	c.MaxAge = 365 * 24 * 60 * 60
	c.Secure = !appengine.IsDevAppServer()
	n.SetBaseCookie(c)
	return n
}

// WithContext http.Request から web.WebContext オブジェクトを取れるようにする
func WithContext(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		c := web.NewContext(w, r)

		// http://blogs.msdn.com/b/ie/archive/2008/07/02/ie8-security-part-v-comprehensive-protection.aspx
		w.Header().Set("X-Content-Type-Options", "nosniff")

		// http://blog.mozilla.com/security/2010/09/08/x-frame-options/
		w.Header().Set("X-Frame-Options", "DENY")

		h.ServeHTTP(w, c.Request)
	}
	return http.HandlerFunc(fn)
}
