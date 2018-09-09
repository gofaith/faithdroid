package faithdroid

type TouchEvent struct {
	X, Y   float64
	Action string
}

func (v *FBaseView) OnTouch(f func(TouchEvent)) {
	fnId := NewToken()
	GlobalVars.EventHandlersMap[fnId] = func(s string) string {
		te := TouchEvent{}
		UnJson(s, &te)
		f(te)
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnTouch", fnId)
}
