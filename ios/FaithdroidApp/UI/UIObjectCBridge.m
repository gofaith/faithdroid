//
//  UIObjectCBridge.m
//  FaithdroidApp
//
//  Created by steven zacker on 6/28/19.
//  Copyright © 2019 steven zacker. All rights reserved.
//
#import "ui.h"
@implementation UIObjectCBridge{
}

- (instancetype)initWithParm:(UIViewController*)act
{
    self = [super init];
    if (self) {
        activity = act;
        viewmap = [NSMutableDictionary dictionary];
    }
    return self;
}

- (NSString * _Nonnull)getAttr:(long)vid attr:(long)attr {
    
    return @"hello";
}

- (void)new:(long)name vid:(long)vid {
    
}

- (void)runOnUIThread:(long)fnID {
    
}

- (void)setAttr:(long)vid attr:(long)attr value:(NSString * _Nullable)value {
    
}

- (void)show:(long)vid {
    NSString* key = [NSString stringWithFormat:@"%ld",vid];
    FView* fview = [viewmap objectForKey:key];
    
}

@end
