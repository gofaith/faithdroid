package io.github.stevenzack.faithdroidapp.FViews

import android.graphics.drawable.Drawable
import android.widget.ImageView
import io.github.stevenzack.faithdroidapp.UI.*


class FImageView(controller: UIController) : FView(), AttrSettable, AttrGettable {
    var v: ImageView

    init {
        parentController = controller
        v = ImageView(parentController!!.activity)
        view = v
    }

  override  fun getAttr(attr: String): String {
        val str = getUniversalAttr(attr)
        if (str != null) {
            return str
        }
        when (attr) {

        }// ------------------------------------------
        return ""
    }

   override fun setAttr(attr: String, value: String) {
        if (setUniversalAttr(attr, value)) {
            return
        }
        when (attr) {
            // -------------------------------------------------------------------
            "Src" -> Toolkit.file2Drawable(parentController!!, value, {d->
                    v.setImageDrawable(d)
            })
            "ScaleType" -> v.setScaleType(parseScaleType(value))
        }
    }

    private fun parseScaleType(value: String): ImageView.ScaleType {
        when (value) {
            "CenterCrop" -> return ImageView.ScaleType.CENTER_CROP
            "CenterInside" -> return ImageView.ScaleType.CENTER_INSIDE
            "FitCenter" -> return ImageView.ScaleType.FIT_CENTER
            "FitStart" -> return ImageView.ScaleType.FIT_START
            "FitEnd" -> return ImageView.ScaleType.FIT_END
            "FitXY" -> return ImageView.ScaleType.FIT_XY
            "Matrix" -> return ImageView.ScaleType.MATRIX
            else -> return ImageView.ScaleType.CENTER
        }
    }
}
