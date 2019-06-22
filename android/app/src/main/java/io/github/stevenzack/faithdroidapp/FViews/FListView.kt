package io.github.stevenzack.faithdroidapp.FViews

import android.view.View
import android.view.ViewGroup
import androidx.recyclerview.widget.RecyclerView
import faithdroid.Faithdroid
import io.github.stevenzack.faithdroidapp.UI.AttrGettable
import io.github.stevenzack.faithdroidapp.UI.AttrSettable
import io.github.stevenzack.faithdroidapp.UI.FView
import io.github.stevenzack.faithdroidapp.UI.UIController
import org.json.JSONObject
import org.json.JSONTokener


class FListView(uiController: UIController, lm: RecyclerView.LayoutManager) : FView(), AttrGettable, AttrSettable {
    private val v: RecyclerView
    private val adapter: MyAdapter
    var fnOnGetView: String? = null
    var fnOnBindData: String? = null
    var fnOnGetCount: String? = null
    var onItemClick: String? = null
    var onItemLongClick: String? = null

    init {
        parentController = uiController
        v = RecyclerView(parentController!!.activity)
        view = v
        v.setLayoutManager(lm)
        adapter = MyAdapter(this)
        v.setAdapter(adapter)
    }

    override  fun getAttr(attr: String): String {
        val str = getUniversalAttr(attr)
        if (str != null) {
            return str
        }
        when (attr) {

        }
        return ""
    }

    override fun setAttr(attr: String, value: String) {
        if (setUniversalAttr(attr, value)) {
            return
        }
        when (attr) {
            // --------------------------------------------
            "OnGetView" -> fnOnGetView = value
            "OnBindData" -> fnOnBindData = value
            "OnGetCount" -> fnOnGetCount = value
            "NotifyDataSetChanged" -> v.getAdapter()!!.notifyDataSetChanged()
            "OnItemClick" -> onItemClick = value
            "OnItemLongClick" -> onItemLongClick = value
        }
        return
    }

    inner class MyAdapter(private val parentListView: FListView) :
        RecyclerView.Adapter<MyAdapter.ViewHolder>() {
        override fun getItemCount(): Int {
            if (parentListView.fnOnGetCount == null) {
                return 0
            }
            try {
                return Integer.parseInt(Faithdroid.triggerEventHandler(parentListView.fnOnGetCount, ""))
            } catch (e: Exception) {
                e.printStackTrace()
            }

            return 0
        }

        private val mItemClickListener = object : OnItemClickListener {
            override fun onItemClick(view: View) {
                if (parentListView.onItemClick == null || parentListView.onItemClick == "")
                    return
                val pos = parentListView.v.getChildAdapterPosition(view)
                Faithdroid.triggerEventHandler(parentListView.onItemClick, pos.toString())
            }

            override fun onItemLongClick(view: View) {
                if (parentListView.onItemLongClick == null || parentListView.onItemLongClick == "")
                    return
                val pos = parentListView.v.getChildAdapterPosition(view)
                Faithdroid.triggerEventHandler(parentListView.onItemLongClick, pos.toString())
            }
        }

        inner class ViewHolder(var fView: FView, var str: String) : RecyclerView.ViewHolder(fView.view!!)

        override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): ViewHolder {
            if (parentListView.fnOnGetView == null) {
                return ViewHolder(null!!, "")
            }
            val viewGot = Faithdroid.triggerEventHandler(parentListView.fnOnGetView, "")
            val `object` = JSONTokener(viewGot).nextValue() as JSONObject
            val v = parentListView.parentController!!.viewmap.get(`object`.getString("VID"))
            v!!.view!!.setLayoutParams(FFrameLayout.parseLP(v))
            val vh = ViewHolder(v, viewGot)
            vh.itemView.setOnClickListener(object : View.OnClickListener{
                override fun onClick(v: View) {
                    mItemClickListener.onItemClick(v)
                }
            })
            vh.itemView.setOnLongClickListener(object : View.OnLongClickListener {
                override fun onLongClick(v: View): Boolean {
                    mItemClickListener.onItemLongClick(v)
                    return true
                }
            })
            return vh
        }

        override fun onBindViewHolder(holder: ViewHolder, position: Int) {
            if (parentListView.fnOnBindData == null) {
                return
            }
            try {
                val jo = JSONObject()
                jo.put("Str", holder.str)
                jo.put("Position", position)
                Faithdroid.triggerEventHandler(parentListView.fnOnBindData, jo.toString())
            } catch (e: Exception) {
                e.printStackTrace()
            }

        }
    }

    interface OnItemClickListener {
        fun onItemClick(view: View)
        fun onItemLongClick(view: View)
    }
}
