//
//  ScrollMapperBibleJumpToViewModel.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-14.
//  Copyright © 2020 Kuang. All rights reserved.
//

import Foundation

class ScrollMapperBibleJumpToViewModel: ObservableObject {
    struct ScopeItem: Identifiable {
        let id: UInt
        let title: String
    }
    let scopes: [ScopeItem] = [ScopeItem(id: 0, title: "Old Testament"), ScopeItem(id: 1, title: "New Testament")]
    @Published var selectedScopeIndex: UInt = 0 {
        didSet {
            if selectedScopeIndex != oldValue {
                expandedBook = nil
                setupListData()
            }
        }
    }
    @Published var listData: [Section] = []
    @Published var expandedBook: ScrollMapperBibleBookInfo.BookInfo? = nil
    
    init() {
        setupListData()
    }
    
    struct Section: Identifiable {
        var id = UUID()
        var title: String
        var books: [Item]
    }
    
    struct Item: Identifiable {
        var id = UUID()
        var book: ScrollMapperBibleBookInfo.BookInfo
    }
    
    private func setupListData() {
        var listData: [Section] = []
        var currentCategory: String = ""
        var items: [Item] = []
        _ = ScrollMapperBibleBookInfo.books.compactMap { (bookInfo) -> ScrollMapperBibleBookInfo.BookInfo? in
            if (bookInfo.otnt == "OT") && (selectedScopeIndex == 1) { return nil }
            if (bookInfo.otnt == "NT") && (selectedScopeIndex == 0) { return nil }
            if bookInfo.category != currentCategory {
                if currentCategory.count > 0 {
                    listData.append(Section(title: currentCategory, books: items))
                    items = []
                }
                currentCategory = bookInfo.category
            }
            items.append(Item(book: bookInfo))
            return nil
        }
        listData.append(Section(title: currentCategory, books: items))
        self.listData = listData
    }
    
    func expand(book: ScrollMapperBibleBookInfo.BookInfo) {
        if expandedBook == book {
            expandedBook = nil
        }
        else {
            expandedBook = book
        }
    }
    
    func jumpTo(book: ScrollMapperBibleBookInfo.BookInfo, chapter: Int) {
        let cid = book.order * 1_000 + chapter
        scrollMapperBiblePublishers.publishCurrentChapter(cid: cid)
    }
}
