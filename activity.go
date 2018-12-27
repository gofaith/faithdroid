package faithdroid

import "fmt"

//Activity
type Activity struct {
	UI                                                        string
	MyIntent                                                  Intent
	OnCreateFn, OnResumeFn, OnPauseFn, OnStartFn, OnDestroyFn func()
}

func (a Activity) GetContext() string {
	return "Activity"
}
func (a *Activity) GetMyActivity() *Activity {
	return a
}

//MainActivity
type MainActivity struct {
	Activity
}

func (a *MainActivity) GetMyActivity() *Activity {
	return &a.Activity
}

//IActivity
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
	if a.OnCreateFn != nil {
		a.OnCreateFn()
	}
}
func (a *Activity) OnResume() {
	if a.OnResumeFn != nil {
		a.OnResumeFn()
	}
}
func (a *Activity) OnPause() {
	if a.OnPauseFn != nil {
		a.OnPauseFn()
	}
}
func (a *Activity) OnStart() {
	if a.OnStartFn != nil {
		a.OnStartFn()
	}
}
func (a *Activity) OnDestroy() {
	if a.OnDestroyFn != nil {
		a.OnDestroyFn()
	}
}

func Finish(a IActivity) {
	GlobalVars.UIs[a.GetMyActivity().UI].ViewSetAttr("Activity", "Finish", "")
}
func ShowToast(a IActivity, s string) {
	GlobalVars.UIs[a.GetMyActivity().UI].ViewSetAttr("Activity", "ShowToast", s)
}
func RunOnUIThread(a IActivity, f func()) {
	fnId := NewToken()
	fn := func(string) string {
		f()
		return ""
	}
	GlobalVars.EventHandlersMapLock.Lock()
	GlobalVars.EventHandlersMap[fnId] = fn
	GlobalVars.EventHandlersMapLock.Unlock()
	GlobalVars.UIs[a.GetMyActivity().UI].RunUIThread(fnId)
}
func ScanFile(a IActivity, fpath string) {
	GlobalVars.UIs[a.GetMyActivity().UI].ViewSetAttr("Activity", "ScanFile", fpath)
}
func GetExternalStorageDirectory(a IActivity) string {
	return GlobalVars.UIs[a.GetMyActivity().UI].ViewGetAttr("Activity", "ExternalStorageDirectory")
}
func GuessFileName(a IActivity, url string) string {
	return GlobalVars.UIs[a.GetMyActivity().UI].ViewGetAttr("Activity", JsonArray([]string{"GuessFileName", url}))
}
func OpenFile(a IActivity, path string) {
	GlobalVars.UIs[a.GetMyActivity().UI].ViewSetAttr("Activity", "OpenFile", path)
}

type App struct {
	Icon, Title, Version, Package, SourceDir string
}

func GetInstalledApps(a IActivity) []App {
	str := GlobalVars.UIs[a.GetMyActivity().UI].ViewGetAttr("Activity", "InstalledApplications")
	var apps []App
	UnJson(str, &apps)
	return apps
}
func SaveDrawableToPNGFile(a IActivity, drawableID, distPath string) {
	GlobalVars.UIs[a.GetMyActivity().UI].ViewSetAttr("Activity", "SaveDrawableToPNGFile", JsonArray([]string{drawableID, distPath}))
}
func BackPressed(a IActivity) {
	GlobalVars.UIs[a.GetMyActivity().UI].ViewSetAttr("Activity", "BackPressed", "")
}
func GetUniqueID(a IActivity) string {
	return GlobalVars.UIs[a.GetMyActivity().UI].ViewGetAttr("Activity", "UniqueID")
}
func GetPackageName(a IActivity) string {
	return GlobalVars.UIs[a.GetMyActivity().UI].GetPkg()
}
func SetStatusBarColor(a IActivity, color string) {
	GlobalVars.UIs[a.GetMyActivity().UI].ViewSetAttr("Activity", "StatusBarColor", color)
}
func SetNavigationBarColor(a IActivity, color string) {
	GlobalVars.UIs[a.GetMyActivity().UI].ViewSetAttr("Activity", "NavigationBarColor", color)
}
func SetNotFinishFlag(a IActivity, b bool) {
	GlobalVars.UIs[a.GetMyActivity().UI].ViewSetAttr("Activity", "NotFinishFlag", fmt.Sprint(b))
}
func FinishAllActivity(a IActivity) {
	GlobalVars.UIs[a.GetMyActivity().UI].ViewSetAttr("Activity", "FinishAllActivity", "")
}
func KillAll(a IActivity) {
	GlobalVars.UIs[a.GetMyActivity().UI].ViewSetAttr("Activity", "KillAll", "")
}
func SetOnBackPressed(a IActivity, f func() bool) {
	fnID := NewToken()
	fn := func(string) string {
		return fmt.Sprint(f())
	}
	GlobalVars.EventHandlersMapLock.Lock()
	GlobalVars.EventHandlersMap[fnID] = fn
	GlobalVars.EventHandlersMapLock.Unlock()
	GlobalVars.UIs[a.GetMyActivity().UI].ViewSetAttr("Activity", "OnBackPressed", fnID)
}
