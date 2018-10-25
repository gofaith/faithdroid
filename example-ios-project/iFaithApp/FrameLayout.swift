//
//  FrameLayout.swift
//  iFaithApp
//
//  Created by asd on 2018/10/25.
//  Copyright © 2018 asd. All rights reserved.
//

import UIKit
class FFrameLayout: FView {
    public var v:UIView!
    func create(c:UIBridge) -> FView {
        parentBridge = c
        v = UIView()
        v.clipsToBounds = true
        view = v
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
        case "AddView":
            v.addSubview(parentBridge.viewMap[value]!.view)
            parentBridge.viewMap[value]!.execAddedFuncs(self)
        default:
            break
        }
    }
}

