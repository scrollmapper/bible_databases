//
//  ScrollMapperBibleSearchView.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-08.
//  Copyright © 2020 Kuang. All rights reserved.
//

import SwiftUI
import Combine

struct ScrollMapperBibleSearchView: View {
    @ObservedObject private var viewModel: ScrollMapperBibleSearchViewModel
    @Environment(\.presentationMode) var presentationMode
    @State private var showActivityIndicator = false
    @State private var searchTerm = ""
    
    init() {
        self.viewModel = ScrollMapperBibleSearchViewModel()
        
//        UISegmentedControl.appearance().selectedSegmentTintColor = .blue
//        UISegmentedControl.appearance().setTitleTextAttributes([.foregroundColor: UIColor.white], for: .selected)
//        UISegmentedControl.appearance().setTitleTextAttributes([.foregroundColor: UIColor.blue], for: .normal)
        UISegmentedControl.appearance().setTitleTextAttributes([.foregroundColor: UIColor.systemBlue], for: .selected)
    }

    var body: some View {
        ZStack {
            VStack {
                // https://www.swiftkickmobile.com/creating-a-segmented-control-in-swiftui/?utm_source=rss&utm_medium=rss&utm_campaign=creating-a-segmented-control-in-swiftui
                Picker(selection: $viewModel.selectedScopeInt, label: Text("Search Scope")) {
                    ForEach(viewModel.scopes, id: \.rawValue) { scope in
                        Text("\(scope.titleFull)")
                    }
                }
                .pickerStyle(SegmentedPickerStyle())
                
                JKCSSearchBarUIViewRepresentable(binding: $viewModel.searchTerm)
                .placeholder("Input some keywords and search")
                .onSearchButtonClicked { (searchText) in
                    //
                }
                
                List {
                    ForEach(viewModel.listData) { section in
                        Section(header: Text(section.book.title_short)) {
                            ForEach(0..<self.numberOfRowsInSection(section), id: \.self) { (i) in
                                self.itemView(book: section.book, item: section.items[i], index: i)
                            }
                        }
                    }
                }
            }
            
            if showActivityIndicator {
                JKCSActivityIndicatorView().startAnimating()
            }
            else {
                JKCSActivityIndicatorView().stopAnimating()
            }
        }
        .navigationBarTitle(Text("Search"), displayMode: .inline)
    }
    
    private func numberOfRowsInSection(_ section: ScrollMapperBibleSearchViewModel.Section) -> Int {
        if section.book == viewModel.expandedBook {
            return section.items.count
        }
        return min(3, section.items.count)
    }
    
    private func itemView(book: ScrollMapperBibleBookInfo.BookInfo, item: ScrollMapperBibleSearchViewModel.Item, index: Int) -> AnyView {
        if (book == viewModel.expandedBook) || (index < 2) {
            return AnyView(
                Text("\(item.verse.c):\(item.verse.v) \(item.verse.t)")
                    .onTapGesture {
                        self.viewModel.jumpTo(item: item)
                        self.presentationMode.wrappedValue.dismiss()
                }
            )
        }
        return AnyView(
            Text("See more...")
                .onTapGesture {
                    self.viewModel.expandedBook = book
            }
        )
    }
    
    private func itemTapped(_ item: ScrollMapperBibleSearchViewModel.Item) {
        
    }
}

struct ScrollMapperBibleSearchView_Previews: PreviewProvider {
    static var previews: some View {
        ScrollMapperBibleSearchView()
    }
}
