package io.github.gofaith.faithdroid.FViews;

import android.content.ClipData;
import android.content.ClipboardManager;
import android.content.Context;

import faithdroid.Faithdroid;
import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.UIController;

public class FClipboard extends FView implements AttrSettable,AttrGettable {
    public ClipboardManager v;

    public FClipboard(UIController controller) {
        parentController=controller;
        v = (ClipboardManager) parentController.activity.getSystemService(Context.CLIPBOARD_SERVICE);
    }
    @Override
    public String getAttr(String attr) {
        if (attr.startsWith("Item")) {
            try {
                int index = Integer.parseInt(attr.substring("Item".length()));
                return v.getPrimaryClip().getItemAt(index).getText().toString();
            } catch (Exception e) {
                e.printStackTrace();
                return e.toString();
            }
        }
        switch (attr) {
            case "ClipData":
                return v.getPrimaryClip().getItemAt(0).getText().toString();
            case "ClipCount":
                return String.valueOf(v.getPrimaryClip().getItemCount());
        }
        return null;
    }

    @Override
    public void setAttr(String attr, final String value) {
        if (value == null) {
            return;
        }
        switch (attr) {
            case "OnChange":
                v.addPrimaryClipChangedListener(new ClipboardManager.OnPrimaryClipChangedListener() {
                    @Override
                    public void onPrimaryClipChanged() {
                        Faithdroid.triggerEventHandler(value, "");
                    }
                });
                break;
            case "ClipData":
                v.setPrimaryClip(ClipData.newPlainText("new",value));
                break;
        }
    }
}
