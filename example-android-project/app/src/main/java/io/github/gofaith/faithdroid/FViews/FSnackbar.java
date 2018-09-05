package io.github.gofaith.faithdroid.FViews;

import android.support.design.widget.Snackbar;
import android.view.View;

import org.json.JSONArray;
import org.json.JSONTokener;

import faithdroid.Faithdroid;
import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.UIController;

public class FSnackbar extends FView implements AttrSettable, AttrGettable {
    public Snackbar v;

    public FSnackbar(UIController controller, View anchorView) {
        parentController=controller;
        v = Snackbar.make(anchorView, "", Snackbar.LENGTH_SHORT);
    }
    @Override
    public String getAttr(String attr) {
        return null;
    }

    @Override
    public void setAttr(String attr, String value) {
        if (value == null) {
            return;
        }
        switch (attr) {
            case "Text":
                v.setText(value);
                break;
            case "Action":
                try {
                    final JSONArray array = (JSONArray) (new JSONTokener(value).nextValue());
                    final String fnId=array.getString(1);
                    v.setAction(array.getString(0),
                            new View.OnClickListener() {
                                @Override
                                public void onClick(View v) {
                                    Faithdroid.triggerEventHandler(fnId, "");
                                }
                            });
                } catch (Exception e) {
                    e.printStackTrace();
                }
                break;
            case "Show":
                v.show();
                break;
        }
    }
}
