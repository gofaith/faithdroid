package io.github.stevenzack.faithdroidapp.FViews

import androidx.constraintlayout.widget.ConstraintLayout
import io.github.stevenzack.faithdroidapp.UI.AttrGettable
import io.github.stevenzack.faithdroidapp.UI.AttrSettable
import io.github.stevenzack.faithdroidapp.UI.FView
import io.github.stevenzack.faithdroidapp.UI.UIController


class FConstraintLayout(controller: UIController) : FView(), AttrSettable, AttrGettable {
    var v: ConstraintLayout

    init {
        parentController = controller
        v = ConstraintLayout(parentController!!.activity)
        view = v
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
            "Append" -> {
                val vids = value.split(",".toRegex()).dropLastWhile { it.isEmpty() }.toTypedArray()
                for (i in vids.indices) {
                    val f = parentController!!.viewmap[vids[i]]
                    val lp = ConstraintLayout.LayoutParams(f!!.size[0], f!!.size[1])
                    for (entry in f.afterConstraintFuncs) {
                        entry.value.addConstraint(this, lp)
                    }
                    lp.leftMargin = f.margin[0]
                    lp.topMargin = f.margin[1]
                    lp.rightMargin = f.margin[2]
                    lp.bottomMargin = f.margin[3]
                    v.addView(parentController!!.viewmap[vids[i]]!!.view, lp)
                }
            }
        }
    }
}
