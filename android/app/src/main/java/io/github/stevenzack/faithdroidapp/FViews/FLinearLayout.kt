package io.github.stevenzack.faithdroidapp.FViews

import android.widget.LinearLayout
import faithdroid.Faithdroid
import io.github.stevenzack.faithdroidapp.UI.AttrGettable
import io.github.stevenzack.faithdroidapp.UI.AttrSettable
import io.github.stevenzack.faithdroidapp.UI.FView
import io.github.stevenzack.faithdroidapp.UI.UIController
import org.json.JSONArray
import org.json.JSONTokener


class FLinearLayout(c: UIController) : FView(), AttrSettable, AttrGettable {
    private val v: LinearLayout

    init {
        parentController = c
        v = LinearLayout(parentController!!.activity)
        view = v
        v.orientation = LinearLayout.VERTICAL
    }

   override fun getAttr(attr: String): String {
        val str = getUniversalAttr(attr)
        if (str != null) {
            return str
        }
        when (attr) {
            // --------------------------------------------------
            "Orientation" -> return if (v.orientation == LinearLayout.HORIZONTAL)
                "HORIZONTAL"
            else
                "VERTICAL"
        }
        return ""
    }

   override fun setAttr(attr: String, value: String) {
        if (setUniversalAttr(attr, value)) {
            return
        }
        when (attr) {
            // ----------------------------------------------------------------------------
            "Orientation" -> if (value == "HORIZONTAL") {
                v.orientation = LinearLayout.HORIZONTAL
            } else {
                v.orientation = LinearLayout.VERTICAL
            }
            "AddView" -> {
                val f = parentController!!.viewmap.get(value)
                v.addView(f!!.view, parseLP(f))
            }
            "AddViewAt" -> try {
                val array = JSONTokener(value).nextValue() as JSONArray
                val pos = Integer.parseInt(array.getString(0))
                val vid = array.getString(1)
                val f1 = parentController!!.viewmap.get(vid)
                v.addView(f1!!.view, pos, parseLP(f1!!))
            } catch (e: Exception) {
                e.printStackTrace()
            }

            "OnClick" -> v.setOnClickListener { Faithdroid.triggerEventHandler(value, "") }
        }
    }

    companion object {

        fun parseLP(f: FView): LinearLayout.LayoutParams {
            val lp = LinearLayout.LayoutParams(f.size[0], f.size[1])
            lp.gravity = f.layoutGravity
            lp.weight = f.layoutWeight
            lp.leftMargin = f.margin[0]
            lp.topMargin = f.margin[1]
            lp.rightMargin = f.margin[2]
            lp.bottomMargin = f.margin[3]
            return lp
        }
    }
}
