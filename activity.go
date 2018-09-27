package faithdroid

// Activity
type Activity struct {
	UI       string
	MyIntent Intent
}

func (a *Activity) GetMyActivity() *Activity {
	return a
}
func (a Activity) GetContext() string {
	return "Activity"
}

// MainActivity
type MainActivity struct {
	Activity
}

func (a *MainActivity) GetMyActivity() *Activity {
	return &a.Activity
}

// Interface
type IActivity interface {
	GetMyActivity() *Activity
}

// Intent
type Intent struct {
	Action string
	Paths  []string
	Extras map[string]string
}

var IntentActions = struct {
	VIEW, SEND, SEND_MULTIPLE string
}{
	"android.intent.action.VIEW",
	"android.intent.action.SEND",
	"android.intent.action.SEND_MULTIPLE",
}

func (a *Activity) SetIntentAction(s string) {
	a.MyIntent.Action = s
}
func (a *Activity) GetIntentAction() string {
	return a.MyIntent.Action
}
func (a *Activity) AddPath(s string) {
	a.MyIntent.Paths = append(a.MyIntent.Paths, s)
}
func (a *Activity) GetPathArray(s string) string {
	if a.MyIntent.Paths == nil {
		return "[]"
	}
	return JsonArray(a.MyIntent.Paths)
}
func (a *Activity) PutExtra(key, value string) {
	if a.MyIntent.Extras == nil {
		a.MyIntent.Extras = make(map[string]string)
	}
	a.MyIntent.Extras[key] = value
}
func (a *Activity) GetExtraValue(key string) string {
	if a.MyIntent.Extras == nil {
		return ""
	}
	if v, ok := a.MyIntent.Extras[key]; ok {
		return v
	}
	return ""
}
func (a *Activity) GetAllExtras() string {
	if a.MyIntent.Extras == nil {
		return "{}"
	}
	return JsonObject(a.MyIntent.Extras)
}

// --------------------------------------------------------

func (a *Activity) SetUIInterface(u UIController) string {
	uId := NewToken()
	GlobalVars.UIs[uId] = u
	a.UI = uId
	return uId
}
func (a *Activity) OnCreate() {

}
func (a *Activity) OnResume() {
}
func (a *Activity) OnPause() {
}
func (a *Activity) OnStart() {
}
func (a *Activity) OnDestroy() {
}

func Finish(a *Activity) {
	GlobalVars.UIs[a.UI].ViewSetAttr("Activity", "Finish", "")
}
func ShowToast(a *Activity, s string) {
	GlobalVars.UIs[a.UI].ViewSetAttr("Activity", "ShowToast", s)
}
func RunOnUIThread(a *Activity, f func()) {
	fnId := NewToken()
	GlobalVars.EventHandlersMap[fnId] = func(string) string {
		f()
		return ""
	}
	GlobalVars.UIs[a.UI].RunOnUIThread(fnId)
}
func ScanFile(a *Activity, fpath string) {
	GlobalVars.UIs[a.UI].ViewSetAttr("Activity", "ScanFile", fpath)
}
func GetExternalStorageDirectory(a *Activity) string {
	return GlobalVars.UIs[a.UI].ViewGetAttr("Activity", "ExternalStorageDirectory")
}
func GuessFileName(a *Activity, url string) string {
	return GlobalVars.UIs[a.UI].ViewGetAttr("Activity", JsonArray([]string{"GuessFileName", url}))
}
