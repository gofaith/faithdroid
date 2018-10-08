package faithdroid

func (a *MainActivity) OnCreate() {
	var data []string
	var counter int
	LinearLayout(a).Size(-2, -2).DeferShow().Append(
		Button(a).Size(-2, -1).Text("add Item").OnClick(func() {
			counter++
			data = append(data, SPrintf(counter))
			vl := GetListViewByItemId("vl")
			vl.NotifyDataSetChanged()
		}),
		VListView(a,
			func(lv *FListView) IView {
				return LinearLayout(a).Size(-2, -1).Horizontal().Append(
					TextView(a).TextColor(Colors.Red).SetItemId(lv, "title").Size(-1, -1),
					Button(a).Size(-2, -1).SetItemId(lv, "bt"))
			},
			func(vh *ViewHolder, pos int) {
				title := vh.GetTextViewByItemId("title")
				title.Text(data[pos])
				bt := vh.GetButtonByItemId("bt")
				bt.Text(data[pos])
			},
			func() int {
				return len(data)
			},
		).Size(-2, -2).SetId("vl"))
}
