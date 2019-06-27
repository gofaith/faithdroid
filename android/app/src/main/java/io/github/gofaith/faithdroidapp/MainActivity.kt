package io.github.gofaith.faithdroidapp

import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.widget.FrameLayout
import faithdroid.MainWindow
import io.github.gofaith.faithdroidapp.UI.UIKotlinBridge

class MainActivity : AppCompatActivity() {

    private lateinit var w: MainWindow
    private lateinit var ui: UIKotlinBridge

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)
        val rootView = findViewById<FrameLayout>(R.id.ctn)
        w = MainWindow()
        ui = UIKotlinBridge(this,rootView)
        w.setUI(ui)
        w.onCreate()
    }
}
