package io.github.gofaith.faithdroid.FViews;

import android.widget.ProgressBar;

import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.UIController;

public class FProgressBar extends FView implements AttrSettable, AttrGettable {
    public ProgressBar v;
    public FProgressBar(UIController uiController){
        parentController = uiController;
        v = new ProgressBar(parentController.activity);
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
        }
        return "";
    }

    @Override
    public void setAttr(String attr, String value) {
        if (setUniversalAttr(attr, value)) {
            return;
        }
        switch (attr) {
            // -------------------------------------------------------------------
        }
    }
}
