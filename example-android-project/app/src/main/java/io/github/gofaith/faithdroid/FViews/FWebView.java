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
