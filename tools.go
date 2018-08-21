package faithdroid

import (
	"github.com/StevenZack/tools/strToolkit"
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
