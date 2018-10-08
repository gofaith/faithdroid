package io.github.gofaith.faithdroid.UI;

import android.app.Activity;
import android.content.ContentUris;
import android.content.Context;
import android.content.Intent;
import android.database.Cursor;
import android.graphics.Color;
import android.graphics.PorterDuff;
import android.graphics.drawable.Drawable;
import android.net.Uri;
import android.os.Build;
import android.os.Environment;
import android.provider.DocumentsContract;
import android.provider.MediaStore;
import android.support.annotation.RequiresApi;
import android.support.v4.content.ContextCompat;
import android.support.v4.content.FileProvider;
import android.support.v4.graphics.drawable.IconCompat;
import android.support.v4.view.ViewCompat;
import android.util.Log;
import android.util.TypedValue;
import android.view.Menu;
import android.view.MenuItem;
import android.webkit.MimeTypeMap;

import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;

import java.io.BufferedInputStream;
import java.io.File;
import java.io.FileOutputStream;
import java.io.InputStream;
import java.io.OutputStream;
import java.net.URISyntaxException;
import java.net.URL;
import java.net.URLConnection;

import faithdroid.Faithdroid;
import io.github.gofaith.faithdroid.BuildConfig;
import io.github.gofaith.faithdroid.R;

import static android.content.ContentValues.TAG;
import static android.webkit.URLUtil.isFileUrl;

