package io.github.gofaith.faithdroid.UI;

import android.view.Menu;
import android.view.MenuItem;

import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;

public class Toolkit {
    public static void parseMenu(UIController uiController,Menu menu, JSONArray array) throws JSONException {
        for (int i=0;i<array.length();i++) {
            JSONObject object = array.getJSONObject(i);
            if (!object.has("MySubMenu")||object.isNull("MySubMenu")) {
                MenuItem item = menu.add(object.getString("MyTitle"));
                if (!object.isNull("MyOnClick") && !object.getString("MyOnClick").equals("")) {
                    uiController.menuItemsOnClickMap.put(item, object.getString("MyOnClick"));
                }
                continue;
            }
            JSONArray subMenu=object.getJSONArray("MySubMenu");
            parseMenu(uiController,menu.addSubMenu(object.getString("MyTitle")),subMenu);
        }
    }
}
