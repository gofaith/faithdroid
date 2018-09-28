package io.github.gofaith.faithdroid.FViews;

import android.graphics.drawable.Drawable;
import android.view.View;
import android.widget.ImageView;

import faithdroid.Faithdroid;
import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.Toolkit;
import io.github.gofaith.faithdroid.UI.UIController;

public class FImageView extends FView implements AttrSettable, AttrGettable {
    public ImageView v;

    public FImageView(UIController controller) {
        parentController=controller;
        v = new ImageView(parentController.activity);
        view=v;
    }
    @Override
    public String getAttr(String attr) {
        String str = getUniversalAttr(attr);
        if (str != null) {
            return str;
        }
        switch (attr){
            // ------------------------------------------
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
            case "Src":
                Drawable drawable = Toolkit.file2Drawable(parentController.activity, value);
                if (drawable == null) {
                    return;
                }
                v.setImageDrawable(drawable);
                break;
        }
    }
}
