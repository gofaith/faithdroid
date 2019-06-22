package io.github.stevenzack.faithdroidapp.UI

import android.os.Bundle
import android.util.Log
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.fragment.app.Fragment
import faithdroid.Faithdroid
import io.github.stevenzack.faithdroidapp.FViews.FViewPager


class DemoObjectFragment : Fragment() {
    private var fviewPager: FViewPager? = null

    fun setFviewPager(fviewPager: FViewPager): DemoObjectFragment {
        this.fviewPager = fviewPager
        return this
    }

    override fun onCreateView(inflater: LayoutInflater, container: ViewGroup?, savedInstanceState: Bundle?): View? {
        val args = arguments
        val i = args!!.getInt(ARG_OBJECT)
        if (fviewPager == null) {
            return null
        }
        if (fviewPager!!.onCreateView != null && !fviewPager!!.onCreateView.equals("")) {
            Log.d("TAG", "onCreateView: 1")
            val vid = Faithdroid.triggerEventHandler(fviewPager!!.onCreateView, i.toString())
            if (fviewPager!!.parentController!!.viewmap.containsKey(vid)) {
                Log.d("TAG", "onCreateView: 2")
                val cview = fviewPager!!.parentController!!.viewmap.get(vid)
                if (cview != null) {
                    Log.d("TAG", "onCreateView: 3")
                    return cview!!.view
                }
            }
        }
        if (fviewPager!!.pages.size <= i) {
            return null
        }
        val p = fviewPager!!.pages.get(i)
        return fviewPager!!.parentController!!.viewmap.get(Faithdroid.triggerEventHandler(p.vid, ""))!!.view
    }

    companion object {
        val ARG_OBJECT = "object"
    }
}
