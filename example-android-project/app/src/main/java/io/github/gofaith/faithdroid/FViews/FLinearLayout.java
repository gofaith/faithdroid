package io.github.gofaith.faithdroid.FViews;

import android.graphics.Color;
import android.util.Log;
import android.view.View;
import android.widget.LinearLayout;

import faithdroid.Faithdroid;
import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.UIController;

public class FLinearLayout extends FView implements AttrSettable,AttrGettable {
    private LinearLayout v;
    public FLinearLayout(UIController c){
        parentController =c;
        v = new LinearLayout(parentController.activity);
        view=v;
        v.setOrientation(LinearLayout.VERTICAL);
    }
    @Override
    public String getAttr(String attr) {
        if(attr==null)
            return "";
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
                // --------------------------------------------------
            case "Orientation":
                if (v.getOrientation()==LinearLayout.HORIZONTAL)
                    return "HORIZONTAL";
                else
                    return "VERTICAL";
        }
        return "";
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
                // ----------------------------------------------------------------------------
            case "Orientation":
                if (value.equals("HORIZONTAL")) {
                    v.setOrientation(LinearLayout.HORIZONTAL);
                }else {
                    v.setOrientation(LinearLayout.VERTICAL);
                }
                break;
            case "AddView":
                final String childVid=value;
                FView f=parentController.viewmap.get(childVid);
                LinearLayout.LayoutParams lp = new LinearLayout.LayoutParams(f.size[0], f.size[1]);
                lp.gravity=f.layoutGravity;
                lp.weight=f.layoutWeight;
                lp.leftMargin = f.margin[0];
                lp.topMargin = f.margin[1];
                lp.rightMargin = f.margin[2];
                lp.bottomMargin = f.margin[3];
                v.addView(f.view,lp);
                break;
            case "OnClick":
                v.setOnClickListener(new View.OnClickListener() {
                    @Override
                    public void onClick(View view) {
                        Faithdroid.triggerEventHandler(value,"");
                    }
                });
                break;
        }
    }
}
