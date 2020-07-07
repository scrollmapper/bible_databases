//
//  ScrollMapperBiblePublishers.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-07.
//  Copyright © 2020 Kuang. All rights reserved.
//

import Foundation
import Combine

// https://medium.com/flawless-app-stories/problem-solving-with-combine-swift-4751885fda77

let scrollMapperBiblePublishers = ScrollMapperBiblePublishers()

class ScrollMapperBiblePublishers {
    let translationSubject = CurrentValueSubject<ScrollMapperBibleVersion.BibleVersion, Never>(ScrollMapperBibleVersion.BibleVersion.KJV)
    let tranlationPublisher: AnyPublisher<ScrollMapperBibleVersion.BibleVersion, Never>
    
    init() {
        tranlationPublisher = translationSubject.eraseToAnyPublisher()
    }
}
