package io.github.stevenzack.faithdroidapp.FViews

import android.content.DialogInterface
import androidx.appcompat.app.AlertDialog
import faithdroid.Faithdroid
import io.github.stevenzack.faithdroidapp.UI.AttrGettable
import io.github.stevenzack.faithdroidapp.UI.AttrSettable
import io.github.stevenzack.faithdroidapp.UI.FView
import io.github.stevenzack.faithdroidapp.UI.UIController
import org.json.JSONArray
import org.json.JSONTokener

class FAlertDialog(controller: UIController) : FView(), AttrSettable, AttrGettable {
    var v: AlertDialog.Builder
    var di: DialogInterface? = null

    init {
        parentController = controller
        v = AlertDialog.Builder(parentController!!.activity)
    }

   override fun getAttr(attr: String): String {
        return ""
    }

   override fun setAttr(attr: String, value: String) {
        if (value == null)
            return
        when (attr) {
            "Title" -> v.setTitle(value)
            "View" -> if (parentController!!.viewmap.containsKey(value)) {
                v.setView(parentController!!.viewmap.get(value)!!.view)
            }
            "PositiveButton" -> try {
                val array = JSONTokener(value).nextValue() as JSONArray
                val fnId = array.getString(1)
                v.setPositiveButton(
                    array.getString(0),
                    DialogInterface.OnClickListener { dialog, which -> Faithdroid.triggerEventHandler(fnId, "") })
            } catch (e: Exception) {
                e.printStackTrace()
            }

            "NegativeButton" -> try {
                val array = JSONTokener(value).nextValue() as JSONArray
                val fnId = array.getString(1)
                v.setNegativeButton(
                    array.getString(0),
                    DialogInterface.OnClickListener { dialog, which -> Faithdroid.triggerEventHandler(fnId, "") })
            } catch (e: Exception) {
                e.printStackTrace()
            }

            "Show" -> di = v.show()
            "Dismiss" -> if (di != null) {
                di!!.dismiss()
            }
        }
    }
}
