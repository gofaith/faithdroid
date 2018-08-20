package io.github.gofaith.faithdroid;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.widget.FrameLayout;

import io.github.gofaith.faithdroid.UI.UIController;

public class MainActivity extends AppCompatActivity {
    private FrameLayout rootview;
    private faithdroid.MainActivity a;
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        rootview = findViewById(R.id.main_ctn);
        a=new faithdroid.MainActivity();
        a.setUIInterface(new UIController(this, rootview));
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
}
