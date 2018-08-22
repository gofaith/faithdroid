package faithdroid

import (
	"fmt"
	"github.com/StevenZack/tools/encodingToolkit"
	"github.com/StevenZack/tools/netToolkit"
	"github.com/StevenZack/tools/strToolkit"
	"os"
)

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
func cacheNetFile(url, cacheDir string, callback func(string)) {
	f := getrpath(cacheDir) + newToken()
	e := netToolkit.DownloadFile(url, f)
	if e != nil {
		fmt.Println(e)
		return
	}
	callback(f)
}
