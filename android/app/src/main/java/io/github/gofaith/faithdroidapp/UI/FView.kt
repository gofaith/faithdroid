package io.github.gofaith.faithdroidapp.UI

import android.view.View

open class FView(val className: Int, val vid:Long, val parentUIBridge:UIKotlinBridge) {
    var view: View? = null
}