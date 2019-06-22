package io.github.stevenzack.faithdroidapp.FViews

import android.content.ContentValues.TAG
import android.content.pm.PackageManager
import android.os.Build
import android.util.Log
import androidx.core.app.ActivityCompat
import io.github.stevenzack.faithdroidapp.UI.UIController


object FPermission {
    fun getAttr(parentController: UIController, attr: String): String? {
        val strs = attr.split(":".toRegex()).dropLastWhile { it.isEmpty() }.toTypedArray()
        val method = strs[0]
        val arg = strs[1]
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.M) {
            when (method) {
                "CheckSelfPermission" -> return if (parentController.activity.checkSelfPermission(arg) === PackageManager.PERMISSION_GRANTED) "true" else "false"
            }
        }
        return null
    }

    fun setAttr(parentController: UIController, attr: String, value: String?) {
        if (value == null) {
            return
        }
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.M) {
            when (attr) {
                "RequestPermissions" -> {
                    val strs = value.split(":".toRegex()).dropLastWhile { it.isEmpty() }.toTypedArray()
                    // fnId=strs[0] , permissions=strs[1].split("|") , requestCode=strs[2]
                    val requestCode = Integer.valueOf(strs[2])
                    val ps = strs[1].split("#".toRegex()).dropLastWhile { it.isEmpty() }.toTypedArray()
                    parentController.onPermissionResults.put(requestCode, strs[0])
                    Log.d(TAG, "setAttr: $value")
                    ActivityCompat.requestPermissions(
                        parentController.activity,
                        ps,
                        requestCode
                    )
                }
            }
        }
    }
}
