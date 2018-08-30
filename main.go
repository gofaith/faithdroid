package faithdroid

type MainActivity struct {
	Activity
}

var strs = []string{
	"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve",
}

func (m *MainActivity) OnCreate() {
	a := &m.Activity
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
		}).show()
}
*/
