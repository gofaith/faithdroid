package faithdroid

type MainActivity struct {
	Activity
}

func (m *MainActivity) OnCreate() {
	a := &m.Activity
	img := `http://img.ithome.com/newsuploadfiles/focus/7d739df2-570f-4e0d-b419-60ad2df697f8.jpg`
	aliimg := `http://xinyu-alog.oss-cn-hangzhou.aliyuncs.com/article/20180831/478a2bb35b674839b55066c8168a474e3333333333333.jpg`
	data := `<!DOCTYPE html>
<html>
<head>
	<title></title>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width,user-scalable=yes,initial-scale=1.0">
	<style type="text/css">
		.container {
    position: absolute;width: 100%;
    top: 50%;
    left: 50%;
    -moz-transform: translateX(-50%) translateY(-50%);
    -webkit-transform: translateX(-50%) translateY(-50%);
    transform: translateX(-50%) translateY(-50%);
}
	</style>
</head>
<body style="background-color: #000;padding: 0;margin: 0;">
	<a style="position: fixed;bottom: 10px;left: 10px;z-index: 5;" href="` + aliimg + `">下载</a>
	<div class="container"><img width="100%" id="img" src="https://www.baidu.com/img/bd_logo1.png" onclick="this.src='` + img + `'"></div>
</body>
</html>`
	ms := make(map[string]func(string) bool)
	ms[img] = func(url string) bool {
		ShowToast(a, "clicked")
		return false
	}
	ms[aliimg] = func(url string) bool {
		ShowToast(a, "clicked2")
		return true
	}
	WebView(a).Size(-2, -2).JavaScriptEnabled(true).HandleUrl(ms).LoadHtmlData(data).OnDownload(func(url string) {
		StartUriIntent(a, url)
	}).Show()
}

/* ListView example
type MainActivity struct {
	Activity
}


var strs = []string{
	"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve",
}

func (m *MainActivity) OnCreate() {
	a := &m.Activity
	linearlayout(a).deferShow().append(
		button(a).text("change").onClick(func() {
			strs = strs[:3]
			getListViewByItemId("lv").notifyDataSetChanged()
		}),
		vlistview(a,
			func(lv *FListView) IView {
				return linearlayout(a).append(
					button(a).setItemId(lv, "bt"),
					textview(a).setItemId(lv, "txt"))
			},
			func(vh *ViewHolder, pos int) {
				vh.getButtonByItemId("bt").text(strs[pos])
				vh.getTextViewByItemId("txt").text(strs[pos])
			},
			func() int {
				return len(strs)
			}).setId("lv"),
	)
}
*/
