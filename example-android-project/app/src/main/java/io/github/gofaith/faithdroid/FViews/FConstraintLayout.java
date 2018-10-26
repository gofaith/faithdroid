package io.github.gofaith.faithdroid.FViews;

import android.support.constraint.ConstraintLayout;
import android.util.Log;

import java.util.Map;

import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.UIController;

public class FConstraintLayout extends FView implements AttrSettable,AttrGettable {
    public ConstraintLayout v;
    public FConstraintLayout(UIController controller) {
        parentController = controller;
        v = new ConstraintLayout(parentController.activity);
        view = v;
    }
    @Override
    public String getAttr(String attr) {
        String str = getUniversalAttr(attr);
        if (str != null) {
            return str;
        }
        switch (attr) {

        }
        return null;
    }

    @Override
    public void setAttr(String attr, String value) {
        if (setUniversalAttr(attr, value)) {
            return;
        }
        switch (attr) {
            case "Append":
                String[] vids = value.split(",");
                for (int i=0;i<vids.length;i++) {
                    FView f= parentController.viewmap.get(vids[i]);
                    ConstraintLayout.LayoutParams lp = new ConstraintLayout.LayoutParams(f.size[0], f.size[1]);
                    for (Map.Entry<String, ConstraintInterface> entry : f.afterConstraintFuncs.entrySet()) {
                        entry.getValue().addConstraint(this,lp);
                    }
                    v.addView(parentController.viewmap.get(vids[i]).view,lp);
                }
                break;
        }
    }
}
