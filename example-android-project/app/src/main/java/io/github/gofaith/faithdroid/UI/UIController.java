package io.github.gofaith.faithdroid.UI;

import android.Manifest;
import android.content.ClipData;
import android.content.ClipboardManager;
import android.content.Context;
import android.content.Intent;
import android.content.pm.ApplicationInfo;
import android.content.pm.PackageManager;
import android.graphics.Bitmap;
import android.graphics.Color;
import android.graphics.drawable.Drawable;
import android.media.MediaScannerConnection;
import android.net.Uri;
import android.os.Build;
import android.os.Bundle;
import android.os.Environment;
import android.os.Parcelable;
import android.provider.Settings;
import android.support.design.widget.Snackbar;
import android.support.v4.content.LocalBroadcastManager;
import android.support.v4.view.ViewCompat;
import android.support.v7.app.AppCompatActivity;
import android.support.v7.widget.LinearLayoutManager;
import android.util.Log;
import android.view.MenuItem;
import android.webkit.URLUtil;
import android.widget.FrameLayout;
import android.widget.Toast;

import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;
import org.json.JSONTokener;

import java.net.URISyntaxException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.Iterator;
import java.util.List;
import java.util.Map;

import faithdroid.Faithdroid;
import io.github.gofaith.faithdroid.CoreService;
import io.github.gofaith.faithdroid.FViews.FAlertDialog;
import io.github.gofaith.faithdroid.FViews.FBottomNav;
import io.github.gofaith.faithdroid.FViews.FButton;
import io.github.gofaith.faithdroid.FViews.FCheckBox;
import io.github.gofaith.faithdroid.FViews.FClipboard;
import io.github.gofaith.faithdroid.FViews.FConstraintLayout;
import io.github.gofaith.faithdroid.FViews.FEditText;
import io.github.gofaith.faithdroid.FViews.FFab;
import io.github.gofaith.faithdroid.FViews.FFrameLayout;
import io.github.gofaith.faithdroid.FViews.FHScrollView;
import io.github.gofaith.faithdroid.FViews.FImageView;
import io.github.gofaith.faithdroid.FViews.FLinearLayout;
import io.github.gofaith.faithdroid.FViews.FListView;
import io.github.gofaith.faithdroid.FViews.FPermission;
import io.github.gofaith.faithdroid.FViews.FPopupMenu;
import io.github.gofaith.faithdroid.FViews.FProgressBar;
import io.github.gofaith.faithdroid.FViews.FRadioButton;
import io.github.gofaith.faithdroid.FViews.FRadioGroup;
import io.github.gofaith.faithdroid.FViews.FService;
import io.github.gofaith.faithdroid.FViews.FSnackbar;
import io.github.gofaith.faithdroid.FViews.FSpace;
import io.github.gofaith.faithdroid.FViews.FSpinner;
import io.github.gofaith.faithdroid.FViews.FSwitch;
import io.github.gofaith.faithdroid.FViews.FTabLayout;
import io.github.gofaith.faithdroid.FViews.FTextView;
import io.github.gofaith.faithdroid.FViews.FToolbar;
import io.github.gofaith.faithdroid.FViews.FVScrollView;
import io.github.gofaith.faithdroid.FViews.FValueAnimator;
import io.github.gofaith.faithdroid.FViews.FView;
import io.github.gofaith.faithdroid.FViews.FViewPager;
import io.github.gofaith.faithdroid.FViews.FWebView;
import io.github.gofaith.faithdroid.R;
import io.github.gofaith.faithdroid.SingleInstanceActivity;
import io.github.gofaith.faithdroid.SingleTaskActivity;
import io.github.gofaith.faithdroid.SingleTopActivity;
import io.github.gofaith.faithdroid.StandardActivity;

import static android.content.ContentValues.TAG;
import static io.github.gofaith.faithdroid.UI.Toolkit.openFile;

