//
//  ScrollMapperBibleCrossReferenceViewModel.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-16.
//  Copyright © 2020 Kuang. All rights reserved.
//

import Foundation
import Combine

class ScrollMapperBibleCrossReferenceViewModel: ObservableObject {
    @Published var translation: ScrollMapperBibleVersion.BibleVersion = .KJV {
        didSet {
            if translation != oldValue {
                setupList()
            }
        }
    }
    var translationSubscriber: AnyCancellable? = nil
    
    var vid: Int {
        didSet {
            if vid != oldValue {
                setupList()
            }
        }
    }
    
    init(vid: Int) {
        self.vid = vid
        subscribe()
        DispatchQueue.main.async {
            self.setupList()
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
        let id = UUID()
        let rank: Int
        let sv: Int
        let ev: Int
        let items: [Item]
    }
    
    struct Item: Identifiable {
        let id = UUID()
        let verse: ScrollMapperBibleText.BibleText
    }
    
    @Published var listData: [Section] = []
    
    private func setupList() {
        var list: [Section] = []
        guard let result = ScrollMapperBibleCrossReference(version: translation, vid: vid)?.result else { return }
        for row in result {
            var verses: [Item] = []
            for verse in row.verses {
                verses.append(Item(verse: verse))
            }
            list.append(Section(rank: row.r, sv: row.sv, ev: row.ev, items: verses))
        }
        self.listData = list
    }
}
