package io.github.gofaith.faithdroid;

import android.content.Intent;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.util.Log;
import android.view.Menu;
import android.view.MenuItem;
import android.widget.FrameLayout;

import org.json.JSONArray;
import org.json.JSONTokener;

import java.util.List;

import faithdroid.Activity;
import faithdroid.Faithdroid;
import io.github.gofaith.faithdroid.UI.UIController;

import static io.github.gofaith.faithdroid.UI.Toolkit.parseMenu;

public class StandardActivity extends AppCompatActivity {
    private FrameLayout ctn;
    private Activity a;
    private String fnId;
    private UIController uiController;
    private String TAG=this.getClass().getName();

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_standard);
        fnId = getIntent().getStringExtra("MyFnId");
        ctn = findViewById(R.id.standard_ctn);
        a=new Activity();
        uiController=new UIController(this, ctn,a);
        String uId=a.setUIInterface(uiController);

        handleIntent();

        a.onCreate();
        Faithdroid.triggerEventHandler(fnId, uId);
    }
    private void handleIntent() {
        Intent intent=getIntent();
        a.setIntentAction(intent.getAction());
        List<String> ps = UIController.parsePaths(this, intent);
        for (int j = 0; j < ps.size(); j++) {
            a.addPath(ps.get(j));
        }
        Bundle bundle=intent.getExtras();
        if (bundle != null) {
            for (String key : bundle.keySet()) {
                a.putExtra(key, String.valueOf(bundle.get(key)));
            }
        }
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
