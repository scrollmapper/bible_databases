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
    private var isNavigationRoot = true
    @State private var showActivityIndicator = false
    @State private var showAlert = false
    @State private var pushLoadedView = false
    @State private var pushConfirmedView = false
    @State private var alert: ScrollMapperBibleTextViewAlert = .none {
        didSet {
            switch alert {
            case .confirmPush, .confirmToggle, .confirmPop, .failedLoading(_):
                showAlert = true
            default:
                break
            }
        }
    }
    func onAlertActionConfirm(alert: ScrollMapperBibleTextViewAlert) {
        switch alert {
        case .confirmPush:
            pushConfirmedView.toggle()
        default:
            break
        }
    }
    
    init(viewTitle: String) {
        self.viewModel = ScrollMapperBibleTextViewModel(viewTitle: viewTitle)
    }

    var body: some View {
        bodyView
    }
    
    private var bodyView: AnyView {
        if self.isNavigationRoot {
            return AnyView(
                NavigationView {
                    baseView
                }
                .navigationViewStyle(StackNavigationViewStyle()) // to prevent it from showing as split view on iPad
            )
        }
        else {
            return baseView
        }
    }
    
    private var baseView: AnyView {
        AnyView(
            ZStack {
                List {
                    ForEach(viewModel.listData) { section in
                        Section(header: Text(section.title)) {
                            ForEach(section.items) { item in
                                self.itemView(item: item)
                            }
                        }
                    }
                }
                .alert(isPresented: $showAlert) {
                    return alert.alert(delegate: self)
                }
                
                NavigationLink(destination: ScrollMapperBibleTextViewChildView2(info: self.viewModel.somethingLoaded), isActive: self.$pushLoadedView) {
                    Text("")
                }.hidden()
                
                NavigationLink(destination: ScrollMapperBibleTextViewChildView1(title: "Confirmed Push Child"), isActive: self.$pushConfirmedView) {
                    Text("")
                }.hidden()
                
                if showActivityIndicator {
                    JKCSActivityIndicatorView().startAnimating()
                }
                else {
                    JKCSActivityIndicatorView().stopAnimating()
                }
            }.navigationBarTitle(Text(viewModel.viewTitle))
        )
    }
    
    private func itemView(item: ScrollMapperBibleTextViewModel.Item) -> AnyView {
        if item.navigationLink {
            let navigationLink = NavigationLink(destination: self.itemDestinationView(item: item)) {
                HStack {
                    self.itemImage(item)
                    VStack(alignment: .leading) {
                        Text(item.title)
                        self.itemDetailView(item)
                    }
                }
            }
            return AnyView(navigationLink)
        }
        else {
            let view = AnyView(
                HStack {
                    self.itemImage(item)
                    VStack(alignment: .leading) {
                        Text(item.title)
                        self.itemDetailView(item)
                    }
                }
            )
            .gesture(TapGesture().onEnded({ (_) in
                self.itemTapped(item)
            }))
            return AnyView(view)
        }
    }
    
    private func itemImage(_ item: ScrollMapperBibleTextViewModel.Item) -> Image {
        return Image(imageName: item.imageName ?? JKCSImageName.system(systemName: "list.bullet"))
    }
    
    private func itemDetailView(_ item: ScrollMapperBibleTextViewModel.Item) -> Text? {
        guard let detail = item.detail else {
            return nil
        }
        return Text(detail).font(.subheadline).foregroundColor(.gray)
    }
    
    private func itemTapped(_ item: ScrollMapperBibleTextViewModel.Item) {
        switch item.title {
        case ScrollMapperBibleTextViewModel.itemLoadAndPushTitle:
            showActivityIndicator = true
            viewModel.loadSomething { (result) in
                self.showActivityIndicator = false
                switch result {
                case .failure(let error):
                    self.alert = .failedLoading(error: error)
                case .success(_):
                    self.pushLoadedView = true
                }
            }
            return
        case ScrollMapperBibleTextViewModel.itemAlertAndPushTitle:
            alert = .confirmPush
        case ScrollMapperBibleTextViewModel.itemAlertAndPopTitle:
            alert = .confirmPop
        default:
            print("*** \(item.title) tapped")
            return
        }
    }
    
    private func itemDestinationView(item: ScrollMapperBibleTextViewModel.Item) -> AnyView {
        switch item.title {
        case ScrollMapperBibleTextViewModel.itemImmediatePushTitle:
            return AnyView(ScrollMapperBibleTextViewChildView1(title: "Immediate Push Child"))
        default:
            return AnyView(EmptyView())
        }
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
    case confirmPush
    case confirmToggle(turnOn: Bool)
    case confirmPop
    case failedLoading(error: JKCSError)
    
    func alert<T: ScrollMapperBibleTextViewAlertActionDelegate>(delegate: T) -> Alert {
        var title: String
        var message: String
        switch self {
        case .none:
            return Alert(title: Text("NONE"))
        case .confirmPush:
            title = "PUSH"
            message = "Proceed?"
        case .confirmToggle(let turnOn):
            title = "TOGGLE"
            message = "Turn it \(turnOn ? "on" : "off")?"
        case .confirmPop:
            title = "BACK"
            message = "Proceed?"
        case .failedLoading(let error):
            title = "FAILED LOADING"
            message = error.message
            return Alert(title: Text(title), message: Text(message), dismissButton: .default(Text("OK")))
        }
        return Alert(title: Text(title), message: Text(message), primaryButton: .destructive(Text("Yes")) {
            delegate.onAlertActionConfirm(alert: self)
        }, secondaryButton: .cancel())
    }
}

// MARK: - Child Views

struct ScrollMapperBibleTextViewChildView1: View {
    var title: String
    var body: some View {
        Text(title)
    }
}

struct ScrollMapperBibleTextViewChildView2: View {
    var info: [String : String]
    var body: some View {
        ScrollView {
            Text(self.stringifyInfo())
        }
    }
    
    init(info: [String : String]? = nil) {
        guard let info = info else {
            self.info = [:]
            return
        }
        self.info = info
    }
    
    private func stringifyInfo() -> String {
        do {
            let data = try JSONSerialization.data(withJSONObject: info, options: .prettyPrinted)
            return String(data: data, encoding: .utf8) ?? "<UTF8 decoding failed>"
        } catch {
            return "<JSON serialization failed>"
        }
    }
}
