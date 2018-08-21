package io.github.gofaith.faithdroid.FViews;

import android.support.v7.app.AppCompatActivity;
import android.widget.Button;

import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.UIController;

public class FButton extends FView implements AttrSettable,AttrGettable {
    private Button v;
    public FButton(UIController c){
        parrentController=c;
        v = new Button(parrentController.activity);
        view=v;
    }
    @Override
    public String getAttr(String attr) {
        switch (attr) {
            case "Text":
                return v.getText().toString();
            case "Enabled":
                return String.valueOf(v.isEnabled());
        }
        return "";
    }

    @Override
    public void setAttr(String attr, String value) {
        if (value==null)
            return;
        switch (attr) {
            case "Size":
                parseSize(parrentController.activity,v,value);
                break;
            case "Text":
                v.setText(value);
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
