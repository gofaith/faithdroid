package io.github.stevenzack.faithdroidapp.FViews

import androidx.appcompat.widget.PopupMenu
import faithdroid.Faithdroid
import io.github.stevenzack.faithdroidapp.UI.*
import org.json.JSONArray
import org.json.JSONTokener


class FPopupMenu(controller: UIController, anchorView: FView) : FView(), AttrSettable, AttrGettable {
    var v: PopupMenu

    init {
        parentController = controller
        v = PopupMenu(parentController!!.activity, anchorView.view!!)
    }

   override fun getAttr(attr: String): String{
        return ""
    }

   override fun setAttr(attr: String, value: String) {
        if (value == null) {
            return
        }
        when (attr) {
            "Menus" -> try {
                val array = JSONTokener(value).nextValue() as JSONArray
                Toolkit.parseMenu(parentController!!, v.menu, array)
                v.setOnMenuItemClickListener(PopupMenu.OnMenuItemClickListener { item ->
                    if (parentController!!.menuItemsOnClickMap.containsKey(item)) {
                        Faithdroid.triggerEventHandler(parentController!!.menuItemsOnClickMap.get(item), "")
                        return@OnMenuItemClickListener true
                    }
                    false
                })
            } catch (e: Exception) {
                e.printStackTrace()
            }

            "Show" -> v.show()
            "Dismiss" -> v.dismiss()
        }
    }
}
