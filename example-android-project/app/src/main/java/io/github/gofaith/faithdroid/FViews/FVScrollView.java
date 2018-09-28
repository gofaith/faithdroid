package io.github.gofaith.faithdroid.FViews;

import android.graphics.drawable.GradientDrawable;
import android.view.View;
import android.view.ViewGroup;
import android.widget.HorizontalScrollView;
import android.widget.LinearLayout;
import android.widget.ScrollView;

import org.json.JSONArray;
import org.json.JSONTokener;

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
        view=v;
        linearLayout = new LinearLayout(parentController.activity);
        linearLayout.setOrientation(LinearLayout.VERTICAL);
        v.addView(linearLayout, ViewGroup.LayoutParams.MATCH_PARENT, ViewGroup.LayoutParams.WRAP_CONTENT);
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

            case "AddView":
                final String childVid=value;
                FView f=parentController.viewmap.get(childVid);
                linearLayout.addView(f.view,FLinearLayout.parseLP(f));
                break;
            case "AddViewAt":
                try {
                    JSONArray array = (JSONArray) (new JSONTokener(value).nextValue());
                    int pos = Integer.parseInt(array.getString(0));
                    String vid = array.getString(1);
                    FView f1 = parentController.viewmap.get(vid);
                    v.addView(f1.view,pos,FLinearLayout.parseLP(f1));
                } catch (Exception e) {
                    e.printStackTrace();
                }
                break;
        }
    }
}
