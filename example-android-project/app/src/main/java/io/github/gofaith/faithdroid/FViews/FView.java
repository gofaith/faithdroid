package io.github.gofaith.faithdroid.FViews;

import android.graphics.Color;
import android.graphics.drawable.Drawable;
import android.os.Build;
import android.support.v4.view.ViewCompat;
import android.support.v7.app.AppCompatActivity;
import android.util.Log;
import android.view.Gravity;
import android.view.View;
import android.view.ViewGroup;
import android.widget.FrameLayout;
import android.widget.LinearLayout;

import org.json.JSONArray;
import org.json.JSONTokener;


import java.util.ArrayList;

import io.github.gofaith.faithdroid.UI.Toolkit;
import io.github.gofaith.faithdroid.UI.UIController;

public class FView {
    public String className,vID,TAG="FView";
    public View view;
    protected UIController parentController;
    public int[] margin=new int[4];
    public int layoutGravity;
    public float layoutWeight;
    public int[] size = new int[]{-2, -2};
    protected void parseSize(String value) {
        long width,height;
        try {
            JSONArray array=(JSONArray)(new JSONTokener(value).nextValue());
            width = array.getLong(0);
            height = array.getLong(1);
        } catch (Exception e) {
            e.printStackTrace();
            return;
        }
        ViewGroup.LayoutParams p = view.getLayoutParams();
        if (p == null) {
            p=new ViewGroup.LayoutParams(ViewGroup.LayoutParams.WRAP_CONTENT, ViewGroup.LayoutParams.WRAP_CONTENT);
        }
        if (width == -1) {
            p.width= ViewGroup.LayoutParams.WRAP_CONTENT;
        } else if (width == -2) {
            p.width= ViewGroup.LayoutParams.MATCH_PARENT;
        }else{
            p.width = (int) dp2pixel(parentController.activity, width);
        }
        if (height == -1) {
            p.height = ViewGroup.LayoutParams.WRAP_CONTENT;
        } else if (height == -2) {
            p.height = ViewGroup.LayoutParams.MATCH_PARENT;
        }else{
            p.height = (int) dp2pixel(parentController.activity, height);
        }
        size[0]=p.width;
        size[1]=p.height;
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
        Drawable draw = Toolkit.file2Drawable(parentController.activity,value);
        if (draw == null) {
            return;
        }
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.JELLY_BEAN) {
            view.setBackground(draw);
        }else {
            view.setBackgroundDrawable(draw);
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
            int left= (int) dp2pixel(parentController.activity,array.getLong(0));
            int top= (int) dp2pixel(parentController.activity,array.getLong(1));
            int right= (int) dp2pixel(parentController.activity,array.getLong(2));
            int bottom= (int) dp2pixel(parentController.activity,array.getLong(3));
            view.setPadding(left,top,right,bottom);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    void setMargin(String value) {
        try {
            JSONArray array = (JSONArray) (new JSONTokener(value).nextValue());
            margin[0] =(int)dp2pixel(parentController.activity, array.getInt(0));
            margin[1] =(int)dp2pixel(parentController.activity, array.getInt(1));
            margin[2] =(int)dp2pixel(parentController.activity, array.getInt(2));
            margin[3] =(int)dp2pixel(parentController.activity, array.getInt(3));
//            if (view.getLayoutParams() instanceof LinearLayout.LayoutParams) {
//                LinearLayout.MarginLayoutParams params = new LinearLayout.MarginLayoutParams(view.getLayoutParams());
//                params.setMargins((int) (dp2pixel(parentController.activity, array.getInt(0))),
//                        (int) (dp2pixel(parentController.activity, array.getInt(2))),
//                        (int) (dp2pixel(parentController.activity, array.getInt(1))),
//                        (int) (dp2pixel(parentController.activity, array.getInt(3))));
//                view.setLayoutParams(params);
//            } else if (view.getLayoutParams() instanceof FrameLayout.LayoutParams) {
//                FrameLayout.MarginLayoutParams params = new FrameLayout.MarginLayoutParams(view.getLayoutParams());
//                params.setMargins((int) (dp2pixel(parentController.activity, array.getInt(0))),
//                        (int) (dp2pixel(parentController.activity, array.getInt(2))),
//                        (int) (dp2pixel(parentController.activity, array.getInt(1))),
//                        (int) (dp2pixel(parentController.activity, array.getInt(3))));
//                view.setLayoutParams(params);
//            }
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    void setLayoutGravity(String value) {
        try {
            layoutGravity = Integer.parseInt(value);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
    void setElevation(String value) {
        try {
            float f = Float.parseFloat(value);
            ViewCompat.setElevation(view,f);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    void setLayoutWeight(String value) {
        try {
            layoutWeight = Float.parseFloat(value);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
