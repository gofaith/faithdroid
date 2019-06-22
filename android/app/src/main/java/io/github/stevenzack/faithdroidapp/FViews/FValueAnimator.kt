package io.github.stevenzack.faithdroidapp.FViews

import android.animation.ValueAnimator
import android.content.ContentValues.TAG
import android.util.Log
import faithdroid.Faithdroid
import io.github.stevenzack.faithdroidapp.UI.AttrGettable
import io.github.stevenzack.faithdroidapp.UI.AttrSettable
import io.github.stevenzack.faithdroidapp.UI.FView
import io.github.stevenzack.faithdroidapp.UI.UIController
import org.json.JSONArray
import org.json.JSONTokener


class FValueAnimator(uiController: UIController) : FView(), AttrGettable, AttrSettable {
    var v: ValueAnimator

    init {
        parentController = uiController
        v = ValueAnimator()
    }

    override fun getAttr(attr: String): String{
        return ""
    }

    override fun setAttr(attr: String, value: String) {
        if (attr == null || value == null) {
            return
        }
        Log.d(TAG, "setAttr: $attr,$value")
        when (attr) {
            "OfFloat" -> try {
                val array = JSONTokener(value).nextValue() as JSONArray
                val vs = FloatArray(array.length())
                for (i in 0 until array.length()) {
                    vs[i] = array.getDouble(i).toFloat()
                }
                v.setFloatValues(*vs)
            } catch (e: Exception) {
                e.printStackTrace()
            }

            "OfInt" -> try {
                val array = JSONTokener(value).nextValue() as JSONArray
                val vs = IntArray(array.length())
                for (i in 0 until array.length()) {
                    vs[i] = array.getInt(i)
                }
                v.setIntValues(*vs)
            } catch (e: Exception) {
                e.printStackTrace()
            }

            "Duration" -> try {
                val d = java.lang.Long.parseLong(value)
                v.duration = d
            } catch (e: Exception) {
                e.printStackTrace()
            }

            "ValueChangedListener" -> v.addUpdateListener { animation ->
                Faithdroid.triggerEventHandler(
                    value,
                    animation.animatedValue.toString()
                )
            }
            "Start" -> v.start()
        }
    }
}