public class UIController implements faithdroid.UIController{
    private final faithdroid.Activity factivity;
    public AppCompatActivity activity;
    public HashMap<String, FView> viewmap = new HashMap<>();
    public FrameLayout rootView;
    public String optionMenus;
    public Map<MenuItem, String> menuItemsOnClickMap = new HashMap<>();
    public Map<Integer, String> onPermissionResults = new HashMap<>();
    public Map<Integer, OnActivityResultListener> onActivityResults = new HashMap<>();
    public List<Runnable> onDestroyEvent = new ArrayList<>();
    public Map<String, Drawable> drawableMap = new HashMap<>();
    private int FILE_SELECT_CODE=7125;

    public boolean notFinishFlag;
    public String onBackgPressedFn;

    public UIController(AppCompatActivity a, FrameLayout main_ctn,faithdroid.Activity factivity) {
        this.activity=a;
        this.rootView =main_ctn;
        this.factivity=factivity;
    }

    public static List<String> parsePaths(Context context, Intent intent) {
        String action=intent.getAction();
        Bundle extras=intent.getExtras();
        List<String> paths = new ArrayList<>();
        if (Intent.ACTION_SEND.equals(action)) {
            Uri uri = (Uri) extras.getParcelable(Intent.EXTRA_STREAM);
            if (uri==null)return paths;
            String path= null;
            try {
                path = Toolkit.getPath(context,uri);
                paths.add(path);
                return paths;
            } catch (URISyntaxException e) {
                e.printStackTrace();
            }
        }else if (Intent.ACTION_VIEW.equals(action)){
            Uri uri=intent.getData();
            if (uri==null)return paths;
            String path= null;
            try {
                path = Toolkit.getPath(context,uri);
                paths.add(path);
                return paths;
            } catch (URISyntaxException e) {
                e.printStackTrace();
            }
        }else if (Intent.ACTION_SEND_MULTIPLE.equals(action)) {
            ArrayList<Parcelable> list = intent.getParcelableArrayListExtra(Intent.EXTRA_STREAM);
            if (list == null || list.size() == 0) {
                return paths;
            }
            for (Parcelable parcelable: intent.getParcelableArrayListExtra(Intent.EXTRA_STREAM)) {
                Uri uri = (Uri) parcelable;
                if (uri==null)return paths;
                String path= null;
                try {
                    path = Toolkit.getPath(context,uri);
                    paths.add(path);
                } catch (URISyntaxException e) {
                    e.printStackTrace();
                    continue;
                }
            }
            return paths;
        }
        return paths;
    }

    @Override
    public faithdroid.Activity getCurrentFActivity() {
        return factivity;
    }

    @Override
    public String getPkg() {
        return activity.getPackageName();
    }

