package io.github.stevenzack.faithdroidapp.UI

import android.graphics.Color
import android.graphics.drawable.Drawable
import android.os.Build
import android.view.MotionEvent
import android.view.View
import android.view.ViewGroup
import androidx.core.view.ViewCompat

import androidx.constraintlayout.widget.ConstraintLayout
import faithdroid.Faithdroid
import io.github.stevenzack.faithdroidapp.FViews.FConstraintLayout
import io.github.stevenzack.faithdroidapp.UI.Toolkit.dp2pixel
import io.github.stevenzack.faithdroidapp.UI.Toolkit.pixel2dp
import org.json.JSONArray
import org.json.JSONObject
import org.json.JSONTokener

open class FView {
    var className:String = ""
    var vID:String = ""
    var view: View? = null
    var parentController:UIController? = null

    //layout
    val margin:MutableList<Int> = mutableListOf(0,0,0,0)
    var layoutGravity:Int = 0
    var layoutWeight:Float = 0f
    val size:MutableList<Int> = mutableListOf(-2,-2)
    private var onclickFn:String = ""

    val afterConstraintFuncs: MutableMap<String, ConstraintInterface> = mutableMapOf()
    interface ConstraintInterface {
        fun addConstraint(parent: FConstraintLayout, lp: ConstraintLayout.LayoutParams)
    }

    fun getUniversalAttr(attr: String?): String? {
        if (attr == null)
            return ""
        when (attr) {
            "Visibility" -> return getVisibility()
            "X" -> return getX()
            "Y" -> return getY()
            "Width" -> return getWidth()
            "Height" -> return getHeight()
            "PivotX" -> return getPivotX()
            "PivotY" -> return getPivotY()
            "ScaleX" -> return getScaleX()
            "ScaleY" -> return getScaleY()
            "Rotation" -> return getRotation()
        }
        return null
    }

