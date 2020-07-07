//
//  ScrollMapperBibleText.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-06.
//  Copyright © 2020 Kuang. All rights reserved.
//

import Foundation
import SQLite3

public class ScrollMapperBibleText: ScrollMapperBibleModelBase {
    public typealias StructType = BibleText
    
    public struct BibleText {
        var id: Int = 0
        var b: Int = 0
        var c: Int = 0
        var v: Int = 0
        var t: String = ""
    }
    
    public lazy var result: [StructType] = {
        return getResult()
    }()
    
    public required init?(version: ScrollMapperBibleVersion.BibleVersion, book: ScrollMapperBibleBookInfo.BibleBook, chapter: Range<Int>, verses: Range<Int>) {
        let statement = "SELECT * FROM"
        super.init(statement: statement)
    }
    
    public required init?(statement: String) {
        fatalError("init(statement:) has not been implemented")
    }
    
    private func getResult() -> [StructType] {
        guard let queryStatement =  queryStatement else {
            return []
        }
        var result: [StructType] = []
        while sqlite3_step(queryStatement) == SQLITE_ROW {
            let id = sqlite3_column_int(queryStatement, 0)
            let b = sqlite3_column_int(queryStatement, 1)
            let c = sqlite3_column_int(queryStatement, 2)
            let v = sqlite3_column_int(queryStatement, 3)
            guard let t = sqlite3_column_text(queryStatement, 4) else {
                continue
            }
            let row = StructType(id: Int(id), b: Int(b), c: Int(c), v: Int(v), t: String(cString: t))
            result.append(row)
        }
        return result.sorted {$0.id < $1.id }
    }
    
    public static func test() {
        print("testScrollMapperBibleCrossReference")
        let _ = ScrollMapperBibleCrossReference(statement: "SELECT * FROM cross_reference WHERE vid = 1001001")?.result.map {
            print("vid: \($0.vid), r: \($0.r), sv: \($0.sv), ev: \($0.ev)")
        }
    }
}
