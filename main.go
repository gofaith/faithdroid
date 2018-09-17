package faithdroid

import (
	"fmt"
	"net/http"
)

var (
	isRunning bool
)

func (m *MainActivity) OnCreate() {
	a := &m.Activity
	LinearLayout(a).Size(-2, -2).Append(
		Button(a).Text("start").OnClick(func() {
			StartService(a, func() {
				http.HandleFunc("/", home)
				e := http.ListenAndServe(":8888", nil)
				if e != nil {
					fmt.Println(e)
					return
				}
			})
		}),
		Button(a).Text("stop"),
	).Show()
}
func home(w http.ResponseWriter, r *http.Request) {
	s := `<!DOCTYPE html> 
<html> 
<head> 
<title></title> 
<meta charset="utf-8"> 
<meta name='viewport' content='width=device-width,user-scalable=yes,initial-scale=1.0'> 
<style type="text/css">
	@keyframes fadein{
		from{opacity: 0;}
		to{opacity: 1;}
	}
	.loader {
	    border: 8px solid #f3f3f3; /* Light grey */
	    border-top: 8px solid #3498db; /* Blue */
	    border-radius: 50%;
	    width: 60px;
	    height: 60px;
	    position: fixed;
	    left: 50%;top: 50%;
	    margin-left: -30px;margin-top: -30px;
	    animation: spin 2s linear infinite;
	}
	@keyframes spin {
	    0% { transform: rotate(0deg); }
	    100% { transform: rotate(360deg); }
	}
</style>
</head> 
<body style='background-color: #000;padding: 0;margin: 0;'> 
<a style='position: fixed;bottom: 10px;left: 10px;background-color: #333;border-radius: 50%;padding-left: 2px;padding-right: 2px;' href='http://xinyu-alog.oss-cn-hangzhou.aliyuncs.com/article/20180831/478a2bb35b674839b55066c8168a474e3333333333333.jpg'>
	<svg t="1537146663812" class="icon" style="" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1859" xmlns:xlink="http://www.w3.org/1999/xlink" width="32" height="32"><defs><style type="text/css"></style></defs><path d="M213.333333 853.333333h597.333334v-85.333333H213.333333m597.333334-384h-170.666667V128H384v256H213.333333l298.666667 298.666667 298.666667-298.666667z" fill="#e6e6e6" p-id="1860"></path></svg>
</a> 
	<div class="loader"></div>
	<img width='100%' style="opacity: 0;" onload="document.getElementsByClassName('loader')[0].style.display='none';this.style.animation='fadein 1s';this.style.opacity='1';" id='img' src='http://xinyu-alog.oss-cn-hangzhou.aliyuncs.com/article/20180831/478a2bb35b674839b55066c8168a474e3333333333333.jpg'> 
</body> 
</html>`
	fmt.Fprintln(w, s)
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
