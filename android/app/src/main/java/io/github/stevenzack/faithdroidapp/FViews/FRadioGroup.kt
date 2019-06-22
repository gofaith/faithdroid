package io.github.stevenzack.faithdroidapp.FViews

import android.content.ContentValues.TAG
import android.util.Log
import android.widget.RadioGroup
import faithdroid.Faithdroid
import io.github.stevenzack.faithdroidapp.UI.AttrGettable
import io.github.stevenzack.faithdroidapp.UI.AttrSettable
import io.github.stevenzack.faithdroidapp.UI.FView
import io.github.stevenzack.faithdroidapp.UI.UIController


class FRadioGroup(controller: UIController) : FView(), AttrGettable, AttrSettable {
    private val v: RadioGroup
    private val idMap = HashMap<Int, String>()

    init {
        parentController = controller
        v = RadioGroup(parentController!!.activity)
        view = v
        v.orientation = RadioGroup.VERTICAL
    }

   override fun getAttr(attr: String): String {
        val str = getUniversalAttr(attr)
        if (str != null) {
            return str
        }
        when (attr) {
            // --------------------------------------------------
            "Orientation" -> return if (v.orientation == RadioGroup.HORIZONTAL)
                "HORIZONTAL"
            else
                "VERTICAL"
            "CheckedRadioButtonId" -> {
                Log.d(TAG, "getAttr: checked id=" + v.checkedRadioButtonId)
                return idMap[v.checkedRadioButtonId]!!
            }
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
                v.orientation = RadioGroup.HORIZONTAL
            } else {
                v.orientation = RadioGroup.VERTICAL
            }
            "AddView" -> {
                val f = parentController!!.viewmap.get(value)
                val lp = RadioGroup.LayoutParams(f!!.size[0], f!!.size[1])
                lp.gravity = f!!.layoutGravity
                lp.weight = f!!.layoutWeight
                lp.leftMargin = f!!.margin[0]
                lp.topMargin = f!!.margin[1]
                lp.rightMargin = f!!.margin[2]
                lp.bottomMargin = f!!.margin[3]
                v.addView(f!!.view, lp)
                idMap[f!!.view!!.getId()] = f!!.vID
            }
            "OnCheckedChange" -> v.setOnCheckedChangeListener { group, checkedId ->
                val checkedVid = idMap[checkedId]
                Faithdroid.triggerEventHandler(value, checkedVid)
            }
        }
    }
}
