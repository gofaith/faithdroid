package io.github.stevenzack.faithdroidapp.FViews

import android.graphics.Color
import android.graphics.drawable.Drawable
import com.google.android.material.tabs.TabLayout
import io.github.stevenzack.faithdroidapp.R
import io.github.stevenzack.faithdroidapp.UI.*
import org.json.JSONArray
import org.json.JSONObject
import org.json.JSONTokener


class FTabLayout(controller: UIController) : FView(), AttrSettable, AttrGettable {
    var v: TabLayout
    var tabsList: MutableList<String> = ArrayList()

    init {
        parentController = controller
        v = TabLayout(parentController!!.activity)
        view = v
        setElevation("4")
        v.setTabTextColors(Color.parseColor("#dddddd"), Color.WHITE)
        v.setSelectedTabIndicatorColor(parentController!!.activity.getResources().getColor(R.color.colorAccent))
        v.setBackgroundColor(parentController!!.activity.getResources().getColor(R.color.colorPrimary))
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
            "TabTextColors" -> try {
                val array = JSONTokener(value).nextValue() as JSONArray
                v.setTabTextColors(Color.parseColor(array.getString(0)), Color.parseColor(array.getString(1)))
            } catch (e: Exception) {
                e.printStackTrace()
            }

            "TabIndicatorColor" -> try {
                v.setSelectedTabIndicatorColor(Color.parseColor(value))
            } catch (e: Exception) {
                e.printStackTrace()
            }

            "AddTab" -> try {
                val `object` = JSONTokener(value).nextValue() as JSONObject
                val icon = `object`.getString("Icon")
                val text = `object`.getString("Text")
                val tab = v.newTab()
                tabsList.add(text)
                tab.setText(text)
                if (icon != "") {
                    Toolkit.file2Drawable(parentController!!, icon, {drawable->
                            tab.setIcon(drawable)
                    })
                }
                v.addTab(tab)
            } catch (e: Exception) {
                e.printStackTrace()
            }

            "ViewPager" -> {
                val o = parentController!!.viewmap.get(value) ?:return
                val fViewPager = o as FViewPager
                v.setupWithViewPager(fViewPager!!.v)
            }
        }
    }
}
