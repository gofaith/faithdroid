package io.github.gofaith.faithdroid;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.util.Log;
import android.view.Menu;
import android.view.MenuItem;
import android.widget.FrameLayout;

import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;
import org.json.JSONTokener;

import faithdroid.Faithdroid;
import io.github.gofaith.faithdroid.FViews.FToolbar;
import io.github.gofaith.faithdroid.UI.UIController;

import static io.github.gofaith.faithdroid.UI.Toolkit.parseMenu;

public class MainActivity extends AppCompatActivity {
    private FrameLayout rootview;
    private faithdroid.MainActivity a;
    private String TAG=this.getClass().getName();
    private UIController uiController;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        rootview = findViewById(R.id.main_ctn);
        a=new faithdroid.MainActivity();
        uiController=new UIController(this, rootview);
        a.setUIInterface(uiController);
        a.onCreate();
    }
    @Override
    protected void onStart() {
        super.onStart();
        a.onStart();
    }

    @Override
    protected void onPause() {
        super.onPause();
        a.onPause();
    }

    @Override
    protected void onResume() {
        super.onResume();
        a.onResume();
    }

    @Override
    protected void onDestroy() {
        super.onDestroy();
        a.onDestroy();
    }

    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        if (uiController.optionMenus== null || uiController.optionMenus == "" || uiController.optionMenus == "[]") {
            return false;
        }
        try {
            Log.d(TAG, "onCreateOptionsMenu: "+uiController.optionMenus);
            JSONArray array = (JSONArray) (new JSONTokener(uiController.optionMenus).nextValue());
            parseMenu(uiController,menu,array);
            return true;
        } catch (Exception e) {
            e.printStackTrace();
        }
        return false;
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item) {
        if (uiController.menuItemsOnClickMap.containsKey(item)) {
            Faithdroid.triggerEventHandler(uiController.menuItemsOnClickMap.get(item), "");
            return true;
        }
        return false;
    }
}
