//
//  ScrollMapperBibleTranslationsView.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-07.
//  Copyright © 2020 Kuang. All rights reserved.
//

import SwiftUI
import Combine

struct ScrollMapperBibleTranslationsView: View, ScrollMapperBibleTranslationsViewAlertActionDelegate {
    @ObservedObject private var viewModel: ScrollMapperBibleTranslationsViewModel
    @Environment(\.presentationMode) var presentationMode
    @State private var showActivityIndicator = false
    @State private var showAlert = false
    @State private var alert: ScrollMapperBibleTranslationsViewAlert = .none {
        didSet {
            switch alert {
            case .confirmSwitch:
                showAlert = true
            default:
                break
            }
        }
    }
    func onAlertActionConfirm(alert: ScrollMapperBibleTranslationsViewAlert) {
        switch alert {
        case .confirmSwitch(let target):
            self.showActivityIndicator = true
            DispatchQueue.main.async {
                self.viewModel.switchTranslation(to: target)
                self.presentationMode.wrappedValue.dismiss()
                self.showActivityIndicator = false
            }
        default:
            break
        }
    }
    
    init() {
        self.viewModel = ScrollMapperBibleTranslationsViewModel()
    }

    var body: some View {
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
            .disabled(showActivityIndicator)
            .alert(isPresented: $showAlert) {
                return alert.alert(delegate: self)
            }
            
            if showActivityIndicator {
                JKCSActivityIndicatorView().startAnimating()
            }
            else {
                JKCSActivityIndicatorView().stopAnimating()
            }
        }
        .navigationBarTitle(Text("Translations"), displayMode: .inline)
    }
    
    private func itemView(item: ScrollMapperBibleTranslationsViewModel.Item) -> AnyView {
        let view = AnyView(
            HStack {
                self.itemImage(item)
                VStack(alignment: .leading) {
                    Text(item.title)
                    self.itemDetailView(item)
                }
                if (item.title == viewModel.translation) {
                    Spacer()
                    Image(systemName: "checkmark")
                }
            }
        )
        .gesture(TapGesture().onEnded({ (_) in
            self.itemTapped(item)
        }))
        return AnyView(view)
    }
    
    private func itemImage(_ item: ScrollMapperBibleTranslationsViewModel.Item) -> Image {
        return Image(imageName: item.imageName ?? JKCSImageName.system(systemName: "list.bullet"))
    }
    
    private func itemDetailView(_ item: ScrollMapperBibleTranslationsViewModel.Item) -> Text? {
        guard let detail = item.detail else {
            return nil
        }
        return Text(detail).font(.subheadline).foregroundColor(.gray)
    }
    
    private func itemTapped(_ item: ScrollMapperBibleTranslationsViewModel.Item) {
        if (item.title == viewModel.translation) {
            return
        }
        alert = .confirmSwitch(target: item.title)
    }
}

struct ScrollMapperBibleTranslationsView_Previews: PreviewProvider {
    static var previews: some View {
        ScrollMapperBibleTranslationsView()
    }
}

// MARK: - Alert

protocol ScrollMapperBibleTranslationsViewAlertActionDelegate {
    func onAlertActionConfirm(alert: ScrollMapperBibleTranslationsViewAlert)
}

enum ScrollMapperBibleTranslationsViewAlert {
    case none
    case confirmSwitch(target: String)
    
    func alert<T: ScrollMapperBibleTranslationsViewAlertActionDelegate>(delegate: T) -> Alert {
        var title: String
        var message: String
        switch self {
        case .none:
            return Alert(title: Text("NONE"))
        case .confirmSwitch(let target):
            title = "SWITCH TRANSLATION"
            message = "Do you want to switch to \(target)?"
        }
        return Alert(title: Text(title), message: Text(message), primaryButton: .destructive(Text("Yes")) {
            delegate.onAlertActionConfirm(alert: self)
        }, secondaryButton: .cancel())
    }
}
