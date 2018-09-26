package faithdroid

import (
	"github.com/StevenZack/tools/netToolkit"
)

func (m *MainActivity) OnCreate() {
	a := &m.Activity
	LinearLayout(a).DeferShow().Append(
		EditText(a).SetId("et").Size(-2, -1),
		Button(a).Text("save").Size(-2, -1).OnClick(func() {
			RequestPermissions(a, []string{Permissions.WRITE_EXTERNAL_STORAGE}, func([]bool) {
				go func() {
					url := GetEditTextById("et").GetText()
					fdist := GetExternalStorageDirectory(a) + "/Download/" + GuessFileName(a, url)
					e := netToolkit.DownloadFile(url, fdist)
					if e != nil {
						println("download():" + e.Error())
						return
					}
					RunOnUIThread(a, func() {
						ScanFile(a, fdist)
						ShowToast(a, fdist)
					})
				}()
			})
		}),
	)
}
