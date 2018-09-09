package io.github.gofaith.faithdroid.FViews;

import android.os.Build;
import android.support.design.widget.FloatingActionButton;
import android.view.View;
import android.view.ViewGroup;

import faithdroid.Faithdroid;
import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.Toolkit;
import io.github.gofaith.faithdroid.UI.UIController;

public class FFab extends  FView implements AttrGettable, AttrSettable {
    public FloatingActionButton v;

    public FFab(UIController controller) {
        parentController=controller;
        v = new FloatingActionButton(parentController.activity);
        view=v;
        setElevation("8");
    }
    @Override
    public String getAttr(String attr) {
        switch (attr) {
            case "Visibility":
                return getVisibility();
            case "X":
                return String.valueOf(v.getX());
            case "Y":
                return String.valueOf(v.getY());
            case "PivotX":
                return String.valueOf(v.getPivotX());
            case "PivotY":
                return String.valueOf(v.getPivotY());
            case "ScaleX":
                return String.valueOf(v.getScaleX());
            case "ScaleY":
                return String.valueOf(v.getScaleY());
            case "Rotation":
                return String.valueOf(v.getRotation());
                //------------------
        }
        return "";
    }
    @Override
    public void setAttr(String attr, final String value) {
        if (value==null)
            return;
        switch (attr) {
            case "BackgroundColor":
                setBackgroundColor(value);
                break;
            case "Background":
                setBackground(value);
                break;
            case "Size":
                parseSize( value);
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
            // -------------------------------------------------------------------
            case "Icon":
                v.setImageDrawable(Toolkit.file2Drawable(parentController.activity,value));
                break;
            case "OnClick":
                v.setOnClickListener(new View.OnClickListener() {
                    @Override
                    public void onClick(View v) {
                        Faithdroid.triggerEventHandler(value, "");
                    }
                });
                break;
        }
    }
}
