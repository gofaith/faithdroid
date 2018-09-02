package io.github.gofaith.faithdroid;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.Menu;
import android.widget.FrameLayout;

import org.json.JSONArray;
import org.json.JSONTokener;

import faithdroid.Activity;
import faithdroid.Faithdroid;
import io.github.gofaith.faithdroid.UI.UIController;

import static io.github.gofaith.faithdroid.UI.Toolkit.parseMenu;

public class StandardActivity extends AppCompatActivity {
    private FrameLayout ctn;
    private Activity a;
    private String fnId,optionMenuStr;
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_standard);
        fnId = getIntent().getStringExtra("FnId");
        ctn = findViewById(R.id.standard_ctn);
        a=new Activity();
        String uId=a.setUIInterface(new UIController(this, ctn, new UIController.OtherMethods() {
            @Override
            public void setOptionMenu(String string) {
                optionMenuStr=string;
            }
        }));
        a.onCreate();
        Faithdroid.triggerEventHandler(fnId, uId);
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
        if (optionMenuStr == null || optionMenuStr == "" || optionMenuStr == "[]") {
            return false;
        }
        try {
            JSONArray array = (JSONArray) (new JSONTokener(optionMenuStr).nextValue());
            parseMenu(menu,array);
            return true;
        } catch (Exception e) {
            e.printStackTrace();
        }
        return false;
    }
}
