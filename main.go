package faithdroid

func (m *MainActivity) OnCreate() {
	a := &m.Activity
	LinearLayout(a).Size(-2, -2).DeferShow().Append(
		RadioGroup(a).Append(
			RadioButton(a).Text("one").SetId("one"),
			RadioButton(a).Text("two").SetId("two"),
			RadioButton(a).Text("three"),
		).OnCheckedChange(func(vid string) {
			t := ""
			switch vid {
			case GetRadioButtonById("one").VID:
				t = "one"
			case GetRadioButtonById("two").VID:
				t = "two"
			default:
				t = vid
			}
			ShowToast(a, t)
		}),
	)
}
