//
//  BundleExtensions.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-07.
//  Copyright © 2020 Kuang. All rights reserved.
//

import Foundation

public extension Bundle {
    
    func displayName() -> String? {
        return object(forInfoDictionaryKey: "CFBundleDisplayName") as? String ?? object(forInfoDictionaryKey: "CFBundleName") as? String
    }
    
    func releaseVersionNumber() -> String? {
        return object(forInfoDictionaryKey: "CFBundleShortVersionString") as? String
    }
    
    func buildVersionNumber() -> String? {
        return object(forInfoDictionaryKey: "CFBundleVersion") as? String
    }
    
}
