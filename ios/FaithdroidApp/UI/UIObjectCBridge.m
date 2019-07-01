//
//  UIObjectCBridge.m
//  FaithdroidApp
//
//  Created by steven zacker on 6/28/19.
//  Copyright © 2019 steven zacker. All rights reserved.
//
#import "ui.h"
@implementation UIObjectCBridge
@synthesize viewmap;
@synthesize activity;

- ( id _Nonnull )initWithParm:(UIViewController*_Nonnull)act{
    self = [super init];
    if (self) {
        activity = act;
        viewmap = [NSMutableDictionary dictionary];
    }
    return self;
}

- (NSString * _Nonnull)getAttr:(long)vid attr:(long)attr {
    FView* fview = viewmap[getKey(vid)];
    NSString* value = [fview getAttr:attr];
    if (value == nil) {
        return @"";
    }
    return value;
}

- (void)new:(long)name vid:(long)vid {
    switch (name) {
        case 0:{
            FButton* fbutton = [[FButton alloc] initWithParm:vid bridge:self];
            viewmap[getKey(vid)] = fbutton;
            NSLog(@"new");
            break;
        }
        case 1 :
            break;
        default:
            break;
    }
}

- (void)runOnUIThread:(long)fnID {
    
}

- (void)setAttr:(long)vid attr:(long)attr value:(NSString * _Nullable)value {
    FView* fview=viewmap[getKey(vid)];

    [fview setAttr:attr value:value];
}

- (void)show:(long)vid {
    FView* fview = viewmap[getKey(vid)];
    if (fview != nil){
        NSLog(@"fview");
    }
    [activity.view addSubview:fview.view];
}

@end
