package io.github.stevenzack.faithdroidapp.FViews

import android.os.Build
import android.util.TypedValue
import android.view.MenuItem
import androidx.annotation.NonNull
import com.google.android.material.bottomnavigation.BottomNavigationView
import faithdroid.Faithdroid
import io.github.stevenzack.faithdroidapp.UI.*
import org.json.JSONArray
import org.json.JSONTokener


class FBottomNav(controller: UIController) : FView(), AttrGettable, AttrSettable {
    var v: BottomNavigationView

    init {
        parentController = controller
        v = BottomNavigationView(parentController!!.activity)
        view = v
        parseSize("[-2,-1]")

        val outValue = TypedValue()
        parentController!!.activity.getTheme().resolveAttribute(android.R.attr.windowBackground, outValue, true)
        val d = parentController!!.activity.getResources().getDrawable(outValue.resourceId)
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.JELLY_BEAN) {
            v.setBackground(d)
        } else {
            v.setBackgroundDrawable(d)
        }
    }

   override fun getAttr(attr: String): String {
        val str = getUniversalAttr(attr)
        if (str != null) {
            return str
        }
        when (attr) {

        }//----------------------------------------------
        return ""
    }

   override fun setAttr(attr: String, value: String) {
        if (setUniversalAttr(attr, value)) {
            return
        }
        when (attr) {
            // -------------------------------------------------------------------
            "Menu" -> try {
                val `object` = JSONTokener(value).nextValue() as JSONArray
                Toolkit.parseMenu(parentController!!, v.getMenu(), `object`)
                v.setOnNavigationItemSelectedListener(object : BottomNavigationView.OnNavigationItemSelectedListener {
                    override fun onNavigationItemSelected(@NonNull menuItem: MenuItem): Boolean {
                        if (parentController!!.menuItemsOnClickMap.containsKey(menuItem)) {
                            Faithdroid.triggerEventHandler(parentController!!.menuItemsOnClickMap.get(menuItem), "")
                            return true
                        }
                        return false
                    }
                })
            } catch (e: Exception) {
                e.printStackTrace()
            }

            "SelectedItem" -> v.setSelectedItemId(v.getMenu().getItem(Integer.parseInt(value)).getItemId())
        }
    }
}
