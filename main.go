package faithdroid

func (a *MainActivity) OnCreate() {
	Button(a).Text("topleft").LayoutGravity(Gravitys.Left | Gravitys.Top).Show()
	Button(a).Text("topleft").LayoutGravity(Gravitys.Left | Gravitys.CenterVertical).Show()
	Button(a).Text("topleft").LayoutGravity(Gravitys.Left | Gravitys.Bottom).Show()
	Button(a).Text("topleft").LayoutGravity(Gravitys.CenterHorizontal | Gravitys.Bottom).Show()
	Button(a).Text("topleft").LayoutGravity(Gravitys.CenterHorizontal | Gravitys.CenterVertical).Show()
	Button(a).Text("topleft").LayoutGravity(Gravitys.CenterHorizontal | Gravitys.Top).Show()
	Button(a).Text("topleft").LayoutGravity(Gravitys.Right | Gravitys.Top).Show()
	Button(a).Text("topleft").LayoutGravity(Gravitys.Right | Gravitys.CenterVertical).Show()
	Button(a).Text("asdgkj").LayoutGravity(Gravitys.Bottom | Gravitys.Right).Show()
}
