package faithdroid

func (a *MainActivity) OnCreate() {
	FrameLayout(a).DeferShow().Size(-2, -2).Append(
		ImageView(a).Size(-2, -2).ScaleType(ScaleTypes.FitCenter).Src(`https://img-prod-cms-rt-microsoft-com.akamaized.net/cms/api/am/imageFileData/RE2mupw?ver=8a53&q=90&m=6&h=201&w=358&b=%23FFFFFFFF&l=f&o=t&aim=true`).Foreground(Colors.RippleEffect))
}
