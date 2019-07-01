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
@property NSMutableDictionary* _Nonnull viewmap;
@property UIViewController* _Nonnull activity;

- ( id _Nonnull )initWithParm:(UIViewController*_Nonnull)act;
@end

// FView
@interface FView : NSObject{
@public
    long className ;
    long vid;
    UIObjectCBridge* parentUIBridge;
    UIView* view;
}
@property long className;
@property long vid;
@property UIObjectCBridge* _Nonnull parentUIBridge;
@property UIView* _Nonnull view;
-(void)setAttr:(long)attr value:(NSString*_Nullable)value;
-(NSString*_Nullable)getAttr:(long)attr;
@end

// vars

NSString* _Nonnull getKey(long vid);

//FButton
@interface FButton : FView {
@public
    UIButton* _Nonnull v;
}
@property UIButton* _Nonnull v;
- ( id _Nonnull )initWithParm:(long)myid bridge:(UIObjectCBridge *_Nonnull)bridge;
@end

#endif /* ui_h */
