//
//  JKCSUIShared.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-07.
//  Copyright © 2020 Kuang. All rights reserved.
//

import SwiftUI

public struct JKCSActivityIndicatorView: UIViewRepresentable {
    public var uiActivityIndicatorView = UIActivityIndicatorView()
    
    public init() {
        
    }
    
    public func makeUIView(context: Context) -> UIActivityIndicatorView {
        return uiActivityIndicatorView
    }
    
    public func updateUIView(_ uiView: UIActivityIndicatorView, context: Context) {
        
    }
    
    public func startAnimating() -> Self {
        uiActivityIndicatorView.startAnimating()
        return self
    }
    
    public func stopAnimating() -> Self {
        uiActivityIndicatorView.stopAnimating()
        return self
    }
}

public extension Image {
    init(imageName: JKCSImageName) {
        switch imageName {
        case .inBundle(let name):
            self.init(name)
        case .system(let systemName):
            self.init(systemName: systemName)
        }
    }
}

public enum JKCSImageName {
    case inBundle(imageName: String)
    case system(systemName: String)
}
