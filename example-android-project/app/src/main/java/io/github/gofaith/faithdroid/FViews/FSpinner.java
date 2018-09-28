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
        String str = getUniversalAttr(attr);
        if (str != null) {
            return str;
        }
        switch (attr){
            // ------------------------------------------
            case "Enabled":
                return String.valueOf(v.isEnabled());
        }
        return "";
    }

    @Override
    public void setAttr(String attr, final String value) {
        if (setUniversalAttr(attr, value)) {
            return;
        }
        switch (attr){
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
