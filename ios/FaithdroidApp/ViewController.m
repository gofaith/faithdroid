//
//  ViewController.m
//  FaithdroidApp
//
//  Created by steven zacker on 6/28/19.
//  Copyright © 2019 steven zacker. All rights reserved.
//

#import "ViewController.h"

@interface ViewController ()

@end

@implementation ViewController
@synthesize ui;
@synthesize w;
- (void)viewDidLoad {
    [super viewDidLoad];
    // Do any additional setup after loading the
    w = [[FaithdroidMainWindow alloc] init];
    ui = [[UIObjectCBridge alloc] initWithParm:self];
    [w setUI:ui];
    [w onCreate];
}


@end
