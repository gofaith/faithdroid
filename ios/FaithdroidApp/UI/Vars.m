//
//  Vars.m
//  FaithdroidApp
//
//  Created by steven zacker on 6/28/19.
//  Copyright © 2019 steven zacker. All rights reserved.
//
#import "ui.h"

long const ClassButton = 0;
long const ClassText = 1;

NSString* _Nonnull getKey(long vid){
    NSString* str = [NSString stringWithFormat:@"%ld",vid];
    return str;
}
