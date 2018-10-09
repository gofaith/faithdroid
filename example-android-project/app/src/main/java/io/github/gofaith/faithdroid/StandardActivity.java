package io.github.gofaith.faithdroid;

import android.content.Intent;
import android.content.pm.PackageManager;
import android.support.annotation.NonNull;
import android.support.annotation.Nullable;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
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
    private UIController parentUIController;
    private String TAG=this.getClass().getName();

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_standard);
        fnId = getIntent().getStringExtra("MyFnId");
        ctn = findViewById(R.id.standard_ctn);
        a=new Activity();
        parentUIController =new UIController(this, ctn,a);
        String uId=a.setUIInterface(parentUIController);

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
    @Override
    protected void onActivityResult(int requestCode, int resultCode, @Nullable Intent data) {
        if (parentUIController.onActivityResults.containsKey(requestCode)) {
            parentUIController.onActivityResults.get(requestCode).onActivityResult(data);
        }
    }
}
