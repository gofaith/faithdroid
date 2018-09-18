package io.github.gofaith.faithdroid.FViews;

import android.Manifest;
import android.content.Context;
import android.content.pm.PackageManager;
import android.os.Build;
import android.support.v4.app.ActivityCompat;
import android.util.Log;

import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.UIController;

import static android.content.ContentValues.TAG;

public class FPermission  {
    public static String getAttr(UIController parentController,String attr) {
        String[] strs = attr.split(":");
        String method=strs[0],arg=strs[1];
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.M) {
            switch (method) {
                case "CheckSelfPermission":
                    return parentController.activity.checkSelfPermission(arg) == PackageManager.PERMISSION_GRANTED ? "true" : "false";
            }
        }
        return null;
    }

    public static void setAttr(UIController parentController,String attr, String value) {
        if (value == null) {
            return;
        }
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.M) {
            switch (attr) {
                case "RequestPermissions":
                    String[] strs = value.split(":");
                    // fnId=strs[0] , permissions=strs[1].split("|") , requestCode=strs[2]
                    Integer requestCode = Integer.valueOf(strs[2]);
                    String[] ps=strs[1].split("#");
                    parentController.onPermissionResults.put(requestCode, strs[0]);
                    Log.d(TAG, "setAttr: "+value);
                    ActivityCompat.requestPermissions(
                            parentController.activity,
                            ps,
                            requestCode);
                    break;
            }
        }
    }
}
