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
                
                UIKitSearchBar(text: $searchTerm, placeholder: "Input some keywords and search").onSearchButtonClicked {
                    self.searchButtonClicked()
                }
                
                List {
                    ForEach(viewModel.listData) { section in
                        Section(header: Text(section.title)) {
                            ForEach(section.items) { item in
                                self.itemView(item: item)
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
    
    private func itemView(item: ScrollMapperBibleSearchViewModel.Item) -> AnyView {
        return AnyView(
            VStack(alignment: .leading) {
                Text(item.title)
                Text(item.detail ?? "").font(.subheadline).foregroundColor(.gray)
            }
            .gesture(TapGesture().onEnded({
                self.itemTapped(item)
            }))
        )
    }
    
    private func itemTapped(_ item: ScrollMapperBibleSearchViewModel.Item) {
        switch item.title {
        default:
            print("*** \(item.title) tapped")
            return
        }
    }
    
    private func searchButtonClicked() {
        print("*** Searching for ", searchTerm, "...")
    }
}

struct ScrollMapperBibleSearchView_Previews: PreviewProvider {
    static var previews: some View {
        ScrollMapperBibleSearchView()
    }
}
