package io.github.gofaith.faithdroid;

import android.support.v7.app.AppCompatActivity;
import android.view.View;

import java.util.HashMap;

import faithdroid.UIInterface;

public class UIComponent implements UIInterface {
    private AppCompatActivity a;
    private HashMap<String, MyView> viewmap = new HashMap<>();
    public UIComponent(AppCompatActivity a) {
        this.a=a;
    }

    @Override
    public void newView(String s, String s1) {
        switch (s) {
            case "TextView":

                break;
        }
    }

    @Override
    public void viewDoFn(String s, String s1, String s2) {

    }

    @Override
    public String viewGetAttr(String s, String s1) {
        return null;
    }

    @Override
    public void viewSetAttr(String s, String s1, String s2) {

    }

    private class MyView {
        public String viewClass;
        public View view;
    }
}
