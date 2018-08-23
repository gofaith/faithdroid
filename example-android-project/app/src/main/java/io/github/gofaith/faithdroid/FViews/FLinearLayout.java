package io.github.gofaith.faithdroid.FViews;

import android.support.v7.app.AppCompatActivity;
import android.view.View;
import android.widget.LinearLayout;

import faithdroid.Faithdroid;
import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.UIController;

public class FLinearLayout extends FView implements AttrSettable,AttrGettable {
    private LinearLayout v;
    public FLinearLayout(UIController c){
        parrentController=c;
        v = new LinearLayout(parrentController.activity);
        view=v;
        v.setOrientation(LinearLayout.VERTICAL);
    }
    @Override
    public String getAttr(String attr) {
        if(attr==null)
            return "";
        switch (attr) {
            case "Visibility":
                return getVisibility();
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
        if (attr == null || value == null) {
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
                parseSize(parrentController.activity, v, value);
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
                v.addView(parrentController.viewmap.get(childVid).view);
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
