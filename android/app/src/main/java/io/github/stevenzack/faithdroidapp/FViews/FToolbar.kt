package io.github.stevenzack.faithdroidapp.FViews

import android.graphics.Color
import android.graphics.drawable.Drawable
import android.view.View
import androidx.appcompat.widget.Toolbar
import faithdroid.Faithdroid
import io.github.stevenzack.faithdroidapp.R
import io.github.stevenzack.faithdroidapp.UI.*


class FToolbar(controller: UIController) : FView(), AttrSettable, AttrGettable {
    private val v: Toolbar
    private val navigationIconSrc: String? = null

    init {
        parentController = controller
        v = parentController!!.activity.getLayoutInflater().inflate(
            R.layout.my_toolbar,
            parentController!!.rootView,
            false
        ) as Toolbar
        view = v
        setElevation("4")
        parentController!!.activity.setSupportActionBar(v)
        parseSize("[-2,-1]")
    }

    override fun getAttr(attr: String): String {
        val str = getUniversalAttr(attr)
        if (str != null) {
            return str
        }
        when (attr) {
            //----------------------------------------------
            "Title" -> return v.getTitle().toString()
            "SubTitle" -> return v.getSubtitle().toString()
        }
        return ""
    }

    override fun setAttr(attr: String, value: String) {
        if (setUniversalAttr(attr, value)) {
            return
        }
        when (attr) {
            // -------------------------------------------------------------------
            "Title" -> parentController!!.activity.getSupportActionBar()!!.setTitle(value)
            "TitleColor" -> v.setTitleTextColor(Color.parseColor(value))
            "SubTitle" -> parentController!!.activity.getSupportActionBar()!!.setSubtitle(value)
            "SubTitleColor" -> v.setSubtitleTextColor(Color.parseColor(value))
            "Menu" -> parentController!!.optionMenus = value
            "BackNavigation" -> {
                if (value == "true") {
                    v.setNavigationOnClickListener(object : View.OnClickListener {
                        override fun onClick(v: View) {
                            parentController!!.activity.onBackPressed()
                        }
                    })
                }
                parentController!!.activity.getSupportActionBar()!!.setDisplayHomeAsUpEnabled(value == "true")
            }
            "NavigationIcon" -> {
                if (navigationIconSrc != null && value == navigationIconSrc) {
                    return
                }
                Toolkit.file2Drawable(parentController!!, value, {drawable->
                        v.setNavigationIcon(drawable)
                })
            }
            "NavigationOnClick" -> v.setNavigationOnClickListener(object : View.OnClickListener {
                override fun onClick(v: View) {
                    Faithdroid.triggerEventHandler(value, "")
                }
            })
        }
    }
}
