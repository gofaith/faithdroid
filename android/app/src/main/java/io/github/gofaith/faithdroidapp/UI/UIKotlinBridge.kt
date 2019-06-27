package io.github.gofaith.faithdroidapp.UI

import android.widget.FrameLayout
import androidx.appcompat.app.AppCompatActivity
import androidx.core.view.ViewCompat
import faithdroid.UIBridge
import io.github.gofaith.faithdroidapp.FViews.FButton

class UIKotlinBridge(val activity: AppCompatActivity,val rootView: FrameLayout) :UIBridge {
    val viewmap: MutableMap<Long,FView> = mutableMapOf()
    override fun getAttr(p0: Long, p1: Long): String {
        return ""
    }

    override fun setAttr(p0: Long, p1: Long, p2: String?) {

    }

    override fun show(vid: Long) {
        rootView.addView(viewmap.get(vid)?.view)
    }

    override fun new_(className: Long, vid: Long) {
        when (className.toInt()) {
            Vars.button.ordinal ->
                viewmap.put(vid,FButton(vid,this))
        }
        viewmap.get(vid)?.view?.id = ViewCompat.generateViewId()
    }

    override fun runOnUIThread(p0: Long) {

    }

}