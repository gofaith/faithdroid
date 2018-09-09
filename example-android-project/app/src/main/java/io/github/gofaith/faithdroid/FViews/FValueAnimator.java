package io.github.gofaith.faithdroid.FViews;

import android.animation.ValueAnimator;
import android.util.Log;

import org.json.JSONArray;
import org.json.JSONTokener;

import java.util.ArrayList;
import java.util.List;

import faithdroid.Faithdroid;
import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.UIController;

public class FValueAnimator extends FView implements AttrGettable, AttrSettable {
    public ValueAnimator v;

    public FValueAnimator(UIController uiController) {
        parentController=uiController;
        v=new ValueAnimator();
    }
    @Override
    public String getAttr(String attr) {
        return null;
    }

    @Override
    public void setAttr(String attr, final String value) {
        if (attr == null || value == null) {
            return;
        }
        Log.d(TAG, "setAttr: "+attr+","+value);
        switch (attr) {
            case "OfFloat":
                try {
                    JSONArray array = (JSONArray) (new JSONTokener(value).nextValue());
                    float[] vs = new float[array.length()];
                    for (int i=0;i<array.length();i++){
                        vs[i] = (float) array.getDouble(i);
                    }
                    v.setFloatValues(vs);
                } catch (Exception e) {
                    e.printStackTrace();
                }
                break;
            case "OfInt":
                try {
                    JSONArray array = (JSONArray) (new JSONTokener(value).nextValue());
                    int[] vs = new int[array.length()];
                    for (int i=0;i<array.length();i++){
                        vs[i] =  array.getInt(i);
                    }
                    v.setIntValues(vs);
                } catch (Exception e) {
                    e.printStackTrace();
                }
                break;
            case "Duration":
                try {
                    long d = Long.parseLong(value);
                    v.setDuration(d);
                } catch (Exception e) {
                    e.printStackTrace();
                }
                break;
            case "ValueChangedListener":
                v.addUpdateListener(new ValueAnimator.AnimatorUpdateListener() {
                    @Override
                    public void onAnimationUpdate(ValueAnimator animation) {
                        Faithdroid.triggerEventHandler(value, animation.getAnimatedValue().toString());
                    }
                });
                break;
            case "Start":
                v.start();
                break;
        }
    }
}
