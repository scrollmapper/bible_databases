//
//  ScrollMapperBibleChapter.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-06.
//  Copyright © 2020 Kuang. All rights reserved.
//

import Foundation

prefix operator <~
postfix operator ~>

public class ScrollMapperBibleChapter: Comparable {
    public let cid: Int
    public let b: Int // book, 0...66
    public let c: Int // chapter, 1...
    
    public init?(b: Int, c: Int) {
        guard
            let bookInfo = ScrollMapperBibleBookInfo.BookInfo(order: b),
            c > 0 && c <= bookInfo.chapters
        else {
            return nil
        }
        self.b = b
        self.c = c
        cid = b * 1_000 + c
    }
    
    public static func < (lhs: ScrollMapperBibleChapter, rhs: ScrollMapperBibleChapter) -> Bool {
        return lhs.cid < rhs.cid
    }
    
    public static func == (lhs: ScrollMapperBibleChapter, rhs: ScrollMapperBibleChapter) -> Bool {
        return lhs.cid == rhs.cid
    }
    
    public static postfix func ~>(lhs: ScrollMapperBibleChapter) -> ScrollMapperBibleChapter? {
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
        return ScrollMapperBibleChapter(b: b, c: c)
    }
    
    public static prefix func <~(rhs: ScrollMapperBibleChapter) -> ScrollMapperBibleChapter? {
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
        return ScrollMapperBibleChapter(b: b, c: c)
    }
    
    public static func test() {
        print("ScrollMapperBibleChapter ~>")
        var c = ScrollMapperBibleChapter(b: 1, c: 1)
        while c != nil {
            print("book: \(c?.b ?? 0), chapter: \(c?.c ?? 0)")
            c = c!~>
        }
        print("book: \(c?.b ?? 0), chapter: \(c?.c ?? 0)")
        
        print("ScrollMapperBibleChapter <~")
        c = ScrollMapperBibleChapter(b: 66, c: 22)
        while c != nil {
            print("book: \(c?.b ?? 0), chapter: \(c?.c ?? 0)")
            c = <~c!
        }
        print("book: \(c?.b ?? 0), chapter: \(c?.c ?? 0)")
    }
}
