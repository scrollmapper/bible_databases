//
//  ScrollMapperBibleTextView.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-07.
//  Copyright © 2020 Kuang. All rights reserved.
//

import SwiftUI
import Combine

struct ScrollMapperBibleTextView: View {
    @ObservedObject private var viewModel: ScrollMapperBibleTextViewModel
    @Environment(\.colorScheme) var colorScheme
    @State private var showActivityIndicator = false
    @State private var pushJumpToView = false
    @State private var pushSearchView = false
    
    @State var reloadChapter = false
    
    init() {
        viewModel = ScrollMapperBibleTextViewModel()
    }

    var body: some View {
        NavigationView {
            ZStack {
                HStack {
                    Spacer().frame(width: 8)
                    createWebView()
                    Spacer().frame(width: 8)
                }
                .alert(isPresented: self.$viewModel.showAlert, content: { () -> Alert in
                    switch self.viewModel.alert {
                    case .verseNumber(let book, let chapter, let verse):
                        return Alert(title: Text("CROSS REFERENCE"), message: Text("Show cross reference for \(book.bookInfo()?.abbreviation ?? "") \(chapter):\(verse)?"), primaryButton: .destructive(Text("YES"), action: {
                            self.viewModel.retrieveCrossReference(book: book, chapter: chapter, verse: verse)
                        }), secondaryButton: .cancel())
                    default:
                        return Alert(title: Text("TO BE ACCOMPLISHED"), message: nil, dismissButton: .destructive(Text("OK")))
                    }
                })
                
                GeometryReader { (geometryProxy) in
                    self.crossRefView(geometryProxy: geometryProxy)
                }
                
                NavigationLink(destination: ScrollMapperBibleJumpToView(), isActive: self.$pushJumpToView) {
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
    
    private func createWebView() -> PeekabooWKWebView {
        viewModel.colorSchemeDidChange(darkMode: colorScheme == .dark)
        let webView = PeekabooWKWebView(viewModel: viewModel, postMessageHandlers: [.wordClickHandler])
            .onClick(delegate: { (message) in
                DispatchQueue.main.async {
                    self.viewModel.checkClickedMessage(message: message)
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
        return webView
    }
    
    private func crossRefView(geometryProxy: GeometryProxy) -> AnyView? {
        guard viewModel.crossReferenceVid != 0 else {
            return nil
        }
        return AnyView(
            ZStack {
                AnyView(
                    Text("")
                        .frame(width: geometryProxy.size.width, height: geometryProxy.size.height)
                        .background(Color.black)
                        .opacity(0.5)
                        .onTapGesture {
                            DispatchQueue.main.async {
                                self.viewModel.dismissCrossReference()
                            }
                    }
                )
                
                VStack {
                    Spacer()
                    HStack {
                        Spacer()
                        ScrollMapperBibleCrossReferenceView(vid: viewModel.crossReferenceVid).frame(width: geometryProxy.size.width * 0.8, height: geometryProxy.size.height * 0.8)
                        Spacer()
                    }
                    Spacer()
                }
            }
        )
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
