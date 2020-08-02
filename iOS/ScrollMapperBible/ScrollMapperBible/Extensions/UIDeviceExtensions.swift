//
//  UIDeviceExtensions.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-18.
//  Copyright © 2020 Kuang. All rights reserved.
//

import UIKit

extension UIDevice {
    var isPad: Bool {
        return userInterfaceIdiom == .pad
    }
    
    var isPhone: Bool {
        return userInterfaceIdiom == .phone
    }
    
    var currentInterfaceOrientation: UIInterfaceOrientation {
        if let interfaceOrientation = UIApplication.shared.windows.first(where: { $0.isKeyWindow })?.windowScene?.interfaceOrientation {
            return interfaceOrientation
        }
        else {
            return .unknown
        }
    }
    
    var isPortrait: Bool {
        return currentInterfaceOrientation.isPortrait
    }
    
    var isLandscape: Bool {
        return currentInterfaceOrientation.isLandscape
    }
}
