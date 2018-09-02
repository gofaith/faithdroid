package faithdroid

type MainActivity struct {
	Activity
}

func (m *MainActivity) OnCreate() {
	a := &m.Activity
	linearlayout(a).size(-2, -2).deferShow().append(
		toolbar(a).size(-2, -1).menu(
			menuItem("title"),
			menuItem("asd").subMenu(
				menuItem("asd-1"),
				menuItem("asd-2").subMenu(
					menuItem("asd-2-1")))),
		button(a).size(-1, -1).text("click"))
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
