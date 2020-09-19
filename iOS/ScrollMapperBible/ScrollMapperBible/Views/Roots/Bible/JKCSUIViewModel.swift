//
//  JKCSUIViewModel.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-08-01.
//  Copyright © 2020 Kuang. All rights reserved.
//

import UIKit
import Combine

let isPad = UIDevice.current.isPad
let isPhone = UIDevice.current.isPhone

class JKCSUIViewModel {
    var isPortrait: Bool {
        UIDevice.current.isPortrait
    }
    var isLandscape: Bool {
        UIDevice.current.isLandscape
    }
    
    var interfaceOrientationDidChangeDelegate: ((UIInterfaceOrientation) -> ())? = nil
    var interfaceOrientationSubscriber: AnyCancellable? = nil
    
    init() {
        interfaceOrientationSubscriber = scrollMapperBibleUIPublishers.interfaceOrientationPublisher.sink(receiveValue: { (orientation) in
            self.interfaceOrientationDidChangeDelegate?(orientation)
        })
    }
    
    @discardableResult
    func onInterfaceOrientationChange(_ closure: @escaping (UIInterfaceOrientation) -> ()) -> Self {
        interfaceOrientationDidChangeDelegate = closure
        return self
    }
    
    deinit {
        interfaceOrientationSubscriber?.cancel()
    }
}
