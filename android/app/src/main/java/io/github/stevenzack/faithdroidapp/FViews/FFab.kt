package io.github.stevenzack.faithdroidapp.FViews

import android.graphics.drawable.Drawable
import android.view.View
import com.google.android.material.floatingactionbutton.FloatingActionButton
import faithdroid.Faithdroid
import io.github.stevenzack.faithdroidapp.UI.*

class FFab(controller: UIController) : FView(), AttrGettable, AttrSettable {
    var v: FloatingActionButton

    init {
        parentController = controller
        v = FloatingActionButton(parentController!!.activity)
        view = v
        setElevation("8")
    }

   override fun getAttr(attr: String): String {
        val str = getUniversalAttr(attr)
        if (str != null) {
            return str
        }
        when (attr) {

        }//------------------
        return ""
    }

   override fun setAttr(attr: String, value: String) {
        if (setUniversalAttr(attr, value)) {
            return
        }
        when (attr) {
            // -------------------------------------------------------------------
            "Icon" -> Toolkit.file2Drawable(parentController!!, value, {d->
                v.setImageDrawable(d)
            })
            "OnClick" -> v.setOnClickListener(object : View.OnClickListener {
                override fun onClick(v: View) {
                    Faithdroid.triggerEventHandler(value, "")
                }
            })
        }
    }
}
