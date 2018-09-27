package io.github.gofaith.faithdroid.FViews;

import android.graphics.drawable.GradientDrawable;
import android.support.v4.view.ViewCompat;
import android.view.View;
import android.view.ViewGroup;
import android.widget.HorizontalScrollView;
import android.widget.LinearLayout;
import android.widget.ScrollView;

import faithdroid.Faithdroid;
import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.UIController;

public class FHScrollView extends FView implements AttrSettable,AttrGettable {
    public HorizontalScrollView v;
    private LinearLayout linearLayout;
    public FHScrollView(UIController controller) {
        parentController = controller;
        v = new HorizontalScrollView(parentController.activity);
        v.setId(ViewCompat.generateViewId());
        view=v;
        linearLayout = new LinearLayout(parentController.activity);
        linearLayout.setId(ViewCompat.generateViewId());
        linearLayout.setOrientation(LinearLayout.HORIZONTAL);
        v.addView(linearLayout, ViewGroup.LayoutParams.WRAP_CONTENT, ViewGroup.LayoutParams.MATCH_PARENT);
    }
    @Override
    public String getAttr(String attr) {
        String str = getUniversalAttr(attr);
        if (str != null) {
            return str;
        }
        switch (attr) {
            // ------------------------------------------
        }
        return "";
    }

    @Override
    public void setAttr(String attr, final String value) {
        if (setUniversalAttr(attr, value)) {
            return;
        }
        switch (attr) {
            // -------------------------------------------------------------------

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
                linearLayout.addView(f.view,lp);
                break;
        }
    }
}
