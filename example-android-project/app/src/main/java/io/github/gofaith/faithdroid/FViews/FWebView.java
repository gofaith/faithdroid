package io.github.gofaith.faithdroid.FViews;

import android.net.Uri;
import android.os.Build;
import android.view.View;
import android.webkit.DownloadListener;
import android.webkit.WebResourceRequest;
import android.webkit.WebView;
import android.webkit.WebViewClient;

import org.json.JSONObject;
import org.json.JSONTokener;

import java.util.HashMap;
import java.util.Iterator;
import java.util.List;
import java.util.Map;

import faithdroid.Faithdroid;
import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.UIController;

public class FWebView extends FView implements AttrGettable, AttrSettable {
    public WebView v;

    public FWebView(UIController controller) {
        this.parentController=controller;
        v = new WebView(parentController.activity);
        view=v;
    }
    @Override
    public String getAttr(String attr) {
        switch (attr) {
            case "Visibility":
                return getVisibility();
            case "X":
                return getX();
            case "Y":
                return getY();
            case "Width":
                return getWidth();
            case "Height":
                return getHeight();
            case "PivotX":
                return getPivotX();
            case "PivotY":
                return getPivotY();
            case "ScaleX":
                return getScaleX();
            case "ScaleY":
                return getScaleY();
            case "Rotation":
                return getRotation();
        }
        return null;
    }

    @Override
    public void setAttr(String attr, final String value) {
        if (value==null)
            return;
        switch (attr) {
            case "BackgroundColor":
                setBackgroundColor(value);
                break;
            case "Background":
                setBackground(value);
                break;
            case "Foreground":
                setForeground(value);
                break;
            case "Size":
                parseSize(value);
                break;
            case "X":
                setX(value);
                break;
            case "Y":
                setY(value);
                break;
            case "PivotX":
                setPivotX(value);
                break;
            case "PivotY":
                setPivotY(value);
                break;
            case "ScaleX":
                setScaleX(value);
                break;
            case "ScaleY":
                setScaleY(value);
                break;
            case "Rotation":
                setRotation(value);
                break;
            case "Visibility":
                setVisibility(value);
                break;
            case "Padding":
                setPadding(value);
                break;
            case "Margin":
                setMargin(value);
                break;
            case "LayoutGravity":
                setLayoutGravity(value);
                break;
            case "Elevation":
                setElevation(value);
                break;
            case "LayoutWeight":
                setLayoutWeight(value);
                break;
            case "OnTouch":
                setOnTouchListener(value);
                break;
            case "OnClick":
                v.setOnClickListener(new View.OnClickListener() {
                    @Override
                    public void onClick(View view) {
                        Faithdroid.triggerEventHandler(value,"");
                    }
                });
                break;
            // -------------------------------------------------------------------
            case "Uri":
                v.loadUrl(value);
                break;
            case "Focusable":
                    v.setFocusable(value.equals("true"));
                break;
            case "SupportZoom":
                    v.getSettings().setSupportZoom(value.equals("true"));
                break;
            case "BuiltInZoomControls":
                v.getSettings().setBuiltInZoomControls(value.equals("true"));
                break;
            case "UseWideViewPort":
                v.getSettings().setUseWideViewPort(value.equals("true"));
                break;
            case "AllowContentAccess":
                v.getSettings().setAllowContentAccess(value.equals("true"));
                break;
            case "AllowFileAccess":
                v.getSettings().setAllowFileAccess(value.equals("true"));
                break;
            case "AllowFileAccessFromFileURLs":
                if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.JELLY_BEAN) {
                    v.getSettings().setAllowFileAccessFromFileURLs(value.equals("true"));
                }
                break;
            case "AllowUniversalAccessFromFileURLs":
                if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.JELLY_BEAN) {
                    v.getSettings().setAllowUniversalAccessFromFileURLs(value.equals("true"));
                }
                break;
            case "AppCacheEnabled":
                v.getSettings().setAppCacheEnabled(value.equals("true"));
                break;
            case "BlockNetworkImage":
                v.getSettings().setBlockNetworkImage(value.equals("true"));
                break;
            case "BlockNetworkLoads":
                v.getSettings().setBlockNetworkLoads(value.equals("true"));
                break;
            case "JavaScriptEnabled":
                v.getSettings().setJavaScriptEnabled(value.equals("true"));
                break;
            case "OffscreenPreRaster":
                if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.M) {
                    v.getSettings().setOffscreenPreRaster(value.equals("true"));
                }
                break;
            case "SaveFormData":
                v.getSettings().setSaveFormData(value.equals("true"));
                break;
            case "SupportMultipleWindows":
                v.getSettings().setSupportMultipleWindows(value.equals("true"));
                break;
            case "TextZoom":
                try {
                    v.getSettings().setTextZoom(Integer.parseInt(value));
                } catch (Exception e) {
                    e.printStackTrace();
                }
                break;
            case "UserAgentString":
                v.getSettings().setUserAgentString(value);
                break;
                // -------------------
            case "HandleUrl":
                MyWebViewClient wc=new MyWebViewClient();
                if (!value.equals("")) {
                    try {
                        wc.urlHandlersMap = new HashMap<>();
                        JSONObject object = (JSONObject) (new JSONTokener(value).nextValue());
                        for (Iterator<String> it = object.keys(); it.hasNext(); ) {
                            String key = it.next();
                            wc.urlHandlersMap.put(key, object.getString(key));
                        }
                    } catch (Exception e) {
                        e.printStackTrace();
                    }
                }
                v.setWebViewClient(wc);
                break;
            case "DownloadListener":
                v.setDownloadListener(new DownloadListener() {
                    @Override
                    public void onDownloadStart(String url, String userAgent, String contentDisposition, String mimetype, long contentLength) {
                        Faithdroid.triggerEventHandler(value, url);
                    }
                });
                break;
        }
    }
    class MyWebViewClient extends WebViewClient{
        public HashMap<String,String> urlHandlersMap=null;
        @Override
        public boolean shouldOverrideUrlLoading(WebView view, WebResourceRequest request) {
            if (urlHandlersMap==null) {
                return super.shouldOverrideUrlLoading(view, request);
            }else {
                if (android.os.Build.VERSION.SDK_INT >= android.os.Build.VERSION_CODES.LOLLIPOP) {
                    String url = request.getUrl().toString();
                    if (urlHandlersMap.containsKey(removeQuery(url))) {
                        String fnId = urlHandlersMap.get(removeQuery(url));
                        Faithdroid.triggerEventHandler(fnId, url);
                        return true;
                    }
                }
            }
            return false;
        }
    }

    public static String removeQuery(String url) {
        int i = url.indexOf("?");
        if (i < 0) {
            return url;
        }
        return url.substring(0, i);
    }
}
