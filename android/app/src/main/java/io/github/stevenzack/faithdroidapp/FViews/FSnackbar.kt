package io.github.stevenzack.faithdroidapp.FViews

import android.view.View
import com.google.android.material.snackbar.Snackbar
import faithdroid.Faithdroid
import io.github.stevenzack.faithdroidapp.UI.AttrGettable
import io.github.stevenzack.faithdroidapp.UI.AttrSettable
import io.github.stevenzack.faithdroidapp.UI.FView
import io.github.stevenzack.faithdroidapp.UI.UIController
import org.json.JSONArray
import org.json.JSONTokener


class FSnackbar(controller: UIController, anchorView: View) : FView(), AttrSettable, AttrGettable {
    var v: Snackbar

    init {
        parentController = controller
        v = Snackbar.make(anchorView, "", Snackbar.LENGTH_SHORT)
    }

    override fun getAttr(attr: String): String {
        return ""
    }

    override fun setAttr(attr: String, value: String) {
        if (value == null) {
            return
        }
        when (attr) {
            "Text" -> v.setText(value)
            "Action" -> try {
                val array = JSONTokener(value).nextValue() as JSONArray
                val fnId = array.getString(1)
                v.setAction(array.getString(0),
                    View.OnClickListener { Faithdroid.triggerEventHandler(fnId, "") })
            } catch (e: Exception) {
                e.printStackTrace()
            }

            "Show" -> v.show()
        }
    }
}
