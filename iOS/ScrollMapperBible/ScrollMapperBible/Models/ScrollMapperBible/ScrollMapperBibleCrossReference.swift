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
    private var vid: Int = 0
    
    public struct CrossReference {
        public let vid: Int
        public let r: Int
        public let sv: Int
        public let ev: Int
        public let verses: [ScrollMapperBibleText.BibleText]
    }
    
    public lazy var result: [StructType] = {
        return getResult()
    }()
    
    public required init?(statement: String) {
        super.init(statement: statement)
    }
    
    public init?(version: ScrollMapperBibleVersion.BibleVersion, book: ScrollMapperBibleBookInfo.BibleBook, chapter: Int, verse: Int) {
        vid = book.order() * 1_000_000 + chapter * 1_000 + verse
        let statement = "SELECT cr.vid, cr.r, cr.sv, cr.ev, tbl.id, tbl.b, tbl.c, tbl.v, tbl.t FROM cross_reference AS cr LEFT JOIN \(version.table() ?? "t_kjv") AS tbl ON (cr.ev = 0 AND tbl.id = cr.sv) OR (cr.ev > 0 AND tbl.id >= cr.sv AND tbl.id <= cr.ev) WHERE cr.vid = \(vid) ORDER BY cr.r, cr.sv"
        super.init(statement: statement)
    }
    
    public init?(version: ScrollMapperBibleVersion.BibleVersion, vid: Int) {
        // SELECT cr.vid, cr.r, cr.sv, cr.ev, tbl.b, tbl.c, tbl.v, tbl.t FROM cross_reference AS cr LEFT JOIN t_kjv AS tbl ON (cr.ev = 0 AND tbl.id = cr.sv) OR (cr.ev > 0 AND tbl.id >= cr.sv AND tbl.id <= cr.ev) WHERE cr.vid = 1001001 ORDER BY cr.r, cr.sv
        self.vid = vid
        let statement = "SELECT cr.vid, cr.r, cr.sv, cr.ev, tbl.id, tbl.b, tbl.c, tbl.v, tbl.t FROM cross_reference AS cr LEFT JOIN \(version.table() ?? "t_kjv") AS tbl ON (cr.ev = 0 AND tbl.id = cr.sv) OR (cr.ev > 0 AND tbl.id >= cr.sv AND tbl.id <= cr.ev) WHERE cr.vid = \(vid) ORDER BY cr.r, cr.sv"
        super.init(statement: statement)
    }
    
    private func getResult() -> [StructType] {
        guard let queryStatement =  queryStatement else {
            return []
        }
        var result: [StructType] = []
        var verses: [ScrollMapperBibleText.BibleText] = []
        var currentRank: Int = 0
        var currentSv: Int = 0
        var currentEv: Int = 0
        while sqlite3_step(queryStatement) == SQLITE_ROW {
            let _ = Int(sqlite3_column_int(queryStatement, 0)) // vid
            let r = Int(sqlite3_column_int(queryStatement, 1))
            let sv = Int(sqlite3_column_int(queryStatement, 2)) // sv
            let ev = Int(sqlite3_column_int(queryStatement, 3)) // ev
            let id = Int(sqlite3_column_int(queryStatement, 4))
            let b = Int(sqlite3_column_int(queryStatement, 5))
            let c = Int(sqlite3_column_int(queryStatement, 6))
            let v = Int(sqlite3_column_int(queryStatement, 7))
            guard let t = sqlite3_column_text(queryStatement, 8) else {
                continue
            }
            let text = String(cString: t)
            if (sv != currentSv) || (ev != currentEv) {
                if currentSv != 0 {
                    result.append((StructType(vid: vid, r: currentRank, sv: currentSv, ev: currentEv, verses: verses)))
                    verses = []
                }
                currentRank = r
                currentSv = sv
                currentEv = ev
            }
            verses.append(ScrollMapperBibleText.BibleText(id: id, b: b, c: c, v: v, t: text))
        }
        result.append((StructType(vid: vid, r: currentRank, sv: currentSv, ev: currentEv, verses: verses)))
        return result
    }
}
