package io.github.gofaith.faithdroid.FViews;

import android.graphics.Color;
import android.support.v4.view.ViewCompat;
import android.view.View;
import android.widget.RadioButton;

import faithdroid.Faithdroid;
import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.UIController;

public class FRadioButton extends FView implements AttrSettable, AttrGettable {
    private RadioButton v;

    public FRadioButton(UIController controller) {
        parentController = controller;
        v = new RadioButton(parentController.activity);
        v.setId(ViewCompat.generateViewId());
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
            case "Text":
                return v.getText().toString();
            case "Enabled":
                return String.valueOf(v.isEnabled());
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
            case "Text":
                v.setText(value);
                break;
            case "TextColor":
                try {
                    v.setTextColor(Color.parseColor(value));
                } catch (Exception e) {
                    e.printStackTrace();
                }
                break;
            case "TextSize":
                try {
                    v.setTextSize(dp2pixel(parentController.activity,Float.valueOf(value)));
                } catch (Exception e) {
                    e.printStackTrace();
                    return;
                }
                break;
            case "Enabled":
                if (value.equals("true")) {
                    v.setEnabled(true);
                } else {
                    v.setEnabled(false);
                }
                break;
        }
    }
}
