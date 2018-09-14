package faithdroid

type MainActivity struct {
	Activity
}

func (m *MainActivity) OnCreate() {
	a := &m.Activity
	LinearLayout(a).Size(-2, -2).Append(
		ViewPager(a).SetId("id").LayoutWeight(1).Size(-2, -2).Pages(
			Page(func() IView {
				return TextView(a).Text("one")
			}),
			Page(func() IView {
				return TextView(a).Text("two")
			})),
		BottomNav(a).Menus(
			MenuItem("title1").onClick(func() {
				GetViewPagerById("id").CurrentItem(0, true)
			}),
			MenuItem("title2").onClick(func() {
				GetViewPagerById("id").CurrentItem(1, true)
			})),
	).Show()
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
