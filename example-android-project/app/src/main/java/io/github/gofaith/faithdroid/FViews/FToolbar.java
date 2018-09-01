package io.github.gofaith.faithdroid.FViews;

import android.graphics.Color;
import android.support.v7.widget.Toolbar;

import io.github.gofaith.faithdroid.R;
import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.UIController;

public class FToolbar extends FView implements AttrSettable,AttrGettable {
    private Toolbar v;

    public FToolbar(UIController controller) {
        parrentController=controller;
        v = (Toolbar) parrentController.activity.getLayoutInflater().inflate(R.layout.my_toolbar, parrentController.rootView, false);
        view=v;
        parrentController.activity.setSupportActionBar(v);
    }
    @Override
    public String getAttr(String attr) {
        switch (attr) {
            case "Visibility":
                return getVisibility();
                //----------------------------------------------
        }
        return "";
    }

    @Override
    public void setAttr(String attr, String value) {
        if (value == null) {
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
                parseSize(parrentController.activity,v,value);
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
            case "Title":
                parrentController.activity.getSupportActionBar().setTitle(value);
                break;
            case "SubTitle":
                parrentController.activity.getSupportActionBar().setSubtitle(value);
                break;
        }
    }
}
