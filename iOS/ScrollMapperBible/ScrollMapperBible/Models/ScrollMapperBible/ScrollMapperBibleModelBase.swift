//
//  ScrollMapperBibleModelBase.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-06.
//  Copyright © 2020 Kuang. All rights reserved.
//

import Foundation

public class ScrollMapperBibleModelBase: JKCSSQLiteQuery {
    let resourceName = "bible-sqlite"
    let resourceExt = "db"
    
    public required init?(statement: String) {
        super.init(resourceName: resourceName, resourceExt: resourceExt, statement: statement)
    }
}
