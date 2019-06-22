package io.github.stevenzack.faithdroidapp.FViews

import android.graphics.Color
import android.widget.RadioButton
import io.github.stevenzack.faithdroidapp.UI.AttrGettable
import io.github.stevenzack.faithdroidapp.UI.AttrSettable
import io.github.stevenzack.faithdroidapp.UI.FView
import io.github.stevenzack.faithdroidapp.UI.Toolkit.dp2pixel
import io.github.stevenzack.faithdroidapp.UI.UIController


class FRadioButton(controller: UIController) : FView(), AttrSettable, AttrGettable {
    private val v: RadioButton

    init {
        parentController = controller
        v = RadioButton(parentController!!.activity)
        view = v
    }

  override  fun getAttr(attr: String): String {
        val str = getUniversalAttr(attr)
        if (str != null) {
            return str
        }
        when (attr) {
            // ------------------------------------------
            "Text" -> return v.text.toString()
            "Enabled" -> return v.isEnabled.toString()
        }
        return ""
    }

   override fun setAttr(attr: String, value: String) {
        if (setUniversalAttr(attr, value)) {
            return
        }
        when (attr) {
            // -------------------------------------------------------------------
            "Text" -> v.text = value
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

            "Enabled" -> if (value == "true") {
                v.isEnabled = true
            } else {
                v.isEnabled = false
            }
        }
    }
}
