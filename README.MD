# 全新的Go语言Android GUI框架
# New Android GUI lib for GO 

![logo](https://avatars2.githubusercontent.com/u/38978951?s=200&v=4)

# [Donate捐赠](https://jywjl.github.io/donate.html)

- [中文文档](https://github.com/gofaith/faithdroid/wiki/%E5%BF%AB%E9%80%9F%E5%BC%80%E5%A7%8B)
- [Wiki](https://github.com/gofaith/faithdroid/wiki)

---
## 快速入门 Get started

1.基础示例 'Hello world' code :

```Go
package faithdroid

func (a *MainActivity) OnCreate() {
	LinearLayout(a).DeferShow().Size(-2, -2).Append(
		TextView(a).SetId("text").Text("text view"),
		Button(a).Size(-1, -1).Text("change text").OnClick(func() {
			GetTextViewById("text").Text("text changed")
		}))
}
```
2.结果 Result

![hello](https://github.com/gofaith/faithdroid/blob/master/res/hello.jpeg?raw=true)

---

## Explore more

- [中文文档](https://github.com/gofaith/faithdroid/wiki/%E7%9B%AE%E5%BD%95)
- [Wiki](https://github.com/gofaith/faithdroid/wiki)
