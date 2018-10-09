package io.github.gofaith.faithdroid.FViews;

import android.graphics.drawable.Drawable;
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
        String str = getUniversalAttr(attr);
        if (str != null) {
            return str;
        }
        switch (attr){
                //------------------
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
            case "Icon":
                Toolkit.file2Drawable(parentController, value, new Toolkit.OnDrawableReadyListener() {
                    @Override
                    public void onDrawableReady(Drawable drawable) {
                        v.setImageDrawable(drawable);
                    }
                });
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
