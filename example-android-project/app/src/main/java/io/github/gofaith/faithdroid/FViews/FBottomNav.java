package io.github.gofaith.faithdroid.FViews;

import android.graphics.drawable.Drawable;
import android.os.Build;
import android.support.annotation.NonNull;
import android.support.design.widget.BottomNavigationView;
import android.util.TypedValue;
import android.view.MenuItem;
import android.view.View;

import org.json.JSONArray;
import org.json.JSONObject;
import org.json.JSONTokener;

import faithdroid.Faithdroid;
import io.github.gofaith.faithdroid.R;
import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.Toolkit;
import io.github.gofaith.faithdroid.UI.UIController;

public class FBottomNav extends FView implements AttrGettable,AttrSettable {
    public BottomNavigationView v;

    public FBottomNav(UIController controller) {
        parentController = controller;
        v = new BottomNavigationView(parentController.activity);
        view=v;
        parseSize("[-2,-1]");

        TypedValue outValue = new TypedValue();
        parentController.activity.getTheme().resolveAttribute(android.R.attr.windowBackground, outValue, true);
        Drawable d= parentController.activity.getResources().getDrawable(outValue.resourceId);
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.JELLY_BEAN) {
            v.setBackground(d);
        }else {
            v.setBackgroundDrawable(d);
        }
    }
    @Override
    public String getAttr(String attr) {
        String str = getUniversalAttr(attr);
        if (str != null) {
            return str;
        }
        switch (attr){
            //----------------------------------------------
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
            case "Menu":
                try {
                    JSONArray object = (JSONArray) (new JSONTokener(value).nextValue());
                    Toolkit.parseMenu(parentController, v.getMenu(), object);
                    v.setOnNavigationItemSelectedListener(new BottomNavigationView.OnNavigationItemSelectedListener() {
                        @Override
                        public boolean onNavigationItemSelected(@NonNull MenuItem menuItem) {
                            if (parentController.menuItemsOnClickMap.containsKey(menuItem)) {
                                Faithdroid.triggerEventHandler(parentController.menuItemsOnClickMap.get(menuItem), "");
                                return true;
                            }
                            return false;
                        }
                    });
                } catch (Exception e) {
                    e.printStackTrace();
                }
        }
    }
}
