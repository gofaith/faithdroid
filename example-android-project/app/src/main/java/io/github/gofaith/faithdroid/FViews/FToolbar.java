package io.github.gofaith.faithdroid.FViews;

import android.graphics.drawable.Drawable;
import android.support.v7.widget.Toolbar;
import android.view.View;

import faithdroid.Faithdroid;
import io.github.gofaith.faithdroid.R;
import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.Toolkit;
import io.github.gofaith.faithdroid.UI.UIController;

public class FToolbar extends FView implements AttrSettable,AttrGettable {
    private Toolbar v;
    public FToolbar(UIController controller) {
        parentController =controller;
        v = (Toolbar) parentController.activity.getLayoutInflater().inflate(R.layout.my_toolbar, parentController.rootView, false);
        view=v;
        setElevation("4");
        parentController.activity.setSupportActionBar(v);
        parseSize("[-2,-1]");
    }
    @Override
    public String getAttr(String attr) {
        String str = getUniversalAttr(attr);
        if (str != null) {
            return str;
        }
        switch (attr){
                //----------------------------------------------
        }
        return "";
    }

    @Override
    public void setAttr(String attr, final String value) {
        if (setUniversalAttr(attr, value)) {
            return;
        }
        switch (attr){
            // -------------------------------------------------------------------
            case "Title":
                parentController.activity.getSupportActionBar().setTitle(value);
                break;
            case "SubTitle":
                parentController.activity.getSupportActionBar().setSubtitle(value);
                break;
            case "Menu":
                parentController.optionMenus=value;
                break;
            case "BackNavigation":
                if (value.equals("true")) {
                    v.setNavigationOnClickListener(new View.OnClickListener() {
                        @Override
                        public void onClick(View v) {
                            parentController.activity.onBackPressed();
                        }
                    });
                }
                parentController.activity.getSupportActionBar().setDisplayHomeAsUpEnabled(value.equals("true"));
                break;
            case "NavigationIcon":
                Toolkit.file2Drawable(parentController, value, new Toolkit.OnDrawableReadyListener() {
                    @Override
                    public void onDrawableReady(Drawable drawable) {
                        v.setNavigationIcon(drawable);
                    }
                });
                break;
            case "NavigationOnClick":
                v.setNavigationOnClickListener(new View.OnClickListener() {
                    @Override
                    public void onClick(View v) {
                        Faithdroid.triggerEventHandler(value, "");
                    }
                });
                break;
        }
    }
}
