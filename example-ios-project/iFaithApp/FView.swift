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
        view.layoutMargins = UIEdgeInsets.zero
    }
    func getUniversalAttr(attr:String) -> String? {
        return ""
    }
    func setUniversalAttr(attr:String,value:String) -> Bool {
        switch attr {
        case "Size":
            parseSize(value)
        case "BackgroundColor":
            setBackgroundColor(value)
        case "LayoutGravity":
            setGravity(value)
        case "Top2TopOf":
            afterAddedFuncs[attr] = {(p:FView)->Void in
                if value == "_Parent_"{
                    self.view.topAnchor.constraint(
                        equalTo: p.view.topAnchor,
                        constant: self.view.layoutMargins.top).isActive = true
                    return
                }
                self.view.topAnchor.constraint(
                    equalTo: self.parentBridge.viewMap[FaithdroidGetVidById(value)]!.view!.topAnchor,
                    constant: self.view.layoutMargins.top).isActive = true
            }
        case "Top2BottomOf":
            afterAddedFuncs[attr] = {(p:FView)->Void in
                if value == "_Parent_"{
                    self.view.topAnchor.constraint(
                        equalTo: p.view.bottomAnchor,
                        constant: self.view.layoutMargins.top).isActive = true
                    return
                }
                self.view.topAnchor.constraint(
                    equalTo: self.parentBridge.viewMap[FaithdroidGetVidById(value)]!.view!.bottomAnchor,
                    constant: self.view.layoutMargins.top).isActive = true
            }
        case "Bottom2TopOf":
            afterAddedFuncs[attr] = {(p:FView)->Void in
                if value == "_Parent_"{
                    self.view.bottomAnchor.constraint(
                        equalTo: p.view.topAnchor,
                        constant: self.view.layoutMargins.bottom).isActive = true
                    return
                }
                self.view.bottomAnchor.constraint(
                    equalTo: self.parentBridge.viewMap[FaithdroidGetVidById(value)]!.view!.topAnchor,
                    constant: self.view.layoutMargins.bottom).isActive = true
            }
        case "Bottom2BottomOf":
            afterAddedFuncs[attr] = {(p:FView)->Void in
                if value == "_Parent_"{
                    self.view.bottomAnchor.constraint(
                        equalTo: p.view.bottomAnchor,
                        constant: self.view.layoutMargins.bottom).isActive = true
                    return
                }
                self.view.bottomAnchor.constraint(
                    equalTo: self.parentBridge.viewMap[FaithdroidGetVidById(value)]!.view!.bottomAnchor,
                    constant: self.view.layoutMargins.bottom).isActive = true
            }
        case "Left2LeftOf":
            afterAddedFuncs[attr] = {(p:FView)->Void in
                if value == "_Parent_"{
                    self.view.leftAnchor.constraint(
                        equalTo: p.view.leftAnchor,
                        constant: self.view.layoutMargins.left).isActive = true
                    return
                }
                self.view.leftAnchor.constraint(
                    equalTo: self.parentBridge.viewMap[FaithdroidGetVidById(value)]!.view!.leftAnchor,
                    constant: self.view.layoutMargins.left).isActive = true
            }
        case "Left2RightOf":
            afterAddedFuncs[attr] = {(p:FView)->Void in
                if value == "_Parent_"{
                    self.view.leftAnchor.constraint(
                        equalTo: p.view.rightAnchor,
                        constant: self.view.layoutMargins.left).isActive = true
                    return
                }
                self.view.leftAnchor.constraint(
                    equalTo: self.parentBridge.viewMap[FaithdroidGetVidById(value)]!.view!.rightAnchor,
                    constant: self.view.layoutMargins.left).isActive = true
            }
        case "Right2RightOf":
            afterAddedFuncs[attr] = {(p:FView)->Void in
                if value == "_Parent_"{
                    self.view.rightAnchor.constraint(
                        equalTo: p.view.rightAnchor,
                        constant: self.view.layoutMargins.right).isActive = true
                    return
                }
                self.view.rightAnchor.constraint(
                    equalTo: self.parentBridge.viewMap[FaithdroidGetVidById(value)]!.view!.rightAnchor,
                    constant: self.view.layoutMargins.right).isActive = true
            }
        case "Right2LeftOf":
            afterAddedFuncs[attr] = {(p:FView)->Void in
                if value == "_Parent_"{
                    self.view.rightAnchor.constraint(
                        equalTo: p.view.leftAnchor,
                        constant: self.view.layoutMargins.right).isActive = true
                    return
                }
                self.view.rightAnchor.constraint(
                    equalTo: self.parentBridge.viewMap[FaithdroidGetVidById(value)]!.view!.leftAnchor,
                    constant: self.view.layoutMargins.right).isActive = true
            }
        case "CenterX":
            afterAddedFuncs[attr]={(p:FView)->Void in
                self.view.centerXAnchor.constraint(equalTo: p.view.centerXAnchor).isActive = true
            }
        case "CenterY":
            afterAddedFuncs[attr]={(p:FView)->Void in
                self.view.centerYAnchor.constraint(equalTo: p.view.centerYAnchor).isActive = true
            }
        case "WidthPercent":
            afterAddedFuncs[attr]={(p:FView)->Void in
                self.view.widthAnchor.constraint(equalTo: p.view.widthAnchor, multiplier: CGFloat.init(Float(value)!)).isActive = true
            }
        case "HeightPercent":
            afterAddedFuncs[attr]={(p:FView)->Void in
                self.view.heightAnchor.constraint(equalTo: p.view.heightAnchor, multiplier: CGFloat.init(Float(value)!)).isActive = true
            }
        default:
            return false
        }
        return true
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
