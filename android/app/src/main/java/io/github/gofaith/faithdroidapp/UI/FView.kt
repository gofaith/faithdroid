package io.github.gofaith.faithdroidapp.UI

import android.graphics.Color
import android.graphics.drawable.Drawable
import android.os.Build
import android.view.MotionEvent
import android.view.View
import android.view.ViewGroup
import androidx.appcompat.app.AppCompatActivity
import faithdroid.Faithdroid
import org.json.JSONArray
import org.json.JSONObject
import org.json.JSONTokener

open class FView(val className: Int, val vid:Long, val parentUIBridge:UIKotlinBridge) {
    var view: View? = null
    var width:Int = -2
    var height:Int = -2
    open fun setAttr(attr: Long, value: String):Boolean {
        when (attr) {

        }
        return false;
    }
    open fun getAttr(attr: Long):String? {
        when (attr){

        }
        return null;
    }
    open fun parseSize(value: String) {
        val w: Long
        val h: Long
        try {
            val array = JSONTokener(value).nextValue() as JSONArray
            w = array.getLong(0)
            h = array.getLong(1)
        } catch (e: Exception) {
            e.printStackTrace()
            return
        }

        var p: ViewGroup.LayoutParams? = view!!.getLayoutParams()
        if (p == null) {
            p = ViewGroup.LayoutParams(ViewGroup.LayoutParams.WRAP_CONTENT, ViewGroup.LayoutParams.WRAP_CONTENT)
        }
        if (w.equals(-1)) {
            p.width = ViewGroup.LayoutParams.WRAP_CONTENT
        } else if (w.equals(-2)) {
            p.width = ViewGroup.LayoutParams.MATCH_PARENT
        } else {
            p.width = dp2pixel(parentUIBridge.activity, w.toFloat()).toInt()
        }
        if (h.equals(-1)) {
            p.height = ViewGroup.LayoutParams.WRAP_CONTENT
        } else if (h.equals(-2)) {
            p.height = ViewGroup.LayoutParams.MATCH_PARENT
        } else {
            p.height = dp2pixel(parentUIBridge.activity, h.toFloat()).toInt()
        }
        width = p.width
        height = p.height
    }

    fun dp2pixel(activity: AppCompatActivity, dps: Float): Float {
        return dps * activity.getResources().getDisplayMetrics().density
    }

    fun pixel2dp(activity: AppCompatActivity, pxs: Float): Float {
        return pxs / activity.getResources().getDisplayMetrics().density
    }

    internal fun setBackgroundColor(value: String?) {
        if (value == null)
            return
        if (value == "#0000000") {
            view.setBackgroundColor(Color.TRANSPARENT)
            return
        }
        try {
            view.setBackgroundColor(Color.parseColor(value))
        } catch (e: Exception) {
            e.printStackTrace()
        }

    }

