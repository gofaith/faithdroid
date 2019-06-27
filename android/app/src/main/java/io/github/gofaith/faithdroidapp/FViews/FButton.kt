package io.github.gofaith.faithdroidapp.FViews

import androidx.appcompat.widget.AppCompatButton
import io.github.gofaith.faithdroidapp.UI.AttrSetGettable
import io.github.gofaith.faithdroidapp.UI.FView
import io.github.gofaith.faithdroidapp.UI.UIKotlinBridge
import io.github.gofaith.faithdroidapp.UI.Vars

class FButton : FView,AttrSetGettable {
    val v:AppCompatButton
    constructor(vid:Long,ui:UIKotlinBridge) : super(Vars.button.ordinal,vid,ui) {
        v = AppCompatButton(ui.activity)
        view=v
    }
    override fun setAttr(attr: Long, value: String) {

    }

    override fun getAttr(attr: Long): String {
        TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    }

}