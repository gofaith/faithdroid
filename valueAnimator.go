package faithdroid

type FValueAnimator struct {
	FBase
}

func ValueAnimator(a IActivity) *FValueAnimator {
	v := &FValueAnimator{}
	v.VID = NewToken()
	v.ClassName = "ValueAnimator"
	v.UI = a.GetMyActivity().UI
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.ViewMap[v.VID] = v
	return v
}
func (v *FValueAnimator) OfFloat(fs ...float64) *FValueAnimator {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OfFloat", JsonArray(fs))
	return v
}
func (v *FValueAnimator) OfInt(fs ...int) *FValueAnimator {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OfInt", JsonArray(fs))
	return v
}
func (v *FValueAnimator) Duration(ms int64) *FValueAnimator {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Duration", SPrintf(ms))
	return v
}
func (v *FValueAnimator) OnValueChanged(f func(string)) *FValueAnimator {
	fnId := NewToken()
	fn:=func(s string) string {
		f(s)
		return ""
	}
	GlobalVars.EventHandlersMapLock.Lock()
	GlobalVars.EventHandlersMap[fnId] = fn
	GlobalVars.EventHandlersMapLock.Unlock()
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ValueChangedListener", fnId)
	return v
}
func (v *FValueAnimator) Start() *FValueAnimator {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Start", "")
	return v
}
