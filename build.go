package faithdroid

type FBuild struct {
	FBase
}

func Build(a IActivity) *FBuild {
	f := &FBuild{}
	f.UI = a.GetMyActivity().UI
	return f
}
func (v *FBuild) Model() string {
	return GlobalVars.UIs[v.UI].ViewGetAttr("Activity", "Build.MODEL")
}
