//
//  ScrollMapperBibleFacilities.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-11.
//  Copyright © 2020 Kuang. All rights reserved.
//

import Foundation
import UIKit

class ScrollMapperBibleFacilities {
    // build HTML header for dynamic type and responsive design
    // https://riptutorial.com/ios/example/24811/matching-dynamic-type-font-size-in-wkwebview
    // BUT THIS DOES NOT WORK
    static func buildHTMLHeader(bodyContent: String) -> String {
        // Get preferred dynamic type font sizes for html styles
        let bodySize = UIFont.preferredFont(forTextStyle: UIFont.TextStyle.body).pointSize
        let h1Size = UIFont.preferredFont(forTextStyle: UIFont.TextStyle.title1).pointSize
        let h2Size = UIFont.preferredFont(forTextStyle: UIFont.TextStyle.title2).pointSize
        let h3Size = UIFont.preferredFont(forTextStyle: UIFont.TextStyle.title3).pointSize
            
        // On iPad, landscape text is larger than preferred font size
        var portraitMultiplier = CGFloat(1.0)
        var landscapeMultiplier = CGFloat(0.5)
        
        // iPhone text is shrunken
        if UIDevice.current.model.range(of: "iPhone") != nil {
            portraitMultiplier = CGFloat(3.0)
            landscapeMultiplier = CGFloat(1.5)
        }
            
        var htmlString = ""
        htmlString += "<html>\n"
        htmlString += "  <head>\n"
        htmlString += "    <style>\n"
        htmlString += "      @media all and (orientation:portrait) {img {max-width: 90%; height: auto;}\n"
        htmlString += "      p, li { font: -apple-system-body; font-family: Georgia, serif; font-size:calc(\(bodySize * portraitMultiplier)px + 1.0vw); font-weight: normal; }\n"
        htmlString += "      h1 { font: -apple-system-headine; font-family: Verdana, sans-serif; font-size:calc(\(h1Size * portraitMultiplier)px + 1.0vw); font-weight: bold; }\n"
        htmlString += "      h2 { font: -apple-system-headine; font-family: Verdana, sans-serif; font-size:calc(\(h2Size * portraitMultiplier)px + 1.0vw); font-weight: bold; }\n"
        htmlString += "      h3, h4 { font: -apple-system-headine; font-family: Verdana, sans-serif; font-size:calc(\(h3Size * portraitMultiplier)px + 1.0vw); font-weight: bold; } }\n"
        htmlString += "      @media all and (orientation:landscape) {img {max-width: 65%; height: auto;}\n"
        htmlString += "      p, li { font: -apple-system-body; font-family: Georgia, serif; font-size:calc(\(bodySize * landscapeMultiplier)px + 1.0vw); font-weight: normal; }\n"
        htmlString += "      h1 { font: -apple-system-headine; font-family: Verdana, sans-serif; font-size:calc(\(h1Size * landscapeMultiplier)px + 1.0vw); font-weight: bold; }\n"
        htmlString += "      h2 { font: -apple-system-headine; font-family: Verdana, sans-serif; font-size:calc(\(h2Size * landscapeMultiplier)px + 1.0vw); font-weight: bold; }\n"
        htmlString += "      h3, h4 { font: -apple-system-headine; font-family: Verdana, sans-serif; font-size:calc(\(h3Size * landscapeMultiplier)px + 1.0vw); font-weight: bold; } }\n"
        htmlString += "    </style>\n"
        htmlString += "  </head>\n"
        htmlString += "  <body>\n"
        htmlString += "    <meta name=\"viewport\" content=\"width: device-width\">\n"
        htmlString += "      \(bodyContent)\n"
        htmlString += "    </meta>\n"
        htmlString += "  </body>\n"
        htmlString += "</html>\n"
            
        return htmlString
    }
}
