package io.github.stevenzack.faithdroidapp.FViews

import android.content.ContentValues.TAG
import android.os.Bundle
import android.util.Log
import androidx.fragment.app.Fragment
import androidx.fragment.app.FragmentManager
import androidx.fragment.app.FragmentStatePagerAdapter
import androidx.viewpager.widget.ViewPager
import faithdroid.Faithdroid
import io.github.stevenzack.faithdroidapp.UI.*
import org.json.JSONArray
import org.json.JSONTokener


class FViewPager(controller: UIController) : FView(), AttrSettable, AttrGettable {
    private var adapter: FaithCollectionPagerAdapter? = null
    var v: ViewPager
    var pages: MutableList<FPage> = ArrayList()
    var onCreateView: String? = null
    var onGetCount: String? = null
    var onPageSelected: String? = null
    var bindFTabLayout: FTabLayout? = null

    init {
        parentController = controller
        v = ViewPager(parentController!!.activity)
        v.addOnPageChangeListener(object : ViewPager.OnPageChangeListener {
            override fun onPageScrolled(i: Int, v: Float, i1: Int) {

            }

            override fun onPageSelected(i: Int) {
                if (onPageSelected != null) {
                    Faithdroid.triggerEventHandler(onPageSelected, i.toString())
                }
            }

            override fun onPageScrollStateChanged(i: Int) {

            }
        })
        view = v
        parseSize("[-1,-1]")
    }

   override fun getAttr(attr: String): String {
        val str = getUniversalAttr(attr)
        if (str != null) {
            return str
        }
        when (attr) {

        }// --------------------------------------------------
        return ""
    }

   override fun setAttr(attr: String, value: String) {
        if (setUniversalAttr(attr, value)) {
            return
        }
        when (attr) {
            // ----------------------------------------------------------------------------
            "Pages" -> {
                Log.d(TAG, "setAttr: $value")
                try {
                    val array = JSONTokener(value).nextValue() as JSONArray
                    for (i in 0 until array.length()) {
                        val `object` = array.getJSONObject(i)
                        val fPage = FPage()
                        fPage.vid = `object`.getString("VID")
                        pages.add(fPage)
                    }
                    if (adapter == null) {
                        adapter =
                            FaithCollectionPagerAdapter(parentController!!.activity.getSupportFragmentManager(), this)
                        v.setAdapter(adapter)
                    }
                } catch (e: Exception) {
                    e.printStackTrace()
                }

            }
            "TabLayout" -> {
                val o = parentController!!.viewmap.get(value) ?: return
                val fTabLayout = o as FTabLayout
                bindFTabLayout = fTabLayout
                fTabLayout.v.setupWithViewPager(v)
            }
            "OnCreateView" -> {
                onCreateView = value
                if (adapter == null && onGetCount != null) {
                    adapter = FaithCollectionPagerAdapter(parentController!!.activity.getSupportFragmentManager(), this)
                    v.setAdapter(adapter)
                }
            }
            "OnGetCount" -> {
                onGetCount = value
                if (adapter == null && onCreateView != null) {
                    adapter = FaithCollectionPagerAdapter(parentController!!.activity.getSupportFragmentManager(), this)
                    v.setAdapter(adapter)
                }
            }
            "CurrentItem" -> try {
                val array = JSONTokener(value).nextValue() as JSONArray
                val index = array.getInt(0)
                val `is` = array.getInt(1)
                v.setCurrentItem(index, `is` == 1)
            } catch (e: Exception) {
                e.printStackTrace()
            }

            "OnPageSelected" -> onPageSelected = value
        }
    }

    // ------------------
    inner class FPage {
        var vid: String? = null
    }

    internal inner class FaithCollectionPagerAdapter(fm: FragmentManager, private val fviewPager: FViewPager) :
        FragmentStatePagerAdapter(fm) {
        override fun getCount():Int{
                if (fviewPager.pages.size == 0 && onGetCount != null && onCreateView != null) {
                    try {
                        return Integer.parseInt(Faithdroid.triggerEventHandler(onGetCount, ""))
                    } catch (e: Exception) {
                    }

                }
                return fviewPager.pages.size
            }

        override fun getItem(i: Int): Fragment {
            val fragment = DemoObjectFragment().setFviewPager(fviewPager)
            val args = Bundle()
            // Our object is just an integer :-P
            args.putInt(DemoObjectFragment.ARG_OBJECT, i)
            fragment.setArguments(args)
            return fragment
        }

        override fun getPageTitle(position: Int): CharSequence {
            return if (bindFTabLayout == null || bindFTabLayout!!.tabsList.size <= position) {
                "Page $position"
            } else bindFTabLayout!!.tabsList[position]
        }
    }
}
