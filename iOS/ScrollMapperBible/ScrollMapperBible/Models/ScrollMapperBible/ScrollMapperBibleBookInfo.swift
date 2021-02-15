//
//  ScrollMapperBibleBookInfo.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-06.
//  Copyright © 2020 Kuang. All rights reserved.
//

import Foundation
import SQLite3

public class ScrollMapperBibleBookInfo: ScrollMapperBibleModelBase {
    public static let books: [BookInfo] = {
        if let books = ScrollMapperBibleBookInfo()?.result {
            return books
        }
        return []
    }()
    
    public enum BibleOTNT: CaseIterable {
        case OT, NT
        
        public var title: String {
            switch self {
            case .OT:
                return "Old Testament"
            case .NT:
                return "New Testament"
            }
        }
        
        public var chapters: Int {
            switch self {
            case .OT:
                return 39
            case .NT:
                return 27
            }
        }
    }
    
    public enum BibleBook: Int, CaseIterable {
        case Genesis = 1
        case Exodus = 2
        case Leviticus = 3
        case Numbers = 4
        case Deuteronomy = 5
        
        case Joshua = 6
        case Judges = 7
        case Ruth = 8
        case Samuel1 = 9
        case Samuel2 = 10
        case Kings1 = 11
        case Kings2 = 12
        case Chronicles1 = 13
        case Chronicles2 = 14
        case Ezra = 15
        case Nehemiah = 16
        case Esther = 17
        
        case Job = 18
        case Psalms = 19
        case Proverbs = 20
        case Ecclesiastes = 21
        case Song = 22
        
        case Isaiah = 23
        case Jeremiah = 24
        case Lamentations = 25
        case Ezekiel = 26
        case Daniel = 27
        
        case Hosea = 28
        case Joel = 29
        case Amos = 30
        case Obadiah = 31
        case Jonah = 32
        case Micah = 33
        case Nahum = 34
        case Habakkuk = 35
        case Zephaniah = 36
        case Haggai = 37
        case Zechariah = 38
        case Malachi = 39
        
        case Matthew = 40
        case Mark = 41
        case Luke = 42
        case John = 43
        case Acts = 44
        
        case Romans = 45
        case Corinthians1 = 46
        case Corinthians2 = 47
        case Galatians = 48
        case Ephesians = 49
        case Philippians = 50
        case Colossians = 51
        case Thessalonians1 = 52
        case Thessalonians2 = 53
        case Timothy1 = 54
        case Timothy2 = 55
        case Titus = 56
        case Philemon = 57
        
        case Hebrews = 58
        case James = 59
        case Peter1 = 60
        case Peter2 = 61
        case John1 = 62
        case John2 = 63
        case John3 = 64
        case Jude = 65
        case Revelation = 66
        
        public func bookInfo() -> BookInfo? {
            let book = ScrollMapperBibleBookInfo.books.first { $0.order == self.order() }
            return book
        }
        
        public func order() -> Int {
            return self.rawValue
        }
    }
    
    public typealias StructType = BookInfo
    
    public struct BookInfo: Equatable {
        var order: Int = 0
        var title_short: String = ""
        var title_full: String = ""
        var abbreviation: String = ""
        var category: String = ""
        var otnt: String = ""
        var chapters: Int = 0
        
        init(order: Int, title_short: String, title_full: String, abbreviation: String, category: String, otnt: String, chapters: Int) {
            self.order = order
            self.title_short = title_short
            self.title_full = title_full
            self.abbreviation = abbreviation
            self.category = category
            self.otnt = otnt
            self.chapters = chapters
        }
        
        init?(order: Int) {
            guard let book = ScrollMapperBibleBookInfo.books.first(where: { (bookInfo) -> Bool in
                bookInfo.order == order
            }) else { return nil }
            self.order = book.order
            title_short = book.title_short
            title_full = book.title_full
            abbreviation = book.abbreviation
            category = book.category
            otnt = book.otnt
            chapters = book.chapters
        }
        
        public static func ==(lhs: Self, rhs: Self) -> Bool {
            return lhs.order == rhs.order
        }
    }
    
    public lazy var result: [StructType] = {
        return getResult()
    }()
    
    public required init?(statement: String = "SELECT * FROM book_info ORDER BY \"order\"") {
        super.init(statement: statement)
    }
    
    private func getResult() -> [StructType] {
        guard let queryStatement =  queryStatement else {
            return []
        }
        var result: [StructType] = []
        while sqlite3_step(queryStatement) == SQLITE_ROW {
            let order = sqlite3_column_int(queryStatement, 0)
            guard
                let title_short = sqlite3_column_text(queryStatement, 1),
                let title_full = sqlite3_column_text(queryStatement, 2),
                let abbreviation = sqlite3_column_text(queryStatement, 3),
                let category = sqlite3_column_text(queryStatement, 4),
                let otnt = sqlite3_column_text(queryStatement, 5)
            else {
                continue
            }
            let chapters = sqlite3_column_int(queryStatement, 6)
            let row = StructType(order: Int(order), title_short: String(cString: title_short), title_full: String(cString: title_full), abbreviation: String(cString: abbreviation), category: String(cString: category), otnt: String(cString: otnt), chapters: Int(chapters))
            result.append(row)
        }
        return result
    }
    
    public static func test() {
        print("testScrollMapperBibleBookInfo")
        _ = ScrollMapperBibleBookInfo()?.result.map {
            print("\($0.order), \($0.title_short), \($0.title_full), \($0.abbreviation), \($0.category), \($0.otnt)")
        }
    }
}
