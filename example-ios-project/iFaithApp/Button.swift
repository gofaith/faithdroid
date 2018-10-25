//
//  Button.swift
//  iFaithApp
//
//  Created by asd on 2018/10/25.
//  Copyright © 2018 asd. All rights reserved.
//

import UIKit
class FButton: FView {
    public var v:UIButton!
    func create(c:UIBridge) -> FView {
        parentBridge = c
        v = UIButton(type: .system)
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
        case "Text":
            return v.title(for: .normal)
        default:
            return ""
        }
    }
    override func setAttr(attr: String, value: String) {
        if setUniversalAttr(attr: attr, value: value) {
            return
        }
        switch attr {
        case "Text":
            v.setTitle(value, for: .normal)
        default:
            break
        }
    }
}