    @Override
    public void newView(String vName, String vID) {
        FView v = null;
        switch (vName) {
            case "Activity":
                newActivity(vID);
                return;
            case "LinearLayout":
                v= new FLinearLayout(this);
                break;
            case "FrameLayout":
                v= new FFrameLayout(this);
                break;
            case "TextView":
                v= new FTextView(this);
                break;
            case "Button":
                v= new FButton(this);
                break;
            case "EditText":
                v= new FEditText(this);
                break;
            case "VListView":
                LinearLayoutManager llm = new LinearLayoutManager(activity);
                llm.setOrientation(LinearLayoutManager.VERTICAL);
                v= new FListView(this, llm);
                break;
            case "HListView":
                LinearLayoutManager hllm = new LinearLayoutManager(activity);
                hllm.setOrientation(LinearLayoutManager.HORIZONTAL);
                v= new FListView(this, hllm);
                break;
            case "Toolbar":
                v= new FToolbar(this);
                break;
            case "TabLayout":
                v= new FTabLayout(this);
                break;
            case "ViewPager":
                v= new FViewPager(this);
                break;
            case "PopupMenu":
                try {
                    JSONArray array = (JSONArray) (new JSONTokener(vID).nextValue());
                    FPopupMenu popupMenu = new FPopupMenu(this, viewmap.get(array.getString(1)));
                    popupMenu.className = "PopupMenu";
                    popupMenu.vID = array.getString(0);
                    viewmap.put(popupMenu.vID, popupMenu);
                } catch (Exception e) {
                    e.printStackTrace();
                }
                return;
            case "Snackbar":
                v= new FSnackbar(this, rootView);
                break;
            case "AlertDialog":
                v= new FAlertDialog(this);
                break;
            case "Fab":
                v= new FFab(this);
                break;
            case "Space":
                v= new FSpace(this);
                break;
            case "ImageView":
                v= new FImageView(this);
                break;
            case "WebView":
                v= new FWebView(this);
                break;
            case "Switch":
                v= new FSwitch(this);
                break;
            case "ValueAnimator":
                v= new FValueAnimator(this);
                break;
            case "Spinner":
                v = new FSpinner(this);
                break;
            case "VScrollView":
                v = new FVScrollView(this);
                break;
            case "HScrollView":
                v = new FHScrollView(this);
                break;
            case "BottomNav":
                v = new FBottomNav(this);
                break;
            case "Clipboard":
                v = new FClipboard(this);
                break;
            case "RadioGroup":
                v = new FRadioGroup(this);
                break;
            case "RadioButton":
                v=new FRadioButton(this);
                break;
            case "CheckBox":
                v=new FCheckBox(this);
                break;
            case "ConstraintLayout":
                v = new FConstraintLayout(this);
                break;
            case "ProgressBar":
                v = new FProgressBar(this);
                break;
        }
        v.className =vName;
        v.vID=vID;
        viewmap.put(vID, v);
        if (v.view!=null){
            v.view.setId(ViewCompat.generateViewId());
        }
    }

