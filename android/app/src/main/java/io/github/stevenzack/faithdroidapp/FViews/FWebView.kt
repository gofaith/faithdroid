package io.github.stevenzack.faithdroidapp.FViews

import android.content.ContentValues.TAG
import android.os.Build
import android.util.Log
import android.view.ViewGroup
import android.webkit.WebResourceRequest
import android.webkit.WebView
import android.webkit.WebViewClient
import faithdroid.Faithdroid
import io.github.stevenzack.faithdroidapp.UI.AttrGettable
import io.github.stevenzack.faithdroidapp.UI.AttrSettable
import io.github.stevenzack.faithdroidapp.UI.FView
import io.github.stevenzack.faithdroidapp.UI.UIController
import org.json.JSONObject
import org.json.JSONTokener


class FWebView(controller: UIController) : FView(), AttrGettable, AttrSettable {
    var v: WebView? = null

    init {
        this.parentController = controller
        v = WebView(parentController!!.activity)
        view = v
        v!!.settings.defaultTextEncodingName = "UTF -8"
        parentController!!.onDestoryEvent.add( {
            if (v != null) {
                val parent = v!!.parent as ViewGroup
                parent?.removeView(v)
                v!!.stopLoading()
                v!!.settings.javaScriptEnabled = false
                v!!.clearHistory()
                v!!.clearView()
                v!!.removeAllViews()
                v!!.destroy()
                v = null
            }
            parentController!!.viewmap.remove(vID)
        })
    }

    override fun getAttr(attr: String): String{
        val str = getUniversalAttr(attr)
        if (str != null) {
            return str
        }
        when (attr) {

        }
        return ""
    }

    override fun setAttr(attr: String, value: String) {
        if (setUniversalAttr(attr, value)) {
            return
        }
        when (attr) {
            // -------------------------------------------------------------------
            "Uri" -> v!!.loadUrl(value)
            "Focusable" -> v!!.isFocusable = value == "true"
            "SupportZoom" -> v!!.settings.setSupportZoom(value == "true")
            "BuiltInZoomControls" -> v!!.settings.builtInZoomControls = value == "true"
            "UseWideViewPort" -> v!!.settings.useWideViewPort = value == "true"
            "AllowContentAccess" -> v!!.settings.allowContentAccess = value == "true"
            "AllowFileAccess" -> v!!.settings.allowFileAccess = value == "true"
            "AllowFileAccessFromFileURLs" -> if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.JELLY_BEAN) {
                v!!.settings.allowFileAccessFromFileURLs = value == "true"
            }
            "AllowUniversalAccessFromFileURLs" -> if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.JELLY_BEAN) {
                v!!.settings.allowUniversalAccessFromFileURLs = value == "true"
            }
            "AppCacheEnabled" -> v!!.settings.setAppCacheEnabled(value == "true")
            "BlockNetworkImage" -> v!!.settings.blockNetworkImage = value == "true"
            "BlockNetworkLoads" -> v!!.settings.blockNetworkLoads = value == "true"
            "JavaScriptEnabled" -> v!!.settings.javaScriptEnabled = value == "true"
            "OffscreenPreRaster" -> if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.M) {
                v!!.settings.offscreenPreRaster = value == "true"
            }
            "SaveFormData" -> v!!.settings.saveFormData = value == "true"
            "SupportMultipleWindows" -> v!!.settings.setSupportMultipleWindows(value == "true")
            "TextZoom" -> try {
                v!!.settings.textZoom = Integer.parseInt(value)
            } catch (e: Exception) {
                e.printStackTrace()
            }

            "UserAgentString" -> v!!.settings.setUserAgentString(value)
            // -------------------
            "HandleUrl" -> {
                val wc = MyWebViewClient()
                if (value != "") {
                    try {
                        wc.urlHandlersMap = HashMap()
                        val `object` = JSONTokener(value).nextValue() as JSONObject
                        Log.d(TAG, "setAttr: HandleUrl$value")
                        val it = `object`.keys()
                        while (it.hasNext()) {
                            val key = it.next()
                            wc.urlHandlersMap!![key] = `object`.getString(key)
                        }
                    } catch (e: Exception) {
                        e.printStackTrace()
                    }

                }
                v!!.webViewClient = wc
            }
            "DownloadListener" -> v!!.setDownloadListener { url, userAgent, contentDisposition, mimetype, contentLength ->
                Faithdroid.triggerEventHandler(
                    value,
                    url
                )
            }
            "LoadHtmlData" -> {
                val a = value.replace("#", "%23")
                v!!.loadData(a, "text/html; charset=UTF-8", null)
            }
        }
    }

    internal inner class MyWebViewClient : WebViewClient() {
        var urlHandlersMap: HashMap<String, String>? = null
        override fun shouldOverrideUrlLoading(view: WebView, request: WebResourceRequest): Boolean {
            if (urlHandlersMap == null) {
                return super.shouldOverrideUrlLoading(view, request)
            } else {
                Log.d(TAG, "shouldOverrideUrlLoading: " + urlHandlersMap!!.size)
                if (android.os.Build.VERSION.SDK_INT >= android.os.Build.VERSION_CODES.LOLLIPOP) {
                    val url = request.url.toString()
                    if (urlHandlersMap!!.containsKey(removeQuery(url))) {
                        val fnId = urlHandlersMap!![removeQuery(url)]
                        val s = Faithdroid.triggerEventHandler(fnId, url)
                        if (s != null && s == "true")
                            return true
                    } else if (urlHandlersMap!!.containsKey(url)) {
                        val fnId = urlHandlersMap!![url]
                        val s = Faithdroid.triggerEventHandler(fnId, url)
                        if (s != null && s == "true")
                            return true
                    }
                }
            }
            return false
        }

        override fun shouldOverrideUrlLoading(view: WebView, url: String): Boolean {
            if (urlHandlersMap == null) {
                return super.shouldOverrideUrlLoading(view, url)
            } else {
                Log.d(TAG, "shouldOverrideUrlLoading2: " + urlHandlersMap!!.size)
                if (android.os.Build.VERSION.SDK_INT >= android.os.Build.VERSION_CODES.LOLLIPOP) {
                    if (urlHandlersMap!!.containsKey(removeQuery(url))) {
                        val fnId = urlHandlersMap!![removeQuery(url)]
                        val s = Faithdroid.triggerEventHandler(fnId, url)
                        if (s != null && s == "true")
                            return true
                    } else if (urlHandlersMap!!.containsKey(url)) {
                        val fnId = urlHandlersMap!![url]
                        val s = Faithdroid.triggerEventHandler(fnId, url)
                        if (s != null && s == "true")
                            return true
                    }
                }
            }
            return false
        }
    }

    companion object {

        fun removeQuery(url: String): String {
            val i = url.indexOf("?")
            return if (i < 0) {
                url
            } else url.substring(0, i)
        }
    }
}
