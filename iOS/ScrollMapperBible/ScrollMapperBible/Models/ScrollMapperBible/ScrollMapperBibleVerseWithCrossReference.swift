//
//  ScrollMapperBibleVerseWithCrossReference.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-11.
//  Copyright © 2020 Kuang. All rights reserved.
//

import Foundation
import SQLite3

public class ScrollMapperBibleVerseWithCrossReference: ScrollMapperBibleModelBase {
    public typealias StructType = VerseWithCrossReference
    
    public struct VerseWithCrossReference {
        var id: Int = 0
        var b: Int = 0
        var c: Int = 0
        var v: Int = 0
        var t: String = ""
        var cr: [CrossReference] = []
    }
    
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
    
    public required init?(version: ScrollMapperBibleVersion.BibleVersion, book: ScrollMapperBibleBookInfo.BibleBook, chapter: Int) {
        guard let table = version.table() else {
            return nil
        }
        let statement = "SELECT bk.id, bk.b, bk.c, bk.v, bk.t, cr.r, cr.sv, cr.ev FROM \(table) as bk LEFT JOIN cross_reference as cr ON bk.id = cr.vid WHERE bk.b = \(book.order()) AND bk.c = \(chapter) ORDER BY bk.id, cr.r"
        super.init(statement: statement)
    }
    
    private func getResult() -> [StructType] {
        guard let queryStatement = queryStatement else {
            return []
        }
        var result: [StructType] = []
        var crossReferences: [CrossReference] = []
        var current_vid: Int = 0
        var current_b: Int = 0
        var current_c: Int = 0
        var current_v: Int = 0
        var current_t: String = ""
        while sqlite3_step(queryStatement) == SQLITE_ROW {
            let vid = Int(sqlite3_column_int(queryStatement, 0))
            let b = Int(sqlite3_column_int(queryStatement, 1))
            let c = Int(sqlite3_column_int(queryStatement, 2))
            let v = Int(sqlite3_column_int(queryStatement, 3))
            guard let text = sqlite3_column_text(queryStatement, 4) else {
                continue
            }
            let t = String(cString: text)
            let cr_r = Int(sqlite3_column_int(queryStatement, 5))
            let cr_sv = Int(sqlite3_column_int(queryStatement, 6))
            let cr_ev = Int(sqlite3_column_int(queryStatement, 7))
            if vid != current_vid {
                if current_vid != 0 {
                    result.append(VerseWithCrossReference(id: current_vid, b: current_b, c: current_c, v: current_v, t: current_t, cr: crossReferences))
                    crossReferences = []
                }
                current_vid = vid
                current_b = b
                current_c = c
                current_v = v
                current_t = t
            }
            crossReferences.append(CrossReference(vid: vid, r: cr_r, sv: cr_sv, ev: cr_ev))
        }
        result.append(VerseWithCrossReference(id: current_vid, b: current_b, c: current_c, v: current_v, t: current_t, cr: crossReferences))
        
        return result
    }
}
