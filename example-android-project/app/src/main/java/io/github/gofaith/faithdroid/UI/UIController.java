package io.github.gofaith.faithdroid.UI;

import android.content.Intent;
import android.support.design.widget.Snackbar;
import android.support.v7.app.AppCompatActivity;
import android.support.v7.widget.LinearLayoutManager;
import android.view.MenuItem;
import android.widget.FrameLayout;
import android.widget.Toast;

import org.json.JSONArray;
import org.json.JSONObject;
import org.json.JSONTokener;

import java.util.HashMap;
import java.util.Map;

import faithdroid.Activity;
import faithdroid.Faithdroid;
import io.github.gofaith.faithdroid.FViews.FAlertDialog;
import io.github.gofaith.faithdroid.FViews.FButton;
import io.github.gofaith.faithdroid.FViews.FEditText;
import io.github.gofaith.faithdroid.FViews.FFab;
import io.github.gofaith.faithdroid.FViews.FFrameLayout;
import io.github.gofaith.faithdroid.FViews.FImageView;
import io.github.gofaith.faithdroid.FViews.FLinearLayout;
import io.github.gofaith.faithdroid.FViews.FListView;
import io.github.gofaith.faithdroid.FViews.FPopupMenu;
import io.github.gofaith.faithdroid.FViews.FSnackbar;
import io.github.gofaith.faithdroid.FViews.FSpace;
import io.github.gofaith.faithdroid.FViews.FTabLayout;
import io.github.gofaith.faithdroid.FViews.FTextView;
import io.github.gofaith.faithdroid.FViews.FToolbar;
import io.github.gofaith.faithdroid.FViews.FView;
import io.github.gofaith.faithdroid.FViews.FViewPager;
import io.github.gofaith.faithdroid.SingleInstanceActivity;
import io.github.gofaith.faithdroid.SingleTaskActivity;
import io.github.gofaith.faithdroid.SingleTopActivity;
import io.github.gofaith.faithdroid.StandardActivity;

public class UIController implements faithdroid.UIController{
    public AppCompatActivity activity;
    public HashMap<String, FView> viewmap = new HashMap<>();
    public FrameLayout rootView;
    public String optionMenus;
    public Map<MenuItem, String> menuItemsOnClickMap = new HashMap<>();
    public UIController(AppCompatActivity a, FrameLayout main_ctn) {
        this.activity=a;
        this.rootView =main_ctn;
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
        }
        v.className =vName;
        v.vID=vID;
        viewmap.put(vID, v);
    }

    private void newActivity(String activityConfig) {
        try {
            JSONObject object = (JSONObject) (new JSONTokener(activityConfig).nextValue());
            String launchMode = object.getString("LaunchMode");
            String genViewFnId = object.getString("FnId");
            if (launchMode == null || genViewFnId == null) {
                return;
            }
            Intent intent = new Intent();
            intent.putExtra("FnId", genViewFnId);
            if (launchMode.equals("SingleInstance")) {
                intent.setClass(activity, SingleInstanceActivity.class);
            } else if (launchMode.equals("SingleTask")) {
                intent.setClass(activity, SingleTaskActivity.class);
            } else if (launchMode.equals("SingleTop")) {
                intent.setClass(activity, SingleTopActivity.class);
            } else {
                intent.setClass(activity, StandardActivity.class);
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
        FView v = viewmap.get(vID);
        return ((AttrGettable) v).getAttr(attr);
    }

    @Override
    public void viewSetAttr(String vID, String attr, String value) {
        if (vID.equals("Activity")) {
            activitySet(attr, value);
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
        FFrameLayout.addToframeLayout(rootView,v);
    }

    private void activitySet(String attr, String value) {
        switch (attr) {
            case "ShowToast":
                Toast.makeText(activity, value, Toast.LENGTH_SHORT).show();
                break;
            case "Finish":
                activity.finish();
                break;
        }
    }
}
