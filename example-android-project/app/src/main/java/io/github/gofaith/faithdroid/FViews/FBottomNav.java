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
        switch (attr) {
            case "Visibility":
                return getVisibility();
            case "X":
                return getX();
            case "Y":
                return getY();
            case "Width":
                return getWidth();
            case "Height":
                return getHeight();
            case "PivotX":
                return getPivotX();
            case "PivotY":
                return getPivotY();
            case "ScaleX":
                return getScaleX();
            case "ScaleY":
                return getScaleY();
            case "Rotation":
                return getRotation();
            //----------------------------------------------
        }
        return "";
    }

    @Override
    public void setAttr(String attr, final String value) {
        if (value == null) {
            return;
        }
        switch (attr) {
            case "BackgroundColor":
                setBackgroundColor(value);
                break;
            case "Background":
                setBackground(value);
                break;
            case "Foreground":
                setForeground(value);
                break;
            case "Size":
                parseSize(value);
                break;
            case "X":
                setX(value);
                break;
            case "Y":
                setY(value);
                break;
            case "PivotX":
                setPivotX(value);
                break;
            case "PivotY":
                setPivotY(value);
                break;
            case "ScaleX":
                setScaleX(value);
                break;
            case "ScaleY":
                setScaleY(value);
                break;
            case "Rotation":
                setRotation(value);
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
            case "OnTouch":
                setOnTouchListener(value);
                break;
            case "OnClick":
                v.setOnClickListener(new View.OnClickListener() {
                    @Override
                    public void onClick(View view) {
                        Faithdroid.triggerEventHandler(value, "");
                    }
                });
                break;
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
