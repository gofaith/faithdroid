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
import io.github.gofaith.faithdroid.UI.DemoObjectFragment;
import io.github.gofaith.faithdroid.UI.UIController;

public class FViewPager extends FView implements AttrSettable, AttrGettable {
    private FaithCollectionPagerAdapter adapter;
    public ViewPager v;
    public List<FPage> pages = new ArrayList<>();
    public String onCreateView,onGetCount,onPageSelected;
    public FTabLayout bindFTabLayout;
    public FViewPager(UIController controller) {
        parentController = controller;
        v = new ViewPager(parentController.activity);
        v.addOnPageChangeListener(new ViewPager.OnPageChangeListener() {
            @Override
            public void onPageScrolled(int i, float v, int i1) {

            }

            @Override
            public void onPageSelected(int i) {
                if (onPageSelected != null) {
                    Faithdroid.triggerEventHandler(onPageSelected, String.valueOf(i));
                }
            }

            @Override
            public void onPageScrollStateChanged(int i) {

            }
        });
        view=v;
        parseSize("[-1,-1]");
    }

    @Override
    public String getAttr(String attr) {
        String str = getUniversalAttr(attr);
        if (str != null) {
            return str;
        }
        switch (attr){
            // --------------------------------------------------
        }
        return "";
    }

    @Override
    public void setAttr(String attr, final String value) {
        if (setUniversalAttr(attr, value)) {
            return;
        }
        switch (attr){
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
                bindFTabLayout=fTabLayout;
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
            case "OnPageSelected":
                onPageSelected = value;
                break;
        }
    }
    // ------------------
   public class FPage{
        public String vid;
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
            if (bindFTabLayout==null||bindFTabLayout.tabsList.size()<=position){
                return "Page "+position;
            }
            return bindFTabLayout.tabsList.get(position);
        }
    }
}