    private void newActivity(String activityConfig) {
        try {
            JSONObject object = (JSONObject) (new JSONTokener(activityConfig).nextValue());
            String launchMode = object.getString("MyLaunchMode");
            String genViewFnId = object.getString("MyFnId");
            if (launchMode == null || genViewFnId == null) {
                return;
            }
            Intent intent = new Intent();
            intent.putExtra("MyFnId", genViewFnId);
            if (launchMode.equals("SingleInstance")) {
                intent.setClass(activity, SingleInstanceActivity.class);
            } else if (launchMode.equals("SingleTask")) {
                intent.setClass(activity, SingleTaskActivity.class);
            } else if (launchMode.equals("SingleTop")) {
                intent.setClass(activity, SingleTopActivity.class);
            } else {
                intent.setClass(activity, StandardActivity.class);
            }
            JSONObject intentObj = object.getJSONObject("MyIntent");
            String gAction=intentObj.getString("Action");
            if (gAction!=null&&!gAction.equals("")) {
                intent.setAction(intentObj.getString("Action"));
            }
            JSONObject extraMap = intentObj.getJSONObject("Extras");
            if (extraMap != null) {
                for (Iterator<String> it = extraMap.keys(); it.hasNext(); ) {
                    String key = it.next();
                    intent.putExtra(key, extraMap.getString(key));
                }
            }
            activity.startActivity(intent);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    @Override
    public void runUIThread(final String s) {
        activity.runOnUiThread(new Runnable() {
            @Override
            public void run() {
                Faithdroid.triggerEventHandler(s,"");
            }
        });
    }

    @Override
    public String viewGetAttr(String vID, String attr) {
        if (vID.equals("Service")) {
            return FService.getAttr(activity,attr);
        }
        if (vID.equals("Permission")) {
            return FPermission.getAttr(this, attr);
        }
        if (vID.equals("Activity")) {
            return activityGet(attr);
        }
        FView v = viewmap.get(vID);
        return ((AttrGettable) v).getAttr(attr);
    }


    @Override
    public void viewSetAttr(String vID, String attr, String value) {
        if (vID.equals("Activity")) {
            activitySet(attr, value);
            return;
        } else if (vID.equals("Service")) {
            FService.setAttr(activity,attr,value);
            return;
        } else if (vID.equals("Permission")) {
            FPermission.setAttr(this, attr, value);
            return;
        }
        FView v = viewmap.get(vID);
        ((AttrSettable)(v)).setAttr(attr,value);
    }

    @Override
    public void append2RootView(String vID) {
        FView v = viewmap.get(vID);
        if (v == null) {
            return;
        }
        rootView.addView(v.view, FFrameLayout.parseLP(v));
    }

    private String activityGet(String attr) {
        if (attr.startsWith("[\"GuessFileName\",")) {
            try {
                JSONArray array = (JSONArray) (new JSONTokener(attr).nextValue());
                String url = array.getString(1);
                return URLUtil.guessFileName(url, null, null);
            } catch (Exception e) {
                e.printStackTrace();
                return "";
            }
        }
        switch (attr) {
            case "Build.MODEL":
                return Build.MODEL;
            case "ExternalStorageDirectory":
                return Environment.getExternalStorageDirectory().getAbsolutePath();
            case "InstalledApplications":
                PackageManager pm = activity.getPackageManager();
                List<ApplicationInfo> list = pm.getInstalledApplications(PackageManager.MATCH_UNINSTALLED_PACKAGES);
                JSONArray array = new JSONArray();
                try{
                    for (int i=0;i<list.size();i++) {
                        JSONObject m = new JSONObject();
                        String fdrawableID = "fdrawable://"+Faithdroid.newToken();
                        drawableMap.put(fdrawableID, pm.getApplicationIcon(list.get(i)).getCurrent());
                        m.put("Icon", fdrawableID);
                        m.put("Title", (String) pm.getApplicationLabel(list.get(i)));
                        m.put("Package", list.get(i).packageName);
                        m.put("SourceDir", list.get(i).sourceDir);
                        String versionname = "";
                        try {
                            versionname = pm.getPackageInfo(list.get(i).packageName, 0).versionName;
                        } catch (PackageManager.NameNotFoundException e) {
                            e.printStackTrace();
                        }
                        m.put("Version", versionname);
                        array.put(m);
                    }
                    String finalStr=array.toString();
                    Log.d(TAG, "activityGet: finalStr="+finalStr);
                    return finalStr;
                } catch (final JSONException e) {
                    activity.runOnUiThread(new Runnable() {
                        @Override
                        public void run() {
                            Toast.makeText(activity,e.toString(),Toast.LENGTH_SHORT).show();
                        }
                    });
                    e.printStackTrace();
                    return "[]";
                }
            case "UniqueID":
                String android_id = Settings.Secure.getString(activity.getContentResolver(),
                        Settings.Secure.ANDROID_ID);
                return android_id;
        }
        return null;
    }
    public interface OnActivityResultListener{
        void onActivityResult(Intent intent);
    }
    private void activitySet(String attr, String value) {
        switch (attr) {
            case "ShowToast":
                Toast.makeText(activity, value, Toast.LENGTH_SHORT).show();
                break;
            case "Finish":
                activity.finish();
                break;
            case "StartUriIntent":
                Intent intent = new Intent();
                intent.setData(Uri.parse(value));
                intent.setAction(Intent.ACTION_VIEW);
                activity.startActivity(intent);
                break;
            case "ScanFile":
                MediaScannerConnection.scanFile(activity,new String[]{value},null,null);
                break;
            case "OpenFile":
                openFile(activity,value);
                break;
            case "SaveDrawableToPNGFile":
                try {
                    JSONArray array = (JSONArray) (new JSONTokener(value).nextValue());
                    String did=array.getString(0);
                    String path = array.getString(1);
                    Toolkit.saveDrawable(drawableMap.get(did), Bitmap.CompressFormat.PNG,path);
                } catch (Exception e) {
                    e.printStackTrace();
                }
                break;
            case "OpenFileChooser":// [ type : "*/*" , allowMultiple : "true" , callback : "218hxjgs861h9cb1298" ]
                try {
                    Intent intent1 = new Intent(Intent.ACTION_GET_CONTENT);
                    final JSONArray array = (JSONArray) (new JSONTokener(value).nextValue());
                    intent1.setType(array.getString(0));
                    intent1.addCategory(Intent.CATEGORY_OPENABLE);
                    intent1.putExtra(Intent.EXTRA_ALLOW_MULTIPLE, array.getString(1).equals("true"));
                    activity.startActivityForResult(Intent.createChooser(intent1, ""),FILE_SELECT_CODE);
                    final String fnID = array.getString(2);
                    onActivityResults.put(FILE_SELECT_CODE, new OnActivityResultListener() {
                        @Override
                        public void onActivityResult(Intent data) {
                            if (data == null) {
                                return;
                            }
                            Uri uri = data.getData();
                            JSONArray array1 = new JSONArray();
                            try {
                                if (uri != null) {
                                    String path = null;
                                    if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.KITKAT) {
                                        path = Toolkit.getPathByUri4kitkat(activity, uri);
                                    }else {
                                        path=Toolkit.getPathByUriBelowKitkat(activity,uri);
                                    }
                                    if (path != null) {
                                        array1.put(path);
                                    }
                                } else {
                                    if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.JELLY_BEAN) {
                                        ClipData clipData = data.getClipData();
                                        if (clipData != null) {
                                            for (int i=0;i<clipData.getItemCount();i++) {
                                                ClipData.Item item = clipData.getItemAt(i);
                                                Uri uri1=item.getUri();
                                                String url = null;
                                                if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.KITKAT) {
                                                    url = Toolkit.getPathByUri4kitkat(activity, uri1);
                                                }else {
                                                    url=Toolkit.getPathByUriBelowKitkat(activity,uri1);
                                                }
                                                if (url!=null)
                                                    array1.put(url);
                                            }
                                        }
                                    }else {
                                        Toast.makeText(activity,"No file manager installed",Toast.LENGTH_LONG).show();
                                    }
                                }
                                if (array1.length() == 0) {
                                    Toast.makeText(activity,"Failed",Toast.LENGTH_LONG).show();
                                }else {
                                    Faithdroid.triggerEventHandler(fnID, array1.toString());
                                }
                            }catch (Exception e){
                                e.printStackTrace();
                                Toast.makeText(activity,e.toString(),Toast.LENGTH_LONG).show();
                            }
                        }
                    });
                } catch (Exception e) {
                    e.printStackTrace();
                    Toast.makeText(activity,e.toString(),Toast.LENGTH_LONG).show();
                }
                break;
            case "BackPressed":
                activity.onBackPressed();
                break;
            case "StatusBarColor":
                if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.LOLLIPOP) {
                    try {
                        activity.getWindow().setStatusBarColor(Color.parseColor(value));
                    }catch (Exception e){
                        e.printStackTrace();
                    }
                }
                break;
            case "NavigationBarColor":
                try{
                    if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.LOLLIPOP) {
                        activity.getWindow().setNavigationBarColor(Color.parseColor(value));
                    }
                }catch (Exception e){
                    e.printStackTrace();
                }
                break;
            case "NotFinishFlag":
                notFinishFlag = value.equals("true");
                break;
            case "FinishAllActivity":
                Intent i = new Intent("uibro");
                i.putExtra("action", "quit");
                LocalBroadcastManager.getInstance(activity).sendBroadcast(i);
                break;
            case "KillAll":
                CoreService.killAll();
                break;
            case "OnBackPressed":
                onBackgPressedFn = value;
                break;
        }
    }
}
