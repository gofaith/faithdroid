//
//  TextView.swift
//  iFaithApp
//
//  Created by asd on 2018/10/26.
//  Copyright © 2018 asd. All rights reserved.
//

import UIKit

class FTextView: FView {
    public var v:UILabel!
    func create(c:UIBridge) -> FView {
        parentBridge = c
        v = UILabel()
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
            return v.text
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
            v.text = value
        default:
            break
        }
    }
}
