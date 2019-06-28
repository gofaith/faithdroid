//
//  FView.h
//  FaithdroidApp
//
//  Created by steven zacker on 6/28/19.
//  Copyright © 2019 steven zacker. All rights reserved.
//

#ifndef FView_h
#define FView_h
#import <Foundation/Foundation.h>
#import <UIKit/UIKit.h>
#import "Vars.h"
#import "UIObjectCBridge.h"
@interface FView : NSObject{
@private
    long className ;
    long vid;
    UIObjectCBridge* parentUIBridge;
    UIView* view;
}
@property long className;
@property long vid;
@property UIObjectCBridge* parentUIBridge;
@property UIView* view;
@end
#endif /* FView_h */
