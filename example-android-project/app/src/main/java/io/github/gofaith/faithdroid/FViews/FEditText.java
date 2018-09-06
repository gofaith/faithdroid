package io.github.gofaith.faithdroid.FViews;

import android.graphics.Color;
import android.text.InputType;
import android.widget.EditText;

import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.UIController;

public class FEditText extends FView implements AttrGettable, AttrSettable {
    public EditText v;

    public FEditText(UIController controller) {
        parentController=controller;
        v = new EditText(parentController.activity);
        view=v;
    }
    @Override
    public String getAttr(String attr) {
        switch (attr) {
            case "Visibility":
                return getVisibility();
            // ------------------------------------------
            case "Text":
                return v.getText().toString();
        }
        return "";
    }

    @Override
    public void setAttr(String attr, String value) {
        if (value==null)
            return;
        switch (attr) {
            case "BackgroundColor":
                setBackgroundColor(value);
                break;
            case "Background":
                setBackground(value);
                break;
            case "Size":
                parseSize(value);
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
            case "LayoutWeight":
                setLayoutWeight(value);
                break;
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
