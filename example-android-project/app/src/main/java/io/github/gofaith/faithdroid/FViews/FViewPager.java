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
    public String onCreateView,onGetCount;
    public FViewPager(UIController controller) {
        parentController = controller;
        v = new ViewPager(parentController.activity);
        v.setId(ViewCompat.generateViewId());
        view=v;
        parseSize("[-1,-1]");
    }

    @Override
    public String getAttr(String attr) {
        if(attr==null)
            return "";
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
            // --------------------------------------------------
        }
        return "";
    }

    @Override
    public void setAttr(String attr, final String value) {
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
            case "Foreground":
                setForeground(value);
                break;
            case "Size":
                parseSize( value);
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
                        Faithdroid.triggerEventHandler(value,"");
                    }
                });
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
                    if(adapter==null) {
                        adapter = new FaithCollectionPagerAdapter(parentController.activity.getSupportFragmentManager(), this);
                        v.setAdapter(adapter);
                    }
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
            case "OnCreateView":
                onCreateView = value;
                if (adapter==null&&onGetCount!=null) {
                    adapter = new FaithCollectionPagerAdapter(parentController.activity.getSupportFragmentManager(), this);
                    v.setAdapter(adapter);
                }
                break;
            case "OnGetCount":
                onGetCount=value;
                if (adapter==null&&onCreateView!=null) {
                    adapter = new FaithCollectionPagerAdapter(parentController.activity.getSupportFragmentManager(), this);
                    v.setAdapter(adapter);
                }
                break;
            case "CurrentItem":
                try {
                    JSONArray array = (JSONArray) (new JSONTokener(value).nextValue());
                    int index = array.getInt(0);
                    int is = array.getInt(1);
                    v.setCurrentItem(index,is==1);
                } catch (Exception e) {
                    e.printStackTrace();
                }
                break;
        }
    }
    // ------------------
    class FPage{
        String vid;
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
            if (fviewPager.pages.size()==0&&onGetCount!=null&&onCreateView!=null) {
                try {
                    return Integer.parseInt(Faithdroid.triggerEventHandler(onGetCount, ""));
                } catch (Exception e) {
                }
            }
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
            if (fviewPager.onCreateView != null && !fviewPager.onCreateView.equals("")) {
                Log.d("TAG", "onCreateView: 1");
                String vid = Faithdroid.triggerEventHandler(fviewPager.onCreateView, String.valueOf(i));
                if (fviewPager.parentController.viewmap.containsKey(vid)) {
                    Log.d("TAG", "onCreateView: 2");
                    FView cview = fviewPager.parentController.viewmap.get(vid);
                    if (cview != null) {
                        Log.d("TAG", "onCreateView: 3");
                        return cview.view;
                    }
                }
            }
            if (fviewPager.pages.size() <= i) {
                return null;
            }
            FPage p = fviewPager.pages.get(i);
            View rootView=fviewPager.parentController.viewmap.get(Faithdroid.triggerEventHandler(p.vid,"")).view;
            return rootView;
        }
    }
}
