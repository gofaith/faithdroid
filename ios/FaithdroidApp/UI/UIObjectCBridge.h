//
//  UIObjectCBridge.h
//  FaithdroidApp
//
//  Created by steven zacker on 6/28/19.
//  Copyright © 2019 steven zacker. All rights reserved.
//

#ifndef UIObjectCBridge_h
#define UIObjectCBridge_h

#import <Foundation/Foundation.h>
#import <UIKit/UIKit.h>
#import "../Faithdroid.framework/Headers/Faithdroid.h"

@interface UIObjectCBridge : NSObject<FaithdroidUIBridge>{
@public
    NSMutableDictionary* viewmap;
    UIViewController* activity;
}
@end
#endif /* UIObjectCBridge_h */
