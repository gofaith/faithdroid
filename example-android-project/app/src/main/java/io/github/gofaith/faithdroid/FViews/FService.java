package io.github.gofaith.faithdroid.FViews;

import android.content.Context;
import android.content.Intent;

import io.github.gofaith.faithdroid.CoreService;
import io.github.gofaith.faithdroid.R;
import io.github.gofaith.faithdroid.UI.UIController;

public class FService extends FView  {
    public FService(UIController uiController) {
        parentController = uiController;
    }
    public static String getAttr(Context context, String attr) {
        return null;
    }
    public static void setAttr(Context context, String attr, final String value) {
        if (value == null) {
            return;
        }
        switch (attr) {
            case "OnCreate":
                CoreService.onCreate = value;
                break;
            case "QuitButton":
                CoreService.quitButton = value;
                break;
            case "OnQuit":
                CoreService.onQuitClick =value;
                break;
            case "Title":
                CoreService.ntf_title=value;
                if (CoreService.notification!=null){
                    CoreService.getInstance().handler.post(new Runnable() {
                        @Override
                        public void run() {
                            CoreService.notification.contentView.setTextViewText(R.id.ntf_title,value);
                            CoreService.getInstance().startForeground(1,CoreService.notification);
                        }
                    });
                }
                break;
            case "SubTitle":
                CoreService.ntf_subtitle = value;
                if (CoreService.notification != null) {
                    CoreService.getInstance().handler.post(new Runnable() {
                        @Override
                        public void run() {
                            CoreService.notification.contentView.setTextViewText(R.id.ntf_subtitle,value);
                            CoreService.getInstance().startForeground(1,CoreService.notification);
                        }
                    });
                }
                break;
            case "Show":
                Intent intent = new Intent(context, CoreService.class);
                context.startService(intent);
                break;
            case "FinishAllActivity":
                CoreService.getInstance().finishAllActivity();
                break;
            case "KillAll":
                CoreService.killAll();
                break;
            case "Stop":
                CoreService.getInstance().stopService();
                break;
        }
    }
}
