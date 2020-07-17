//
//  ScrollMapperBibleWebViewColor.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-16.
//  Copyright © 2020 Kuang. All rights reserved.
//

import Foundation
import UIKit

class WebViewColor {
    enum Color {
        case undefined
        case black(darkMode: Bool)
        case white(darkMode: Bool)
        
        var color: UIColor {
            switch self {
            case .undefined:
                return .clear
            case .black(let darkMode):
                return darkMode ? .white : .black
            case .white(let darkMode):
                return darkMode ? .black : .white
            }
        }
        var string: String {
            switch self {
            case .undefined:
                return "undefined"
            case .black(let darkMode):
                return darkMode ? "white" : "black"
            case .white(let darkMode):
                return darkMode ? "black" : "white"
            }
        }
    }
    static var black: Color {
        var color: Color = .black(darkMode: false)
        if #available(iOS 13, *) {
            let _ = UIColor.init { (traitCollection) -> UIColor in
                color =  Color.black(darkMode: traitCollection.userInterfaceStyle == .dark)
                return .black
            }
        }
        return color
    }
    static var white: Color {
        var color: Color = .white(darkMode: false)
        if #available(iOS 13, *) {
            let _ = UIColor.init { (traitCollection) -> UIColor in
                color =  Color.white(darkMode: traitCollection.userInterfaceStyle == .dark)
                return .black
            }
        }
        return color
    }
}
