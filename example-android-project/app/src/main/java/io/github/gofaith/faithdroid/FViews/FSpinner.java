package io.github.gofaith.faithdroid.FViews;

import android.view.View;
import android.widget.AdapterView;
import android.widget.ArrayAdapter;
import android.widget.Spinner;

import org.json.JSONArray;
import org.json.JSONTokener;

import java.util.ArrayList;
import java.util.List;

import faithdroid.Faithdroid;
import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.UIController;

public class FSpinner extends FView implements AttrGettable,AttrSettable {
    public Spinner v;
    private List<String> data_list = new ArrayList<>();
    private ArrayAdapter<String> adapter =null;
    private int selected=-1;
    public FSpinner(UIController controller) {
        parentController = controller;
        v = new Spinner(parentController.activity);
        view=v;
        adapter = new ArrayAdapter<>(parentController.activity, android.R.layout.simple_spinner_dropdown_item, data_list);
        v.setAdapter(adapter);
    }
    @Override
    public String getAttr(String attr) {
        switch (attr) {
            case "Visibility":
                return getVisibility();
            case "X":
                return getX();
            case "Y":
                return getY();
            case "Width":
                return getWidth();
            case "Height":
                return getHeight();
            case "PivotX":
                return getPivotX();
            case "PivotY":
                return getPivotY();
            case "ScaleX":
                return getScaleX();
            case "ScaleY":
                return getScaleY();
            case "Rotation":
                return getRotation();
            // ------------------------------------------
            case "Enabled":
                return String.valueOf(v.isEnabled());
        }
        return "";
    }

    @Override
    public void setAttr(String attr, final String value) {
        if (value==null)
            return;
        switch (attr) {
            case "BackgroundColor":
                setBackgroundColor(value);
                break;
            case "Background":
                setBackground(value);
                break;
            case "Foreground":
                setForeground(value);
                break;
            case "Size":
                parseSize(value);
                break;
            case "X":
                setX(value);
                break;
            case "Y":
                setY(value);
                break;
            case "PivotX":
                setPivotX(value);
                break;
            case "PivotY":
                setPivotY(value);
                break;
            case "ScaleX":
                setScaleX(value);
                break;
            case "ScaleY":
                setScaleY(value);
                break;
            case "Rotation":
                setRotation(value);
                break;
            case "Visibility":
                setVisibility(value);
                break;
            case "Padding":
                setPadding(value);
                break;
            case "Margin":
                setMargin(value);
                break;
            case "LayoutGravity":
                setLayoutGravity(value);
                break;
            case "Elevation":
                setElevation(value);
                break;
            case "LayoutWeight":
                setLayoutWeight(value);
                break;
            case "OnTouch":
                setOnTouchListener(value);
                break;
            case "OnClick":
                v.setOnClickListener(new View.OnClickListener() {
                    @Override
                    public void onClick(View view) {
                        Faithdroid.triggerEventHandler(value, "");
                    }
                });
                break;
            // -------------------------------------------------------------------
            case "NotifyDataSetChanged":
                adapter.notifyDataSetChanged();
                break;
            case "List":
                try {
                    JSONArray array = (JSONArray) (new JSONTokener(value).nextValue());
                    data_list.clear();
                    for (int i=0;i<array.length();i++) {
                        data_list.add(array.getString(i));
                    }
                    adapter.notifyDataSetChanged();
                } catch (Exception e) {
                    e.printStackTrace();
                }
                break;
            case "ListAdd":
                data_list.add(value);
                adapter.notifyDataSetChanged();
                break;
            case "ListRemove":
                data_list.remove(Integer.parseInt(value));
                adapter.notifyDataSetChanged();
                break;
            case "OnItemClick":
                v.setOnItemSelectedListener(new AdapterView.OnItemSelectedListener() {
                    @Override
                    public void onItemSelected(AdapterView<?> parent, View view, int position, long id) {
                        Faithdroid.triggerEventHandler(value, String.valueOf(position));
                    }

                    @Override
                    public void onNothingSelected(AdapterView<?> parent) {

                    }
                });
                break;
            case "Enabled":
                v.setEnabled(value.equals("true"));
        }
    }
}
