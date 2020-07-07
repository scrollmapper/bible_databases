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
    
    public required init?(statement: String) {
        super.init(statement: statement)
    }
    
    public init?(version: ScrollMapperBibleVersion.BibleVersion, vidStart: Int, vidEnd: Int) {
        guard let table = version.table() else {
            return nil
        }
        let statement = "SELECT * FROM \(table) WHERE id >= \(vidStart) AND id <= \(vidEnd)"
        super.init(statement: statement)
    }
    
    public init?(version: ScrollMapperBibleVersion.BibleVersion, book: ScrollMapperBibleBookInfo.BibleBook, chapter: Int) {
        guard let table = version.table() else {
            return nil
        }
        let statement = "SELECT * FROM \(table) WHERE b = \(book.order()) AND c = \(chapter)"
        super.init(statement: statement)
    }
    
    public init?(version: ScrollMapperBibleVersion.BibleVersion, book: ScrollMapperBibleBookInfo.BibleBook, chapter: Int, verseFrom: Int) {
        guard let table = version.table() else {
            return nil
        }
        let statement = "SELECT * FROM \(table) WHERE b = \(book.order()) AND c = \(chapter) AND v >= \(verseFrom)"
        super.init(statement: statement)
    }
    
    public init?(version: ScrollMapperBibleVersion.BibleVersion, book: ScrollMapperBibleBookInfo.BibleBook, chapter: Int, verseUntil: Int) {
        guard let table = version.table() else {
            return nil
        }
        let statement = "SELECT * FROM \(table) WHERE b = \(book.order()) AND c = \(chapter) AND v < \(verseUntil)"
        super.init(statement: statement)
    }
    
    public init?(version: ScrollMapperBibleVersion.BibleVersion, book: ScrollMapperBibleBookInfo.BibleBook, chapter: Int, verseThrough: Int) {
        guard let table = version.table() else {
            return nil
        }
        let statement = "SELECT * FROM \(table) WHERE b = \(book.order()) AND c = \(chapter) AND v <= \(verseThrough)"
        super.init(statement: statement)
    }
    
    public init?(version: ScrollMapperBibleVersion.BibleVersion, book: ScrollMapperBibleBookInfo.BibleBook, chapter: Int, verseRange: Range<Int>) {
        guard let table = version.table() else {
            return nil
        }
        let statement = "SELECT * FROM \(table) WHERE b = \(book.order()) AND c = \(chapter) AND v >= \(verseRange.lowerBound) AND v < \(verseRange.upperBound)"
        super.init(statement: statement)
    }
    
    public init?(version: ScrollMapperBibleVersion.BibleVersion, book: ScrollMapperBibleBookInfo.BibleBook, chapter: Int, verseRangeClosed: ClosedRange<Int>) {
        guard let table = version.table() else {
            return nil
        }
        let statement = "SELECT * FROM \(table) WHERE b = \(book.order()) AND c = \(chapter) AND v >= \(verseRangeClosed.lowerBound) AND v <= \(verseRangeClosed.upperBound)"
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
        return result.sorted {$0.id < $1.id }
    }
    
    public static func test() {
        print("testScrollMapperBibleText 63001001~64001015")
        let _ = ScrollMapperBibleText(version: .KJV, vidStart: 63001001, vidEnd: 64001015)?.result.map {
            print("\($0.id), (\($0.b), \($0.c), \($0.v)), t: \($0.t)")
        }
        
        print("testScrollMapperBibleText [Matthew 28]")
        let _ = ScrollMapperBibleText(version: .KJV, book: .Matthew, chapter: 28)?.result.map {
            print("\($0.id), (\($0.b), \($0.c), \($0.v)), t: \($0.t)")
        }
        
        print("testScrollMapperBibleText [Matthew 28:16-]")
        let _ = ScrollMapperBibleText(version: .KJV, book: .Matthew, chapter: 28, verseFrom: 16)?.result.map {
            print("\($0.id), (\($0.b), \($0.c), \($0.v)), t: \($0.t)")
        }
        
        print("testScrollMapperBibleText [Matthew 28:-<16]")
        let _ = ScrollMapperBibleText(version: .KJV, book: .Matthew, chapter: 28, verseUntil: 16)?.result.map {
            print("\($0.id), (\($0.b), \($0.c), \($0.v)), t: \($0.t)")
        }
        
        print("testScrollMapperBibleText [Matthew 28:-15]")
        let _ = ScrollMapperBibleText(version: .KJV, book: .Matthew, chapter: 28, verseThrough: 15)?.result.map {
            print("\($0.id), (\($0.b), \($0.c), \($0.v)), t: \($0.t)")
        }
        
        print("testScrollMapperBibleText [Matthew 28:1-<16]")
        let _ = ScrollMapperBibleText(version: .KJV, book: .Matthew, chapter: 28, verseRange: 1..<16)?.result.map {
            print("\($0.id), (\($0.b), \($0.c), \($0.v)), t: \($0.t)")
        }
        
        print("testScrollMapperBibleText [Matthew 28:1-15]")
        let _ = ScrollMapperBibleText(version: .KJV, book: .Matthew, chapter: 28, verseRangeClosed: 1...15)?.result.map {
            print("\($0.id), (\($0.b), \($0.c), \($0.v)), t: \($0.t)")
        }
    }
}
