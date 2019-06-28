//
//  AttrSetGettable.h
//  FaithdroidApp
//
//  Created by steven zacker on 6/28/19.
//  Copyright © 2019 steven zacker. All rights reserved.
//

#ifndef AttrSetGettable_h
#define AttrSetGettable_h

#import <Foundation/Foundation.h>

@protocol AttrSetGettable <NSObject>

@required
-(void)setAttr:(long)attr value:(NSString * _Nullable)value;
-(NSString * _Nonnull)getAttr:(long)attr;

@end


#endif /* AttrSetGettable_h */
