//
//  UIDeviceExtensions.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-18.
//  Copyright © 2020 Kuang. All rights reserved.
//

import UIKit

extension UIDevice {
    func isPad() -> Bool {
        return userInterfaceIdiom == .pad
    }
    
    func isPhone() -> Bool {
        return userInterfaceIdiom == .phone
    }
    
    func currentInterfaceOrientation() -> UIInterfaceOrientation {
        if let interfaceOrientation = UIApplication.shared.windows.first(where: { $0.isKeyWindow })?.windowScene?.interfaceOrientation {
            return interfaceOrientation
        }
        else {
            return .unknown
        }
    }
    
    func isPortrait() -> Bool {
        return currentInterfaceOrientation().isPortrait
    }
    
    func isLandscape() -> Bool {
        return currentInterfaceOrientation().isLandscape
    }
}
