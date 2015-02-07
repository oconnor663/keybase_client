//
//  KBProgressOverlayView.h
//  Keybase
//
//  Created by Gabriel on 1/19/15.
//  Copyright (c) 2015 Gabriel Handford. All rights reserved.
//

@import Foundation;
@import AppKit;

#import <YOLayout/YOLayout.h>
#import "KBActivityIndicatorView.h"

@interface KBProgressOverlayView : YONSView

@property BOOL hidesWhenStopped;
@property (nonatomic) BOOL animating;

- (void)startAnimating;
- (void)stopAnimating;

@end
