package io.github.stevenzack.faithdroidapp.UI

import kotlinx.serialization.Serializable

@Serializable
data class FMyIntent(val action:String,val paths:List<String>,val extras:Map<String,String>)

@Serializable
data class FActivityConfig(val myLaunchMode:String,val myFnId:String,val myIntent: FMyIntent)