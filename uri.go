package faithdroid

func UriAssets(a string) string {
	S := ""
	if a[:1] == "/" {
		S = "file:///android_asset" + a
	} else {
		S = "file:///android_asset/" + a
	}
	return S
}
