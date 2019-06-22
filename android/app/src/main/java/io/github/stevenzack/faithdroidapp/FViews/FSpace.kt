package io.github.stevenzack.faithdroidapp.FViews

import android.widget.Space
import io.github.stevenzack.faithdroidapp.UI.AttrGettable
import io.github.stevenzack.faithdroidapp.UI.AttrSettable
import io.github.stevenzack.faithdroidapp.UI.FView
import io.github.stevenzack.faithdroidapp.UI.UIController

class FSpace(controller: UIController) : FView(), AttrSettable, AttrGettable {
    var v: Space

    init {
        parentController = controller
        v = Space(controller.activity)
        view = v
        layoutWeight = 1F
    }

   override fun getAttr(attr: String): String {
        val str = getUniversalAttr(attr)
        if (str != null) {
            return str
        }
        when (attr) {

        }
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
