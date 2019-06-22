package io.github.stevenzack.faithdroidapp.FViews

import android.view.View
import android.widget.AdapterView
import android.widget.ArrayAdapter
import android.widget.Spinner
import faithdroid.Faithdroid
import io.github.stevenzack.faithdroidapp.UI.AttrGettable
import io.github.stevenzack.faithdroidapp.UI.AttrSettable
import io.github.stevenzack.faithdroidapp.UI.FView
import io.github.stevenzack.faithdroidapp.UI.UIController
import org.json.JSONArray
import org.json.JSONTokener

import java.util.ArrayList

class FSpinner(controller: UIController) : FView(), AttrGettable, AttrSettable {
    var v: Spinner
    private val data_list = ArrayList<String>()
    private var adapter: ArrayAdapter<String>? = null
    private val selected = -1

    init {
        parentController = controller
        v = Spinner(parentController!!.activity)
        view = v
        adapter = ArrayAdapter(parentController!!.activity, android.R.layout.simple_spinner_dropdown_item, data_list)
        v.adapter = adapter
    }

  override  fun getAttr(attr: String): String {
        val str = getUniversalAttr(attr)
        if (str != null) {
            return str
        }
        when (attr) {
            // ------------------------------------------
            "Enabled" -> return v.isEnabled.toString()
            "SelectedIndex" -> return v.selectedItemPosition.toString()
        }
        return ""
    }

   override fun setAttr(attr: String, value: String) {
        if (setUniversalAttr(attr, value)) {
            return
        }
        when (attr) {
            // -------------------------------------------------------------------
            "NotifyDataSetChanged" -> adapter!!.notifyDataSetChanged()
            "List" -> try {
                val array = JSONTokener(value).nextValue() as JSONArray
                data_list.clear()
                for (i in 0 until array.length()) {
                    data_list.add(array.getString(i))
                }
                adapter!!.notifyDataSetChanged()
            } catch (e: Exception) {
                e.printStackTrace()
            }

            "ListAdd" -> {
                data_list.add(value)
                adapter!!.notifyDataSetChanged()
            }
            "ListRemove" -> {
                data_list.removeAt(Integer.parseInt(value))
                adapter!!.notifyDataSetChanged()
            }
            "OnItemClick" -> v.onItemSelectedListener = object : AdapterView.OnItemSelectedListener {
                override fun onItemSelected(parent: AdapterView<*>, view: View, position: Int, id: Long) {
                    Faithdroid.triggerEventHandler(value, position.toString())
                }

                override fun onNothingSelected(parent: AdapterView<*>) {

                }
            }
            "Enabled" -> v.isEnabled = value == "true"
        }
    }
}
