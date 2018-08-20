package io.github.gofaith.faithdroid.UI;

import android.support.v7.app.AppCompatActivity;
import android.widget.FrameLayout;
import android.widget.Space;

import java.util.HashMap;

import faithdroid.UIInterface;
import io.github.gofaith.faithdroid.FViews.FButton;
import io.github.gofaith.faithdroid.FViews.FLinearLayout;
import io.github.gofaith.faithdroid.FViews.FTextView;
import io.github.gofaith.faithdroid.FViews.FView;

public class UIController implements UIInterface {
    public AppCompatActivity activity;
    public HashMap<String, FView> viewmap = new HashMap<>();
    public FrameLayout rootView;
    public UIController(AppCompatActivity a, FrameLayout main_ctn) {
        this.activity=a;
        this.rootView =main_ctn;
    }

    @Override
    public void newView(String vName, String vID) {
        FView v = null;
        v.className =vName;
        v.vID=vID;
        switch (vName) {
            case "LinearLayout":
                FLinearLayout fLinearLayout = new FLinearLayout(this);
                viewmap.put(vID, fLinearLayout);
                break;
            case "TextView":
                FTextView textView = new FTextView(this);
                viewmap.put(vID , textView);
                break;
            case "Button":
                FButton button = new FButton(this);
                viewmap.put(vID, button);
                break;
            default:
                Space space = new Space(activity);
                v.view=space;
                break;
        }
        viewmap.put(vID, v);
    }
    @Override
    public String viewGetAttr(String vID, String attr) {
        FView v = viewmap.get(vID);
        return ((AttrGettable) v.view).getAttr(attr);
    }

    @Override
    public void viewSetAttr(String vID, String attr, String value) {
        FView v = viewmap.get(vID);
        ((AttrSettable)(v.view)).setAttr(attr,value);
    }

    @Override
    public void showOnRootView(String vID) {
        FView v = viewmap.get(vID);
        rootView.addView(v.view);
    }
}
