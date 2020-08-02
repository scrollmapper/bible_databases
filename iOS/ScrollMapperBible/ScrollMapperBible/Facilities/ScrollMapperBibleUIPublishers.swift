//
//  ScrollMapperBibleUIPublishers.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-08-01.
//  Copyright © 2020 Kuang. All rights reserved.
//

import UIKit
import Combine

let scrollMapperBibleUIPublishers = ScrollMapperBibleUIPublishers()

class ScrollMapperBibleUIPublishers {
    private let interfaceOrientationSubject: CurrentValueSubject<UIInterfaceOrientation, Never>
    let interfaceOrientationPublisher: AnyPublisher<UIInterfaceOrientation, Never>
    
    init() {
        let interfaceOrientation = UIDevice.current.currentInterfaceOrientation
        interfaceOrientationSubject = CurrentValueSubject<UIInterfaceOrientation, Never>(interfaceOrientation)
        interfaceOrientationPublisher = interfaceOrientationSubject.eraseToAnyPublisher()
        
        NotificationCenter.default.addObserver(self, selector: #selector(interfaceOrientationDidChange(notification:)), name: UIDevice.orientationDidChangeNotification, object: nil)
    }
    
    @objc private func interfaceOrientationDidChange(notification: Notification) {
        let interfaceOrientation = UIDevice.current.currentInterfaceOrientation
        interfaceOrientationSubject.send(interfaceOrientation)
    }
    
    deinit {
        NotificationCenter.default.removeObserver(self)
    }
}
