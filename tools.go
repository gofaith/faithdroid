package faithdroid

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func NewToken() string {
	ct := time.Now().UnixNano()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(ct, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	return token
}
func JsonArray(i interface{}) string {
	b, e := json.Marshal(i)
	if e != nil {
		return "[]"
	}
	return string(b)
}
func JsonObject(i interface{}) string {
	b, e := json.Marshal(i)
	if e != nil {
		return "{}"
	}
	return string(b)
}
func UnJson(str string, v interface{}) {
	e := json.Unmarshal([]byte(str), v)
	if e != nil {
		fmt.Println("UnJson() err:", e)
	}
}
func Getrpath(s string) string {
	if len(s) < 1 {
		return ""
	}
	sp := string(os.PathSeparator)
	if s[len(s)-1:] == sp {
		return s
	}
	return s + sp
}
func DownloadNetFile(url, downloadDir string, callback func(string)) {
	f := Getrpath(downloadDir) + NewToken()
	e := DownloadFile(url, f)
	if e != nil {
		fmt.Println(e)
		return
	}
	callback(f)
}
func CacheNetFile(url, cacheDir string, callback func(string)) {
	f := Getrpath(cacheDir) + Url2cachePath(url)
	if _, e := os.Stat(f); e != nil {
		e = DownloadFile(url, f)
		if e != nil {
			fmt.Println(e)
			return
		}
	}
	callback(f)
}
func EndsWith(s, suffix string) bool {
	if len(suffix) > len(s) {
		return false
	}
	if s[len(s)-len(suffix):] == suffix {
		return true
	}
	return false
}
func Url2cachePath(url string) string {
	rUrl := GetRealUrl(url)
	s := strings.Replace(rUrl, "://", "/", -1)
	if EndsWith(s, "/") {
		return s[:len(s)-1]
	}
	return s
}
func GetRealUrl(url string) string {
	for i := 0; i < len(url); i++ {
		item := url[i : i+1]
		if item == "?" || item == "#" {
			return url[:i]
		}
	}
	return url
}
func SPrintf(a ...interface{}) string {
	return fmt.Sprint(a...)
}
func XPrintf(a int) string {
	return fmt.Sprintf("%x", a)
}
func a2i(a string) (int, error) {
	return strconv.Atoi(a)
}
func a2f(a string) float64 {
	f, e := strconv.ParseFloat(a, 32)
	if e != nil {
		fmt.Println(e)
		return 0
	}
	return float64(f)
}
func DownloadFile(url, fdist string) error {
	rp, e := http.Get(url)
	if e != nil {
		return e
	}
	defer rp.Body.Close()
	f, e := WriteFile(fdist)
	if e != nil {
		return e
	}
	defer f.Close()
	_, e = io.Copy(f, rp.Body)
	return e
}
func WriteFile(f string) (*os.File, error) {
	e := os.MkdirAll(GetDirOfFile(f), 0755)
	if e != nil {
		fmt.Println(e)
		return nil, e
	}
	file, e := os.OpenFile(f, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	return file, e
}
func GetDirOfFile(f string) string {
	for i := len(f) - 1; i > -1; i-- {
		if f[i:i+1] == string(os.PathSeparator) {
			return f[:i]
		}
	}
	return f
}
