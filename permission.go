package faithdroid

import (
	"strings"
)

func CheckSelfPermission(a *Activity, permission string) bool {
	b := GlobalVars.UIs[a.UI].ViewGetAttr("Permission", "CheckSelfPermission:"+permission)
	if b == "true" {
		return true
	}
	return false
}
func RequestPermissions(a *Activity, perms []string, onResult func([]bool)) {
	fnId := NewToken()
	GlobalVars.EventHandlersMap[fnId] = func(s string) string {
		var bs []bool
		UnJson(s, &bs)
		onResult(bs)
		return ""
	}
	GlobalVars.UIs[a.UI].ViewSetAttr("Permission", "RequestPermissions", fnId+":"+strings.Join(perms, "#")+":"+NewNumToken())
}

var Permissions = struct {
	CAMERA                 string
	WRITE_EXTERNAL_STORAGE string
	READ_EXTERNAL_STORAGE  string
	READ_PHONE_STATE       string
	READ_CALENDAR          string
	WRITE_CALENDAR         string
	READ_CONTACTS          string
	WRITE_CONTACTS         string
	GET_ACCOUNTS           string
	ACCESS_FINE_LOCATION   string
	ACCESS_COARSE_LOCATION string
	RECORD_AUDIO           string
	CALL_PHONE             string
	READ_CALL_LOG          string
	WRITE_CALL_LOG         string
	ADD_VOICEMAIL          string
	USE_SIP                string
	PROCESS_OUTGOING_CALLS string
	BODY_SENSORS           string
	SEND_SMS               string
	RECEIVE_SMS            string
	READ_SMS               string
	RECEIVE_WAP_PUSH       string
	RECEIVE_MMS            string
}{
	"android.permission.CAMERA",
	"android.permission.WRITE_EXTERNAL_STORAGE",
	"android.permission.READ_EXTERNAL_STORAGE",
	"android.permission.READ_PHONE_STATE",
	"android.permission.READ_CALENDAR",
	"android.permission.WRITE_CALENDAR",
	"android.permission.READ_CONTACTS",
	"android.permission.WRITE_CONTACTS",
	"android.permission.GET_ACCOUNTS",
	"android.permission.ACCESS_FINE_LOCATION",
	"android.permission.ACCESS_COARSE_LOCATION",
	"android.permission.RECORD_AUDIO",
	"android.permission.CALL_PHONE",
	"android.permission.READ_CALL_LOG",
	"android.permission.WRITE_CALL_LOG",
	"android.permission.ADD_VOICEMAIL",
	"android.permission.USE_SIP",
	"android.permission.PROCESS_OUTGOING_CALLS",
	"android.permission.BODY_SENSORS",
	"android.permission.SEND_SMS",
	"android.permission.RECEIVE_SMS",
	"android.permission.READ_SMS",
	"android.permission.RECEIVE_WAP_PUSH",
	"android.permission.RECEIVE_MMS",
}
