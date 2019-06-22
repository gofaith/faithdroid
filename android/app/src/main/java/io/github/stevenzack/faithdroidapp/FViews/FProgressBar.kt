package io.github.stevenzack.faithdroidapp.FViews

import android.widget.ProgressBar
import io.github.stevenzack.faithdroidapp.UI.AttrGettable
import io.github.stevenzack.faithdroidapp.UI.AttrSettable
import io.github.stevenzack.faithdroidapp.UI.FView
import io.github.stevenzack.faithdroidapp.UI.UIController


class FProgressBar(uiController: UIController) : FView(), AttrSettable, AttrGettable {
    var v: ProgressBar

    init {
        parentController = uiController
        v = ProgressBar(parentController!!.activity)
        view = v
    }

   override fun getAttr(attr: String): String {
        val str = getUniversalAttr(attr)
        if (str != null) {
            return str
        }
        when (attr) {

        }// ------------------------------------------
        return ""
    }

   override fun setAttr(attr: String, value: String) {
        if (setUniversalAttr(attr, value)) {
            return
        }
        when (attr) {

        }// -------------------------------------------------------------------
    }
}
