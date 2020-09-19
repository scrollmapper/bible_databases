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
        Text("Copyright")
    }
    
}

struct ScrollMapperBibleCopyrightAndLicenseView_Previews: PreviewProvider {
    static var previews: some View {
        ScrollMapperBibleCopyrightAndLicenseView()
    }
}