public class Toolkit {
    public static void parseMenu(final UIController uiController, Menu menu, JSONArray array) throws JSONException {
        for (int i=0;i<array.length();i++) {
            JSONObject object = array.getJSONObject(i);
            if (!object.has("MySubMenu")||object.isNull("MySubMenu")) {
                final MenuItem item = menu.add(0, ViewCompat.generateViewId(),i,object.getString("MyTitle"));
                if (!object.isNull("MyOnClick") && !object.getString("MyOnClick").equals("")) {
                    uiController.menuItemsOnClickMap.put(item, object.getString("MyOnClick"));
                }
                if (!object.isNull("MyIcon") && !object.getString("MyIcon").equals("")) {
                    String mIcon=object.getString("MyIcon");
                    if (mIcon.startsWith("http")) {
                        final File file = new File("/data/data/" + uiController.getPkg() + "/cacheDir/" + Faithdroid.url2cachePath(mIcon));
                        if (file.exists()) {
                            item.setIcon(file2Drawable(uiController.activity, "file://" + file.getAbsolutePath()));
                        }else{
                            downloadFile(mIcon, file.getAbsolutePath(), new DownloadedListener() {
                                @Override
                                public void onFailed(String error) {
                                    Log.d(TAG, "onFailed: "+error);
                                }
                                @Override
                                public void onSucceed(String fpath) {
                                    uiController.activity.runOnUiThread(new Runnable() {
                                        @Override
                                        public void run() {
                                            item.setIcon(file2Drawable(uiController.activity, "file://" + file.getAbsolutePath()));
                                        }
                                    });
                                }
                            });
                        }
                    }else {
                        item.setIcon(file2Drawable(uiController.activity, object.getString("MyIcon")));
                    }
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
        } else if (value.equals("RippleEffect")) {
            return activity.getResources().getDrawable(R.drawable.ripple);
        }
        return null;
    }
    public static String getPath(Context context, Uri uri) throws URISyntaxException {
        if ("content".equalsIgnoreCase(uri.getScheme())) {
            String[] projection = { "_data" };
            Cursor cursor = null;
            try {
                cursor = context.getContentResolver().query(uri, projection, null, null, null);
                int column_index = cursor.getColumnIndexOrThrow("_data");
                if (cursor.moveToFirst()) {
                    return cursor.getString(column_index);
                }
            } catch (Exception e) {
                e.printStackTrace();
            }
        } else if ("file".equalsIgnoreCase(uri.getScheme())) {
            return uri.getPath();
        }
        return null;
    }
    public static String getPathByUriBelowKitkat(Context context, Uri data) {
        String filename=null;
        if (data.getScheme().toString().compareTo("content") == 0) {
            Cursor cursor = context.getContentResolver().query(data, new String[] { "_data" }, null, null, null);
            if (cursor.moveToFirst()) {
                filename = cursor.getString(0);
            }
        } else if (data.getScheme().toString().compareTo("file") == 0) {// file:///开头的uri
            filename = data.toString();
            filename = data.toString().replace("file://", "");// 替换file://
            if (!filename.startsWith("/mnt")) {// 加上"/mnt"头
                filename += "/mnt";
            }
        }
        return filename;
    }
    // 专为Android4.4设计的从Uri获取文件绝对路径，以前的方法已不好使
    @RequiresApi(api = Build.VERSION_CODES.KITKAT)
    public static String getPathByUri4kitkat(final Context context, final Uri uri) {
        final boolean isKitKat = Build.VERSION.SDK_INT >= Build.VERSION_CODES.KITKAT;
        // DocumentProvider
        if (isKitKat && DocumentsContract.isDocumentUri(context, uri)) {
            if (isExternalStorageDocument(uri)) {// ExternalStorageProvider
                final String docId = DocumentsContract.getDocumentId(uri);
                final String[] split = docId.split(":");
                final String type = split[0];
                if ("primary".equalsIgnoreCase(type)) {
                    return Environment.getExternalStorageDirectory() + "/" + split[1];
                }
            } else if (isDownloadsDocument(uri)) {// DownloadsProvider
                final String id = DocumentsContract.getDocumentId(uri);
                final Uri contentUri = ContentUris.withAppendedId(Uri.parse("content://downloads/public_downloads"),
                        Long.valueOf(id));
                return getDataColumn(context, contentUri, null, null);
            } else if (isMediaDocument(uri)) {// MediaProvider
                final String docId = DocumentsContract.getDocumentId(uri);
                final String[] split = docId.split(":");
                final String type = split[0];
                Uri contentUri = null;
                if ("image".equals(type)) {
                    contentUri = MediaStore.Images.Media.EXTERNAL_CONTENT_URI;
                } else if ("video".equals(type)) {
                    contentUri = MediaStore.Video.Media.EXTERNAL_CONTENT_URI;
                } else if ("audio".equals(type)) {
                    contentUri = MediaStore.Audio.Media.EXTERNAL_CONTENT_URI;
                }
                final String selection = "_id=?";
                final String[] selectionArgs = new String[] { split[1] };
                return getDataColumn(context, contentUri, selection, selectionArgs);
            }
        } else if ("content".equalsIgnoreCase(uri.getScheme())) {// MediaStore
            // (and
            // general)
            return getDataColumn(context, uri, null, null);
        } else if ("file".equalsIgnoreCase(uri.getScheme())) {// File
            return uri.getPath();
        }
        return null;
    }
    public static String getDataColumn(Context context, Uri uri, String selection, String[] selectionArgs) {
        Cursor cursor = null;
        final String column = "_data";
        final String[] projection = { column };
        try {
            cursor = context.getContentResolver().query(uri, projection, selection, selectionArgs, null);
            if (cursor != null && cursor.moveToFirst()) {
                final int column_index = cursor.getColumnIndexOrThrow(column);
                return cursor.getString(column_index);
            }
        } finally {
            if (cursor != null)
                cursor.close();
        }
        return null;
    }
    public static boolean isExternalStorageDocument(Uri uri) {
        return "com.android.externalstorage.documents".equals(uri.getAuthority());
    }
    public static boolean isDownloadsDocument(Uri uri) {
        return "com.android.providers.downloads.documents".equals(uri.getAuthority());
    }
    public static boolean isMediaDocument(Uri uri) {
        return "com.android.providers.media.documents".equals(uri.getAuthority());
    }

    interface DownloadedListener {
        void onFailed(String error);
        void onSucceed(String fpath);
    }

    public static void openFile(Context context,String path) {
        Intent intent = new Intent(Intent.ACTION_VIEW);
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.N) {
            /* Android N 写法*/
            intent.setFlags(Intent.FLAG_GRANT_READ_URI_PERMISSION);
            Uri contentUri = FileProvider.getUriForFile(context, BuildConfig.APPLICATION_ID + ".fileProvider", new File(path));
            intent.setDataAndType(contentUri, getMimeType(path));
        } else {
            /* Android N之前的老版本写法*/
            intent.setDataAndType(Uri.fromFile(new File(path)), getMimeType(path));
            intent.setFlags(Intent.FLAG_ACTIVITY_NEW_TASK);
        }
        context.startActivity(intent);
    }

    public static String getMimeType(String url) {
        String type = null;
        String extension = MimeTypeMap.getFileExtensionFromUrl(url);
        if (extension != null) {
            type = MimeTypeMap.getSingleton().getMimeTypeFromExtension(extension);
        }
        return type;
    }
    public static void downloadFile(final String url_str, final String fdist, final DownloadedListener callback){
        new Thread(new Runnable() {
            @Override
            public void run() {
                File f=new File(fdist);
                String dir=f.getAbsolutePath().substring(0,f.getAbsolutePath().length()-f.getName().length());
                new File(dir).mkdirs();
                int count;
                try {
                    URL url = new URL(url_str);
                    URLConnection conection = url.openConnection();
                    conection.connect();
                    InputStream input = new BufferedInputStream(url.openStream(), 8192);
                    OutputStream output = new FileOutputStream(fdist);
                    byte data[] = new byte[1024];
                    long total = 0;
                    while ((count = input.read(data)) != -1) {
                        total += count;
//                publishProgress("" + (int) ((total * 100) / lenghtOfFile));
                        output.write(data, 0, count);
                    }
                    output.flush();
                    output.close();
                    input.close();
                    callback.onSucceed(fdist);
                } catch (Exception e) {
                    Log.e("Error: ", e.getMessage());
                    callback.onFailed(e.toString());
                }
            }
        }).start();
    }
}
