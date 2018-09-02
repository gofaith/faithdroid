package io.github.gofaith.faithdroid.UI;

import android.view.Menu;
import android.view.MenuItem;

import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;

public class Toolkit {
    public static void parseMenu(Menu menu, JSONArray array) throws JSONException {
        for (int i=0;i<array.length();i++) {
            JSONObject object = array.getJSONObject(i);
            if (object.isNull("MySubMenu")) {
                MenuItem item = menu.add(object.getString("MyTitle"));
                continue;
            }
            JSONArray subMenu=object.getJSONArray("MySubMenu");
            parseMenu(menu.addSubMenu(object.getString("MyTitle")),subMenu);
        }
    }
}
