package io.github.gofaith.faithdroidapp.FViews

import androidx.appcompat.widget.AppCompatButton
import io.github.gofaith.faithdroidapp.UI.FView
import io.github.gofaith.faithdroidapp.UI.UIKotlinBridge
import io.github.gofaith.faithdroidapp.UI.ClassNames

class FButton : FView {
    val v:AppCompatButton
    constructor(vid:Long,ui:UIKotlinBridge) : super(ClassNames.button.ordinal,vid,ui) {
        v = AppCompatButton(ui.activity)
        view=v
    }

    override fun setAttr(attr: Long, value: String): Boolean {
        if(super.setAttr(attr, value)){
            return true;
        }
        return false;
    }

    override fun getAttr(attr: Long): String? {
        var str = super.getAttr(attr)
        if (str != null) {
            return str;
        }

        return ""
    }
}