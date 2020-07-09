//
//  ScrollMapperBibleSearchViewModel.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-08.
//  Copyright © 2020 Kuang. All rights reserved.
//

import Foundation
import Combine

let scrollMapperBibleSceneStorageKeySearchScope = "268C5DE5-D9D3-491C-9147-C6217DBF876E"

class ScrollMapperBibleSearchViewModel: ScrollMapperBibleViewModelBase {
    override init() {
        super.init()
        
        if let searchScopeInt = UserDefaults.standard.value(forKey: scrollMapperBibleSceneStorageKeySearchScope) as? Int,
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
                UserDefaults.standard.set(selectedScopeInt, forKey: scrollMapperBibleSceneStorageKeySearchScope)
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
