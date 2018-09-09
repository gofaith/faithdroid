package io.github.gofaith.faithdroid.UI;

import android.app.Activity;
import android.content.ContentUris;
import android.content.Context;
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
import android.support.v4.graphics.drawable.IconCompat;
import android.util.TypedValue;
import android.view.Menu;
import android.view.MenuItem;

import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;

import java.net.URISyntaxException;

import io.github.gofaith.faithdroid.R;

public class Toolkit {
    public static void parseMenu(UIController uiController,Menu menu, JSONArray array) throws JSONException {
        for (int i=0;i<array.length();i++) {
            JSONObject object = array.getJSONObject(i);
            if (!object.has("MySubMenu")||object.isNull("MySubMenu")) {
                MenuItem item = menu.add(object.getString("MyTitle"));
                if (!object.isNull("MyOnClick") && !object.getString("MyOnClick").equals("")) {
                    uiController.menuItemsOnClickMap.put(item, object.getString("MyOnClick"));
                }
                if (!object.isNull("MyIcon") && !object.getString("MyIcon").equals("")) {
                    item.setIcon(file2Drawable(uiController.activity,object.getString("MyIcon")));
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
//            if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.HONEYCOMB) {
//                // If we're running on Honeycomb or newer, then we can use the Theme's
//                // selectableItemBackground to ensure that the View has a pressed state
//                TypedValue outValue = new TypedValue();
//                activity.getTheme().resolveAttribute(R.attr.selectableItemBackground, outValue, true);
//                Drawable d= activity.getResources().getDrawable(outValue.resourceId);
//                return  d;
//            }
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
}
