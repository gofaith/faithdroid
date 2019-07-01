//
//  ViewController.h
//  FaithdroidApp
//
//  Created by steven zacker on 6/28/19.
//  Copyright © 2019 steven zacker. All rights reserved.
//

#import <UIKit/UIKit.h>
#import "UI/ui.h"
#import "Faithdroid.framework/Headers/Faithdroid.h"
@interface ViewController : UIViewController{
@public
    UIObjectCBridge* _Nullable ui;
    FaithdroidMainWindow* _Nullable w;
}
@property UIObjectCBridge* _Nullable ui;
@property FaithdroidMainWindow* _Nullable w;
@end

