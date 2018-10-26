package io.github.gofaith.faithdroid.FViews;

import android.graphics.Color;
import android.graphics.drawable.Drawable;
import android.os.Build;
import android.support.constraint.ConstraintLayout;
import android.support.v4.view.ViewCompat;
import android.support.v7.app.AppCompatActivity;
import android.util.Log;
import android.view.Gravity;
import android.view.MotionEvent;
import android.view.View;
import android.view.ViewGroup;
import android.widget.FrameLayout;
import android.widget.LinearLayout;

import org.json.JSONArray;
import org.json.JSONObject;
import org.json.JSONTokener;


import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;

import faithdroid.Faithdroid;
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
    public Map<String, ConstraintInterface> afterConstraintFuncs = new HashMap<>();
    interface ConstraintInterface{
        void addConstraint(FConstraintLayout parent, ConstraintLayout.LayoutParams lp);
    }
    protected String getUniversalAttr(String attr) {
        if(attr==null)
            return "";
        switch (attr) {
            case "Visibility":
                return getVisibility();
            case "X":
                return getX();
            case "Y":
                return getY();
            case "Width":
                return getWidth();
            case "Height":
                return getHeight();
            case "PivotX":
                return getPivotX();
            case "PivotY":
                return getPivotY();
            case "ScaleX":
                return getScaleX();
            case "ScaleY":
                return getScaleY();
            case "Rotation":
                return getRotation();
        }
        return null;
    }

    protected boolean setUniversalAttr(String attr, final String value) {
        if (value == null) {
            return true;
        }
        switch (attr) {
            case "BackgroundColor":
                setBackgroundColor(value);
                break;
            case "Background":
                setBackground(value);
                break;
            case "Foreground":
                setForeground(value);
                break;
            case "Size":
                parseSize(value);
                break;
            case "X":
                setX(value);
                break;
            case "Y":
                setY(value);
                break;
            case "PivotX":
                setPivotX(value);
                break;
            case "PivotY":
                setPivotY(value);
                break;
            case "ScaleX":
                setScaleX(value);
                break;
            case "ScaleY":
                setScaleY(value);
                break;
            case "Rotation":
                setRotation(value);
                break;
            case "Visibility":
                setVisibility(value);
                break;
            case "Padding":
                setPadding(value);
                break;
            case "Margin":
                setMargin(value);
                break;
            case "LayoutGravity":
                setLayoutGravity(value);
                break;
            case "Elevation":
                setElevation(value);
                break;
            case "LayoutWeight":
                setLayoutWeight(value);
                break;
            case "OnTouch":
                setOnTouchListener(value);
                break;
            case "OnClick":
                view.setOnClickListener(new View.OnClickListener() {
                    @Override
                    public void onClick(View view) {
                        Faithdroid.triggerEventHandler(value, "");
                    }
                });
                break;
            case "Clickable":
                view.setClickable(value.equals("true"));
                break;
                // constraint
            case "Top2TopOf":
                afterConstraintFuncs.put(attr, new ConstraintInterface() {
                    @Override
                    public void addConstraint(FConstraintLayout parent, ConstraintLayout.LayoutParams lp) {
                        if (value.equals("_Parent_")) {
                            lp.topToTop = parent.view.getId();
                            return;
                        }
                        lp.topToTop = parentController.viewmap.get(Faithdroid.getVidById(value)).view.getId();
                    }
                });
                break;
            case "Top2BottomOf":
                afterConstraintFuncs.put(attr, new ConstraintInterface() {
                    @Override
                    public void addConstraint(FConstraintLayout parent, ConstraintLayout.LayoutParams lp) {
                        if (value.equals("_Parent_")) {
                            lp.bottomToBottom = parent.view.getId();
                            return;
                        }
                        lp.topToBottom = parentController.viewmap.get(Faithdroid.getVidById(value)).view.getId();
                    }
                });
                break;
            case "Bottom2BottomOf":
                afterConstraintFuncs.put(attr, new ConstraintInterface() {
                    @Override
                    public void addConstraint(FConstraintLayout parent, ConstraintLayout.LayoutParams lp) {
                        if (value.equals("_Parent_")) {
                            lp.bottomToBottom = parent.view.getId();
                            return;
                        }
                        lp.bottomToBottom = parentController.viewmap.get(Faithdroid.getVidById(value)).view.getId();
                    }
                });
                break;
            case "Bottom2TopOf":
                afterConstraintFuncs.put(attr, new ConstraintInterface() {
                    @Override
                    public void addConstraint(FConstraintLayout parent, ConstraintLayout.LayoutParams lp) {
                        if (value.equals("_Parent_")) {
                            lp.bottomToTop = parent.view.getId();
                            return;
                        }
                        lp.bottomToTop = parentController.viewmap.get(Faithdroid.getVidById(value)).view.getId();
                    }
                });
                break;
            case "Left2LeftOf":
                afterConstraintFuncs.put(attr, new ConstraintInterface() {
                    @Override
                    public void addConstraint(FConstraintLayout parent, ConstraintLayout.LayoutParams lp) {
                        if (value.equals("_Parent_")) {
                            lp.leftToLeft = parent.view.getId();
                            return;
                        }
                        lp.leftToLeft = parentController.viewmap.get(Faithdroid.getVidById(value)).view.getId();
                    }
                });
                break;
            case "Right2RightOf":
                afterConstraintFuncs.put(attr, new ConstraintInterface() {
                    @Override
                    public void addConstraint(FConstraintLayout parent, ConstraintLayout.LayoutParams lp) {
                        if (value.equals("_Parent_")) {
                            lp.rightToRight = parent.view.getId();
                            return;
                        }
                        lp.rightToRight = parentController.viewmap.get(Faithdroid.getVidById(value)).view.getId();
                    }
                });
                break;
            case "Left2RightOf":
                afterConstraintFuncs.put(attr, new ConstraintInterface() {
                    @Override
                    public void addConstraint(FConstraintLayout parent, ConstraintLayout.LayoutParams lp) {
                        if (value.equals("_Parent_")) {
                            lp.leftToRight = parent.view.getId();
                            return;
                        }
                        lp.leftToRight = parentController.viewmap.get(Faithdroid.getVidById(value)).view.getId();
                    }
                });
                break;
            case "Right2LeftOf":
                afterConstraintFuncs.put(attr, new ConstraintInterface() {
                    @Override
                    public void addConstraint(FConstraintLayout parent, ConstraintLayout.LayoutParams lp) {
                        if (value.equals("_Parent_")) {
                            lp.rightToLeft = parent.view.getId();
                            return;
                        }
                        lp.rightToLeft = parentController.viewmap.get(Faithdroid.getVidById(value)).view.getId();
                    }
                });
                break;
            case "CenterX":
                setUniversalAttr("Left2LeftOf", "_Parent_");
                setUniversalAttr("Right2RightOf", "_Parent_");
                break;
            case "CenterY":
                setUniversalAttr("Top2TopOf", "_Parent_");
                setUniversalAttr("Bottom2BottomOf", "_Parent_");
                break;
            default:
                return false;
        }
        return true;
    }
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
            setLayoutWeight("1");
        }else{
            p.width = (int) dp2pixel(parentController.activity, width);
        }
        if (height == -1) {
            p.height = ViewGroup.LayoutParams.WRAP_CONTENT;
        } else if (height == -2) {
            p.height = ViewGroup.LayoutParams.MATCH_PARENT;
            setLayoutWeight("1");
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
        Log.d(TAG, "setBackgroundColor: "+value);
        if (value==null)
            return;
        if (value.equals("#0000000")) {
            view.setBackgroundColor(Color.TRANSPARENT);
            return;
        }
        try {
            view.setBackgroundColor(Color.parseColor(value));
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
    void setBackground(String value) {
        if (value==null)
            return;
        Toolkit.file2Drawable(parentController, value, new Toolkit.OnDrawableReadyListener() {
            @Override
            public void onDrawableReady(Drawable draw) {
                if (draw == null) {
                    return;
                }
                if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.JELLY_BEAN) {
                    view.setBackground(draw);
                }else {
                    view.setBackgroundDrawable(draw);
                }
            }
        });

        if (value.equals("RippleEffect")) {
            view.setClickable(true);
        }
    }

    void setForeground(String value) {
        if (value == null) {
            return;
        }
        Toolkit.file2Drawable(parentController, value, new Toolkit.OnDrawableReadyListener() {
            @Override
            public void onDrawableReady(Drawable draw) {
                if (draw == null) {
                    return;
                }
                if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.M) {
                    view.setForeground(draw);
                }
            }
        });

        if (value.equals("RippleEffect")) {
            view.setClickable(true);
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
            float f = dp2pixel(parentController.activity, Float.parseFloat(value));
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

    void setX(String value) {
        try {
            float f = dp2pixel(parentController.activity, Float.parseFloat(value));
            view.setX(f);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
    void setY(String value) {
        try {
            float f = dp2pixel(parentController.activity, Float.parseFloat(value));
            view.setY(f);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
    void setPivotX(String value) {
        try {
            float f = dp2pixel(parentController.activity, Float.parseFloat(value));
            view.setPivotX(f);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
    void setPivotY(String value) {
        try {
            float f = dp2pixel(parentController.activity, Float.parseFloat(value));
            view.setPivotY(f);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
    void setScaleX(String value) {
        try {
            float f = dp2pixel(parentController.activity, Float.parseFloat(value));
            view.setScaleX(f);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
    void setScaleY(String value) {
        try {
            float f = dp2pixel(parentController.activity, Float.parseFloat(value));
            view.setScaleY(f);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    void setRotation(String value) {
        try {
            float f = dp2pixel(parentController.activity, Float.parseFloat(value));
            view.setRotation(f);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
    String getX(){
        return String.valueOf(pixel2dp(parentController.activity, view.getX()));
    }
    String getY(){
        return String.valueOf(pixel2dp(parentController.activity, view.getY()));
    }
    String getWidth(){
        return String.valueOf(pixel2dp(parentController.activity, view.getWidth()));
    }
    String getHeight(){
        return String.valueOf(pixel2dp(parentController.activity, view.getHeight()));
    }
    String getPivotX(){
        return String.valueOf(pixel2dp(parentController.activity, view.getPivotX()));
    }
    String getPivotY() {
        return String.valueOf(pixel2dp(parentController.activity, view.getPivotY()));
    }
    String getScaleX(){
        return String.valueOf(pixel2dp(parentController.activity, view.getScaleX()));
    }
    String getScaleY(){
        return String.valueOf(pixel2dp(parentController.activity, view.getScaleY()));
    }
    String getRotation(){
        return String.valueOf(view.getRotation());
    }

    void setOnTouchListener(final String value){
        view.setOnTouchListener(new View.OnTouchListener() {
            @Override
            public boolean onTouch(View v, MotionEvent event) {
                try {
                    String action = "";
                    switch (event.getAction()) {
                        case MotionEvent.ACTION_MOVE:
                            action = "Move";
                            break;
                        case MotionEvent.ACTION_DOWN:
                            action = "Down";
                            break;
                        case MotionEvent.ACTION_UP:
                            action = "Up";
                            break;
                    }
                    JSONObject object = new JSONObject();
                    object.put("Action", action);
                    object.put("X", pixel2dp(parentController.activity,event.getX()));
                    object.put("Y", pixel2dp(parentController.activity,event.getY()));
                    Faithdroid.triggerEventHandler(value, object.toString());
                } catch (Exception e) {
                    e.printStackTrace();
                    return false;
                }
                return true;
            }
        });
    }

}
