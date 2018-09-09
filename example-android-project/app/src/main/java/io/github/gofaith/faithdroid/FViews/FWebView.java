package io.github.gofaith.faithdroid.FViews;

import android.webkit.WebView;

import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.UIController;

public class FWebView extends FView implements AttrGettable, AttrSettable {
    public WebView v;

    public FWebView(UIController controller) {
        this.parentController=controller;
        v = new WebView(parentController.activity);
        view=v;
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
        }
        return null;
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
            // -------------------------------------------------------------------
            case "Uri":
                v.loadUrl(value);
                break;
            case "Focusable":
                if (value.equals("true")) {
                    v.setFocusable(true);
                }else {
                    v.setFocusable(false);
                }
                break;
        }
    }
}
