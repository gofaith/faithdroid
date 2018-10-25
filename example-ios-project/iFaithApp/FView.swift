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
        afterAddedFuncs["Gravity"] = {(parent:FView) -> Void in
            if parent.className == "FrameLayout"{
                self.view.topAnchor.constraint(equalTo: parent.view.topAnchor).isActive = true
                self.view.leftAnchor.constraint(equalTo: parent.view.leftAnchor).isActive = true
            }else if parent.className == "LinearLayout"{
                if parent.otherInfo == "Vertical"{
                    self.view.leftAnchor.constraint(equalTo: parent.view.leftAnchor).isActive = true
                }else{
                    self.view.topAnchor.constraint(equalTo: parent.view.topAnchor).isActive = true
                }
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
        case "LayoutGravity":
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
        let ss = value.replacingOccurrences(of: "[", with: "", options: .literal, range: nil).replacingOccurrences(of: "]", with: "", options: .literal, range: nil).split(separator: ",")
        let width = Int(ss[0])!
        let height =  Int(ss[1])!
        size[0] = width
        size[1] = height
        if width == -2 {
            afterAddedFuncs["width"]={(parent:FView)->Void in
                self.view.widthAnchor.constraint(equalTo: parent.view.widthAnchor).isActive = true
            }
        }else if width == -1 {
            afterAddedFuncs.removeValue(forKey: "width")
        }else if width > 0{
            afterAddedFuncs["width"] = {(parent:FView)->Void in
                self.view.widthAnchor.constraint(equalToConstant: CGFloat(width)).isActive = true
            }
        }
        
        if height == -2{
            afterAddedFuncs["height"]={(parent:FView)->Void in
                self.view.heightAnchor.constraint(equalTo: parent.view.heightAnchor).isActive = true
            }
        }else if height == -1 {
            afterAddedFuncs.removeValue(forKey: "height")
        }else if height > 0{
            afterAddedFuncs["height"] = {(parent:FView)->Void in
                self.view.heightAnchor.constraint(equalToConstant: CGFloat(height)).isActive = true
            }
        }
    }
    func setGravity(_ g:String)  {
        switch g {
        case "3"://left
            afterAddedFuncs["Gravity"] = {(parent:FView)->Void in
                self.view.leftAnchor.constraint(equalTo: parent.view.leftAnchor).isActive = true
                if parent.className == "FrameLayout"{
                    self.view.topAnchor.constraint(equalTo: parent.view.topAnchor).isActive = true
                }
            }
        case "5"://Right
            afterAddedFuncs["Gravity"] = {(parent:FView)->Void in
                self.view.rightAnchor.constraint(equalTo: parent.view.rightAnchor).isActive = true
                if parent.className == "FrameLayout"{
                    self.view.topAnchor.constraint(equalTo: parent.view.topAnchor).isActive = true
                }
            }
        case "48"://top
            afterAddedFuncs["Gravity"] = {(p:FView)->Void in
                self.view.topAnchor.constraint(equalTo: p.view.topAnchor).isActive = true
                if p.className == "FrameLayout"{
                    self.view.leftAnchor.constraint(equalTo: p.view.leftAnchor).isActive = true
                }
            }
        case "80"://bottom
            afterAddedFuncs["Gravity"] = {(p:FView)->Void in
                self.view.bottomAnchor.constraint(equalTo: p.view.bottomAnchor).isActive = true
                if p.className == "FrameLayout"{
                    self.view.leftAnchor.constraint(equalTo: p.view.leftAnchor).isActive  = true
                }
            }
        case "17"://center
            afterAddedFuncs["Gravity"] = {(p:FView)->Void in
                if p.className  == "FrameLayout"{
                    self.view.centerXAnchor.constraint(equalTo: p.view.centerXAnchor).isActive = true
                    self.view.centerYAnchor.constraint(equalTo: p.view.centerYAnchor).isActive = true
                }else if p.className == "LinearLayout"{
                    if p.otherInfo == "Vertical"{
                        self.view.centerXAnchor.constraint(equalTo: p.view.centerXAnchor).isActive = true
                    }else{
                        self.view.centerYAnchor.constraint(equalTo: p.view.centerYAnchor).isActive = true
                    }
                }
            }
        case "1"://HCENTER
            afterAddedFuncs["Gravity"] = {(p:FView)->Void in
                self.view.centerXAnchor.constraint(equalTo: p.view.centerXAnchor).isActive = true
                if p.className == "FrameLayout"{
                    self.view.topAnchor.constraint(equalTo: p.view.topAnchor).isActive  = true
                }
            }
        case "16"://VCenter
            afterAddedFuncs["Gravity"] = {(p:FView)->Void in
                self.view.centerYAnchor.constraint(equalTo: p.view.centerYAnchor).isActive = true
                if p.className == "FrameLayout"{
                    self.view.leftAnchor.constraint(equalTo: p.view.leftAnchor).isActive = true
                }
            }
        case "51"://topLeft
            afterAddedFuncs["Gravity"] = {(parent:FView) -> Void in
                self.view.topAnchor.constraint(equalTo: parent.view.topAnchor).isActive = true
                self.view.leftAnchor.constraint(equalTo: parent.view.leftAnchor).isActive = true
            }
        case "49"://topCenter
            afterAddedFuncs["Gravity"] = {(parent:FView) -> Void in
                self.view.topAnchor.constraint(equalTo: parent.view.topAnchor).isActive = true
                self.view.centerXAnchor.constraint(equalTo: parent.view.centerXAnchor).isActive = true
            }
        case "53"://topRight
            afterAddedFuncs["Gravity"] = {(parent:FView) -> Void in
                self.view.topAnchor.constraint(equalTo: parent.view.topAnchor).isActive = true
                self.view.rightAnchor.constraint(equalTo: parent.view.rightAnchor).isActive = true
            }
        case "19"://LeftCenter
            afterAddedFuncs["Gravity"] = {(parent:FView) -> Void in
                self.view.centerYAnchor.constraint(equalTo: parent.view.centerYAnchor).isActive = true
                self.view.leftAnchor.constraint(equalTo: parent.view.leftAnchor).isActive = true
            }
        case "21"://rightCenter
            afterAddedFuncs["Gravity"] = {(parent:FView) -> Void in
                self.view.centerYAnchor.constraint(equalTo: parent.view.centerYAnchor).isActive = true
                self.view.rightAnchor.constraint(equalTo: parent.view.rightAnchor).isActive = true
            }
        case "83"://bottomLeft
            afterAddedFuncs["Gravity"] = {(parent:FView) -> Void in
                self.view.bottomAnchor.constraint(equalTo: parent.view.bottomAnchor).isActive = true
                self.view.leftAnchor.constraint(equalTo: parent.view.leftAnchor).isActive = true
            }
        case "81"://bottomCenter
            afterAddedFuncs["Gravity"] = {(parent:FView) -> Void in
                self.view.bottomAnchor.constraint(equalTo: parent.view.bottomAnchor).isActive = true
                self.view.centerXAnchor.constraint(equalTo: parent.view.centerXAnchor).isActive = true
            }
        case "85"://rightBottom
            afterAddedFuncs["Gravity"] = {(parent:FView) -> Void in
                self.view.rightAnchor.constraint(equalTo: parent.view.rightAnchor).isActive = true
                self.view.bottomAnchor.constraint(equalTo: parent.view.bottomAnchor).isActive = true
            }
        default:
            break
        }
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
