package io.github.gofaith.faithdroid.FViews;

import android.graphics.Color;
import android.graphics.drawable.Drawable;
import android.os.Build;
import android.support.v4.view.ViewCompat;
import android.support.v7.app.AppCompatActivity;
import android.view.View;
import android.view.ViewGroup;

import com.squareup.picasso.Picasso;

import org.json.JSONArray;
import org.json.JSONTokener;

import java.io.IOException;

import io.github.gofaith.faithdroid.UI.UIController;

public class FView {
    public String className,vID;
    public View view;
    public UIController parrentController;
    public void parseSize(AppCompatActivity activity, View v, String value) {
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
    public static float dp2pixel(AppCompatActivity activity, float dps) {
        float pxs = dps *activity.getResources().getDisplayMetrics().density;
        return pxs;
    }
    public static float pixel2dp(AppCompatActivity activity,float pxs) {
        float dps = pxs/activity.getResources().getDisplayMetrics().density;
        return dps;
    }

    public  void setBackgroundColor(View v, String value) {
        if (value==null)
            return;
        try {
            v.setBackgroundColor(Color.parseColor(value));
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
    public  void setBackground(View v, String value) {
        if (value==null)
            return;
        if (value.startsWith("file://")) {
            String path=value.substring("file://".length());
            Drawable draw=Drawable.createFromPath(path);
            if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.JELLY_BEAN) {
                v.setBackground(draw);
            }else {
                v.setBackgroundDrawable(draw);
            }
        } else if (value.startsWith("assets://")) {
//            Drawable d = Drawable.createFromStream(getAssets().open("Cloths/btn_no.png"), null);
            String path = value.substring("assets://".length());
            try {
                Drawable drawable = Drawable.createFromStream(parrentController.activity.getAssets().open(path), null);
                if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.JELLY_BEAN) {
                    v.setBackground(drawable);
                }else {
                    v.setBackgroundDrawable(drawable);
                }
            } catch (Exception e) {
                e.printStackTrace();
                return;
            }
        }
    }
}
