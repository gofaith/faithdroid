//
//  UIBridge.swift
//  iFaithApp
//
//  Created by asd on 2018/10/25.
//  Copyright © 2018 asd. All rights reserved.
//

import UIKit
import Faithdroid
class UIBridge: NSObject, FaithdroidUIControllerProtocol {
    public var factivity:FaithdroidActivity!
    public var activity:ViewController!
    public var viewMap = [String:FView]()
    public var rootView:FFrameLayout!
    func create(viewController:ViewController,rootView:UIView,factivity:FaithdroidActivity!) -> UIBridge {
        self.activity = viewController
        self.rootView = FFrameLayout()
        self.rootView.parentBridge = self
        self.rootView.v = rootView
        self.rootView.view = rootView
        self.factivity = factivity
        self.rootView.className = "FrameLayout"
        self.rootView.vID = "root"
        viewMap["root"] = self.rootView
        return self
    }
    
    func newView(_ viewName: String!, vid VID: String!) {
        print("newView:"+viewName)
        var v:FView!
        switch viewName {
        case "Button":
            v = FButton().create(c:self)
        default:
            break
        }
        v.className = viewName
        v.vID = VID
        viewMap[VID] = v
    }
    func append2RootView(_ VID: String!) {
        if let v = viewMap[VID]{
            rootView.v.addSubview(v.view)
            v.execAddedFuncs(rootView)
        }
    }
    
    func getCurrentFActivity() -> FaithdroidActivity! {
        return factivity
    }
    
    func getPkg() -> String! {
        let a = Bundle.main.object(forInfoDictionaryKey: "CFBundleName") as! String
        return a
    }
    
    
    func runUIThread(_ fnID: String!) {
        
    }
    
    func viewGetAttr(_ VID: String!, attr: String!) -> String! {
        if let v = viewMap[VID] {
            return v.getAttr(attr: attr)
        }
        return ""
    }
    
    func viewSetAttr(_ VID: String!, attr: String!, value: String!) {
        print("set"+attr+":"+value)
        if let v = viewMap[VID] {
            v.setAttr(attr: attr, value: value)
        }
    }
    
    
}
