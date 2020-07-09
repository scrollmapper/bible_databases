//
//  ScrollMapperBibleCrossReference.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-06.
//  Copyright © 2020 Kuang. All rights reserved.
//

import Foundation
import SQLite3

public class ScrollMapperBibleCrossReference: ScrollMapperBibleModelBase {
    public typealias StructType = CrossReference
    
    public struct CrossReference {
        var vid: Int = 0
        var r: Int = 0
        var sv: Int = 0
        var ev: Int = 0
    }
    
    public lazy var result: [StructType] = {
        return getResult()
    }()
    
    public required init?(statement: String) {
        super.init(statement: statement)
    }
    
    public init?(book: ScrollMapperBibleBookInfo.BibleBook, chapter: Int, verse: Int) {
        let vid = book.order() * 1_000_000 + chapter * 1_000 + verse
        let statement = "SELECT * FROM cross_reference WHERE vid = \(vid)"
        super.init(statement: statement)
    }
    
    public init?(vid: Int) {
        let statement = "SELECT * FROM cross_reference WHERE vid = \(vid)"
        super.init(statement: statement)
    }
    
    private func getResult() -> [StructType] {
        guard let queryStatement =  queryStatement else {
            return []
        }
        var result: [StructType] = []
        while sqlite3_step(queryStatement) == SQLITE_ROW {
            let vid = sqlite3_column_int(queryStatement, 0)
            let r = sqlite3_column_int(queryStatement, 1)
            let sv = sqlite3_column_int(queryStatement, 2)
            let ev = sqlite3_column_int(queryStatement, 3)
            let row = StructType(vid: Int(vid), r: Int(r), sv: Int(sv), ev: Int(ev))
            result.append(row)
        }
        return result.sorted { $0.r < $1.r }
    }
    
    public static func test() {
        print("testScrollMapperBibleCrossReference 1001001")
        _ = ScrollMapperBibleCrossReference(statement: "SELECT * FROM cross_reference WHERE vid = 1001001")?.result.compactMap { crossReference -> CrossReference? in
            print("vid: \(crossReference.vid), r: \(crossReference.r), sv: \(crossReference.sv), ev: \(crossReference.ev)")
            return nil
        }
        
        print("Cross references that run across books")
        _ = ScrollMapperBibleCrossReference(statement: "SELECT * FROM cross_reference WHERE ev > 0")?.result.compactMap { crossReference -> CrossReference? in
            let bookStart = crossReference.sv / 1_000_000
            let bookEnd = crossReference.ev / 1_000_000
            if bookEnd > bookStart {
                print("vid: \(crossReference.vid), r: \(crossReference.r), sv: \(crossReference.sv), ev: \(crossReference.ev)")
                return crossReference
            }
            return nil
        }
    }
}
