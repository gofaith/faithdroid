package io.github.gofaith.faithdroid.FViews;

import android.support.v7.app.AppCompatActivity;
import android.view.View;
import android.view.ViewGroup;

import org.json.JSONArray;
import org.json.JSONTokener;

import io.github.gofaith.faithdroid.UI.UIController;

public class FView {
    public String className,vID;
    public View view;
    public UIController parrentController;
    public static void parseSize(AppCompatActivity activity, View v, String value) {
        long width,height;
        try {
            JSONArray array=(JSONArray)(new JSONTokener(value).nextValue());
            width = array.getLong(0);
            height = array.getLong(1);
        } catch (Exception e) {
            e.printStackTrace();
            return;
        }
        ViewGroup.LayoutParams p = v.getLayoutParams();
        if (p == null) {
            p=new ViewGroup.LayoutParams(ViewGroup.LayoutParams.WRAP_CONTENT, ViewGroup.LayoutParams.WRAP_CONTENT);
        }
        if (width == -1) {
            p.width= ViewGroup.LayoutParams.WRAP_CONTENT;
        } else if (width == -2) {
            p.width= ViewGroup.LayoutParams.MATCH_PARENT;
        }else{
            p.width = (int) dp2pixel(activity, width);
        }
        if (height == -1) {
            p.height = ViewGroup.LayoutParams.WRAP_CONTENT;
        } else if (height == -2) {
            p.height = ViewGroup.LayoutParams.MATCH_PARENT;
        }else{
            p.height = (int) dp2pixel(activity, height);
        }
        v.setLayoutParams(p);
    }
    public static float dp2pixel(AppCompatActivity activity,float dps) {
        float pxs = dps *activity.getResources().getDisplayMetrics().density;
        return pxs;
    }
    public static float pixel2dp(AppCompatActivity activity,float pxs) {
        float dps = pxs/activity.getResources().getDisplayMetrics().density;
        return dps;
    }
}
