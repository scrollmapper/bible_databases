//
//  ScrollMapperBibleMainView.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-07.
//  Copyright © 2020 Kuang. All rights reserved.
//

import SwiftUI

struct ScrollMapperBibleMainView: View {
    @ObservedObject var viewModel = ScrollMapperBibleMainViewModel()

    var body: some View {
        TabView {
            ScrollMapperBibleTextView()
                .tabItem {
                    Image(systemName: "book")
                    Text("Bible")
                }
            ScrollMapperBibleSettingsView()
                .tabItem {
                    Image(systemName: "gear")
                    Text("Settings")
                }
        }
    }
    
}

struct ScrollMapperBibleMainView_Previews: PreviewProvider {
    static var previews: some View {
        ScrollMapperBibleMainView()
    }
}
