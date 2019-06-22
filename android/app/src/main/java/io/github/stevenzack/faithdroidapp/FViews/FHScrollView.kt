package io.github.stevenzack.faithdroidapp.FViews

import android.view.ViewGroup
import android.widget.HorizontalScrollView
import android.widget.LinearLayout
import io.github.stevenzack.faithdroidapp.UI.AttrGettable
import io.github.stevenzack.faithdroidapp.UI.AttrSettable
import io.github.stevenzack.faithdroidapp.UI.FView
import io.github.stevenzack.faithdroidapp.UI.UIController
import org.json.JSONArray
import org.json.JSONTokener

class FHScrollView(controller: UIController) : FView(), AttrSettable, AttrGettable {
    var v: HorizontalScrollView
    private val linearLayout: LinearLayout

    init {
        parentController = controller
        v = HorizontalScrollView(parentController!!.activity)
        view = v
        linearLayout = LinearLayout(parentController!!.activity)
        linearLayout.setOrientation(LinearLayout.HORIZONTAL)
        v.addView(linearLayout, ViewGroup.LayoutParams.WRAP_CONTENT, ViewGroup.LayoutParams.MATCH_PARENT)
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
            // -------------------------------------------------------------------

            "AddView" -> {
                val f = parentController!!.viewmap.get(value)
                linearLayout.addView(f!!.view, FLinearLayout.parseLP(f))
            }
            "AddViewAt" -> try {
                val array = JSONTokener(value).nextValue() as JSONArray
                val pos = Integer.parseInt(array.getString(0))
                val vid = array.getString(1)
                val f1 = parentController!!.viewmap.get(vid)
                v.addView(f1!!.view, pos, FLinearLayout.parseLP(f1))
            } catch (e: Exception) {
                e.printStackTrace()
            }

        }
    }
}
