//
//  ViewController.swift
//  iFaithApp
//
//  Created by asd on 2018/10/25.
//  Copyright © 2018 asd. All rights reserved.
//

import UIKit
import Faithdroid
class ViewController: UIViewController {
    private var a:FaithdroidMainActivity!
    private var parentBridge:UIBridge!
    override func viewDidLoad() {
        super.viewDidLoad()
        // Do any additional setup after loading the view, typically from a nib.
        a = FaithdroidMainActivity()
        parentBridge = UIBridge().create(viewController: self, rootView: addContainer(), factivity: nil)
        a.setUIInterface(parentBridge)
        a.onCreate()
    }

    func addContainer()->UIView {
        let navHeight = Int(UIApplication.shared.statusBarFrame.height) + Int(navigationController!.navigationBar.frame.height)
        let rootView = UIView()
        rootView.clipsToBounds = true
        let rWidth = Int(self.view.frame.width)
        let rHeight = Int(self.view.frame.height) - navHeight
        rootView.frame = CGRect(x: 0, y: navHeight, width: rWidth, height: rHeight)
        rootView.backgroundColor = UIColor.white
        self.view.addSubview(rootView)
        return rootView
    }
}

