package io.github.stevenzack.faithdroidapp.UI

import android.content.ContentUris
import android.content.Context
import android.content.Intent
import android.database.Cursor
import android.graphics.Bitmap
import android.graphics.Canvas
import android.graphics.drawable.BitmapDrawable
import android.graphics.drawable.Drawable
import android.net.Uri
import android.os.Build
import android.os.Environment
import android.provider.DocumentsContract
import android.provider.MediaStore
import android.util.Log
import android.view.Menu
import android.view.MenuItem
import android.webkit.MimeTypeMap
import androidx.core.view.ViewCompat
import faithdroid.Faithdroid
import org.json.JSONArray
import org.json.JSONException

import java.io.*
import java.net.URISyntaxException
import java.net.URL

import android.content.ContentValues.TAG
import androidx.annotation.NonNull
import androidx.annotation.RequiresApi
import androidx.appcompat.app.AppCompatActivity
import androidx.core.content.ContextCompat
import androidx.core.content.FileProvider
import io.github.stevenzack.faithdroidapp.BuildConfig
import io.github.stevenzack.faithdroidapp.R

object Toolkit {
    @Throws(JSONException::class)
    fun parseMenu(uiController: UIController, menu: Menu, array: JSONArray) {
        for (i in 0 until array.length()) {
            val `object` = array.getJSONObject(i)
            if (!`object`.has("MySubMenu") || `object`.isNull("MySubMenu")) {
                val item = menu.add(0, ViewCompat.generateViewId(), i, `object`.getString("MyTitle"))
                if (!`object`.isNull("MyOnClick") && `object`.getString("MyOnClick") != "") {
                    uiController.menuItemsOnClickMap[item] = `object`.getString("MyOnClick")
                }
                if (!`object`.isNull("MyIcon") && `object`.getString("MyIcon") != "") {
                    val mIcon = `object`.getString("MyIcon")
                    file2Drawable(uiController, mIcon, {d->
                        item.icon=d
                    })
                }
                if (!`object`.isNull("MyShowAsAction") && `object`.getString("MyShowAsAction") != "") {
                    item.setShowAsAction(MenuItem.SHOW_AS_ACTION_IF_ROOM)
                }
                continue
            }
            val subMenu = `object`.getJSONArray("MySubMenu")
            parseMenu(uiController, menu.addSubMenu(`object`.getString("MyTitle")), subMenu)
        }
    }


    fun file2Drawable(parentViewController: UIController, value: String, listener: (Drawable)->Unit) {
        if (value.startsWith("http")) {
            val file = File("/data/data/" + parentViewController.pkg + "/cacheDir/" + Faithdroid.url2cachePath(value))
            if (file.exists()) {
                file2Drawable(parentViewController, "file://" + file.absolutePath, listener)
            } else {
                downloadFile(value, file.absolutePath, object : DownloadedListener {
                    override fun onFailed(error: String) {
                        Log.d(TAG, "onFailed: $error")
                    }

                    override fun onSucceed(fpath: String) {
                        parentViewController.activity.runOnUiThread {
                            file2Drawable(
                                parentViewController,
                                "file://$fpath",
                                listener
                            )
                        }
                    }
                })
            }
        } else if (value.startsWith("file://")) {
            val path = value.substring("file://".length)
            val draw = Drawable.createFromPath(path)
            listener(draw!!)
        } else if (value.startsWith("assets://")) {
            //            Drawable d = Drawable.createFromStream(getAssets().open("Cloths/btn_no.png"), null);
            val path = value.substring("assets://".length)
            try {
                val drawable = Drawable.createFromStream(parentViewController.activity.assets.open(path), null)
                listener(drawable!!)
            } catch (e: Exception) {
                e.printStackTrace()
            }

        } else if (value.startsWith("drawable://")) {
            val src: Int
            when (value.substring("drawable://".length)) {
                "add" -> src = R.drawable.add
                else -> src = R.mipmap.ic_launcher_round
            }
            listener(ContextCompat.getDrawable(parentViewController.activity, src)!!)
        } else if (value == "RippleEffect") {
            listener(parentViewController.activity.resources.getDrawable(R.drawable.ripple))
        } else if (value.startsWith("fdrawable://")) {
            listener(parentViewController.drawableMap.get(value)!!)
        } else if (value == "RadiusCorner") {
            listener(parentViewController.activity.resources.getDrawable(R.drawable.radius_corner))
        }
    }

