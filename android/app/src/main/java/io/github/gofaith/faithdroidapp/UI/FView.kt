package io.github.gofaith.faithdroidapp.UI

import android.view.View

open class FView(val className: Int, val vid:Long, val parentUIBridge:UIKotlinBridge) {
    var view: View? = null

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
}