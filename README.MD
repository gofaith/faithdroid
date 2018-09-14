# Faithdroid :Android GUI lib for GO

1.'Hello world' code :

```Go
package faithdroid

func (m *MainActivity) OnCreate() {
	a := &m.Activity
	LinearLayout(a).Size(-2, -2).Append(
		TextView(a).Text("text view").SetId("text"),
		Button(a).Text("change text").OnClick(func() {
			GetTextViewById("text").Text("clicked")
		}),
	).Show()
}
```
2.Result
