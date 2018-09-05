package io.github.gofaith.faithdroid.FViews;

import android.view.MenuItem;
import android.widget.PopupMenu;

import org.json.JSONArray;
import org.json.JSONTokener;

import faithdroid.Faithdroid;
import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.Toolkit;
import io.github.gofaith.faithdroid.UI.UIController;

public class FPopupMenu extends FView implements AttrSettable, AttrGettable {
    public PopupMenu v;

    public FPopupMenu(UIController controller, FView anchorView) {
        parentController=controller;
        v = new PopupMenu(parentController.activity, anchorView.view);
    }
    @Override
    public String getAttr(String attr) {
        return null;
    }

    @Override
    public void setAttr(String attr, String value) {
        if (value == null) {
            return;
        }
        switch (attr) {
            case "Menus":
                try {
                    JSONArray array = (JSONArray) (new JSONTokener(value).nextValue());
                    Toolkit.parseMenu(parentController, v.getMenu(), array);
                    v.setOnMenuItemClickListener(new PopupMenu.OnMenuItemClickListener() {
                        @Override
                        public boolean onMenuItemClick(MenuItem item) {
                            if (parentController.menuItemsOnClickMap.containsKey(item)) {
                                Faithdroid.triggerEventHandler(parentController.menuItemsOnClickMap.get(item), "");
                                return true;
                            }
                            return false;
                        }
                    });
                } catch (Exception e) {
                    e.printStackTrace();
                }
                break;
            case "Show":
                v.show();
                break;
            case "Dismiss":
                v.dismiss();
                break;
        }
    }
}