    internal fun setBackground(value: String?) {
        if (value == null)
            return
        Toolkit.file2Drawable(parentUIBridge, value, object : Toolkit.OnDrawableReadyListener() {
            fun onDrawableReady(draw: Drawable?) {
                if (draw == null) {
                    return
                }
                if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.JELLY_BEAN) {
                    view.setBackground(draw)
                } else {
                    view.setBackgroundDrawable(draw)
                }
            }
        })
        if (value == "RippleEffect") {
            view.setClickable(true)
        }
    }

    internal fun setForeground(value: String?) {
        if (value == null) {
            return
        }
        Toolkit.file2Drawable(parentUIBridge, value, object : Toolkit.OnDrawableReadyListener() {
            fun onDrawableReady(draw: Drawable?) {
                if (draw == null) {
                    return
                }
                if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.M) {
                    view.setForeground(draw)
                }
            }
        })

        if (value == "RippleEffect") {
            view.setClickable(true)
        }
    }

    internal fun setVisibility(value: String) {
        var vsb = View.VISIBLE
        if (value == "INVISIBLE") {
            vsb = View.INVISIBLE
        } else if (value == "GONE") {
            vsb = View.GONE
        }
        view.setVisibility(vsb)
    }

    internal fun getVisibility(): String {
        val vsb = view.getVisibility()
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
            val left = dp2pixel(parentUIBridge.activity, array.getLong(0).toFloat()).toInt()
            val top = dp2pixel(parentUIBridge.activity, array.getLong(1).toFloat()).toInt()
            val right = dp2pixel(parentUIBridge.activity, array.getLong(2).toFloat()).toInt()
            val bottom = dp2pixel(parentUIBridge.activity, array.getLong(3).toFloat()).toInt()
            view.setPadding(left, top, right, bottom)
        } catch (e: Exception) {
            e.printStackTrace()
        }

    }

    internal fun setMargin(value: String) {
        try {
            val array = JSONTokener(value).nextValue() as JSONArray
            margin[0] = dp2pixel(parentUIBridge.activity, array.getInt(0).toFloat()).toInt()
            margin[1] = dp2pixel(parentUIBridge.activity, array.getInt(1).toFloat()).toInt()
            margin[2] = dp2pixel(parentUIBridge.activity, array.getInt(2).toFloat()).toInt()
            margin[3] = dp2pixel(parentUIBridge.activity, array.getInt(3).toFloat()).toInt()
            //            if (view.getLayoutParams() instanceof LinearLayout.LayoutParams) {
            //                LinearLayout.MarginLayoutParams params = new LinearLayout.MarginLayoutParams(view.getLayoutParams());
            //                params.setMargins((int) (dp2pixel(parentUIBridge.activity, array.getInt(0))),
            //                        (int) (dp2pixel(parentUIBridge.activity, array.getInt(2))),
            //                        (int) (dp2pixel(parentUIBridge.activity, array.getInt(1))),
            //                        (int) (dp2pixel(parentUIBridge.activity, array.getInt(3))));
            //                view.setLayoutParams(params);
            //            } else if (view.getLayoutParams() instanceof FrameLayout.LayoutParams) {
            //                FrameLayout.MarginLayoutParams params = new FrameLayout.MarginLayoutParams(view.getLayoutParams());
            //                params.setMargins((int) (dp2pixel(parentUIBridge.activity, array.getInt(0))),
            //                        (int) (dp2pixel(parentUIBridge.activity, array.getInt(2))),
            //                        (int) (dp2pixel(parentUIBridge.activity, array.getInt(1))),
            //                        (int) (dp2pixel(parentUIBridge.activity, array.getInt(3))));
            //                view.setLayoutParams(params);
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
            val f = dp2pixel(parentUIBridge.activity, java.lang.Float.parseFloat(value))
            ViewCompat.setElevation(view, f)
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
            val f = dp2pixel(parentUIBridge.activity, java.lang.Float.parseFloat(value))
            view.setX(f)
        } catch (e: Exception) {
            e.printStackTrace()
        }

    }

    internal fun setY(value: String) {
        try {
            val f = dp2pixel(parentUIBridge.activity, java.lang.Float.parseFloat(value))
            view.setY(f)
        } catch (e: Exception) {
            e.printStackTrace()
        }

    }

    internal fun setPivotX(value: String) {
        try {
            val f = dp2pixel(parentUIBridge.activity, java.lang.Float.parseFloat(value))
            view.setPivotX(f)
        } catch (e: Exception) {
            e.printStackTrace()
        }

    }

    internal fun setPivotY(value: String) {
        try {
            val f = dp2pixel(parentUIBridge.activity, java.lang.Float.parseFloat(value))
            view.setPivotY(f)
        } catch (e: Exception) {
            e.printStackTrace()
        }

    }

    internal fun setScaleX(value: String) {
        try {
            val f = dp2pixel(parentUIBridge.activity, java.lang.Float.parseFloat(value))
            view.setScaleX(f)
        } catch (e: Exception) {
            e.printStackTrace()
        }

    }

    internal fun setScaleY(value: String) {
        try {
            val f = dp2pixel(parentUIBridge.activity, java.lang.Float.parseFloat(value))
            view.setScaleY(f)
        } catch (e: Exception) {
            e.printStackTrace()
        }

    }

    internal fun setRotation(value: String) {
        try {
            val f = dp2pixel(parentUIBridge.activity, java.lang.Float.parseFloat(value))
            view.setRotation(f)
        } catch (e: Exception) {
            e.printStackTrace()
        }

    }

    internal fun getX(): String {
        return pixel2dp(parentUIBridge.activity, view.getX()).toString()
    }

    internal fun getY(): String {
        return pixel2dp(parentUIBridge.activity, view.getY()).toString()
    }

    internal fun getWidth(): String {
        return pixel2dp(parentUIBridge.activity, view.getWidth().toFloat()).toString()
    }

    internal fun getHeight(): String {
        return pixel2dp(parentUIBridge.activity, view.getHeight().toFloat()).toString()
    }

    internal fun getPivotX(): String {
        return pixel2dp(parentUIBridge.activity, view.getPivotX()).toString()
    }

    internal fun getPivotY(): String {
        return pixel2dp(parentUIBridge.activity, view.getPivotY()).toString()
    }

    internal fun getScaleX(): String {
        return pixel2dp(parentUIBridge.activity, view.getScaleX()).toString()
    }

    internal fun getScaleY(): String {
        return pixel2dp(parentUIBridge.activity, view.getScaleY()).toString()
    }

    internal fun getRotation(): String {
        return view.getRotation().toString()
    }

    internal fun setOnTouchListener(value: String) {
        view.setOnTouchListener(View.OnTouchListener { v, event ->
            try {
                var action = ""
                when (event.action) {
                    MotionEvent.ACTION_MOVE -> action = "Move"
                    MotionEvent.ACTION_DOWN -> action = "Down"
                    MotionEvent.ACTION_UP -> action = "Up"
                }
                val `object` = JSONObject()
                `object`.put("Action", action)
                `object`.put("X", pixel2dp(parentUIBridge.activity, event.x).toDouble())
                `object`.put("Y", pixel2dp(parentUIBridge.activity, event.y).toDouble())
                Faithdroid.triggerEventHandler(value, `object`.toString())
            } catch (e: Exception) {
                e.printStackTrace()
                return@OnTouchListener false
            }

            true
        })
    }
}