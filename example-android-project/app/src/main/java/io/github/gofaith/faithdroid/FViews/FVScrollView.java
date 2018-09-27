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

public class FVScrollView extends FView implements AttrSettable,AttrGettable {
    public ScrollView v;
    private LinearLayout linearLayout;
    public FVScrollView(UIController controller) {
        parentController = controller;
        v = new ScrollView(parentController.activity);
        v.setId(ViewCompat.generateViewId());
        view=v;
        linearLayout = new LinearLayout(parentController.activity);
        linearLayout.setOrientation(LinearLayout.VERTICAL);
        linearLayout.setId(ViewCompat.generateViewId());
        v.addView(linearLayout, ViewGroup.LayoutParams.MATCH_PARENT, ViewGroup.LayoutParams.WRAP_CONTENT);
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
