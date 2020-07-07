//
//  ScrollMapperBibleSettingsViewModel.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-07.
//  Copyright © 2020 Kuang. All rights reserved.
//

import Foundation
import Combine

class ScrollMapperBibleSettingsViewModel: ObservableObject {
    var translationSubscriber: AnyCancellable? = nil
    var translation: ScrollMapperBibleVersion.BibleVersion = .KJV {
        didSet {
            if translation != oldValue {
                setupListData()
            }
        }
    }
    
    init() {
        subscribe()
        
        DispatchQueue.main.async {
            self.setupListData()
        }
    }
    
    deinit {
        unsubscribe()
    }
    
    func subscribe() {
        translationSubscriber = scrollMapperBiblePublishers.tranlationPublisher.sink(receiveValue: { [weak self] (translation) in
            self?.translation = translation
        })
    }
    
    func unsubscribe() {
        translationSubscriber?.cancel()
        translationSubscriber = nil
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
    
    static let itemTranslationsTitle = "Translations"
    
    static let itemAcknowledgementsTitle = "Acknowledgements"
    static let itemCopyrightAndLicenseTitle = "Copyright \u{00A9} And License"
    static let itemVersionTitle = "Version" // Unclickable
    
    @Published var listData: [Section] = []
    
    private func setupListData() {
        var sfSymbolAcknowledgements: String
        if #available(iOS 14.0, *) {
            sfSymbolAcknowledgements = "lasso.sparkles"
        }
        else {
            sfSymbolAcknowledgements = "wand.and.stars"
        }
        
        var listData: [Section] = []
        listData.append(
            Section(title: "PREFERENCES", items: [
                Item(title: ScrollMapperBibleSettingsViewModel.itemTranslationsTitle, detail: translation.rawValue, imageName: JKCSImageName.system(systemName: "t.circle"), navigationLink: true)
            ])
        )
        listData.append(
            Section(title: "ABOUT", items: [
                Item(title: ScrollMapperBibleSettingsViewModel.itemAcknowledgementsTitle, imageName: JKCSImageName.system(systemName: sfSymbolAcknowledgements), navigationLink: true),
                Item(title: ScrollMapperBibleSettingsViewModel.itemCopyrightAndLicenseTitle, imageName: JKCSImageName.system(systemName: "doc.plaintext"), navigationLink: true),
                Item(title: ScrollMapperBibleSettingsViewModel.itemVersionTitle, detail: version(), imageName: JKCSImageName.system(systemName: "arkit")),
            ])
        )
        self.listData = listData
    }
    
    private func version() -> String {
        return "\(Bundle.main.releaseVersionNumber() ?? "0.0") (\(Bundle.main.buildVersionNumber() ?? "0"))"
    }
}