    fun setUniversalAttr(attr: String, value: String?): Boolean {
        if (value == null) {
            return true
        }
        when (attr) {
            "BackgroundColor" -> setBackgroundColor(value)
            "Background" -> setBackground(value)
            "Foreground" -> setForeground(value)
            "Size" -> parseSize(value)
            "X" -> setX(value)
            "Y" -> setY(value)
            "PivotX" -> setPivotX(value)
            "PivotY" -> setPivotY(value)
            "ScaleX" -> setScaleX(value)
            "ScaleY" -> setScaleY(value)
            "Rotation" -> setRotation(value)
            "Visibility" -> setVisibility(value)
            "Padding" -> setPadding(value)
            "Margin" -> setMargin(value)
            "LayoutGravity" -> setLayoutGravity(value)
            "Elevation" -> setElevation(value)
            "LayoutWeight" -> setLayoutWeight(value)
            "OnTouch" -> setOnTouchListener(value)
            "OnClick" -> {
                if (onclickFn == null) {
                    view!!.setOnClickListener({ Faithdroid.triggerEventHandler(onclickFn, "") })
                }
                onclickFn = value
            }
            "Clickable" -> view!!.setClickable(value == "true")
            // constraint
            "Top2TopOf" -> afterConstraintFuncs[attr] = object : ConstraintInterface {
                override fun addConstraint(parent: FConstraintLayout, lp: ConstraintLayout.LayoutParams) {
                    if (value == "_Parent_") {
                        lp.topToTop = parent.view!!.getId()
                        return
                    }
                    lp.topToTop = parentController!!.viewmap.get(Faithdroid.getVidById(value))!!.view!!.getId()
                }
            }
            "Top2BottomOf" -> afterConstraintFuncs[attr] = object : ConstraintInterface {
                override fun addConstraint(parent: FConstraintLayout, lp: ConstraintLayout.LayoutParams) {
                    if (value == "_Parent_") {
                        lp.bottomToBottom = parent.view!!.getId()
                        return
                    }
                    lp.topToBottom = parentController!!.viewmap.get(Faithdroid.getVidById(value))!!.view!!.getId()
                }
            }
            "Bottom2BottomOf" -> afterConstraintFuncs[attr] = object : ConstraintInterface {
                override fun addConstraint(parent: FConstraintLayout, lp: ConstraintLayout.LayoutParams) {
                    if (value == "_Parent_") {
                        lp.bottomToBottom = parent.view!!.getId()
                        return
                    }
                    lp.bottomToBottom = parentController!!.viewmap.get(Faithdroid.getVidById(value))!!.view!!.getId()
                }
            }
            "Bottom2TopOf" -> afterConstraintFuncs[attr] = object : ConstraintInterface {
                override fun addConstraint(parent: FConstraintLayout, lp: ConstraintLayout.LayoutParams) {
                    if (value == "_Parent_") {
                        lp.bottomToTop = parent.view!!.getId()
                        return
                    }
                    lp.bottomToTop = parentController!!.viewmap.get(Faithdroid.getVidById(value))!!.view!!.getId()
                }
            }
            "Left2LeftOf" -> afterConstraintFuncs[attr] = object : ConstraintInterface {
                override fun addConstraint(parent: FConstraintLayout, lp: ConstraintLayout.LayoutParams) {
                    if (value == "_Parent_") {
                        lp.leftToLeft = parent.view!!.getId()
                        return
                    }
                    lp.leftToLeft = parentController!!.viewmap.get(Faithdroid.getVidById(value))!!.view!!.getId()
                }
            }
            "Right2RightOf" -> afterConstraintFuncs[attr] = object : ConstraintInterface {
                override fun addConstraint(parent: FConstraintLayout, lp: ConstraintLayout.LayoutParams) {
                    if (value == "_Parent_") {
                        lp.rightToRight = parent.view!!.getId()
                        return
                    }
                    lp.rightToRight = parentController!!.viewmap.get(Faithdroid.getVidById(value))!!.view!!.getId()
                }
            }
            "Left2RightOf" -> afterConstraintFuncs[attr] = object : ConstraintInterface {
                override fun addConstraint(parent: FConstraintLayout, lp: ConstraintLayout.LayoutParams) {
                    if (value == "_Parent_") {
                        lp.leftToRight = parent.view!!.getId()
                        return
                    }
                    lp.leftToRight = parentController!!.viewmap.get(Faithdroid.getVidById(value))!!.view!!.getId()
                }
            }
            "Right2LeftOf" -> afterConstraintFuncs[attr] = object : ConstraintInterface {
                override fun addConstraint(parent: FConstraintLayout, lp: ConstraintLayout.LayoutParams) {
                    if (value == "_Parent_") {
                        lp.rightToLeft = parent.view!!.getId()
                        return
                    }
                    lp.rightToLeft = parentController!!.viewmap.get(Faithdroid.getVidById(value))!!.view!!.getId()
                }
            }
            "CenterX" -> {
                setUniversalAttr("Left2LeftOf", "_Parent_")
                setUniversalAttr("Right2RightOf", "_Parent_")
            }
            "CenterY" -> {
                setUniversalAttr("Top2TopOf", "_Parent_")
                setUniversalAttr("Bottom2BottomOf", "_Parent_")
            }
            "WidthPercent" -> afterConstraintFuncs[attr] = object : ConstraintInterface {
                override fun addConstraint(parent: FConstraintLayout, lp: ConstraintLayout.LayoutParams) {
                    if (value == "1") {
                        lp.width = -1//matchParent
                        return
                    }
                    lp.width = 0
                    lp.matchConstraintPercentWidth = java.lang.Float.parseFloat(value)
                }
            }
            "HeightPercent" -> afterConstraintFuncs[attr] = object : ConstraintInterface {
                override fun addConstraint(parent: FConstraintLayout, lp: ConstraintLayout.LayoutParams) {
                    if (value == "1") {
                        lp.height = -1//matchParent
                        return
                    }
                    lp.height = 0
                    lp.matchConstraintPercentHeight = java.lang.Float.parseFloat(value)
                }
            }
            else -> return false
        }
        return true
    }
    fun parseSize(value: String) {
        val width: Int
        val height: Int
        try {
            val array = JSONTokener(value).nextValue() as JSONArray
            width = array.getInt(0)
            height = array.getInt(1)
        } catch (e: Exception) {
            e.printStackTrace()
            return
        }

        var p: ViewGroup.LayoutParams? = view!!.getLayoutParams()
        if (p == null) {
            p = ViewGroup.LayoutParams(ViewGroup.LayoutParams.WRAP_CONTENT, ViewGroup.LayoutParams.WRAP_CONTENT)
        }
        if (width == -1) {
            p.width = ViewGroup.LayoutParams.WRAP_CONTENT
        } else if (width == -2) {
            p.width = ViewGroup.LayoutParams.MATCH_PARENT
        } else {
            p.width = dp2pixel(parentController!!.activity, width.toFloat()).toInt()
        }
        if (height == -1) {
            p.height = ViewGroup.LayoutParams.WRAP_CONTENT
        } else if (height == -2) {
            p.height = ViewGroup.LayoutParams.MATCH_PARENT
        } else {
            p.height = dp2pixel(parentController!!.activity, height.toFloat()).toInt()
        }
        size[0] = p.width
        size[1] = p.height
    }
    internal fun setBackgroundColor(value: String?) {
        if (value == null)
            return
        if (value == "#0000000") {
            view!!.setBackgroundColor(Color.TRANSPARENT)
            return
        }
        try {
            view!!.setBackgroundColor(Color.parseColor(value))
        } catch (e: Exception) {
            e.printStackTrace()
        }

    }

