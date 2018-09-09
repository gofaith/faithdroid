package faithdroid

import (
	"testing"
)

type Name struct {
}

func (*Name) NewView(viewName string, vID string) {
}
func (*Name) ViewSetAttr(vID string, attr string, value string) {
}
func (*Name) ViewGetAttr(vID string, attr string) string {
	return ""
}
func (*Name) ShowOnRootView(vID string) {
}
func (*Name) RunOnUIThread(fnID string) {
}
func (*Name) GetPkg() string {
	return ""
}
func Test_cacheNetFile(t *testing.T) {
	// CacheNetFile("https://www.baidu.com/img/bd_logo1.png", "/home/asd/Documents/cacheDir", func(string) {})
	// t.Log(Url2cachePath("https://www.baidu.com/img/bd_logo1.png"))
	te := TouchEvent{}
	UnJson(`{"Action":"Up","X":365.9765625,"Y":720.92578125}`, &te)
	t.Log(SPrintf(te))
}
