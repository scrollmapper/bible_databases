//
//  ScrollMapperBiblePublishers.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-07.
//  Copyright © 2020 Kuang. All rights reserved.
//

import Foundation
import Combine

fileprivate let scrollMapperBibleSceneStorageKeyTranslation = "83EB5B53-6982-44A9-A382-DA08AA48ABF1"
fileprivate let scrollMapperBibleSceneStorageKeyCurrentChapterCid = "ED1BD62A-8E2F-4626-9D17-22E78298FE69"

// https://medium.com/flawless-app-stories/problem-solving-with-combine-swift-4751885fda77

let scrollMapperBiblePublishers = ScrollMapperBiblePublishers()

class ScrollMapperBiblePublishers {
    private let translationSubject: CurrentValueSubject<String, Never>
    let tranlationPublisher: AnyPublisher<String, Never>
    
    private let currentChapterCidSubject: CurrentValueSubject<Int, Never>
    let currentChapterCidPublisher: AnyPublisher<Int, Never>
    
    init() {
        let translation: String
        if let savedTranslation = UserDefaults.standard.string(forKey: scrollMapperBibleSceneStorageKeyTranslation) {
            translation = savedTranslation
        }
        else {
            translation = ScrollMapperBibleVersion.BibleVersion.KJV.rawValue
        }
        translationSubject = CurrentValueSubject<String, Never>(translation)
        tranlationPublisher = translationSubject.eraseToAnyPublisher()
        
        let cid = UserDefaults.standard.integer(forKey: scrollMapperBibleSceneStorageKeyCurrentChapterCid)
        currentChapterCidSubject = CurrentValueSubject<Int, Never>(cid)
        currentChapterCidPublisher = currentChapterCidSubject.eraseToAnyPublisher()
    }
    
    func publishTranslation(translation: String) {
        UserDefaults.standard.set(translation, forKey: scrollMapperBibleSceneStorageKeyTranslation)
        translationSubject.send(translation)
    }
    
    func publishCurrentChapter(cid: Int) {
        UserDefaults.standard.set(cid, forKey: scrollMapperBibleSceneStorageKeyCurrentChapterCid)
        currentChapterCidSubject.send(cid)
    }
}
