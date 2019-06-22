package io.github.stevenzack.faithdroidapp.FViews

import android.content.ClipData
import android.content.ClipboardManager
import android.content.Context
import faithdroid.Faithdroid
import io.github.stevenzack.faithdroidapp.UI.AttrGettable
import io.github.stevenzack.faithdroidapp.UI.AttrSettable
import io.github.stevenzack.faithdroidapp.UI.FView
import io.github.stevenzack.faithdroidapp.UI.UIController


class FClipboard(controller: UIController) : FView(), AttrSettable, AttrGettable {
    var v: ClipboardManager

    init {
        parentController = controller
        v = parentController!!.activity.getSystemService(Context.CLIPBOARD_SERVICE) as ClipboardManager
    }

   override fun getAttr(attr: String): String {
        if (attr.startsWith("Item")) {
            try {
                val index = Integer.parseInt(attr.substring("Item".length))
                return v.getPrimaryClip()?.getItemAt(index)?.getText().toString()
            } catch (e: Exception) {
                e.printStackTrace()
                return e.toString()
            }

        }
        when (attr) {
            "ClipData" -> return v.getPrimaryClip()?.getItemAt(0)?.getText().toString()
            "ClipCount" -> return v.getPrimaryClip()?.getItemCount().toString()
        }
        return ""
    }

override    fun setAttr(attr: String, value: String) {
        if (value == null) {
            return
        }
        when (attr) {
            "OnChange" -> v.addPrimaryClipChangedListener(object : ClipboardManager.OnPrimaryClipChangedListener {
                override fun onPrimaryClipChanged() {
                    Faithdroid.triggerEventHandler(value, "")
                }
            })
            "ClipData" -> v.setPrimaryClip(ClipData.newPlainText("new", value))
        }
    }
}
