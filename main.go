package faithdroid

type MainActivity struct {
	Activity
}

func (m *MainActivity) OnCreate() {
	a := &m.Activity
	img := `http://img.ithome.com/newsuploadfiles/focus/7d739df2-570f-4e0d-b419-60ad2df697f8.jpg?r=131809578775293415`
	FrameLayout(a).Size(-2, -2).PaddingAll(50).BackgroundColor(Colors.White).Append(
		ImageView(a).Src(img).BackgroundColor(Colors.White).Foreground(Colors.RippleEffect).Elevation(16).Size(-2, -2).MarginAll(30)).Show()
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
