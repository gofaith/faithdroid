package io.github.stevenzack.faithdroidapp.UI

import android.content.Intent
import android.graphics.drawable.Drawable
import android.view.MenuItem
import android.widget.FrameLayout
import androidx.appcompat.app.AppCompatActivity
import faithdroid.*
import faithdroid.UIController
import io.github.stevenzack.faithdroidapp.FViews.FFrameLayout
import kotlinx.serialization.json.Json

class UIController(
    val activity: AppCompatActivity,
    val rootView:FrameLayout,
    val factivity: Activity) : UIController {

    //vars
    val viewmap:MutableMap<String,FView> = mutableMapOf();
    val menuItemsOnClickMap:MutableMap<MenuItem,String> = mutableMapOf()
    val onPermissionResults:MutableMap<Int,String> = mutableMapOf()
    val onActivityResults:MutableMap<Int,OnActivityResultListener> = mutableMapOf()
    val onDestoryEvent:MutableList<()->Unit> = mutableListOf()
    val drawableMap: MutableMap<String, Drawable> = mutableMapOf()

    var optionMenus: String = ""
    var notFinishFlag: Boolean = false
    var onBackPressed:String = ""

    private val FILE_SELECT_CODE = 7124

    private fun newActivity(activityConfig: String) {
        val config = Json.parse(FActivityConfig.serializer(), activityConfig)

        val intent = Intent()
        intent.putExtra("MyFnId",config.myFnId)
        intent.setClass(activity, MainActivity::class.java)
        if (config.myIntent.action!=""){
            intent.setAction(config.myIntent.action)
        }
        for (k in config.myIntent.extras) {
            intent.putExtra(k.key,k.value)
        }

        activity.startActivity(intent)
    }

    override fun append2RootView(p0: String?) {
        val v=viewmap.get(p0)
        rootView.addView(v!!.view, FFrameLayout.parseLP(v))
    }

    override fun runUIThread(p0: String?) {
        activity.runOnUiThread {
            Faithdroid.triggerEventHandler(p0,"")
        }
    }

    override fun viewGetAttr(vID: String?, attr: String?): String {
        return when (vID) {
            else -> (viewmap.get(vID) as AttrGettable).getAttr(attr!!)
        }
    }

    override fun getCurrentFActivity(): Activity {
        return factivity
    }

    override fun getPkg(): String {
        return activity.packageName
    }

    override fun viewSetAttr(vID: String?, attr: String?, value: String?) {
        when (vID) {
            else -> {
                val v = viewmap.get(vID)
                (v as AttrSettable).setAttr(attr!!,value!!)
            }
        }
    }

    override fun newView(p0: String?, p1: String?) {
        var v:FView
        when (p0) {
            "TextView"->{

            }
        }
    }

    interface OnActivityResultListener {
        fun onActivityResult(intent: Intent)
    }
}
