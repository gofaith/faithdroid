package io.github.gofaith.faithdroid.FViews;

import android.graphics.Color;
import android.support.v4.view.ViewCompat;
import android.text.InputType;
import android.view.View;
import android.widget.EditText;

import faithdroid.Faithdroid;
import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.UIController;

public class FEditText extends FView implements AttrGettable, AttrSettable {
    public EditText v;

    public FEditText(UIController controller) {
        parentController=controller;
        v = new EditText(parentController.activity);
        v.setId(ViewCompat.generateViewId());
        view=v;
    }
    @Override
    public String getAttr(String attr) {
        String str = getUniversalAttr(attr);
        if (str != null) {
            return str;
        }
        switch (attr) {
            // ------------------------------------------
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
        switch (attr) {
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
                    v.setTextSize(dp2pixel(parentController.activity, Float.valueOf(value)));
                } catch (Exception e) {
                    e.printStackTrace();
                    return;
                }
                break;
            case "InputType":
                int it;
                if (value.equals("Text")) {
                    it=InputType.TYPE_CLASS_TEXT;
                } else if (value.equals("Number")) {
                    it = InputType.TYPE_CLASS_NUMBER;
                } else if (value.equals("Password")) {
                    it=InputType.TYPE_CLASS_TEXT|InputType.TYPE_TEXT_VARIATION_WEB_PASSWORD;
                } else  {
                    it = InputType.TYPE_TEXT_VARIATION_VISIBLE_PASSWORD;
                }
                v.setInputType(it);
                break;
        }
    }
}
