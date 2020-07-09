//
//  ScrollMapperBibleTable.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-09.
//  Copyright © 2020 Kuang. All rights reserved.
//

import Foundation
import SQLite3

public class ScrollMapperBibleTable {
    public static func generateChapterCidOrderScript() -> String {
        var script = ""
        script += "CREATE TABLE IF NOT EXISTS \"chapter\" (\n"
        script += "  \"order\" INTEGER NOT NULL UNIQUE,\n"
        script += "  \"cid\" INTEGER NOT NULL UNIQUE,\n"
        script += "  \"b\" INTEGER NOT NULL,\n"
        script += "  \"d\" INTEGER NOT NULL\n"
        script += ");\n"
        
        var bibleChapter = ScrollMapperBibleChapter.BibleChapter(order: 1, b: 1, c: 1)
        var order: Int = 1
        while bibleChapter != nil {
            script += "INSERT INTO chapter VALUES(\(order), \(bibleChapter!.cid), \(bibleChapter!.b), \(bibleChapter!.c));\n"
            bibleChapter = bibleChapter!~>
            order += 1
        }
        return script
    }
}
