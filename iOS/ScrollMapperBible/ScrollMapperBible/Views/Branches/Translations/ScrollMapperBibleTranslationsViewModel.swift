//
//  ScrollMapperBibleTranslationsViewModel.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-07.
//  Copyright © 2020 Kuang. All rights reserved.
//

import Foundation
import Combine

class ScrollMapperBibleTranslationsViewModel: ObservableObject {
    lazy var translations: [ScrollMapperBibleVersion.BibleVersionKey] = {
        let translations = ScrollMapperBibleVersion()?.result ?? []
        return translations
    }()
    
    init() {
        setupListData()
    }
    
    struct Section: Identifiable {
        var id = UUID()
        var title: String
        var items: [Item]
    }
    
    struct Item: Identifiable {
        var id = UUID()
        var title: String
        var detail: String? = nil
        var imageName: JKCSImageName? = nil
        var navigationLink: Bool = false
    }
    
    static let itemAlertAndPopTitle = "Alert And Pop"
    
    @Published var listData: [Section] = []
    
    private func setupListData() {
        var items = [Item]()
        _ = translations.map { (bibleVersionKey) in
            items.append(Item(title: bibleVersionKey.abbreviation, detail: bibleVersionKey.version, imageName: .system(systemName: "book")))
        }
        self.listData = [Section(title: "TRANSLATIONS", items: items)]
    }
}
