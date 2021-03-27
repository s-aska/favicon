package functions

import (
	"fmt"
	"html"
	"html/template"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	humanize "github.com/dustin/go-humanize"
)

var lnToBr = strings.NewReplacer("\n", "<br>")

var rAutolink = regexp.MustCompile(`((?:https?:\/\/)[\w\.-]+(?:\.[\w\.-]+)*(?:\:[0-9]+)?[\w\-\._~:/?#[\]@!\$&'\(\)\*\+,;=\%.]+)|(.)`)

// AutoLink AutoLink
func AutoLink(in string) template.HTML {
	in = strings.TrimSpace(in)
	out := rAutolink.ReplaceAllStringFunc(in, func(s string) string {
		matches := rAutolink.FindStringSubmatch(s)
		if matches[1] != "" {
			linkURL, err := url.Parse(matches[1])
			if err != nil {
				return html.EscapeString(matches[1])
			}
			target := ""
			return fmt.Sprintf(`<a href="%s" target="%s">%s</a>`, html.EscapeString(linkURL.String()), target, html.EscapeString(linkURL.String()))
		}
		return html.EscapeString(matches[2])
	})
	return template.HTML(lnToBr.Replace(out))
}

// LnToBr ...
func LnToBr(s string) template.HTML {
	return template.HTML(lnToBr.Replace(html.EscapeString(strings.TrimSpace(s))))
}

// Commify ...
func Commify(i interface{}) string {
	switch x := i.(type) {
	case int:
		return humanize.Comma(int64(x))
	case int64:
		return humanize.Comma(x)
	}
	return ""
}

// ForURI ...
func ForURI(r *http.Request, k, v string) string {
	q := r.URL.Query()
	q.Set(k, v)
	return "?" + q.Encode()
}

// JST Japan
var jst = time.FixedZone("Asia/Tokyo", 9*60*60)

var funcMap = template.FuncMap{
	"lnToBr":    AutoLink,
	"ForURI":    ForURI,
	"Commify":   Commify,
	"Join":      strings.Join,
	"Split":     strings.Split,
	"HasPrefix": strings.HasPrefix,
	"HasSuffix": strings.HasSuffix,
	"TrimRight": strings.TrimRight,
	"RemoveSpace": func(s string) string {
		return strings.Replace(s, " ", "", -1)
	},
	"Date": func(t time.Time) string {
		if t.IsZero() {
			return ""
		}
		return t.In(jst).Format("2006.01.02")
	},
	"DateTime": func(t time.Time) string {
		if t.IsZero() {
			return ""
		}
		return t.In(jst).Format("2006.01.02 15:04")
	},
	"DateTimeFromEpoch": func(epoch int64) string {
		if epoch == 0 {
			return ""
		}
		return time.Unix(epoch/1000000000, 0).In(jst).Format("2006.01.02 15:04")
	},
	"Format": func(f string, t time.Time) string {
		if t.IsZero() {
			return ""
		}
		return t.In(jst).Format(f)
	},
	"Add": func(src int, plus int) int {
		return src + plus
	},
	"Odd": func(a int) bool {
		return a%2 != 0
	},
	"ContainsAny": func(list []int64, i int64) bool {
		for _, row := range list {
			if row == i {
				return true
			}
		}
		return false
	},
	"uri": func(s string) string {
		return url.QueryEscape(s)
	},
	"raw": func(h string) template.HTML {
		return template.HTML(h)
	},
	"NowDate": func() string {
		return time.Now().In(jst).Format("2006年1月2日")
	},
	"NowDateForInputForm": func() string {
		return time.Now().In(jst).Format("2006-01-02")
	},
	"DateForPoem": func(t time.Time) string {
		if t.IsZero() {
			return ""
		}
		return t.In(jst).Format("2006年1月2日")
	},
}

// NewFuncMap ...
func NewFuncMap() template.FuncMap {
	return funcMap
}
