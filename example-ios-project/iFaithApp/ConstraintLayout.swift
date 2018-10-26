//
//  ConstraintLayout.swift
//  iFaithApp
//
//  Created by asd on 2018/10/26.
//  Copyright © 2018 asd. All rights reserved.
//

import UIKit
class FConstraintLayout: FView {
    public var v:UIView!
    func create(_ c:UIBridge) -> FView {
        parentBridge = c
        v = UIView()
        view = v
        v.clipsToBounds = true
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
        case "Append":
            let vids = value.split(separator: ",")
            for vid in  vids{
                v.addSubview(parentBridge.viewMap[String(vid)]!.view)
            }
            for vid in vids{
                parentBridge.viewMap[String(vid)]!.execAddedFuncs(self)
            }
        default:
            break
        }
    }
}
