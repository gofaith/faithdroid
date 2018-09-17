package io.github.gofaith.faithdroid.FViews;

import android.content.Context;
import android.content.Intent;

import io.github.gofaith.faithdroid.CoreService;
import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.UIController;

public class FService extends FView  {
    public FService(UIController uiController) {
        parentController = uiController;
    }
    public static String getAttr(Context context, String attr) {
        return null;
    }

    public static void setAttr(Context context,String attr, String value) {
        if (value == null) {
            return;
        }
        switch (attr) {
            case "OnCreate":
                CoreService.onCreate = value;
                Intent intent = new Intent(context, CoreService.class);
                context.startService(intent);
                break;
        }
    }
}
