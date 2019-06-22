package io.github.stevenzack.faithdroidapp.FViews

import android.graphics.Color
import android.text.InputFilter
import android.text.InputType
import android.widget.EditText
import io.github.stevenzack.faithdroidapp.UI.AttrGettable
import io.github.stevenzack.faithdroidapp.UI.AttrSettable
import io.github.stevenzack.faithdroidapp.UI.FView
import io.github.stevenzack.faithdroidapp.UI.Toolkit.dp2pixel
import io.github.stevenzack.faithdroidapp.UI.UIController

class FEditText(controller: UIController) : FView(), AttrGettable, AttrSettable {
    var v: EditText

    init {
        parentController = controller
        v = EditText(parentController!!.activity)
        view = v
    }

   override fun getAttr(attr: String): String {
        val str = getUniversalAttr(attr)
        if (str != null) {
            return str
        }
        when (attr) {
            // ------------------------------------------
            "Text" -> return v.getText().toString()
        }
        return ""
    }

   override fun setAttr(attr: String, value: String) {
        if (setUniversalAttr(attr, value)) {
            return
        }
        when (attr) {
            // -------------------------------------------------------------------
            "Text" -> v.setText(value)
            "TextColor" -> try {
                v.setTextColor(Color.parseColor(value))
            } catch (e: Exception) {
                e.printStackTrace()
            }

            "TextSize" -> try {
                v.setTextSize(dp2pixel(parentController!!.activity, java.lang.Float.valueOf(value)))
            } catch (e: Exception) {
                e.printStackTrace()
                return
            }

            "InputType" -> {
                val it: Int
                if (value == "Text") {
                    it = InputType.TYPE_CLASS_TEXT
                } else if (value == "Number") {
                    it = InputType.TYPE_CLASS_NUMBER
                } else if (value == "Password") {
                    it = InputType.TYPE_CLASS_TEXT or InputType.TYPE_TEXT_VARIATION_WEB_PASSWORD
                } else {
                    it = InputType.TYPE_TEXT_VARIATION_VISIBLE_PASSWORD
                }
                v.setInputType(it)
            }
            "MaxLines" -> v.setMaxLines(Integer.parseInt(value))
            "MaxEms" -> v.setMaxEms(Integer.parseInt(value))
            "Hint" -> v.setHint(value)
            "MaxLength" -> v.setFilters(arrayOf<InputFilter>(InputFilter.LengthFilter(Integer.parseInt(value))))
        }
    }
}
