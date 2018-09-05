package io.github.gofaith.faithdroid.FViews;

import android.content.DialogInterface;
import android.support.v7.app.AlertDialog;

import org.json.JSONArray;
import org.json.JSONTokener;

import faithdroid.Faithdroid;
import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.UIController;

public class FAlertDialog extends FView implements AttrSettable, AttrGettable {
    public AlertDialog.Builder v;
    public DialogInterface di;
    public FAlertDialog(UIController controller) {
        parentController=controller;
        v = new AlertDialog.Builder(parentController.activity);
    }
    @Override
    public String getAttr(String attr) {
        return null;
    }

    @Override
    public void setAttr(String attr, String value) {
        if (value==null)
            return;
        switch (attr) {
            case "Title":
                v.setTitle(value);
                break;
            case "View":
                if (parentController.viewmap.containsKey(value)) {
                    v.setView(parentController.viewmap.get(value).view);
                }
                break;
            case "PositiveButton":
                try {
                    JSONArray array = (JSONArray) (new JSONTokener(value).nextValue());
                    final String fnId = array.getString(1);
                    v.setPositiveButton(array.getString(0), new DialogInterface.OnClickListener() {
                        @Override
                        public void onClick(DialogInterface dialog, int which) {
                            Faithdroid.triggerEventHandler(fnId, "");
                        }
                    });
                } catch (Exception e) {
                    e.printStackTrace();
                }
                break;
            case "NegativeButton":
                try {
                    JSONArray array = (JSONArray) (new JSONTokener(value).nextValue());
                    final String fnId = array.getString(1);
                    v.setNegativeButton(array.getString(0), new DialogInterface.OnClickListener() {
                        @Override
                        public void onClick(DialogInterface dialog, int which) {
                            Faithdroid.triggerEventHandler(fnId, "");
                        }
                    });
                } catch (Exception e) {
                    e.printStackTrace();
                }
                break;
            case "Show":
                di = v.show();
                break;
            case "Dismiss":
                if (di!=null) {
                    di.dismiss();
                }
                break;
        }
    }
}
