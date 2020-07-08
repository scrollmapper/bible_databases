//
//  ScrollMapperBibleSearchViewModel.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-08.
//  Copyright © 2020 Kuang. All rights reserved.
//

import Foundation
import Combine

enum ScrollMapperBibleSearchScope: Int, CaseIterable {
    case All = 0
    case OT = 1
    case NT = 2
    
    var titleFull: String {
        switch self {
        case .All:
            return "All"
        case .OT:
            return "Old Testament"
        case .NT:
            return "New Testament"
        }
    }
    
    var titleAbbreviated: String {
        switch self {
        case .All:
            return "All"
        case .OT:
            return "OT"
        case .NT:
            return "NT"
        }
    }
}

class ScrollMapperBibleSearchViewModel: ScrollMapperBibleViewModelBase {
    override init() {
        super.init()
        
        if let searchScopeInt = UserDefaults.standard.value(forKey: scrollMapperBiblePreferencesKeySearchScope) as? Int,
            let _ = ScrollMapperBibleSearchScope(rawValue: searchScopeInt),
            searchScopeInt != self.selectedScopeInt {
            self.selectedScopeInt = searchScopeInt
        }
        else {
            selectedScopeInt = ScrollMapperBibleSearchScope.All.rawValue
        }
        
        setupListData()
    }
    
    override func translationDidChange() {
        super.translationDidChange()
        
        setupListData()
    }
    
    func searchScopeDidChange() {
        setupListData()
    }
    
    let scopes: [ScrollMapperBibleSearchScope] = ScrollMapperBibleSearchScope.allCases
    @Published var selectedScopeInt: Int = 0 {
        didSet {
            if selectedScopeInt != oldValue {
                setupListData()
                UserDefaults.standard.set(selectedScopeInt, forKey: scrollMapperBiblePreferencesKeySearchScope)
            }
        }
    }
    
    struct Section: Identifiable {
        var id = UUID()
        var title: String
        var items: [Item]
    }
    
    struct Item: Identifiable {
        var id = UUID()
        var title: String
        var detail: String
        var isSeeMorePlaceholder: Bool
    }
    
    @Published var listData: [Section] = []
    
    private func setupListData() {
        var listData: [Section] = []
        self.listData = listData
    }
}
