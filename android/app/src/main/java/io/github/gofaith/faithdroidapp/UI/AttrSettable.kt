package io.github.gofaith.faithdroidapp.UI

interface AttrSetGettable {
    fun setAttr(attr:Long,value:String)
    fun getAttr(attr: Long):String
}