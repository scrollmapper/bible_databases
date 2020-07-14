//
//  ScrollMapperBibleViewModelBase.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-07.
//  Copyright © 2020 Kuang. All rights reserved.
//

import Foundation
import Combine

class ScrollMapperBibleViewModelBase {
    var translationSubscriber: AnyCancellable? = nil
    var translation: ScrollMapperBibleVersion.BibleVersion = .KJV {
        didSet {
            if translation != oldValue {
                translationDidChange()
            }
        }
    }
    
    init() {
        subscribe()
    }
    
    deinit {
        unsubscribe()
    }
    
    func subscribe() {
        translationSubscriber = scrollMapperBiblePublishers.tranlationPublisher.sink(receiveValue: { [weak self] (translation) in
            if let bibleVersion = ScrollMapperBibleVersion.BibleVersion(rawValue: translation) {
                self?.translation = bibleVersion
            }
        })
    }
    
    func unsubscribe() {
        translationSubscriber?.cancel()
        translationSubscriber = nil
    }
    
    func translationDidChange() {
        // subclass shall override this method
    }
}
