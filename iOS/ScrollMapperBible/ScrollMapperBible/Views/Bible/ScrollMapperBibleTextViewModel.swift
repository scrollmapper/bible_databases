//
//  ScrollMapperBibleTextViewModel.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-07.
//  Copyright © 2020 Kuang. All rights reserved.
//

import Foundation
import Combine

class ScrollMapperBibleTextViewModel: ObservableObject {
    var viewTitle: String
    
    init(viewTitle: String) {
        self.viewTitle = viewTitle
        
        DispatchQueue.main.async {
            self.setupListData()
        }
    }
    
    struct Section: Identifiable {
        var id = UUID()
        var title: String
        var items: [Item]
    }
    
    struct Item: Identifiable {
        var id = UUID()
        var title: String
        var detail: String? = nil
        var imageName: JKCSImageName? = nil
        var navigationLink: Bool = false
    }
    
    static let itemImmediatePushTitle = "Immediate Push"
    static let itemLoadAndPushTitle = "Load And Push"
    static let itemAlertAndPushTitle = "Alert And Push"
    static let itemAlertAndToggleTitle = "Alert And Toggle"
    static let itemAlertAndPopTitle = "Alert And Pop"
    static let itemVersionTitle = "Version" // Unclickable
    
    @Published var listData: [Section] = []
    
    private func setupListData() {
        var listData: [Section] = []
        listData.append(
            Section(title: "GENERAL", items: [
                Item(title: ScrollMapperBibleTextViewModel.itemImmediatePushTitle, imageName: JKCSImageName.system(systemName: "p.circle"), navigationLink: true),
                Item(title: ScrollMapperBibleTextViewModel.itemLoadAndPushTitle, imageName: JKCSImageName.system(systemName: "l.circle")),
                Item(title: ScrollMapperBibleTextViewModel.itemAlertAndPushTitle, imageName: JKCSImageName.system(systemName: "a.circle")),
                Item(title: ScrollMapperBibleTextViewModel.itemAlertAndToggleTitle, imageName: JKCSImageName.system(systemName: "t.circle")),
                Item(title: ScrollMapperBibleTextViewModel.itemAlertAndPopTitle, imageName: JKCSImageName.system(systemName: "b.circle"))
            ])
        )
        listData.append(
            Section(title: "ABOUT", items: [
                Item(title: ScrollMapperBibleTextViewModel.itemVersionTitle, imageName: JKCSImageName.system(systemName: "v.circle")),
            ])
        )
        self.listData = listData
    }
    
    private func version() -> String {
        return "x.y.z (build)"
    }
    
    var somethingLoaded: [String : String]? = nil
    
    private var loadSomethingRequest: AnyCancellable? = nil
    func loadSomething(completionHandler: @escaping (Result<Any?, JKCSError>) -> ()) {
        loadSomethingRequest = self.mockingModelLoadSomething().receive(on: RunLoop.main).sink(receiveCompletion: { (completion) in
            switch completion {
            case .failure(let error):
                completionHandler(Result.failure(error))
            case .finished:
                self.loadSomethingRequest = nil
                completionHandler(Result.success(nil))
            }
        }, receiveValue: { (response) in
            if let result = response["result"] {
                self.somethingLoaded = result
            }
        })
    }
    
    private func mockingModelLoadSomething() -> AnyPublisher<[String : [String : String]], JKCSError> {
        let success = true
        return Future<[String : [String : String]], JKCSError> { promise in
            DispatchQueue.main.asyncAfter(deadline: .now() + .seconds(2)) {
                let result = success ? Result.success(["result" : ["result" : "success", "hint": "To trigger the 'Failed to load' case, go to viewModel's mockingModelLoadSomething method, change 'let success = true' to 'let success = false'"]]) : Result.failure(JKCSError.customError(message: "Server side failure\n\nTo trigger the 'Success' case, go to viewModel's mockingModelLoadSomething method, change 'let success = false' to 'let success = true'"))
                promise(result)
            }
        }.eraseToAnyPublisher()
    }
}
