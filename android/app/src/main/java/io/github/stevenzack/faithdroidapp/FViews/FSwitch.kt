package io.github.stevenzack.faithdroidapp.FViews

import android.graphics.Color
import android.widget.CompoundButton
import androidx.appcompat.widget.SwitchCompat
import faithdroid.Faithdroid
import io.github.stevenzack.faithdroidapp.UI.AttrGettable
import io.github.stevenzack.faithdroidapp.UI.AttrSettable
import io.github.stevenzack.faithdroidapp.UI.FView
import io.github.stevenzack.faithdroidapp.UI.Toolkit.dp2pixel
import io.github.stevenzack.faithdroidapp.UI.UIController


class FSwitch(controller: UIController) : FView(), AttrSettable, AttrGettable {
    var v: SwitchCompat

    init {
        parentController = controller
        v = SwitchCompat(parentController!!.activity)
        view = v
    }

  override  fun getAttr(attr: String): String {
        val str = getUniversalAttr(attr)
        if (str != null) {
            return str
        }
        when (attr) {
            // ------------------------------------------
            "Checked" -> return v.isChecked().toString()
            "Enabled" -> return v.isEnabled().toString()
            "TextOn" -> return v.getTextOn().toString()
            "TextOff" -> return v.getTextOff().toString()
        }
        return ""
    }

  override  fun setAttr(attr: String, value: String) {
        if (setUniversalAttr(attr, value)) {
            return
        }
        when (attr) {
            // -------------------------------------------------------------------
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

            "Enabled" -> if (value == "true") {
                v.setEnabled(true)
            } else {
                v.setEnabled(false)
            }
            "TextOn" -> {
                v.setShowText(true)
                v.setTextOn(value)
            }
            "TextOff" -> {
                v.setShowText(true)
                v.setTextOff(value)
            }
            "OnCheckedChangeListener" -> v.setOnCheckedChangeListener(CompoundButton.OnCheckedChangeListener { buttonView, isChecked ->
                Faithdroid.triggerEventHandler(
                    value,
                    isChecked.toString()
                )
            })
        }
    }
}
