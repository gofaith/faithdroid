package io.github.gofaith.faithdroid.FViews;

import android.widget.Space;

import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.UIController;

public class FSpace extends FView implements AttrSettable, AttrGettable {
    public Space v;

    public FSpace(UIController controller) {
        parentController=controller;
        v = new Space(controller.activity);
        view=v;
        layoutWeight=1;
    }
    @Override
    public String getAttr(String attr) {
        if (attr == null) {
            return "";
        }
        switch (attr) {
            case "Visibility":
                return getVisibility();
        }
        return "";
    }

    @Override
    public void setAttr(String attr, String value) {
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
        }
    }
}
