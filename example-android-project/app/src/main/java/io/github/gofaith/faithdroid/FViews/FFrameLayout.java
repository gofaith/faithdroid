package io.github.gofaith.faithdroid.FViews;

import android.view.View;
import android.widget.FrameLayout;
import android.widget.LinearLayout;

import faithdroid.Faithdroid;
import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.UIController;

public class FFrameLayout extends FView implements AttrGettable, AttrSettable {
    public FrameLayout v;

    public FFrameLayout(UIController controller) {
        parentController = controller;
        v = new FrameLayout(parentController.activity);
        view=v;
    }
    @Override
    public String getAttr(String attr) {
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
            // --------------------------------------------------
        }
        return null;
    }

    @Override
    public void setAttr(String attr, final String value) {

        if (attr == null || value == null) {
            return;
        }
        switch (attr) {
            case "BackgroundColor":
                setBackgroundColor(value);
                break;
            case "Background":
                setBackground(value);
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
            // ----------------------------------------------------------------------------
            case "AddView":
                final String childVid = value;
                FView f = parentController.viewmap.get(childVid);
                addToframeLayout(v,f);
                break;
            case "OnClick":
                v.setOnClickListener(new View.OnClickListener() {
                    @Override
                    public void onClick(View view) {
                        Faithdroid.triggerEventHandler(value, "");
                    }
                });
                break;
        }
    }

    public static void addToframeLayout(FrameLayout frameLayout, FView f) {
        FrameLayout.LayoutParams lp = new FrameLayout.LayoutParams(f.size[0], f.size[1]);
        lp.gravity=f.layoutGravity;
        lp.leftMargin = f.margin[0];
        lp.topMargin = f.margin[1];
        lp.rightMargin = f.margin[2];
        lp.bottomMargin = f.margin[3];
        frameLayout.addView(f.view,lp);
    }
}
