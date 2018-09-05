package io.github.gofaith.faithdroid.FViews;

import android.widget.FrameLayout;

import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;

public class FFrameLayout extends FView implements AttrGettable, AttrSettable {
    public FrameLayout v;

    @Override
    public String getAttr(String attr) {
        return null;
    }

    @Override
    public void setAttr(String attr, String value) {

    }

    public static void addToframeLayout(FrameLayout frameLayout, FView f) {
        FrameLayout.LayoutParams lp = new FrameLayout.LayoutParams(f.size[0], f.size[1]);
        lp.gravity=f.layoutGravity;
        lp.leftMargin = f.margin[0];
        lp.topMargin = f.margin[1];
        lp.rightMargin = f.margin[2];
        lp.bottomMargin = f.margin[3];
        frameLayout.addView(f.view,lp);
    }
}
