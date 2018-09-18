package io.github.gofaith.faithdroid;

import android.content.Intent;
import android.content.pm.PackageManager;
import android.support.annotation.NonNull;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.Menu;
import android.view.MenuItem;
import android.widget.FrameLayout;

import org.json.JSONArray;
import org.json.JSONTokener;

import java.util.List;

import faithdroid.Faithdroid;
import io.github.gofaith.faithdroid.UI.UIController;

import static io.github.gofaith.faithdroid.UI.Toolkit.parseMenu;

public class MainActivity extends AppCompatActivity {
    private FrameLayout rootview;
    private faithdroid.MainActivity a;
    private String TAG=this.getClass().getName();
    private UIController parentUIController;
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        rootview = findViewById(R.id.main_ctn);
        a=new faithdroid.MainActivity();
        parentUIController =new UIController(this, rootview,null);
        a.setUIInterface(parentUIController);

        handleIntent();
        a.onCreate();
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
        for (int i = 0; i< parentUIController.onDestroyEvent.size(); i++) {
            runOnUiThread(parentUIController.onDestroyEvent.get(i));
        }
    }

    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        if (parentUIController.optionMenus== null || parentUIController.optionMenus == "" || parentUIController.optionMenus == "[]") {
            return false;
        }
        try {
            JSONArray array = (JSONArray) (new JSONTokener(parentUIController.optionMenus).nextValue());
            parseMenu(parentUIController,menu,array);
            return true;
        } catch (Exception e) {
            e.printStackTrace();
        }
        return false;
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item) {
        if (parentUIController.menuItemsOnClickMap.containsKey(item)) {
            Faithdroid.triggerEventHandler(parentUIController.menuItemsOnClickMap.get(item), "");
            return true;
        }
        return false;
    }

    @Override
    public void onRequestPermissionsResult(int requestCode, @NonNull String[] permissions, @NonNull int[] grantResults) {
        if (parentUIController.onPermissionResults.containsKey(requestCode)) {
            JSONArray array = new JSONArray();
            for(int i=0;i<grantResults.length;i++) {
                array.put(grantResults[i] == PackageManager.PERMISSION_GRANTED);
            }
            Faithdroid.triggerEventHandler(parentUIController.onPermissionResults.get(requestCode), array.toString());
        }
    }
}
