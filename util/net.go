package util

import (
	"strings"

	"github.com/StevenZack/tools/fileToolkit"

	"github.com/StevenZack/tools/netToolkit"
)

func CacheNetFile(url, pkg string) (string, error) {
	path := "/data/data/" + pkg + "/cacheDir/"
	url = strings.Replace(url, "://", "/", -1)
	if fileToolkit.IsFileExists(path + url) {
		return path + url, nil
	}
	e := netToolkit.DownloadFile(url, path+url)
	if e != nil {
		return "", e
	}
	return path + url, nil
}