    internal fun setBackground(value: String?) {
        if (value == null)
            return
        Toolkit.file2Drawable(parentController!!, value, object : Toolkit.OnDrawableReadyListener {
            override fun onDrawableReady(draw: Drawable?) {
                if (draw == null) {
                    return
                }
                if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.JELLY_BEAN) {
                    view!!.setBackground(draw)
                } else {
                    view!!.setBackgroundDrawable(draw)
                }
            }
        })
        if (value == "RippleEffect") {
            view!!.setClickable(true)
        }
    }

    internal fun setForeground(value: String?) {
        if (value == null) {
            return
        }
        Toolkit.file2Drawable(parentController!!, value, object : Toolkit.OnDrawableReadyListener {
            override fun onDrawableReady(draw: Drawable?) {
                if (draw == null) {
                    return
                }
                if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.M) {
                    view!!.setForeground(draw)
                }
            }
        })

        if (value == "RippleEffect") {
            view!!.setClickable(true)
        }
    }

    internal fun setVisibility(value: String) {
        var vsb = View.VISIBLE
        if (value == "INVISIBLE") {
            vsb = View.INVISIBLE
        } else if (value == "GONE") {
            vsb = View.GONE
        }
        view!!.setVisibility(vsb)
    }

    internal fun getVisibility(): String {
        val vsb = view!!.getVisibility()
        if (vsb == View.VISIBLE) {
            return "VISIBLE"
        } else if (vsb == View.GONE) {
            return "GONE"
        }
        return "INVISIBLE"
    }

    internal fun setPadding(value: String) {
        try {
            val array = JSONTokener(value).nextValue() as JSONArray
            val left = dp2pixel(parentController!!.activity, array.getLong(0).toFloat()).toInt()
            val top = dp2pixel(parentController!!.activity, array.getLong(1).toFloat()).toInt()
            val right = dp2pixel(parentController!!.activity, array.getLong(2).toFloat()).toInt()
            val bottom = dp2pixel(parentController!!.activity, array.getLong(3).toFloat()).toInt()
            view!!.setPadding(left, top, right, bottom)
        } catch (e: Exception) {
            e.printStackTrace()
        }

    }

    internal fun setMargin(value: String) {
        try {
            val array = JSONTokener(value).nextValue() as JSONArray
            margin[0] = dp2pixel(parentController!!.activity, array.getInt(0).toFloat()).toInt()
            margin[1] = dp2pixel(parentController!!.activity, array.getInt(1).toFloat()).toInt()
            margin[2] = dp2pixel(parentController!!.activity, array.getInt(2).toFloat()).toInt()
            margin[3] = dp2pixel(parentController!!.activity, array.getInt(3).toFloat()).toInt()
            //            if (view!!.getLayoutParams() instanceof LinearLayout.LayoutParams) {
            //                LinearLayout.MarginLayoutParams params = new LinearLayout.MarginLayoutParams(view!!.getLayoutParams());
            //                params.setMargins((int) (dp2pixel(parentController!!.activity, array.getInt(0))),
            //                        (int) (dp2pixel(parentController!!.activity, array.getInt(2))),
            //                        (int) (dp2pixel(parentController!!.activity, array.getInt(1))),
            //                        (int) (dp2pixel(parentController!!.activity, array.getInt(3))));
            //                view!!.setLayoutParams(params);
            //            } else if (view!!.getLayoutParams() instanceof FrameLayout.LayoutParams) {
            //                FrameLayout.MarginLayoutParams params = new FrameLayout.MarginLayoutParams(view!!.getLayoutParams());
            //                params.setMargins((int) (dp2pixel(parentController!!.activity, array.getInt(0))),
            //                        (int) (dp2pixel(parentController!!.activity, array.getInt(2))),
            //                        (int) (dp2pixel(parentController!!.activity, array.getInt(1))),
            //                        (int) (dp2pixel(parentController!!.activity, array.getInt(3))));
            //                view!!.setLayoutParams(params);
            //            }
        } catch (e: Exception) {
            e.printStackTrace()
        }

    }

    internal fun setLayoutGravity(value: String) {
        try {
            layoutGravity = Integer.parseInt(value)
        } catch (e: Exception) {
            e.printStackTrace()
        }

    }

    internal fun setElevation(value: String) {
        try {
            val f = dp2pixel(parentController!!.activity, java.lang.Float.parseFloat(value))
            ViewCompat.setElevation(view!!, f)
        } catch (e: Exception) {
            e.printStackTrace()
        }

    }

    internal fun setLayoutWeight(value: String) {
        try {
            layoutWeight = java.lang.Float.parseFloat(value)
        } catch (e: Exception) {
            e.printStackTrace()
        }

    }

    internal fun setX(value: String) {
        try {
            val f = dp2pixel(parentController!!.activity, java.lang.Float.parseFloat(value))
            view!!.setX(f)
        } catch (e: Exception) {
            e.printStackTrace()
        }

    }

    internal fun setY(value: String) {
        try {
            val f = dp2pixel(parentController!!.activity, java.lang.Float.parseFloat(value))
            view!!.setY(f)
        } catch (e: Exception) {
            e.printStackTrace()
        }

    }

    internal fun setPivotX(value: String) {
        try {
            val f = dp2pixel(parentController!!.activity, java.lang.Float.parseFloat(value))
            view!!.setPivotX(f)
        } catch (e: Exception) {
            e.printStackTrace()
        }

    }

    internal fun setPivotY(value: String) {
        try {
            val f = dp2pixel(parentController!!.activity, java.lang.Float.parseFloat(value))
            view!!.setPivotY(f)
        } catch (e: Exception) {
            e.printStackTrace()
        }

    }

    internal fun setScaleX(value: String) {
        try {
            val f = dp2pixel(parentController!!.activity, java.lang.Float.parseFloat(value))
            view!!.setScaleX(f)
        } catch (e: Exception) {
            e.printStackTrace()
        }

    }

    internal fun setScaleY(value: String) {
        try {
            val f = dp2pixel(parentController!!.activity, java.lang.Float.parseFloat(value))
            view!!.setScaleY(f)
        } catch (e: Exception) {
            e.printStackTrace()
        }

    }

    internal fun setRotation(value: String) {
        try {
            val f = dp2pixel(parentController!!.activity, java.lang.Float.parseFloat(value))
            view!!.setRotation(f)
        } catch (e: Exception) {
            e.printStackTrace()
        }

    }
    //get

    internal fun getX(): String {
        return pixel2dp(parentController!!.activity, view!!.getX()).toString()
    }

    internal fun getY(): String {
        return pixel2dp(parentController!!.activity, view!!.getY()).toString()
    }

    internal fun getWidth(): String {
        return pixel2dp(parentController!!.activity, view!!.getWidth().toFloat()).toString()
    }

    internal fun getHeight(): String {
        return pixel2dp(parentController!!.activity, view!!.getHeight().toFloat()).toString()
    }

    internal fun getPivotX(): String {
        return pixel2dp(parentController!!.activity, view!!.getPivotX()).toString()
    }

    internal fun getPivotY(): String {
        return pixel2dp(parentController!!.activity, view!!.getPivotY()).toString()
    }

    internal fun getScaleX(): String {
        return pixel2dp(parentController!!.activity, view!!.getScaleX()).toString()
    }

    internal fun getScaleY(): String {
        return pixel2dp(parentController!!.activity, view!!.getScaleY()).toString()
    }

    internal fun getRotation(): String {
        return view!!.getRotation().toString()
    }

    internal fun setOnTouchListener(value: String) {
        view!!.setOnTouchListener(View.OnTouchListener { v, event ->
            try {
                var action = ""
                when (event.action) {
                    MotionEvent.ACTION_MOVE -> action = "Move"
                    MotionEvent.ACTION_DOWN -> action = "Down"
                    MotionEvent.ACTION_UP -> action = "Up"
                }
                val `object` = JSONObject()
                `object`.put("Action", action)
                `object`.put("X", pixel2dp(parentController!!.activity, event.x).toDouble())
                `object`.put("Y", pixel2dp(parentController!!.activity, event.y).toDouble())
                Faithdroid.triggerEventHandler(value, `object`.toString())
            } catch (e: Exception) {
                e.printStackTrace()
                return@OnTouchListener false
            }

            true
        })
    }

}