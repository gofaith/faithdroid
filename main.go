package faithdroid

import (
	"strings"
)

func (a *MainActivity) OnCreate() {
	LinearLayout(a).DeferShow().Size(-2, -2).Append(
		Button(a).Size(-2, -1).Text("open").OnClick(func() {
			OpenFileChooser(a, "*/*", true, func(as []string) {
				GetTextViewById("txt").Text(strings.Join(as, "\n"))
			})
		}),
		TextView(a).Size(-2, -1).SetId("txt"))
}
