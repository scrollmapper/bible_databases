//
//  ScrollMapperBibleSettingsView.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-07.
//  Copyright © 2020 Kuang. All rights reserved.
//

import SwiftUI
import Combine

struct ScrollMapperBibleSettingsView: View {
    @ObservedObject private var viewModel: ScrollMapperBibleSettingsViewModel
    
    init() {
        self.viewModel = ScrollMapperBibleSettingsViewModel()
    }

    var body: some View {
        NavigationView {
            List {
                ForEach(viewModel.listData) { section in
                    Section(header: Text(section.title)) {
                        ForEach(section.items) { item in
                            self.itemView(item: item)
                        }
                    }
                }
            }
            .navigationBarTitle(Text("Settings"), displayMode: .inline)
        }
        .navigationViewStyle(StackNavigationViewStyle()) // to prevent it from showing as split view on iPad
    }
    
    private func itemView(item: ScrollMapperBibleSettingsViewModel.Item) -> AnyView {
        if item.navigationLink {
            let navigationLink = NavigationLink(destination: self.itemDestinationView(item: item)) {
                HStack {
                    self.itemImage(item).frame(width: 20)
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
                    self.itemImage(item).frame(width: 20)
                    VStack(alignment: .leading) {
                        Text(item.title)
                        self.itemDetailView(item)
                    }
                }
            )
            return AnyView(view)
        }
    }
    
    private func itemImage(_ item: ScrollMapperBibleSettingsViewModel.Item) -> Image {
        return Image(imageName: item.imageName ?? JKCSImageName.system(systemName: "list.bullet"))
    }
    
    private func itemDetailView(_ item: ScrollMapperBibleSettingsViewModel.Item) -> Text? {
        guard let detail = item.detail else {
            return nil
        }
        return Text(detail).font(.subheadline).foregroundColor(.gray)
    }
    
    private func itemDestinationView(item: ScrollMapperBibleSettingsViewModel.Item) -> AnyView {
        switch item.title {
        case ScrollMapperBibleSettingsViewModel.itemTranslationsTitle:
            return AnyView(ScrollMapperBibleTranslationsView())
        case ScrollMapperBibleSettingsViewModel.itemCopyrightAndLicenseTitle:
            return AnyView(ScrollMapperBibleCopyrightAndLicenseView())
        default:
            return AnyView(EmptyView())
        }
    }
}

struct ScrollMapperBibleSettingsView_Previews: PreviewProvider {
    static var previews: some View {
        ScrollMapperBibleSettingsView()
    }
}