    @Throws(URISyntaxException::class)
    fun getPath(context: Context, uri: Uri): String? {
        if ("content".equals(uri.scheme!!, ignoreCase = true)) {
            val projection = arrayOf("_data")
            var cursor: Cursor? = null
            try {
                cursor = context.contentResolver.query(uri, projection, null, null, null)
                val column_index = cursor!!.getColumnIndexOrThrow("_data")
                if (cursor.moveToFirst()) {
                    return cursor.getString(column_index)
                }
            } catch (e: Exception) {
                e.printStackTrace()
            }

        } else if ("file".equals(uri.scheme!!, ignoreCase = true)) {
            return uri.path
        }
        return null
    }

    fun getPathByUriBelowKitkat(context: Context, data: Uri): String? {
        var filename: String? = null
        if (data.scheme!!.toString().compareTo("content") == 0) {
            val cursor = context.contentResolver.query(data, arrayOf("_data"), null, null, null)
            if (cursor!!.moveToFirst()) {
                filename = cursor.getString(0)
            }
        } else if (data.scheme!!.toString().compareTo("file") == 0) {// file:///开头的uri
            filename = data.toString()
            filename = data.toString().replace("file://", "")// 替换file://
            if (!filename.startsWith("/mnt")) {// 加上"/mnt"头
                filename += "/mnt"
            }
        }
        return filename
    }

    // 专为Android4.4设计的从Uri获取文件绝对路径，以前的方法已不好使
    @RequiresApi(api = Build.VERSION_CODES.KITKAT)
    fun getPathByUri4kitkat(context: Context, uri: Uri): String? {
        val isKitKat = Build.VERSION.SDK_INT >= Build.VERSION_CODES.KITKAT
        // DocumentProvider
        if (isKitKat && DocumentsContract.isDocumentUri(context, uri)) {
            if (isExternalStorageDocument(uri)) {// ExternalStorageProvider
                val docId = DocumentsContract.getDocumentId(uri)
                val split = docId.split(":".toRegex()).dropLastWhile { it.isEmpty() }.toTypedArray()
                val type = split[0]
                if ("primary".equals(type, ignoreCase = true)) {
                    return Environment.getExternalStorageDirectory().toString() + "/" + split[1]
                }
            } else if (isDownloadsDocument(uri)) {// DownloadsProvider
                val id = DocumentsContract.getDocumentId(uri)
                val contentUri = ContentUris.withAppendedId(
                    Uri.parse("content://downloads/public_downloads"),
                    java.lang.Long.valueOf(id)
                )
                return getDataColumn(context, contentUri, null, null)
            } else if (isMediaDocument(uri)) {// MediaProvider
                val docId = DocumentsContract.getDocumentId(uri)
                val split = docId.split(":".toRegex()).dropLastWhile { it.isEmpty() }.toTypedArray()
                val type = split[0]
                var contentUri: Uri? = null
                if ("image" == type) {
                    contentUri = MediaStore.Images.Media.EXTERNAL_CONTENT_URI
                } else if ("video" == type) {
                    contentUri = MediaStore.Video.Media.EXTERNAL_CONTENT_URI
                } else if ("audio" == type) {
                    contentUri = MediaStore.Audio.Media.EXTERNAL_CONTENT_URI
                }
                val selection = "_id=?"
                val selectionArgs = arrayOf(split[1])
                return getDataColumn(context, contentUri, selection, selectionArgs)
            }
        } else if ("content".equals(uri.scheme!!, ignoreCase = true)) {// MediaStore
            // (and
            // general)
            return getDataColumn(context, uri, null, null)
        } else if ("file".equals(uri.scheme!!, ignoreCase = true)) {// File
            return uri.path
        }
        return null
    }

    fun getDataColumn(context: Context, uri: Uri?, selection: String?, selectionArgs: Array<String>?): String? {
        var cursor: Cursor? = null
        val column = "_data"
        val projection = arrayOf(column)
        try {
            cursor = context.contentResolver.query(uri!!, projection, selection, selectionArgs, null)
            if (cursor != null && cursor.moveToFirst()) {
                val column_index = cursor.getColumnIndexOrThrow(column)
                return cursor.getString(column_index)
            }
        } catch (e: Exception) {
            e.printStackTrace()
        }

        cursor?.close()
        return null
    }

