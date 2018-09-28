package io.github.gofaith.faithdroid.FViews;

import android.util.Log;
import android.widget.RadioGroup;

import java.util.HashMap;
import java.util.Map;

import faithdroid.Faithdroid;
import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.UIController;

public class FRadioGroup extends FView implements AttrGettable, AttrSettable {
    private RadioGroup v;
    private Map<Integer, String> idMap = new HashMap<>();
    public FRadioGroup(UIController controller) {
        parentController=controller;
        v = new RadioGroup(parentController.activity);
        view=v;
        v.setOrientation(RadioGroup.VERTICAL);
    }
    @Override
    public String getAttr(String attr) {
        String str = getUniversalAttr(attr);
        if (str != null) {
            return str;
        }
        switch (attr){
            // --------------------------------------------------
            case "Orientation":
                if (v.getOrientation()==RadioGroup.HORIZONTAL)
                    return "HORIZONTAL";
                else
                    return "VERTICAL";
            case "CheckedRadioButtonId":
                Log.d(TAG, "getAttr: checked id="+v.getCheckedRadioButtonId());
                return idMap.get(v.getCheckedRadioButtonId());
        }
        return null;
    }

    @Override
    public void setAttr(String attr, final String value) {
        if (setUniversalAttr(attr, value)) {
            return;
        }
        switch (attr){
            // ----------------------------------------------------------------------------
            case "Orientation":
                if (value.equals("HORIZONTAL")) {
                    v.setOrientation(RadioGroup.HORIZONTAL);
                } else {
                    v.setOrientation(RadioGroup.VERTICAL);
                }
                break;
            case "AddView":
                final String childVid = value;
                FView f = parentController.viewmap.get(childVid);
                RadioGroup.LayoutParams lp = new RadioGroup.LayoutParams(f.size[0], f.size[1]);
                lp.gravity = f.layoutGravity;
                lp.weight = f.layoutWeight;
                lp.leftMargin = f.margin[0];
                lp.topMargin = f.margin[1];
                lp.rightMargin = f.margin[2];
                lp.bottomMargin = f.margin[3];
                v.addView(f.view,lp);
                idMap.put(f.view.getId(),f.vID);
                break;
            case "OnCheckedChange":
                v.setOnCheckedChangeListener(new RadioGroup.OnCheckedChangeListener() {
                    @Override
                    public void onCheckedChanged(RadioGroup group, int checkedId) {
                        String checkedVid = idMap.get(checkedId);
                        Faithdroid.triggerEventHandler(value, checkedVid);
                    }
                });
                break;
        }
    }
}
