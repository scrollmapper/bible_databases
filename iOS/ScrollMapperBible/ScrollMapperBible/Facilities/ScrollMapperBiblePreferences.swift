//
//  ScrollMapperBiblePreferences.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-07.
//  Copyright © 2020 Kuang. All rights reserved.
//

import Foundation

fileprivate let scrollMapperBiblePreferencesKeyTranslation = "83EB5B53-6982-44A9-A382-DA08AA48ABF1"

// https://www.hackingwithswift.com/quick-start/swiftui/how-to-use-environmentobject-to-share-data-between-views

class ScrollMapperBiblePreferences: ObservableObject {
    @Published var translation: ScrollMapperBibleVersion.BibleVersion = .KJV {
        didSet {
            if translation != oldValue {
                scrollMapperBiblePublishers.translationSubject.send(translation)
                UserDefaults.standard.set(translation.rawValue, forKey: scrollMapperBiblePreferencesKeyTranslation)
            }
        }
    }
    
    init() {
        if let translationString = UserDefaults.standard.value(forKey: scrollMapperBiblePreferencesKeyTranslation) as? String,
            let translation = ScrollMapperBibleVersion.BibleVersion(rawValue: translationString),
            translation != self.translation {
            self.translation = translation
        }
    }
}
