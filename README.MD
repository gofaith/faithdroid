# Faithdroid - New Android GUI lib for GO

- [Wiki](https://github.com/gofaith/faithdroid/wiki)
- [中文文档](https://github.com/gofaith/faithdroid/wiki/%E7%9B%AE%E5%BD%95)

---
## Get started

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

![hello](https://github.com/gofaith/faithdroid/blob/master/res/hello.jpeg?raw=true)

---

## Explore more

- [Wiki](https://github.com/gofaith/faithdroid/wiki)
- [中文文档](https://github.com/gofaith/faithdroid/wiki/%E7%9B%AE%E5%BD%95)
