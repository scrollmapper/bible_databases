//
//  ScrollMapperBibleTextView.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-07.
//  Copyright © 2020 Kuang. All rights reserved.
//

import SwiftUI
import Combine

struct ScrollMapperBibleTextView: View, ScrollMapperBibleTextViewAlertActionDelegate {
    @ObservedObject private var viewModel: ScrollMapperBibleTextViewModel
    @EnvironmentObject var scrollMapperBiblePreferences: ScrollMapperBiblePreferences
    @State private var showActivityIndicator = false
    @State private var showAlert = false
    @State private var pushJumpToView = false
    @State private var pushSearchView = false
    @State private var alert: ScrollMapperBibleTextViewAlert = .none {
        didSet {
            switch alert {
            case .failed(_):
                showAlert = true
            default:
                break
            }
        }
    }
    func onAlertActionConfirm(alert: ScrollMapperBibleTextViewAlert) {
        switch alert {
        default:
            break
        }
    }
    
    init(viewTitle: String) {
        self.viewModel = ScrollMapperBibleTextViewModel()
    }

    var body: some View {
        NavigationView {
            ZStack {
                List {
                    ForEach(viewModel.listData) { section in
                        Section(header: self.sectionHeaderView(section: section)) {
                            ForEach(section.items) { item in
                                self.itemView(item: item)
                            }
                        }
                    }
                }
                .alert(isPresented: $showAlert) {
                    return alert.alert(delegate: self)
                }
                
                NavigationLink(destination: EmptyView(), isActive: self.$pushJumpToView) {
                    Text("")
                }.hidden()
                
                NavigationLink(destination: ScrollMapperBibleSearchView(), isActive: self.$pushSearchView) {
                    Text("")
                }.hidden()
                
                if showActivityIndicator {
                    JKCSActivityIndicatorView().startAnimating()
                }
                else {
                    JKCSActivityIndicatorView().stopAnimating()
                }
            }
            .navigationBarTitle(Text("\(scrollMapperBiblePreferences.translation.rawValue)"), displayMode: .inline)
            .navigationBarItems(trailing: navigationBarTrailing())
        }
        .navigationViewStyle(StackNavigationViewStyle()) // to prevent it from showing as split view on iPad
    }
    
    private func navigationBarTrailing() -> some View {
        Button(action: {
            self.pushSearchView = true
        }) {
            Image(systemName: "magnifyingglass")
        }
    }
    
    private func sectionHeaderView(section: ScrollMapperBibleTextViewModel.Section) -> AnyView {
        return AnyView(
            HStack {
                Spacer()
                Text(section.title)
                Spacer()
            }
        )
    }
    
    private func itemView(item: ScrollMapperBibleTextViewModel.Item) -> AnyView {
        return AnyView(
            Text(item.text)
        )
    }
}

struct ScrollMapperBibleTextView_Previews: PreviewProvider {
    static var previews: some View {
        ScrollMapperBibleTextView(viewTitle: "Bible")
    }
}

// MARK: - Alert

protocol ScrollMapperBibleTextViewAlertActionDelegate {
    func onAlertActionConfirm(alert: ScrollMapperBibleTextViewAlert)
}

enum ScrollMapperBibleTextViewAlert {
    case none
    case failed(error: JKCSError)
    
    func alert<T: ScrollMapperBibleTextViewAlertActionDelegate>(delegate: T) -> Alert {
        var title: String
        var message: String
        switch self {
        case .none:
            return Alert(title: Text("NONE"))
        case .failed(let error):
            title = "FAILED"
            message = error.message
            return Alert(title: Text(title), message: Text(message), dismissButton: .default(Text("OK")))
        }
    }
}
