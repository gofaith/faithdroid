package faithdroid

import (
	"testing"
)

func Test_cacheNetFile(t *testing.T) {
	cacheNetFile("https://www.baidu.com/img/bd_logo1.png", "/home/asd/Documents/cacheDir", func(string) {})
	// t.Log(url2cachePath("https://www.baidu.com/img/bd_logo1.png"))
}
