package faithdroid

import (
	"fmt"
	"github.com/StevenZack/tools/encodingToolkit"
	"github.com/StevenZack/tools/netToolkit"
	"github.com/StevenZack/tools/strToolkit"
	"os"
	"strings"
)

/*
func twoInt(ps ...interface{}) (int, int, bool) {
	if len(ps) != 2 {
		return -1, -1, false
	}
	v1, ok1 := ps[0].(int)
	v2, ok2 := ps[1].(int)
	return v1, v2, ok1 && ok2
}
func oneString(ps ...interface{}) (string, bool) {
	if len(ps) == 1 {
		if v, ok := ps[0].(string); ok {
			return v, true
		}
	}
	return "", false
}
func multiViews(ps ...interface{}) ([]IView, bool) {
	if len(ps) <= 0 {
		return nil, false
	}
	var ws []IView
	if _, ok := ps[0].(IView); !ok {
		return nil, false
	}
	for _, v := range ps {
		if widget, ok := v.(IView); ok {
			ws = append(ws, widget)
		}
	}
	return ws, true
}
*/
func newToken() string {
	return strToolkit.NewToken()
}
func jsonArray(i interface{}) string {
	return encodingToolkit.JsonArray(i)
}
func getrpath(s string) string {
	if len(s) < 1 {
		return ""
	}
	sp := string(os.PathSeparator)
	if s[len(s)-1:] == sp {
		return s
	}
	return s + sp
}
func downloadNetFile(url, downloadDir string, callback func(string)) {
	f := getrpath(downloadDir) + newToken()
	e := netToolkit.DownloadFile(url, f)
	if e != nil {
		fmt.Println(e)
		return
	}
	callback(f)
}
func cacheNetFile(url, cacheDir string, callback func(string)) {
	f := getrpath(cacheDir) + url2cachePath(url)
	if _, e := os.Stat(f); e != nil {
		e = netToolkit.DownloadFile(url, f)
		if e != nil {
			fmt.Println(e)
			return
		}
	}
	callback(f)
}
func url2cachePath(url string) string {
	rUrl := getRealUrl(url)
	s := strings.Replace(rUrl, "://", "/", -1)
	if strToolkit.EndsWith(s, "/") {
		return s[:len(s)-1]
	}
	return s
}
func getRealUrl(url string) string {
	for i := 0; i < len(url); i++ {
		item := url[i : i+1]
		if item == "?" || item == "#" {
			return url[:i]
		}
	}
	return url
}
func sPrintf(a interface{}) string {
	return fmt.Sprintf("%v", a)
}
func xPrintf(a int) string {
	return fmt.Sprintf("%x", a)
}
