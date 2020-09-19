//
//  ScrollMapperBibleCopyrightAndLicenseView.swift
//  ScrollMapperBible
//
//  Created by John K on 2020-09-19.
//  Copyright © 2020 Kuang. All rights reserved.
//

import SwiftUI

struct ScrollMapperBibleCopyrightAndLicenseView: View {
    @ObservedObject var viewModel = ScrollMapperBibleCopyrightAndLicenseViewModel()

    var body: some View {
        VStack {
            Spacer().frame(height: 20)
            HStack {
                Spacer().frame(width: 8)
                self.bodyView
                Spacer().frame(width: 8)
            }
            Spacer()
        }.navigationBarTitle("Copyright \u{00A9}")
    }
    
    private var bodyView: AnyView {
        var repo: AnyView
        if #available(iOS 14, *), let url = URL(string: viewModel.repo) {
            repo = AnyView(Link(viewModel.repo, destination: url))
        }
        else {
            repo = AnyView(Text(viewModel.repo))
        }
        
        return AnyView(
            VStack{
                HStack {
                    Text("This is an open source project forked from ")
                    Spacer()
                }
                HStack {
                    repo
                    Spacer()
                }
            }
        )
    }
    
}

struct ScrollMapperBibleCopyrightAndLicenseView_Previews: PreviewProvider {
    static var previews: some View {
        ScrollMapperBibleCopyrightAndLicenseView()
    }
}
