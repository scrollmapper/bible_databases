//
//  ScrollMapperBibleTextViewModel.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-07.
//  Copyright © 2020 Kuang. All rights reserved.
//

import Foundation
import Combine

class ScrollMapperBibleTextViewModel: ScrollMapperBibleViewModelBase {
    override init() {
        super.init()
        
        self.setupListData()
    }
    
    override func translationDidChange() {
        super.translationDidChange()
        
        setupListData()
    }
    
    struct Section: Identifiable {
        var id = UUID()
        var bibleChapter: ScrollMapperBibleChapter? = nil
        var verses: [ScrollMapperBibleText.BibleText] = []
        var title: String = ""
        var items: [Item] = [] // each section contains only one item because we join all verses of a chapter into one string
    }
    
    struct Item: Identifiable {
        var id = UUID()
        var text: String = ""
    }
    
    @Published var listData: [Section] = []
    
    private func setupListData() {
        var listData = [Section]()
        var currentChapter = Section()
        var currentChapterText = ""
        var currentB: Int = 0
        var currentC: Int = 0
        // vidStart: 1001001, vidEnd: 66022021
        _ = ScrollMapperBibleText(version: translation, vidStart: 1001001, vidEnd: 66022021)?.result.compactMap({ (bibleText) -> ScrollMapperBibleText.BibleText? in
            if (bibleText.b != currentB) || (bibleText.c != currentC) {
                if (currentB != 0) && (currentC != 0) {
                    currentChapter.bibleChapter = ScrollMapperBibleChapter(b: currentB, c: currentC)
                    let bookInfo = ScrollMapperBibleBookInfo.BookInfo(order: currentB)
                    currentChapter.title = "\(bookInfo?.title_short ?? "<nil>") \(currentC)"
                    let item = Item(text: currentChapterText)
                    currentChapter.items.append(item)
                    listData.append(currentChapter)
                    currentChapter = Section()
                    currentChapterText = ""
                }
                currentB = bibleText.b
                currentC = bibleText.c
            }
            currentChapter.verses.append(bibleText)
            if currentChapterText.count > 0 {
                currentChapterText.append(contentsOf: " ")
            }
            currentChapterText.append(contentsOf: "\(bibleText.v) \(bibleText.t)")
            return nil
        })
        currentChapter.bibleChapter = ScrollMapperBibleChapter(b: currentB, c: currentC)
        let bookInfo = ScrollMapperBibleBookInfo.BookInfo(order: currentB)
        currentChapter.title = "\(bookInfo?.title_short ?? "<nil>") \(currentC)"
        let item = Item(text: currentChapterText)
        currentChapter.items.append(item)
        listData.append(currentChapter)
        
        self.listData = listData
    }
}
