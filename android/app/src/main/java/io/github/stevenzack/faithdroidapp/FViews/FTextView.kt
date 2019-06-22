package io.github.stevenzack.faithdroidapp.FViews

import android.graphics.Color
import android.view.View
import android.widget.TextView
import faithdroid.Faithdroid
import io.github.stevenzack.faithdroidapp.UI.AttrGettable
import io.github.stevenzack.faithdroidapp.UI.AttrSettable
import io.github.stevenzack.faithdroidapp.UI.FView
import io.github.stevenzack.faithdroidapp.UI.Toolkit.dp2pixel
import io.github.stevenzack.faithdroidapp.UI.UIController

class FTextView(c: UIController) : FView(), AttrGettable, AttrSettable {
    private val v: TextView

    init {
        parentController = c
        v = TextView(parentController!!.activity)
        view = v
    }

   override fun getAttr(attr: String): String {
        val str = getUniversalAttr(attr)
        if (str != null) {
            return str
        }
        when (attr) {
            //--------------------------------------------
            "Text" -> return v.text.toString()
        }
        return ""
    }

   override fun setAttr(attr: String, value: String) {
        if (setUniversalAttr(attr, value)) {
            return
        }
        when (attr) {
            // ----------------------------------------------------------
            "Text" -> if (value != null)
                v.text = value
            "TextColor" -> try {
                v.setTextColor(Color.parseColor(value))
            } catch (e: Exception) {
                e.printStackTrace()
            }

            "TextSize" -> try {
                v.textSize = dp2pixel(parentController!!.activity, java.lang.Float.valueOf(value))
            } catch (e: Exception) {
                e.printStackTrace()
                return
            }

            "OnClick" -> v.setOnClickListener { Faithdroid.triggerEventHandler(value, "") }
        }
    }
}
