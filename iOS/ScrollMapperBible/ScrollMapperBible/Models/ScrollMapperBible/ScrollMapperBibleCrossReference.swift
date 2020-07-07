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
    
    public required init?(statement: String = "SELECT * FROM cross_reference") {
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
        print("testScrollMapperBibleCrossReference")
        let _ = ScrollMapperBibleCrossReference(statement: "SELECT * FROM cross_reference WHERE vid = 1001001")?.result.map {
            print("vid: \($0.vid), r: \($0.r), sv: \($0.sv), ev: \($0.ev)")
        }
    }
}
