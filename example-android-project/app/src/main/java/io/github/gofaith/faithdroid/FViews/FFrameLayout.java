package io.github.gofaith.faithdroid.FViews;

import android.view.View;
import android.widget.FrameLayout;
import android.widget.LinearLayout;

import org.json.JSONArray;
import org.json.JSONTokener;

import faithdroid.Faithdroid;
import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.UIController;

public class FFrameLayout extends FView implements AttrGettable, AttrSettable {
    public FrameLayout v;

    public FFrameLayout(UIController controller) {
        parentController = controller;
        v = new FrameLayout(parentController.activity);
        view=v;
    }
    @Override
    public String getAttr(String attr) {
        String str = getUniversalAttr(attr);
        if (str != null) {
            return str;
        }
        switch (attr){
            // --------------------------------------------------
        }
        return null;
    }

    @Override
    public void setAttr(String attr, final String value) {
        if (setUniversalAttr(attr, value)) {
            return;
        }
        switch (attr){
            // ----------------------------------------------------------------------------
            case "AddView":
                final String childVid = value;
                FView f = parentController.viewmap.get(childVid);
                v.addView(f.view,parseLP(f));
                break;
            case "AddViewAt":
                try {
                    JSONArray array = (JSONArray) (new JSONTokener(value).nextValue());
                    int pos = Integer.parseInt(array.getString(0));
                    String vid = array.getString(1);
                    FView f1 = parentController.viewmap.get(vid);
                    v.addView(f1.view,pos,parseLP(f1));
                } catch (Exception e) {
                    e.printStackTrace();
                }
                break;
            case "OnClick":
                v.setOnClickListener(new View.OnClickListener() {
                    @Override
                    public void onClick(View view) {
                        Faithdroid.triggerEventHandler(value, "");
                    }
                });
                break;
        }
    }


    public static FrameLayout.LayoutParams parseLP(FView f) {
        FrameLayout.LayoutParams lp = new FrameLayout.LayoutParams(f.size[0], f.size[1]);
        lp.gravity=f.layoutGravity;
        lp.leftMargin = f.margin[0];
        lp.topMargin = f.margin[1];
        lp.rightMargin = f.margin[2];
        lp.bottomMargin = f.margin[3];
        return lp;
    }
}
