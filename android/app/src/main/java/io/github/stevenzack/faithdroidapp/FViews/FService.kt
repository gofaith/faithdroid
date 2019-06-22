package io.github.stevenzack.faithdroidapp.FViews

import android.content.Context
import android.content.Intent
import io.github.stevenzack.faithdroidapp.R
import io.github.stevenzack.faithdroidapp.UI.FView
import io.github.stevenzack.faithdroidapp.UI.UIController

class FService(uiController: UIController) : FView() {
    init {
        parentController = uiController
    }

    companion object {
        fun getAttr(context: Context, attr: String): String? {
            return null
        }

        fun setAttr(context: Context, attr: String, value: String?) {
            if (value == null) {
                return
            }
            when (attr) {
                "OnCreate" -> CoreService.onCreate = value
                "QuitButton" -> CoreService.quitButton = value
                "OnQuit" -> CoreService.onQuitClick = value
                "Title" -> {
                    CoreService.ntf_title = value
                    if (CoreService.notification != null) {
                        CoreService.getInstance().handler.post(Runnable {
                            CoreService.notification.contentView.setTextViewText(R.id.ntf_title, value)
                            CoreService.getInstance().startForeground(1, CoreService.notification)
                        })
                    }
                }
                "SubTitle" -> {
                    CoreService.ntf_subtitle = value
                    if (CoreService.notification != null) {
                        CoreService.getInstance().handler.post(Runnable {
                            CoreService.notification.contentView.setTextViewText(R.id.ntf_subtitle, value)
                            CoreService.getInstance().startForeground(1, CoreService.notification)
                        })
                    }
                }
                "Show" -> {
                    val intent = Intent(context, CoreService::class.java)
                    context.startService(intent)
                }
                "FinishAllActivity" -> CoreService.getInstance().finishAllActivity()
                "KillAll" -> CoreService.killAll()
                "Stop" -> CoreService.getInstance().stopService()
            }
        }
    }
}
