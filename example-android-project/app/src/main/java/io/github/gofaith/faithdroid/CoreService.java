package io.github.gofaith.faithdroid;

import android.app.Notification;
import android.app.NotificationChannel;
import android.app.NotificationManager;
import android.app.PendingIntent;
import android.app.Service;
import android.content.Intent;
import android.graphics.Color;
import android.os.Build;
import android.os.IBinder;
import android.support.v4.app.NotificationCompat;
import android.widget.RemoteViews;

import faithdroid.Faithdroid;

public class CoreService extends Service {
    public static String onCreate=null;
    public static CoreService instance = null;
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
            NotificationChannel notificationChannel = new NotificationChannel(NOTIFICATION_CHANNEL_ID, "My Notifications", NotificationManager.IMPORTANCE_DEFAULT);
            // Configure the notification channel.
            notificationChannel.setDescription("Channel description");
            notificationChannel.enableLights(true);
            notificationChannel.setLightColor(Color.RED);
//            notificationChannel.setVibrationPattern(new long[]{0, 1000, 500, 1000});
            notificationChannel.enableVibration(false);
            notificationManager.createNotificationChannel(notificationChannel);
        }
        NotificationCompat.Builder builder = new NotificationCompat.Builder(getApplicationContext(),"my_notification_channel");
        long when = System.currentTimeMillis();
        builder.setSmallIcon(R.mipmap.ic_launcher_round);
        Intent notificationIntent = new Intent(getApplicationContext(), MainActivity.class).putExtra("notification", "1");
        PendingIntent contentIntent = PendingIntent.getActivity(getApplicationContext(), 1, notificationIntent, PendingIntent.FLAG_UPDATE_CURRENT);
        builder.setContentIntent(contentIntent);
        builder.setSound(null);
        Notification notification = builder.build();
        notification.when = when;
        RemoteViews remoteViews = new RemoteViews(getApplicationContext().getPackageName(), R.layout.core_notification);
//        listener(remoteViews,getApplicationContext());
        notification.contentView = remoteViews;
        notification.flags |= Notification.FLAG_NO_CLEAR;
        startForeground(1,notification);
    }
}
