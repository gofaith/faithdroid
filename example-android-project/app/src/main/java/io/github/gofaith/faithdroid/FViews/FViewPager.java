package io.github.gofaith.faithdroid.FViews;

import android.os.Bundle;
import android.support.design.widget.TabLayout;
import android.support.v4.app.Fragment;
import android.support.v4.app.FragmentManager;
import android.support.v4.app.FragmentStatePagerAdapter;
import android.support.v4.view.ViewCompat;
import android.support.v4.view.ViewPager;
import android.util.Log;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;

import org.json.JSONArray;
import org.json.JSONObject;
import org.json.JSONTokener;

import java.util.ArrayList;
import java.util.List;

import faithdroid.Faithdroid;
import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.UIController;

public class FViewPager extends FView implements AttrSettable, AttrGettable {
    private FaithCollectionPagerAdapter adapter;
    public ViewPager v;
    List<FPage> pages = new ArrayList<>();
    public FViewPager(UIController controller) {
        parentController = controller;
        v = new ViewPager(parentController.activity);
        v.setId(ViewCompat.generateViewId());
        view=v;
        parseSize(parentController.activity,v,"[-1,-1]");
    }

    @Override
    public String getAttr(String attr) {
        if(attr==null)
            return "";
        switch (attr) {
            case "Visibility":
                return getVisibility();
            // --------------------------------------------------
        }
        return "";
    }

    @Override
    public void setAttr(String attr, String value) {
        if (attr == null || value == null) {
            return;
        }
        switch (attr) {
            case "BackgroundColor":
                setBackgroundColor(value);
                break;
            case "Background":
                setBackground(value);
                break;
            case "Size":
                parseSize(parentController.activity, v, value);
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
            // ----------------------------------------------------------------------------
            case "Pages":
                Log.d(TAG, "setAttr: "+value);
                try {
                    JSONArray array = (JSONArray) (new JSONTokener(value).nextValue());
                    for(int i=0;i<array.length();i++){
                        JSONObject object = array.getJSONObject(i);
                        FPage fPage=new FPage();
                        fPage.vid = object.getString("VID");
                        pages.add(fPage);
                    }
                    adapter = new FaithCollectionPagerAdapter(parentController.activity.getSupportFragmentManager(), this);
                    v.setAdapter(adapter);
                } catch (Exception e) {
                    e.printStackTrace();
                }
                break;
            case "TabLayout":
                FView o = parentController.viewmap.get(value);
                if (o==null)break;
                FTabLayout fTabLayout=(FTabLayout)o;
                fTabLayout.v.setupWithViewPager(v);
                break;
        }
    }
    // ------------------
    class FPage{
        String vid, onSelected;
    }
    class FaithCollectionPagerAdapter extends FragmentStatePagerAdapter {
        private final FViewPager fviewPager;

        public FaithCollectionPagerAdapter(FragmentManager fm, FViewPager fViewPager) {
            super(fm);
            this.fviewPager =fViewPager;
        }
        @Override
        public Fragment getItem(int i) {
            Fragment fragment = new DemoObjectFragment().setFviewPager(fviewPager);
            Bundle args = new Bundle();
            // Our object is just an integer :-P
            args.putInt(DemoObjectFragment.ARG_OBJECT, i);
            fragment.setArguments(args);
            return fragment;
        }
        @Override
        public int getCount() {
            return fviewPager.pages.size();
        }
        @Override
        public CharSequence getPageTitle(int position) {
            return "OBJECT " + (position + 1);
        }
    }
    public static class DemoObjectFragment extends Fragment {
        public static final String ARG_OBJECT = "object";
        private FViewPager fviewPager;

        public DemoObjectFragment setFviewPager(FViewPager fviewPager) {
            this.fviewPager=fviewPager;
            return this;
        }

        @Override
        public View onCreateView(LayoutInflater inflater, ViewGroup container, Bundle savedInstanceState) {
            Bundle args = getArguments();
            int i=args.getInt(ARG_OBJECT);
            FPage p = fviewPager.pages.get(i);
            View rootView=fviewPager.parentController.viewmap.get(Faithdroid.triggerEventHandler(p.vid,"")).view;
            return rootView;
        }
    }
}
