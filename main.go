package faithdroid

func (a *MainActivity) OnCreate() {
	LinearLayout(a).DeferShow().Size(-2, -2).Append(
		EditText(a).Size(-2, -1).Text("steven").InputTypeEnglish().SetId("username").Hint("username").MaxLength(6).MaxLines(1),
		EditText(a).Size(-2, -1).Hint("password").InputTypePassword().SetId("password").MaxLength(12).MaxLines(1),
		Button(a).Size(-2, -1).Text("show").OnClick(func() {
			username := GetEditTextById("username").GetText()
			password := GetEditTextById("password").GetText()
			ShowToast(a, username+":"+password)
		}))
}
