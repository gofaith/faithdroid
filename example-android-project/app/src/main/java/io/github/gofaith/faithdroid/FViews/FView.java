package io.github.gofaith.faithdroid.FViews;

import android.graphics.Color;
import android.graphics.drawable.Drawable;
import android.os.Build;
import android.support.v4.view.ViewCompat;
import android.support.v7.app.AppCompatActivity;
import android.util.Log;
import android.view.View;
import android.view.ViewGroup;
import android.widget.LinearLayout;

import org.json.JSONArray;
import org.json.JSONTokener;


import io.github.gofaith.faithdroid.UI.UIController;

public class FView {
    public String className,vID,TAG="FView";
    public View view;
    protected UIController parrentController;
    protected void parseSize(AppCompatActivity activity, View v, String value) {
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

    void setBackgroundColor(String value) {
        if (value==null)
            return;
        try {
            view.setBackgroundColor(Color.parseColor(value));
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
    void setBackground(String value) {
        if (value==null)
            return;
        if (value.startsWith("file://")) {
            String path=value.substring("file://".length());
            Drawable draw=Drawable.createFromPath(path);
            if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.JELLY_BEAN) {
                view.setBackground(draw);
            }else {
                view.setBackgroundDrawable(draw);
            }
        } else if (value.startsWith("assets://")) {
//            Drawable d = Drawable.createFromStream(getAssets().open("Cloths/btn_no.png"), null);
            String path = value.substring("assets://".length());
            try {
                Drawable drawable = Drawable.createFromStream(parrentController.activity.getAssets().open(path), null);
                if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.JELLY_BEAN) {
                    view.setBackground(drawable);
                }else {
                    view.setBackgroundDrawable(drawable);
                }
            } catch (Exception e) {
                e.printStackTrace();
            }
        }
    }

     void setVisibility(String value) {
        int vsb=View.VISIBLE;
        if (value.equals("INVISIBLE")){
            vsb=View.INVISIBLE;
        } else if (value.equals("GONE")) {
            vsb=View.GONE;
        }
        view.setVisibility(vsb);
    }
     String getVisibility(){
        int vsb=view.getVisibility();
        if (vsb== View.VISIBLE) {
            return "VISIBLE";
        } else if (vsb == View.GONE) {
            return "GONE";
        }
        return "INVISIBLE";
    }

     void setPadding(String value) {
        try {
            JSONArray array = (JSONArray) (new JSONTokener(value).nextValue());
            int left= (int) dp2pixel(parrentController.activity,array.getLong(0));
            int top= (int) dp2pixel(parrentController.activity,array.getLong(1));
            int right= (int) dp2pixel(parrentController.activity,array.getLong(2));
            int bottom= (int) dp2pixel(parrentController.activity,array.getLong(3));
            view.setPadding(left,top,right,bottom);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    void setMargin(String value) {
        try {
            JSONArray array = (JSONArray) (new JSONTokener(value).nextValue());
            ViewGroup.MarginLayoutParams params = new ViewGroup.MarginLayoutParams(view.getLayoutParams());
            params.setMargins(array.getInt(0),array.getInt(1),array.getInt(2),array.getInt(3));
            view.setLayoutParams(params);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    void setLayoutGravity(String value) {
        try {
            LinearLayout.LayoutParams lp = new LinearLayout.LayoutParams(view.getLayoutParams());
            lp.gravity = Integer.parseInt(value);
            view.setLayoutParams(lp);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    void setElevation(String value) {
        try {
            Log.d(TAG, "setElevation: "+value);
            float f = Float.parseFloat(value);
            Log.d(TAG, "setElevation: "+f);
            ViewCompat.setElevation(view,f);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    void setLayoutWeight(String value) {
        try {
            LinearLayout.LayoutParams lp = new LinearLayout.LayoutParams(view.getLayoutParams());
            lp.weight = Float.parseFloat(value);
            view.setLayoutParams(lp);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
