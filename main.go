package faithdroid

type MainActivity struct {
	Activity
}
type M struct {
	VID string
	Bt  *FButton
	Tv  *FTextView
}

var strs = []string{
	"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve",
}

func (m *MainActivity) OnCreate() {
	a := &m.Activity
	hlistview(a, func() string {
		mh := M{}
		l := linearlayout(a).vertical().append(
			button(a).assign(&mh.Bt),
			textview(a).assign(&mh.Tv))
		mh.VID = l.getVID()
		return jsonObject(mh)
	}, func(str string, pos int) {
		mh := M{}
		unJson(str, &mh)
		mh.Bt.text(":" + strs[pos])
		mh.Tv.text(strs[pos])
	}, func() int {
		return len(strs)
	}).show()
}

/* ListView example
type MainActivity struct {
	Activity
}
type MyHolder struct {
	VID string
	Bt  *FButton
	Tv  *FTextView
}

var strs = []string{
	"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve",
}

func (m *MainActivity) OnCreate() {
	a := &m.Activity
	linearlayout(a).size(-2, -2).append(
		vlistview(a, func() string {
			mh := MyHolder{}
			l := linearlayout(a).horizontal().append(button(a).assign(&mh.Bt).textSize(20), textview(a).assign(&mh.Tv).textSize(20))
			mh.VID = l.getVID()
			return jsonObject(mh)
		}, func(str string, pos int) {
			mh := MyHolder{}
			unJson(str, &mh)
			fmt.Println(mh.Bt == nil)
			mh.Bt.text(":" + strs[pos])
			mh.Tv.text(strs[pos])
		}, func() int {
			return len(strs)
		}).size(-2, -2),
	).show()
}
*/
