package io.github.gofaith.faithdroid.FViews;

import android.graphics.Color;
import android.support.v4.view.ViewCompat;
import android.util.Log;
import android.view.View;
import android.widget.LinearLayout;

import faithdroid.Faithdroid;
import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.UIController;

public class FLinearLayout extends FView implements AttrSettable,AttrGettable {
    private LinearLayout v;
    public FLinearLayout(UIController c){
        parentController =c;
        v = new LinearLayout(parentController.activity);
        v.setId(ViewCompat.generateViewId());
        view=v;
        v.setOrientation(LinearLayout.VERTICAL);
    }
    @Override
    public String getAttr(String attr) {
        String str = getUniversalAttr(attr);
        if (str != null) {
            return str;
        }
        switch (attr) {
                // --------------------------------------------------
            case "Orientation":
                if (v.getOrientation()==LinearLayout.HORIZONTAL)
                    return "HORIZONTAL";
                else
                    return "VERTICAL";
        }
        return "";
    }

    @Override
    public void setAttr(String attr, final String value) {
        if (setUniversalAttr(attr, value)) {
            return;
        }
        switch (attr) {
                // ----------------------------------------------------------------------------
            case "Orientation":
                if (value.equals("HORIZONTAL")) {
                    v.setOrientation(LinearLayout.HORIZONTAL);
                }else {
                    v.setOrientation(LinearLayout.VERTICAL);
                }
                break;
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
                v.addView(f.view,lp);
                break;
            case "OnClick":
                v.setOnClickListener(new View.OnClickListener() {
                    @Override
                    public void onClick(View view) {
                        Faithdroid.triggerEventHandler(value,"");
                    }
                });
                break;
        }
    }
}
