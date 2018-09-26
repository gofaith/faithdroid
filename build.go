package faithdroid

type FBuild struct {
	FBase
}

func Build(a *Activity) *FBuild {
	f := &FBuild{}
	f.UI = a.UI
	return f
}
func (v *FBuild) Model() string {
	return GlobalVars.UIs[v.UI].ViewGetAttr("Activity", "Build.MODEL")
}
