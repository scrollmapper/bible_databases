//
//  StringExtensions.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-11.
//  Copyright © 2020 Kuang. All rights reserved.
//

import Foundation

extension String {
    func withHTMLTags(fontSize: Int = 0, color: String = "", sup: Bool = false) -> String {
        var htmlString = "\(self)"
        var styleContent = ""
        if fontSize > 0 {
            if styleContent.count > 0 {
                styleContent += "; "
            }
            styleContent += "font-size:\(fontSize)px"
        }
        if color.count > 0 {
            if styleContent.count > 0 {
                styleContent += "; "
            }
            styleContent += "color:\(color)"
        }
        var style: String
        if styleContent.count > 0 {
            style = " style=\"\(styleContent)\""
        }
        else {
            style = ""
        }
        htmlString = "<span\(style)>\(htmlString)</span>"
        if sup {
            htmlString = "<sup>\(htmlString)</sup>"
        }
        return htmlString
    }
}
