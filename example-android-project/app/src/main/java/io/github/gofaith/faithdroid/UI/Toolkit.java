package io.github.gofaith.faithdroid.UI;

import android.app.Activity;
import android.graphics.drawable.Drawable;
import android.os.Build;
import android.support.v4.content.ContextCompat;
import android.support.v4.graphics.drawable.IconCompat;
import android.view.Menu;
import android.view.MenuItem;

import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;

import io.github.gofaith.faithdroid.R;

public class Toolkit {
    public static void parseMenu(UIController uiController,Menu menu, JSONArray array) throws JSONException {
        for (int i=0;i<array.length();i++) {
            JSONObject object = array.getJSONObject(i);
            if (!object.has("MySubMenu")||object.isNull("MySubMenu")) {
                MenuItem item = menu.add(object.getString("MyTitle"));
                if (!object.isNull("MyOnClick") && !object.getString("MyOnClick").equals("")) {
                    uiController.menuItemsOnClickMap.put(item, object.getString("MyOnClick"));
                }
                if (!object.isNull("MyIcon") && !object.getString("MyIcon").equals("")) {
                    item.setIcon(file2Drawable(uiController.activity,object.getString("MyIcon")));
                }
                if (!object.isNull("MyShowAsAction") && !object.getString("MyShowAsAction").equals("")) {
                    item.setShowAsAction(MenuItem.SHOW_AS_ACTION_IF_ROOM);
                }
                continue;
            }
            JSONArray subMenu=object.getJSONArray("MySubMenu");
            parseMenu(uiController,menu.addSubMenu(object.getString("MyTitle")),subMenu);
        }
    }

    public static Drawable file2Drawable(Activity activity, String value) {
        if (value.startsWith("file://")) {
            String path=value.substring("file://".length());
            Drawable draw=Drawable.createFromPath(path);
            return draw;
        } else if (value.startsWith("assets://")) {
//            Drawable d = Drawable.createFromStream(getAssets().open("Cloths/btn_no.png"), null);
            String path = value.substring("assets://".length());
            try {
                Drawable drawable = Drawable.createFromStream(activity.getAssets().open(path), null);
                return drawable;
            } catch (Exception e) {
                e.printStackTrace();
            }
        } else if (value.startsWith("drawable://")) {
            int src;
            switch (value.substring("drawable://".length())){
                case "add":
                    src=R.drawable.add;
                    break;
                default:
                    src= R.mipmap.ic_launcher_round;
                    break;
            }
            return ContextCompat.getDrawable(activity, src);
        }
        return null;
    }
}
