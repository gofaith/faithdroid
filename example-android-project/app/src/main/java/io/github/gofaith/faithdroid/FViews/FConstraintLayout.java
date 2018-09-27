package io.github.gofaith.faithdroid.FViews;

import android.support.constraint.ConstraintLayout;
import android.support.v4.view.ViewCompat;
import android.widget.LinearLayout;

import java.util.HashMap;
import java.util.Map;

import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.UIController;

public class FConstraintLayout extends FView implements AttrSettable, AttrGettable {
    private ConstraintLayout v;
    private Map<String, Integer> vidToId = new HashMap<>();
    public FConstraintLayout(UIController controller) {
        parentController = controller;
        v = new ConstraintLayout(parentController.activity);
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
            // --------------------------------------------------
        }
        return "";
    }

    @Override
    public void setAttr(String attr, String value) {
        if (setUniversalAttr(attr, value)) {
            return;
        }
        switch (attr) {
            // ----------------------------------------------------------------------------
            case "AddView":
                FView f = parentController.viewmap.get(value);
                v.addView(f.view,parseLP(f));
                break;
            case "Append":
                break;
        }
    }

    private ConstraintLayout.LayoutParams parseLP(FView f) {
        ConstraintLayout.LayoutParams lp = new ConstraintLayout.LayoutParams(f.size[0], f.size[1]);

        lp.leftMargin = f.margin[0];
        lp.topMargin = f.margin[1];
        lp.rightMargin = f.margin[2];
        lp.bottomMargin = f.margin[3];
        if (f.topToTopOf != 0) {
            lp.topToTop = checkIfParent(f.topToTopOf);
        }
        if (f.topToBottomOf != 0) {
            lp.topToBottom = checkIfParent(f.topToBottomOf);
        }
        if (f.bottomToBottomOf != 0) {
            lp.bottomToBottom = checkIfParent(f.bottomToBottomOf);
        }
        if (f.bottomToTopOf != 0) {
            lp.bottomToTop = checkIfParent(f.bottomToTopOf);
        }
        if (f.leftToLeftOf != 0) {
            lp.leftToLeft = checkIfParent(f.leftToLeftOf);
        }
        if (f.leftToRightOf != 0) {
            lp.leftToRight = checkIfParent(f.leftToRightOf);
        }
        if (f.rightToRightOf != 0) {
            lp.rightToRight = checkIfParent(f.rightToRightOf);
        }
        if (f.rightToLeftOf != 0) {
            lp.rightToLeft = checkIfParent(f.rightToLeftOf);
        }
        return lp;
    }

    private int checkIfParent(int i) {
        if (i == -24) {
            return v.getId();
        }
        return i;
    }
}
