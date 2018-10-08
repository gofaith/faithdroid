package io.github.gofaith.faithdroid.UI;

import android.Manifest;
import android.content.ClipboardManager;
import android.content.Context;
import android.content.Intent;
import android.media.MediaScannerConnection;
import android.net.Uri;
import android.os.Build;
import android.os.Bundle;
import android.os.Environment;
import android.os.Parcelable;
import android.support.design.widget.Snackbar;
import android.support.v7.app.AppCompatActivity;
import android.support.v7.widget.LinearLayoutManager;
import android.view.MenuItem;
import android.webkit.URLUtil;
import android.widget.FrameLayout;
import android.widget.Toast;

import org.json.JSONArray;
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
import io.github.gofaith.faithdroid.FViews.FEditText;
import io.github.gofaith.faithdroid.FViews.FFab;
import io.github.gofaith.faithdroid.FViews.FFrameLayout;
import io.github.gofaith.faithdroid.FViews.FHScrollView;
import io.github.gofaith.faithdroid.FViews.FImageView;
import io.github.gofaith.faithdroid.FViews.FLinearLayout;
import io.github.gofaith.faithdroid.FViews.FListView;
import io.github.gofaith.faithdroid.FViews.FPermission;
import io.github.gofaith.faithdroid.FViews.FPopupMenu;
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
import io.github.gofaith.faithdroid.SingleInstanceActivity;
import io.github.gofaith.faithdroid.SingleTaskActivity;
import io.github.gofaith.faithdroid.SingleTopActivity;
import io.github.gofaith.faithdroid.StandardActivity;

import static io.github.gofaith.faithdroid.UI.Toolkit.openFile;

public class UIController implements faithdroid.UIController{
    private final faithdroid.Activity factivity;
    public AppCompatActivity activity;
    public HashMap<String, FView> viewmap = new HashMap<>();
    public FrameLayout rootView;
    public String optionMenus;
    public Map<MenuItem, String> menuItemsOnClickMap = new HashMap<>();
    public Map<Integer, String> onPermissionResults = new HashMap<>();
    public List<Runnable> onDestroyEvent = new ArrayList<>();
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
                FLinearLayout fLinearLayout = new FLinearLayout(this);
                v=fLinearLayout;
                break;
            case "FrameLayout":
                FFrameLayout fFrameLayout = new FFrameLayout(this);
                v=fFrameLayout;
                break;
            case "TextView":
                FTextView textView = new FTextView(this);
                v=textView;
                break;
            case "Button":
                FButton button = new FButton(this);
                v=button;
                break;
            case "EditText":
                FEditText editText = new FEditText(this);
                v=editText;
                break;
            case "VListView":
                LinearLayoutManager llm = new LinearLayoutManager(activity);
                llm.setOrientation(LinearLayoutManager.VERTICAL);
                FListView listView = new FListView(this, llm);
                v=listView;
                break;
            case "HListView":
                LinearLayoutManager hllm = new LinearLayoutManager(activity);
                hllm.setOrientation(LinearLayoutManager.HORIZONTAL);
                FListView hlistView = new FListView(this, hllm);
                v=hlistView;
                break;
            case "Toolbar":
                FToolbar fToolbar = new FToolbar(this);
                v=fToolbar;
                break;
            case "TabLayout":
                FTabLayout fTabLayout = new FTabLayout(this);
                v=fTabLayout;
                break;
            case "ViewPager":
                FViewPager fViewPager = new FViewPager(this);
                v=fViewPager;
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
                FSnackbar fSnackbar = new FSnackbar(this, rootView);
                v=fSnackbar;
                break;
            case "AlertDialog":
                FAlertDialog fAlertDialog = new FAlertDialog(this);
                v=fAlertDialog;
                break;
            case "Fab":
                FFab fFab = new FFab(this);
                v=fFab;
                break;
            case "Space":
                FSpace space = new FSpace(this);
                v=space;
                break;
            case "ImageView":
                FImageView imageView = new FImageView(this);
                v=imageView;
                break;
            case "WebView":
                FWebView webView = new FWebView(this);
                v=webView;
                break;
            case "Switch":
                FSwitch fSwitch = new FSwitch(this);
                v=fSwitch;
                break;
            case "ValueAnimator":
                FValueAnimator valueAnimator = new FValueAnimator(this);
                v=valueAnimator;
                break;
            case "Spinner":
                FSpinner spinner = new FSpinner(this);
                v = spinner;
                break;
            case "VScrollView":
                FVScrollView fvScrollView = new FVScrollView(this);
                v = fvScrollView;
                break;
            case "HScrollView":
                FHScrollView fhScrollView = new FHScrollView(this);
                v = fhScrollView;
                break;
            case "BottomNav":
                FBottomNav fBottomNav = new FBottomNav(this);
                v = fBottomNav;
                break;
            case "Clipboard":
                FClipboard clipboard = new FClipboard(this);
                v = clipboard;
                break;
            case "RadioGroup":
                FRadioGroup fRadioGroup = new FRadioGroup(this);
                v = fRadioGroup;
                break;
            case "RadioButton":
                FRadioButton fRadioButton = new FRadioButton(this);
                v = fRadioButton;
                break;
            case "CheckBox":
                FCheckBox fCheckBox = new FCheckBox(this);
                v = fCheckBox;
                break;
        }
        v.className =vName;
        v.vID=vID;
        viewmap.put(vID, v);
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
    public void runOnUIThread(final String s) {
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
    public void showOnRootView(String vID) {
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
        }
        return null;
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
        }
    }
}
