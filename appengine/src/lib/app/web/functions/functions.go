package functions

import (
	"html"
	"html/template"
	"strconv"
	"strings"
	"time"

	"fmt"

	humanize "github.com/dustin/go-humanize"
)

// Raw ...
func Raw(s string) template.HTML {
	return template.HTML(s)
}

// Add ...
func Add(a int, b int) int {
	return a + b
}

// Has ...
func Has(array []string, target string) bool {
	for _, a := range array {
		if a == target {
			return true
		}
	}
	return false
}

// TimeParse ...
func TimeParse(s string) time.Time {
	t, err := time.Parse("2006-01-02T15:04:05-07:00", s)
	if err != nil {
		return time.Unix(0, 0)
	}
	return t
}

// Atoi ...
func Atoi(s string) int {
	if strings.HasSuffix(s, ".00") {
		s = strings.Replace(s, ".00", "", 1)
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

// YMDHMS ...
func YMDHMS(t time.Time) string {
	return t.Format("2006/01/02 15:04:05")
}

// RelTime ...
func RelTime(t time.Time) string {
	d := t.Unix() - time.Now().Unix()
	if d > 86400 {
		return fmt.Sprintf("あと%d日", d/86400)
	} else if d > 3600 {
		return fmt.Sprintf("あと%d時間", d/3600)
	} else if d > 60 {
		return fmt.Sprintf("あと%d分", d/60)
	}
	return fmt.Sprintf("あと%d秒", d)
}

// RelDuration ...
func RelDuration(d int) string {
	if d > 86400 {
		return fmt.Sprintf("%d日", d/86400)
	} else if d > 3600 {
		return fmt.Sprintf("%d時間", d/3600)
	} else if d > 60 {
		return fmt.Sprintf("%d分", d/60)
	}
	return fmt.Sprintf("%d秒", d)
}

// Commify ...
func Commify(i int) string {
	return humanize.Comma(int64(i))
}

var lnToBr = strings.NewReplacer("\n", "<br>")

// LnToBr ...
func LnToBr(s string) template.HTML {
	return template.HTML(lnToBr.Replace(html.EscapeString(s)))
}

// ToDateTime ...
func ToDateTime(epoch int) string {
	return time.Unix(int64(epoch), 0).In(time.Local).Format("2006/01/02 15:04")
}
