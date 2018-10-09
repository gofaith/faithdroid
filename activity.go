package faithdroid

//Activity
type Activity struct {
	UI       string
	MyIntent Intent
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

}
func (a *Activity) OnResume() {
}
func (a *Activity) OnPause() {
}
func (a *Activity) OnStart() {
}
func (a *Activity) OnDestroy() {
}

func Finish(a IActivity) {
	GlobalVars.UIs[a.GetMyActivity().UI].ViewSetAttr("Activity", "Finish", "")
}
func ShowToast(a IActivity, s string) {
	GlobalVars.UIs[a.GetMyActivity().UI].ViewSetAttr("Activity", "ShowToast", s)
}
func RunOnUIThread(a IActivity, f func()) {
	fnId := NewToken()
	GlobalVars.EventHandlersMap[fnId] = func(string) string {
		f()
		return ""
	}
	GlobalVars.UIs[a.GetMyActivity().UI].RunOnUIThread(fnId)
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
func OpenFileChooser(a IActivity, selectType string, allowMultiple bool, callback func([]string)) {
	if !CheckSelfPermission(a, Permissions.READ_EXTERNAL_STORAGE) {
		RequestPermissions(a, []string{Permissions.READ_EXTERNAL_STORAGE}, func(bs []bool) {
			OpenFileChooser(a, selectType, allowMultiple, callback)
		})
		return
	}
	fnId := NewToken()
	GlobalVars.EventHandlersMap[fnId] = func(s string) string {
		var as []string
		UnJson(s, &as)
		callback(as)
		return ""
	}
	GlobalVars.UIs[a.GetMyActivity().UI].ViewSetAttr("Activity", "OpenFileChooser", JsonArray([]string{
		selectType,
		SPrintf(allowMultiple),
		fnId,
	}))
}
func backPressed(a IActivity) {
	GlobalVars.UIs[a.GetMyActivity().UI].ViewSetAttr("Activity", "BackPressed", "")
}
