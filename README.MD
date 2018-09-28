# 全新的Go语言Android GUI框架
# New Android GUI lib for GO 

- [中文文档](https://github.com/gofaith/faithdroid/wiki/%E7%9B%AE%E5%BD%95)
- [Wiki](https://github.com/gofaith/faithdroid/wiki)

---
## 快速入门 Get started

1.基础示例 'Hello world' code :

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
2.结果 Result

![hello](https://github.com/gofaith/faithdroid/blob/master/res/hello.jpeg?raw=true)

---

## Explore more

- [中文文档](https://github.com/gofaith/faithdroid/wiki/%E7%9B%AE%E5%BD%95)
- [Wiki](https://github.com/gofaith/faithdroid/wiki)
