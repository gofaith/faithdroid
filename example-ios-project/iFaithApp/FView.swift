//
//  FView.swift
//  iFaithApp
//
//  Created by asd on 2018/10/25.
//  Copyright © 2018 asd. All rights reserved.
//

import Faithdroid
import UIKit

class FView {
    public var className:String!
    public var otherInfo:String!
    public var vID:String!
    public var view:UIView!
    public var parentBridge:UIBridge!
    public var size = [-1,-1]
    public var gravity = ""
    public var afterAddedFuncs = [String:(parent: FView)->Void]()
    func execAddedFuncs(_ parent: FView)  {
        for (_,fn) in afterAddedFuncs {
            fn(parent)
        }
    }
    func setupView()  {
        view.translatesAutoresizingMaskIntoConstraints = false
        view.sizeThatFits(CGSize(width: view.frame.size.width, height: view.frame.size.height))
        afterAddedFuncs["gravity"] = {(parent:FView) -> Void in
            if parent.className == "FrameLayout"{
                self.view.topAnchor.constraint(equalTo: parent.view.topAnchor).isActive = true
                self.view.leftAnchor.constraint(equalTo: parent.view.leftAnchor).isActive = true
            }
        }
    }
    func getUniversalAttr(attr:String) -> String? {
        return ""
    }
    func setUniversalAttr(attr:String,value:String) -> Bool {
        switch attr {
        case "Size":
            parseSize(value)
            return true
        case "BackgroundColor":
            setBackgroundColor(value)
            return true
        case "Gravity":
            setGravity(value)
            return true
        default:
            break
        }
        return false
    }
    func setAttr(attr:String,value:String)  {
        
    }
    func getAttr(attr:String) -> String! {
        return ""
    }
    func parseSize(_ value:String) {
        
    }
    func setGravity(_ g:String)  {
    }
    func setBackgroundColor(_ value:String){
        view.backgroundColor = hexStringToUIColor(hex: value)
    }
    func hexStringToUIColor (hex:String) -> UIColor {
        var cString:String = hex.trimmingCharacters(in: .whitespacesAndNewlines).uppercased()
        
        if (cString.hasPrefix("#")) {
            cString.remove(at: cString.startIndex)
        }
        
        if ((cString.count) != 6) {
            return UIColor.gray
        }
        
        var rgbValue:UInt32 = 0
        Scanner(string: cString).scanHexInt32(&rgbValue)
        
        return UIColor(
            red: CGFloat((rgbValue & 0xFF0000) >> 16) / 255.0,
            green: CGFloat((rgbValue & 0x00FF00) >> 8) / 255.0,
            blue: CGFloat(rgbValue & 0x0000FF) / 255.0,
            alpha: CGFloat(1.0)
        )
    }
}
