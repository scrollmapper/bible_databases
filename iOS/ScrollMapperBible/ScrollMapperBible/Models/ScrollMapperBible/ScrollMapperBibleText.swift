//
//  ScrollMapperBibleText.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-06.
//  Copyright © 2020 Kuang. All rights reserved.
//

import Foundation
import SQLite3

public enum ScrollMapperBibleSearchScope: Int, CaseIterable {
    case OT = 1
    case NT = 2
    case All = 3
    
    public var titleFull: String {
        switch self {
        case .All:
            return "All"
        case .OT:
            return "Old Testament"
        case .NT:
            return "New Testament"
        }
    }
    
    public var titleAbbreviated: String {
        switch self {
        case .All:
            return "All"
        case .OT:
            return "OT"
        case .NT:
            return "NT"
        }
    }
}

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
    
    public required init?(statement: String) {
        super.init(statement: statement)
    }
    
    public init?(version: ScrollMapperBibleVersion.BibleVersion, vidStart: Int, vidEnd: Int) {
        guard let table = version.table() else {
            return nil
        }
        let statement = "SELECT * FROM \(table) WHERE id >= \(vidStart) AND id <= \(vidEnd) ORDER BY id"
        super.init(statement: statement)
    }
    
    public init?(version: ScrollMapperBibleVersion.BibleVersion, book: ScrollMapperBibleBookInfo.BibleBook, chapter: Int) {
        guard let table = version.table() else {
            return nil
        }
        let statement = "SELECT * FROM \(table) WHERE b = \(book.order()) AND c = \(chapter) ORDER BY id"
        super.init(statement: statement)
    }
    
    public init?(version: ScrollMapperBibleVersion.BibleVersion, book: ScrollMapperBibleBookInfo.BibleBook, chapter: Int, verseFrom: Int) {
        guard let table = version.table() else {
            return nil
        }
        let statement = "SELECT * FROM \(table) WHERE b = \(book.order()) AND c = \(chapter) AND v >= \(verseFrom) ORDER BY id"
        super.init(statement: statement)
    }
    
    public init?(version: ScrollMapperBibleVersion.BibleVersion, book: ScrollMapperBibleBookInfo.BibleBook, chapter: Int, verseUntil: Int) {
        guard let table = version.table() else {
            return nil
        }
        let statement = "SELECT * FROM \(table) WHERE b = \(book.order()) AND c = \(chapter) AND v < \(verseUntil) ORDER BY id"
        super.init(statement: statement)
    }
    
    public init?(version: ScrollMapperBibleVersion.BibleVersion, book: ScrollMapperBibleBookInfo.BibleBook, chapter: Int, verseThrough: Int) {
        guard let table = version.table() else {
            return nil
        }
        let statement = "SELECT * FROM \(table) WHERE b = \(book.order()) AND c = \(chapter) AND v <= \(verseThrough) ORDER BY id"
        super.init(statement: statement)
    }
    
    public init?(version: ScrollMapperBibleVersion.BibleVersion, book: ScrollMapperBibleBookInfo.BibleBook, chapter: Int, verseRange: Range<Int>) {
        guard let table = version.table() else {
            return nil
        }
        let statement = "SELECT * FROM \(table) WHERE b = \(book.order()) AND c = \(chapter) AND v >= \(verseRange.lowerBound) AND v < \(verseRange.upperBound) ORDER BY id"
        super.init(statement: statement)
    }
    
    public init?(version: ScrollMapperBibleVersion.BibleVersion, book: ScrollMapperBibleBookInfo.BibleBook, chapter: Int, verseRangeClosed: ClosedRange<Int>) {
        guard let table = version.table() else {
            return nil
        }
        let statement = "SELECT * FROM \(table) WHERE b = \(book.order()) AND c = \(chapter) AND v >= \(verseRangeClosed.lowerBound) AND v <= \(verseRangeClosed.upperBound) ORDER BY id"
        super.init(statement: statement)
    }
    
    public init?(version: ScrollMapperBibleVersion.BibleVersion, searchBy term: String, in scope: ScrollMapperBibleSearchScope) {
        guard let table = version.table() else {
            return nil
        }
        let keywords = term.components(separatedBy: .whitespaces).compactMap { (keyword) -> String? in
            if keyword.trimmingCharacters(in: .whitespaces).count == 0 {
                return nil
            }
            return keyword
        }
        guard keywords.count > 0 else {
            return nil
        }
        var statementKeywordsPart = ""
        _ = keywords.compactMap({ (keyword) -> String? in
            if statementKeywordsPart.count > 0 {
                statementKeywordsPart += " AND"
            }
            statementKeywordsPart += " t LIKE '%\(keyword)%'"
            return nil
        })
        var statement = "SELECT * FROM \(table) WHERE" + statementKeywordsPart
        if scope == .OT {
            statement += " AND b < 40"
        }
        else if scope == .NT {
            statement += " AND b > 39"
        }
        statement += " ORDER BY id"
        super.init(statement: statement)
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
        return result
    }
    
    public static func test() {
        print("testScrollMapperBibleText 63001001~64001015")
        _ = ScrollMapperBibleText(version: .KJV, vidStart: 63001001, vidEnd: 64001015)?.result.map {
            print("\($0.id), (\($0.b), \($0.c), \($0.v)), t: \($0.t)")
        }
        
        print("testScrollMapperBibleText [Matthew 28]")
        _ = ScrollMapperBibleText(version: .KJV, book: .Matthew, chapter: 28)?.result.map {
            print("\($0.id), (\($0.b), \($0.c), \($0.v)), t: \($0.t)")
        }
        
        print("testScrollMapperBibleText [Matthew 28:16-]")
        _ = ScrollMapperBibleText(version: .KJV, book: .Matthew, chapter: 28, verseFrom: 16)?.result.map {
            print("\($0.id), (\($0.b), \($0.c), \($0.v)), t: \($0.t)")
        }
        
        print("testScrollMapperBibleText [Matthew 28:-<16]")
        _ = ScrollMapperBibleText(version: .KJV, book: .Matthew, chapter: 28, verseUntil: 16)?.result.map {
            print("\($0.id), (\($0.b), \($0.c), \($0.v)), t: \($0.t)")
        }
        
        print("testScrollMapperBibleText [Matthew 28:-15]")
        _ = ScrollMapperBibleText(version: .KJV, book: .Matthew, chapter: 28, verseThrough: 15)?.result.map {
            print("\($0.id), (\($0.b), \($0.c), \($0.v)), t: \($0.t)")
        }
        
        print("testScrollMapperBibleText [Matthew 28:1-<16]")
        _ = ScrollMapperBibleText(version: .KJV, book: .Matthew, chapter: 28, verseRange: 1..<16)?.result.map {
            print("\($0.id), (\($0.b), \($0.c), \($0.v)), t: \($0.t)")
        }
        
        print("testScrollMapperBibleText [Matthew 28:1-15]")
        _ = ScrollMapperBibleText(version: .KJV, book: .Matthew, chapter: 28, verseRangeClosed: 1...15)?.result.map {
            print("\($0.id), (\($0.b), \($0.c), \($0.v)), t: \($0.t)")
        }
        
        print("testScrollMapperBibleText search for 'in the beginning' in all")
        _ = ScrollMapperBibleText(version: .KJV, searchBy: "in the beginning", in: .All)?.result.map {
            print("\($0.id), (\($0.b), \($0.c), \($0.v)), t: \($0.t)")
        }
        
        print("testScrollMapperBibleText search for 'in the beginning' in OT")
        _ = ScrollMapperBibleText(version: .KJV, searchBy: "in the beginning", in: .OT)?.result.map {
            print("\($0.id), (\($0.b), \($0.c), \($0.v)), t: \($0.t)")
        }
        
        print("testScrollMapperBibleText search for 'in the beginning' in NT")
        _ = ScrollMapperBibleText(version: .KJV, searchBy: "in the beginning", in: .NT)?.result.map {
            print("\($0.id), (\($0.b), \($0.c), \($0.v)), t: \($0.t)")
        }
    }
}
