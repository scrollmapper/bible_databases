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
    
    @State var reloadChapter = false
    
    init() {
        viewModel = ScrollMapperBibleTextViewModel()
    }

    var body: some View {
        NavigationView {
            ZStack {
                PeekabooWKWebView(viewModel: viewModel, postMessageHandlers: [.wordClickHandler])
                .onClick(delegate: { (message) in
                    print("*** onClick: \(message)")
                    if let clicked = message["clicked"] as? [String : Any] {
                        if let word = clicked["word"] as? String {
                            print("*** word clicked: \(word)")
                        }
                        if let sentence = clicked["sentence"] as? String {
                            print("*** sentence clicked: \(sentence)")
                        }
                    }
                })
                .onSwipe(delegate: { (direction) in
                    if direction == .left {
                        self.viewModel.gotoNextChapter()
                    }
                    else if direction == .right {
                        self.viewModel.gotoPreviousChapter()
                    }
                })
                .padding()
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
            .navigationBarTitle(Text("\(viewModel.translation.rawValue)"), displayMode: .inline)
            .navigationBarItems(leading: navigationBarLeading(), trailing: navigationBarTrailing())
        }
        .navigationViewStyle(StackNavigationViewStyle()) // to prevent it from showing as split view on iPad
    }
    
    private func navigationBarLeading() -> some View {
        Button(action: {
            self.pushJumpToView = true
        }) {
            Image(systemName: "map")
        }
    }
    
    private func navigationBarTrailing() -> some View {
        Button(action: {
            self.pushSearchView = true
        }) {
            Image(systemName: "magnifyingglass")
        }
    }
}

struct ScrollMapperBibleTextView_Previews: PreviewProvider {
    static var previews: some View {
        ScrollMapperBibleTextView()
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
