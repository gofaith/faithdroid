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
	// cacheNetFile("https://www.baidu.com/img/bd_logo1.png", "/home/asd/Documents/cacheDir", func(string) {})
	// t.Log(url2cachePath("https://www.baidu.com/img/bd_logo1.png"))
	str := `{"VID":"95cc2b10c1eda70d4950e474efa66cca","Bt":{"VID":"09a4f930f0c4523d684a72aa573ca651","ClassName":"Button"},"Tv":{"VID":"100b6dfa3d28e5a5d55a1aba7de127ec","ClassName":"TextView"}}`
	mh := MyHolder{}
	unJson(str, &mh)
	t.Log(mh.Bt == nil)
}
