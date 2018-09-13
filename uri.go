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
func UriRes(a string) string {
	s := ""
	if a[:1] == "/" {
		s = "file:///android_res" + a
	} else {
		s = "file:///android_res/" + a
	}
	return s
}
