package io.github.gofaith.faithdroidapp

import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import faithdroid.MainWindow
import io.github.gofaith.faithdroidapp.UI.UIKotlinBridge

class MainActivity : AppCompatActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)
        val w = MainWindow()
        val ui = UIKotlinBridge()
        w.setUI(ui)
        w.onCreate()
    }
}