    fun isExternalStorageDocument(uri: Uri): Boolean {
        return "com.android.externalstorage.documents" == uri.authority
    }

    fun isDownloadsDocument(uri: Uri): Boolean {
        return "com.android.providers.downloads.documents" == uri.authority
    }

    fun isMediaDocument(uri: Uri): Boolean {
        return "com.android.providers.media.documents" == uri.authority
    }

    fun saveDrawable(drawable: Drawable, format: Bitmap.CompressFormat, path: String) {
        val bitmap = getBitmapFromDrawable(drawable)
        saveBitmap(bitmap, format, path)
    }

    fun drawableToBitmap(drawable: Drawable?): Bitmap? {
        return if (drawable == null) null else (drawable as BitmapDrawable).bitmap
    }

    @NonNull
    private fun getBitmapFromDrawable(@NonNull drawable: Drawable): Bitmap {
        val bmp = Bitmap.createBitmap(drawable.intrinsicWidth, drawable.intrinsicHeight, Bitmap.Config.ARGB_8888)
        val canvas = Canvas(bmp)
        drawable.setBounds(0, 0, canvas.width, canvas.height)
        drawable.draw(canvas)
        return bmp
    }

    fun saveBitmap(bitmap: Bitmap, format: Bitmap.CompressFormat, path: String) {
        // 创建一个位于SD卡上的文件
        val file = File(path)
        try {
            if (!file.exists()) {
                file.createNewFile()
            }
            var out: FileOutputStream? = null
            // 打开指定文件输出流
            out = FileOutputStream(file)
            // 将位图输出到指定文件
            bitmap.compress(
                format, 100,
                out
            )
            out.close()
        } catch (e: IOException) {
            e.printStackTrace()
        }

    }

    interface DownloadedListener {
        fun onFailed(error: String)
        fun onSucceed(fpath: String)
    }

    fun openFile(context: Context, path: String) {
        val intent = Intent(Intent.ACTION_VIEW)
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.N) {
            /* Android N 写法*/
            intent.flags = Intent.FLAG_GRANT_READ_URI_PERMISSION
            val contentUri =
                FileProvider.getUriForFile(context, BuildConfig.APPLICATION_ID + ".fileProvider", File(path))
            intent.setDataAndType(contentUri, getMimeType(path))
        } else {
            /* Android N之前的老版本写法*/
            intent.setDataAndType(Uri.fromFile(File(path)), getMimeType(path))
            intent.flags = Intent.FLAG_ACTIVITY_NEW_TASK
        }
        context.startActivity(intent)
    }

    fun getMimeType(url: String): String? {
        var type: String? = null
        val extension = MimeTypeMap.getFileExtensionFromUrl(url)
        if (extension != null) {
            type = MimeTypeMap.getSingleton().getMimeTypeFromExtension(extension)
        }
        return type
    }
    fun dp2pixel(activity: AppCompatActivity, dps: Float): Float {
        return dps * activity.getResources().getDisplayMetrics().density
    }

    fun pixel2dp(activity: AppCompatActivity, pxs: Float): Float {
        return pxs / activity.getResources().getDisplayMetrics().density
    }

    fun downloadFile(url_str: String, fdist: String, callback: DownloadedListener) {
        Thread(Runnable {
            val f = File(fdist)
            val dir = f.absolutePath.substring(0, f.absolutePath.length - f.name.length)
            File(dir).mkdirs()
            var count: Int
            try {
                val url = URL(url_str)
                val conection = url.openConnection()
                conection.connect()
                val input = BufferedInputStream(url.openStream(), 8192)
                val output = FileOutputStream(fdist)
                val data = ByteArray(1024)
                var total: Long = 0

                while (true) {
                    count = input.read(data)
                    if (count == -1){
                        break
                    }
                    total += count.toLong()
                    //                publishProgress("" + (int) ((total * 100) / lenghtOfFile));
                    output.write(data, 0, count)
                }
                output.flush()
                output.close()
                input.close()
                callback.onSucceed(fdist)
            } catch (e: Exception) {
                Log.e("Error: ", e.message)
                callback.onFailed(e.toString())
            }
        }).start()
    }
}
