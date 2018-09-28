package io.github.gofaith.faithdroid.FViews;

import android.util.Log;
import android.widget.RadioGroup;

import java.util.HashMap;
import java.util.Map;

import faithdroid.Faithdroid;
import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.UIController;

public class FRadioGroup extends FView implements AttrGettable, AttrSettable {
    private RadioGroup v;
    private Map<Integer, String> idMap = new HashMap<>();
    public FRadioGroup(UIController controller) {
        parentController=controller;
        v = new RadioGroup(parentController.activity);
        view=v;
        v.setOrientation(RadioGroup.VERTICAL);
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
            case "Orientation":
                if (v.getOrientation()==RadioGroup.HORIZONTAL)
                    return "HORIZONTAL";
                else
                    return "VERTICAL";
            case "CheckedRadioButtonId":
                Log.d(TAG, "getAttr: checked id="+v.getCheckedRadioButtonId());
                return idMap.get(v.getCheckedRadioButtonId());
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
            // ----------------------------------------------------------------------------
            case "Orientation":
                if (value.equals("HORIZONTAL")) {
                    v.setOrientation(RadioGroup.HORIZONTAL);
                } else {
                    v.setOrientation(RadioGroup.VERTICAL);
                }
                break;
            case "AddView":
                final String childVid = value;
                FView f = parentController.viewmap.get(childVid);
                RadioGroup.LayoutParams lp = new RadioGroup.LayoutParams(f.size[0], f.size[1]);
                lp.gravity = f.layoutGravity;
                lp.weight = f.layoutWeight;
                lp.leftMargin = f.margin[0];
                lp.topMargin = f.margin[1];
                lp.rightMargin = f.margin[2];
                lp.bottomMargin = f.margin[3];
                v.addView(f.view,lp);
                idMap.put(f.view.getId(),f.vID);
                break;
            case "OnCheckedChange":
                v.setOnCheckedChangeListener(new RadioGroup.OnCheckedChangeListener() {
                    @Override
                    public void onCheckedChanged(RadioGroup group, int checkedId) {
                        String checkedVid = idMap.get(checkedId);
                        Faithdroid.triggerEventHandler(value, checkedVid);
                    }
                });
                break;
        }
    }
}
