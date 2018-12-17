package faithdroid

func (a *MainActivity)OnCreate() {
	TextView(a).Text(GetPackageName(a)).Show()
}