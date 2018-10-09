package io.github.gofaith.faithdroid.FViews;

import android.graphics.drawable.Drawable;
import android.widget.ImageView;

import io.github.gofaith.faithdroid.UI.AttrGettable;
import io.github.gofaith.faithdroid.UI.AttrSettable;
import io.github.gofaith.faithdroid.UI.Toolkit;
import io.github.gofaith.faithdroid.UI.UIController;

public class FImageView extends FView implements AttrSettable, AttrGettable {
    public ImageView v;

    public FImageView(UIController controller) {
        parentController=controller;
        v = new ImageView(parentController.activity);
        view=v;
    }
    @Override
    public String getAttr(String attr) {
        String str = getUniversalAttr(attr);
        if (str != null) {
            return str;
        }
        switch (attr){
            // ------------------------------------------
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
            case "Src":
                Toolkit.file2Drawable(parentController, value, new Toolkit.OnDrawableReadyListener() {
                    @Override
                    public void onDrawableReady(Drawable drawable) {
                        v.setImageDrawable(drawable);
                    }
                });
                break;
            case "ScaleType":
                v.setScaleType(parseScaleType(value));
                break;
        }
    }

    private ImageView.ScaleType parseScaleType(String value) {
        switch (value) {
            case "CenterCrop":
                return ImageView.ScaleType.CENTER_CROP;
            case "CenterInside":
                return ImageView.ScaleType.CENTER_INSIDE;
            case "FitCenter":
                return ImageView.ScaleType.FIT_CENTER;
            case "FitStart":
                return ImageView.ScaleType.FIT_START;
            case "FitEnd":
                return ImageView.ScaleType.FIT_END;
            case "FitXY":
                return ImageView.ScaleType.FIT_XY;
            case "Matrix":
                return ImageView.ScaleType.MATRIX;
            default:
                return ImageView.ScaleType.CENTER;
        }
    }
}
