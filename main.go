package faithdroid

type MainActivity struct {
	Activity
}

func (m *MainActivity) OnCreate() {
	a := &m.Activity
	LinearLayout(a).size(-2, -2).deferShow().append(
		Toolbar(a),
		EditText(a).size(-2, -1).typeEnglish(),
		EditText(a).size(-2, -1).typeText(),
		EditText(a).size(-2, -1).typeNumber(),
		EditText(a).size(-2, -1).typePassword())
	Fab(a).icon("drawable://add").layoutGravity(Gravitys.Bottom | Gravitys.Right).marginAll(16).elevation(8).onClick(func() {
		showToast(a, "clicked")
	}).show()
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
