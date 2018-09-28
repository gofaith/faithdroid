package io.github.gofaith.faithdroid.FViews;

import android.graphics.Color;
import android.view.View;
import android.widget.CheckBox;
import android.widget.CompoundButton;

import faithdroid.Faithdroid;
import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.UIController;

public class FCheckBox extends FView implements AttrGettable, AttrSettable {
    private CheckBox v;

    public FCheckBox(UIController controller) {
        parentController = controller;
        v = new CheckBox(parentController.activity);
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
            case "Checked":
                return String.valueOf(v.isChecked());
            case "Enabled":
                return String.valueOf(v.isEnabled());
            case "Text":
                return v.getText().toString();
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
            case "OnCheckedChangeListener":
                v.setOnCheckedChangeListener(new CompoundButton.OnCheckedChangeListener() {
                    @Override
                    public void onCheckedChanged(CompoundButton buttonView, boolean isChecked) {
                        Faithdroid.triggerEventHandler(value, String.valueOf(isChecked));
                    }
                });
                break;
        }
    }
}
