package io.github.gofaith.faithdroid;

import android.app.Notification;
import android.app.NotificationChannel;
import android.app.NotificationManager;
import android.app.PendingIntent;
import android.app.Service;
import android.content.BroadcastReceiver;
import android.content.Context;
import android.content.Intent;
import android.content.IntentFilter;
import android.graphics.Color;
import android.os.Build;
import android.os.Handler;
import android.os.IBinder;
import android.os.Looper;
import android.support.v4.app.NotificationCompat;
import android.support.v4.content.LocalBroadcastManager;
import android.view.View;
import android.widget.RemoteViews;

import faithdroid.Faithdroid;

public class CoreService extends Service {
    public static String onCreate=null;
    public static String quitButton = null;
    public static String onQuitClick =null;
    private static CoreService instance = null;
    public static Notification notification;
    public static String ntf_title="",ntf_subtitle="";
    public Handler handler=new Handler(Looper.myLooper());

    public CoreService() {
    }
    public static CoreService getInstance(){
        return instance;
    }

    @Override
    public void onCreate() {
        instance = this;
        if (onCreate != null && !onCreate.equals("")) {
            showNtf();
            Faithdroid.triggerEventHandler(onCreate, "");
        }
    }

    @Override
    public int onStartCommand(Intent intent, int flags, int startId) {
        return START_STICKY;
    }

    @Override
    public IBinder onBind(Intent intent) {
        // TODO: Return the communication channel to the service.
        throw new UnsupportedOperationException("Not yet implemented");
    }
    private void showNtf() {
        final String NOTIFICATION_CHANNEL_ID = "my_notification_channel";
        NotificationManager notificationManager = (NotificationManager) getSystemService(NOTIFICATION_SERVICE);
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.O) {
            NotificationChannel notificationChannel = new NotificationChannel(NOTIFICATION_CHANNEL_ID, "My Notifications", NotificationManager.IMPORTANCE_LOW);
            notificationChannel.setSound(null,null);
            notificationManager.createNotificationChannel(notificationChannel);
        }
        NotificationCompat.Builder builder = new NotificationCompat.Builder(getApplicationContext(),"my_notification_channel");
        long when = System.currentTimeMillis();
        builder.setSmallIcon(R.drawable.ntf_small_icon);
        builder.setOngoing(false);
        builder.setPriority(NotificationCompat.PRIORITY_DEFAULT);
        Intent notificationIntent = new Intent(getApplicationContext(), MainActivity.class).putExtra("notification", "1");
        PendingIntent contentIntent = PendingIntent.getActivity(getApplicationContext(), 1, notificationIntent, PendingIntent.FLAG_UPDATE_CURRENT);
        builder.setContentIntent(contentIntent);
//        builder.setSound(null);
        notification = builder.build();
        notification.when = when;
        RemoteViews remoteViews = new RemoteViews(getApplicationContext().getPackageName(), R.layout.core_notification);
        listener(remoteViews,getApplicationContext());
        notification.contentView = remoteViews;
        notification.flags |= Notification.FLAG_NO_CLEAR;
        startForeground(1,notification);
    }

    @Override
    public void onDestroy() {
        super.onDestroy();
    }

    private void listener(RemoteViews remoteViews, Context context) {
        Intent intent2 = new Intent("QuitService");
        PendingIntent pendingIntent2 = PendingIntent.getBroadcast(context, 1, intent2, 0);
        IntentFilter intentFilter = new IntentFilter();
        intentFilter.addAction("QuitService");
        BroadcastReceiver receiver = new BroadcastReceiver() {
            @Override
            public void onReceive(Context context, Intent intent) {
                if (intent.getAction().equals("QuitService")) {
                    Faithdroid.triggerEventHandler(onQuitClick, "");
                }
            }
        };
        context.registerReceiver(receiver, intentFilter);
        if (quitButton!=null) {
            remoteViews.setOnClickPendingIntent(R.id.ntf_quit, pendingIntent2);
            remoteViews.setTextViewText(R.id.ntf_quit,quitButton);
        }else {
            remoteViews.setViewVisibility(R.id.ntf_quit, View.GONE);
        }
        remoteViews.setTextViewText(R.id.ntf_title,ntf_title);
        remoteViews.setTextViewText(R.id.ntf_subtitle,ntf_subtitle);
    }
    public void finishAllActivity(){
        Intent i = new Intent("uibro");
        i.putExtra("action", "quit");
        LocalBroadcastManager.getInstance(CoreService.this).sendBroadcast(i);
    }
    public static void killAll(){
        android.os.Process.killProcess(android.os.Process.myPid());
    }
    public void stopService(){
        stopForeground(true);
        stopSelf();
    }
}
