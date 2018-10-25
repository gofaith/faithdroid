//
//  LinearLayout.swift
//  iFaithApp
//
//  Created by asd on 2018/10/25.
//  Copyright © 2018 asd. All rights reserved.
//

import UIKit
class FLinearLayout: FView {
    public var v:UIView!
    public var lastChild:FView!
    func create(c:UIBridge) -> FView {
        parentBridge = c
        v = UIView()
        v.clipsToBounds = true
        view = v
        otherInfo = "Vertical"
        setupView()
        return self
    }
    override func getAttr(attr: String) -> String! {
        let str = getUniversalAttr(attr: attr)
        if let s = str{
            return s
        }
        switch attr {
        default:
            return ""
        }
    }
    override func setAttr(attr: String, value: String) {
        if setUniversalAttr(attr: attr, value: value) {
            return
        }
        switch attr {
        case "Orientation":
            if value == "VERTICAL"{
                otherInfo = "Vertical"
            }else {
                otherInfo = "Horizontal"
            }
        case "AddView":
            v.addSubview(parentBridge.viewMap[value]!.view)
            if otherInfo == "Vertical"{
                if let lastKid = lastChild{
                    parentBridge.viewMap[value]!.afterAddedFuncs["GravityTop"]={(parent:FView)->Void in
                        let mv = self.parentBridge.viewMap[value]!.view
                        mv!.topAnchor.constraint(equalTo: lastKid.view.bottomAnchor).isActive = true
                    }
                }else{
                    parentBridge.viewMap[value]!.afterAddedFuncs["GravityTop"]={(parent:FView)->Void in
                        let mv = self.parentBridge.viewMap[value]!.view
                        mv!.topAnchor.constraint(equalTo: parent.view.topAnchor).isActive = true
                    }
                }
            }else{
                if let lastKid = lastChild{
                    parentBridge.viewMap[value]!.afterAddedFuncs["GravityLeft"]={(p:FView)->Void in
                        let mv = self.parentBridge.viewMap[value]!.view
                        mv!.leftAnchor.constraint(equalTo: lastKid.view.rightAnchor).isActive = true
                    }
                }else{
                    parentBridge.viewMap[value]!.afterAddedFuncs["GravityLeft"]={(p:FView)->Void in
                        let mv = self.parentBridge.viewMap[value]!.view
                        mv!.leftAnchor.constraint(equalTo: p.view.leftAnchor).isActive = true
                    }
                }
            }
            parentBridge.viewMap[value]!.execAddedFuncs(self)
            lastChild = parentBridge.viewMap[value]
        default:
            break
        }
    }
}

