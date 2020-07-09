//
//  UIKitSearchBar.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-08.
//  Copyright © 2020 Kuang. All rights reserved.
//

import SwiftUI

// https://github.com/axelhodler/swiftuisearchbar

struct UIKitSearchBar: UIViewRepresentable {

    @Binding var text: String
    var placeholder: String
    let keyboardType: UIKeyboardType = .default
    var delegateWrapper = DelegateWrapper()
    var coordinatorWrapper = CoordinatorWrapper()
    
    class DelegateWrapper {
        var onSearchButtonClickedDelegate: (() -> ())? = nil
    }
    
    class CoordinatorWrapper {
        var coordinator: Coordinator? = nil
    }

    class Coordinator: NSObject, UISearchBarDelegate {

        @Binding var text: String
        var onSearchButtonClickedDelegate: (() -> ())? = nil

        init(text: Binding<String>) {
            _text = text
        }
        
        func searchBarTextDidBeginEditing(_ searchBar: UISearchBar) {
            searchBar.setShowsCancelButton(true, animated: true)
        }
        
        func searchBarTextDidEndEditing(_ searchBar: UISearchBar) {
            searchBar.setShowsCancelButton(false, animated: false)
        }

        func searchBar(_ searchBar: UISearchBar, textDidChange searchText: String) {
            text = searchText
        }
        
        func searchBarSearchButtonClicked(_ searchBar: UISearchBar) {
            searchBar.resignFirstResponder()
            if let delegate = onSearchButtonClickedDelegate {
                delegate()
            }
        }
        
        func searchBarCancelButtonClicked(_ searchBar: UISearchBar) {
            searchBar.resignFirstResponder()
        }
    }

    func makeCoordinator() -> UIKitSearchBar.Coordinator {
        if coordinatorWrapper.coordinator == nil {
            coordinatorWrapper.coordinator = Coordinator(text: $text)
            coordinatorWrapper.coordinator?.onSearchButtonClickedDelegate = delegateWrapper.onSearchButtonClickedDelegate
        }
        return coordinatorWrapper.coordinator!
    }

    func makeUIView(context: UIViewRepresentableContext<UIKitSearchBar>) -> UISearchBar {
        let searchBar = UISearchBar(frame: .zero)
        searchBar.delegate = context.coordinator
        searchBar.placeholder = placeholder
        searchBar.searchBarStyle = .minimal
        searchBar.autocapitalizationType = .none
        searchBar.keyboardType = keyboardType
        return searchBar
    }

    func updateUIView(_ uiView: UISearchBar, context: UIViewRepresentableContext<UIKitSearchBar>) {
        uiView.text = text
    }
    
    func onSearchButtonClicked(delegate: @escaping () -> ()) -> UIKitSearchBar {
        delegateWrapper.onSearchButtonClickedDelegate = delegate
        coordinatorWrapper.coordinator?.onSearchButtonClickedDelegate = delegate
        return self
    }
}

struct UIKitSearchBarExampleView {
    let cars = ["Subaru WRX", "Tesla Model 3", "Porsche 911", "Renault Zoe", "DeLorean", "Mitsubishi Lancer", "Audi RS6"]
    @State private var searchText : String = ""

    var body: some View {
        NavigationView {
            VStack {
                UIKitSearchBar(text: $searchText, placeholder: "Search cars")
                List {
                    ForEach(self.cars.filter {
                        self.searchText.isEmpty ? true : $0.lowercased().contains(self.searchText.lowercased())
                    }, id: \.self) { car in
                        Text(car)
                    }
                }.navigationBarTitle(Text("Cars"))
            }
        }
    }
}
