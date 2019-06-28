//
//  ui.h
//  FaithdroidApp
//
//  Created by steven zacker on 6/28/19.
//  Copyright © 2019 steven zacker. All rights reserved.
//

#ifndef ui_h
#define ui_h

#import <Foundation/Foundation.h>
#import <UIKit/UIKit.h>
#import "../Faithdroid.framework/Headers/Faithdroid.h"

// UIObjectCBridge
@interface UIObjectCBridge : NSObject<FaithdroidUIBridge>{
@public
    NSMutableDictionary* viewmap;
    UIViewController* activity;
}
@end

// AttrSetGettable
@protocol AttrSetGettable <NSObject>

@required
-(void)setAttr:(long)attr value:(NSString * _Nullable)value;
-(NSString * _Nonnull)getAttr:(long)attr;

@end

// FView
@interface FView : NSObject{
@private
    long className ;
    long vid;
    UIObjectCBridge* parentUIBridge;
    UIView* view;
}
@property long className;
@property long vid;
@property UIObjectCBridge* _Nullable parentUIBridge;
@property UIView* _Nullable view;
@end

// vars

enum {
    ClassButton,
    ClassText,
};

NSString* _Nonnull getKey(long vid);
#endif /* ui_h */
