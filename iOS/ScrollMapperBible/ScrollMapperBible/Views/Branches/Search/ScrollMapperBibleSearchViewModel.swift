//
//  ScrollMapperBibleSearchViewModel.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-08.
//  Copyright © 2020 Kuang. All rights reserved.
//

import Foundation
import Combine

fileprivate let scrollMapperBibleSceneStorageKeySearchScope = "268C5DE5-D9D3-491C-9147-C6217DBF876E"
fileprivate let scrollMapperBibleSceneStorageKeySearchTerm = "8BD9E40A-A2F3-49B7-9F69-2055AA38EC9E"

class ScrollMapperBibleSearchViewModel: ObservableObject {
    @Published var translation: ScrollMapperBibleVersion.BibleVersion = .KJV {
        didSet {
            if translation != oldValue {
                search()
            }
        }
    }
    var translationSubscriber: AnyCancellable? = nil
    
    let scopes: [ScrollMapperBibleSearchScope] = ScrollMapperBibleSearchScope.allCases
    @Published var selectedScopeInt: Int = 0 {
        didSet {
            var rawValue = selectedScopeInt + 3
            if rawValue > 3 {
                rawValue -= 3
            }
            selectedScope = ScrollMapperBibleSearchScope(rawValue: rawValue)!
            if selectedScopeInt != oldValue {
                UserDefaults.standard.set(selectedScopeInt, forKey: scrollMapperBibleSceneStorageKeySearchScope)
            }
        }
    }
    private var selectedScope: ScrollMapperBibleSearchScope = .All {
        didSet {
            if selectedScope != oldValue {
                expandedBook = nil
                search()
            }
        }
    }
    
    @Published var expandedBook: ScrollMapperBibleBookInfo.BookInfo? = nil
    
    @Published var searchTerm: String = "" {
        didSet {
            search()
            if searchTerm != oldValue {
                UserDefaults.standard.set(searchTerm, forKey: scrollMapperBibleSceneStorageKeySearchTerm)
            }
        }
    }
    
    init() {
        subscribe()
        
        if let searchScopeInt = UserDefaults.standard.value(forKey: scrollMapperBibleSceneStorageKeySearchScope) as? Int,
            let _ = ScrollMapperBibleSearchScope(rawValue: searchScopeInt),
            searchScopeInt != self.selectedScopeInt {
            self.selectedScopeInt = searchScopeInt
        }
        else {
            selectedScopeInt = ScrollMapperBibleSearchScope.All.rawValue
        }
        
        DispatchQueue.main.async {
            if let searchTerm = UserDefaults.standard.string(forKey: scrollMapperBibleSceneStorageKeySearchTerm) {
                self.searchTerm = searchTerm
            }
        }
    }
    
    deinit {
        unsubscribe()
    }
    
    func subscribe() {
        translationSubscriber = scrollMapperBiblePublishers.tranlationPublisher.sink(receiveValue: { (translation) in
            self.translation = ScrollMapperBibleVersion.BibleVersion(rawValue: translation) ?? ScrollMapperBibleVersion.BibleVersion.KJV
        })
    }
    
    func unsubscribe() {
        translationSubscriber?.cancel()
        translationSubscriber = nil
    }
    
    struct Section: Identifiable {
        var id = UUID()
        var book: ScrollMapperBibleBookInfo.BookInfo
        var items: [Item]
    }
    
    struct Item: Identifiable {
        var id = UUID()
        var verse: ScrollMapperBibleText.BibleText
    }
    
    @Published var listData: [Section] = []
    
    private func search() {
        let term = searchTerm.trimmingCharacters(in: .whitespaces)
        guard term.count > 0 else {
            listData = []
            return
        }
        guard let result = ScrollMapperBibleText(version: translation, searchBy: term, in: selectedScope)?.result else { return }
        var currentBook: ScrollMapperBibleBookInfo.BookInfo? = nil
        var items: [Item] = []
        var list: [Section] = []
        for verse in result {
            if verse.b != currentBook?.order {
                if currentBook != nil {
                    list.append(Section(book: currentBook!, items: items))
                    items = []
                }
                currentBook = ScrollMapperBibleBookInfo.BookInfo(order: verse.b)
            }
            items.append(Item(verse: verse))
        }
        if let currentBook = currentBook {
            list.append(Section(book: currentBook, items: items))
        }
        self.listData = list
    }
    
    func jumpTo(item: Item) {
        let cid = item.verse.b * 1_000 + item.verse.c
        scrollMapperBiblePublishers.publishCurrentChapter(cid: cid)
    }
}
