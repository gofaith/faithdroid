package io.github.gofaith.faithdroid.FViews;

import android.view.View;
import android.widget.Space;

import faithdroid.Faithdroid;
import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.UIController;

public class FSpace extends FView implements AttrSettable, AttrGettable {
    public Space v;

    public FSpace(UIController controller) {
        parentController=controller;
        v = new Space(controller.activity);
        view=v;
        layoutWeight=1;
    }
    @Override
    public String getAttr(String attr) {
        String str = getUniversalAttr(attr);
        if (str != null) {
            return str;
        }
        switch (attr){
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
        }
    }
}
