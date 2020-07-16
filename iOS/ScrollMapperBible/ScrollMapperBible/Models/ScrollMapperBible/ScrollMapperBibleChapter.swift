//
//  ScrollMapperBibleChapter.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-06.
//  Copyright © 2020 Kuang. All rights reserved.
//

import Foundation
import SQLite3

prefix operator <~
postfix operator ~>

public class ScrollMapperBibleChapter: ScrollMapperBibleModelBase {
    public static let chapters: [BibleChapter] = {
        if let chapters = ScrollMapperBibleChapter()?.result {
            return chapters
        }
        return []
    }()
    
    public typealias StructType = BibleChapter
    
    public struct BibleChapter: Comparable {
        let order: Int
        let cid: Int
        let b: Int
        let c: Int
        
        init?(order: Int, cid: Int, b: Int, c: Int) {
            guard (order > 0) && (order <= 1189) && (cid > 0) && (b > 0) && (b <= 66) && (c > 0) && (cid == (b * 1_000 + c)) else {
                return nil
            }
            self.order = order
            self.cid = cid
            self.b = b
            self.c = c
        }
        
        init?(order: Int) {
            let bibleChapter = ScrollMapperBibleChapter.chapters.first { $0.order == order }
            if let bibleChapter = bibleChapter {
                self.order = bibleChapter.order
                self.cid = bibleChapter.cid
                self.b = bibleChapter.b
                self.c = bibleChapter.c
            }
            else {
                return nil
            }
        }
        
        init?(b: Int, c: Int) {
            let bibleChapter = ScrollMapperBibleChapter.chapters.first { ($0.b == b) && ($0.c == c) }
            if let bibleChapter = bibleChapter {
                self.order = bibleChapter.order
                self.cid = bibleChapter.cid
                self.b = bibleChapter.b
                self.c = bibleChapter.c
            }
            else {
                return nil
            }
        }
        
        init?(cid: Int) {
            let bibleChapter = ScrollMapperBibleChapter.chapters.first { $0.cid == cid }
            if let bibleChapter = bibleChapter {
                self.order = bibleChapter.order
                self.cid = bibleChapter.cid
                self.b = bibleChapter.b
                self.c = bibleChapter.c
            }
            else {
                return nil
            }
        }
        
        init?(order: Int, cid: Int) {
            guard (order > 0) && (order <= 1189) && (cid > 0) else {
                return nil
            }
            let b = Int(cid / 1_000)
            let c = Int(cid % 1000)
            guard (b > 0) && (b <= 66) && (c > 0) else {
                return nil
            }
            self.order = order
            self.cid = cid
            self.b = b
            self.c = c
        }
        
        init?(order: Int, b: Int, c: Int) {
            guard (order > 0) && (order <= 1189) && (b > 0) && (b <= 66) && (c > 0) else {
                return nil
            }
            let cid = (b * 1_000 + c)
            self.order = order
            self.cid = cid
            self.b = b
            self.c = c
        }
        
        public static func < (lhs: BibleChapter, rhs: BibleChapter) -> Bool {
            return lhs.cid < rhs.cid
        }
        
        public static func == (lhs: BibleChapter, rhs: BibleChapter) -> Bool {
            return lhs.cid == rhs.cid
        }
        
        public static postfix func ~>(lhs: BibleChapter) -> BibleChapter? {
            var b: Int = 0
            var c: Int = 0
            guard let bookInfo = ScrollMapperBibleBookInfo.BookInfo(order: lhs.b) else { return nil }
            if lhs.c < bookInfo.chapters {
                b = lhs.b
                c = lhs.c + 1
            }
            else {
                if lhs.b < 66 {
                    b = lhs.b + 1
                    c = 1
                }
                else {
                    return nil
                }
            }
            return BibleChapter(order: (lhs.order + 1), b: b, c: c)
        }
        
        public static prefix func <~(rhs: BibleChapter) -> BibleChapter? {
            guard rhs.order > 1 else {
                return nil
            }
            var b: Int = 0
            var c: Int = 0
            if rhs.c > 1 {
                b = rhs.b
                c = rhs.c - 1
            }
            else {
                b = rhs.b - 1
                guard let bookInfo = ScrollMapperBibleBookInfo.BookInfo(order: b) else {
                    return nil
                }
                c = bookInfo.chapters
            }
            return BibleChapter(order: (rhs.order - 1), b: b, c: c)
        }
        
        var bibleBook: ScrollMapperBibleBookInfo.BibleBook {
            ScrollMapperBibleBookInfo.BibleBook(rawValue: b)!
        }
    }
    
    public lazy var result: [StructType] = {
        return getResult()
    }()
    
    public required init?(statement: String = "SELECT * FROM chapter ORDER BY cid") {
        super.init(statement: statement)
    }
    
    private func getResult() -> [StructType] {
        guard let queryStatement = queryStatement else {
            return []
        }
        var result: [StructType] = []
        while sqlite3_step(queryStatement) == SQLITE_ROW {
            let order = sqlite3_column_int(queryStatement, 0)
            let cid = sqlite3_column_int(queryStatement, 1)
            let b = sqlite3_column_int(queryStatement, 2)
            let c = sqlite3_column_int(queryStatement, 3)
            guard let row = StructType(order: Int(order), cid: Int(cid), b: Int(b), c: Int(c)) else {
                continue
            }
            result.append(row)
        }
        return result
    }
}
