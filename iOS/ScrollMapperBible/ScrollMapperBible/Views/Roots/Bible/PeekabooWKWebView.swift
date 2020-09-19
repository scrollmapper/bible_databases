//
//  PeekabooWKWebView.swift
//
//  Created by Zhengqian Kuang on 2020-06-26.
//  Copyright © 2019 Zhengqian Kuang. All rights reserved.
//

import SwiftUI
import UIKit
import WebKit
import Combine

enum PeekabooWKWebViewEvent: String {
    case wordClickHandler
}

// A bug in enum with associated value: https://github.com/textmate/swift.tmbundle/issues/35
enum PeekabooWKWebViewLoad {
    case url(_ url: URL)
    case urlString(_ urlString: String)
    case file(_ resourceName: String, _ resourceExt: String)
    case htmlString(htmlString: String)
}

typealias PeekabooWKWebViewMessage = [String : Any]

protocol PeekabooWKWebViewControllerDelegate {
    func reload(htmlString: String)
}

struct PeekabooWKWebView: UIViewRepresentable {
    private var viewModel: ScrollMapperBibleTextViewModel
    private let postMessageHandlers: [PeekabooWKWebViewEvent]?
    private let webView: WKWebView
    private var viewController = PeekabooWKWebViewController()
    
    class PeekabooWKWebViewController: UIViewController, WKScriptMessageHandler {
        var delegate: PeekabooWKWebViewControllerDelegate? = nil
        var viewModel: ScrollMapperBibleTextViewModel? = nil {
            didSet {
                currentChapterUpdatedSubscriber?.cancel()
                currentChapterUpdatedSubscriber = viewModel?.currentChapterUpdatedPublisher.sink(receiveValue: { (htmlString) in
                    self.delegate?.reload(htmlString: htmlString)
                })
            }
        }
        var currentChapterUpdatedSubscriber: AnyCancellable? = nil
        
        var onClickDelegate: ((PeekabooWKWebViewMessage) -> ())? = nil
        var onSwipeDelegate: ((UISwipeGestureRecognizer.Direction) -> ())? = nil
        
        deinit {
            currentChapterUpdatedSubscriber?.cancel()
            currentChapterUpdatedSubscriber = nil
        }
        
        @objc func didSwipe(swipeGestureRecognizer: UISwipeGestureRecognizer) {
            onSwipeDelegate?(swipeGestureRecognizer.direction)
        }
        
        func userContentController(_ userContentController: WKUserContentController, didReceive message: WKScriptMessage) {
            if message.name == PeekabooWKWebViewEvent.wordClickHandler.rawValue {
                guard
                    let body = message.body as? [String: Any]
                else {
                    return
                }
                onClickDelegate?(body)
            }
        }
    }

    init(viewModel: ScrollMapperBibleTextViewModel, postMessageHandlers: [PeekabooWKWebViewEvent]? = nil) {
        self.viewModel = viewModel
        
        self.postMessageHandlers = postMessageHandlers
        let config = WKWebViewConfiguration()
        if let postMessagehandlers = self.postMessageHandlers {
            for postMessagehandler in postMessagehandlers {
                config.userContentController.add(viewController, name: postMessagehandler.rawValue)
            }
        }
        webView = WKWebView(frame: .zero, configuration: config)
        webView.isOpaque = false
        let leftSwipeGestureRecognizer = UISwipeGestureRecognizer(target: viewController, action: #selector(viewController.didSwipe(swipeGestureRecognizer:)))
        leftSwipeGestureRecognizer.direction = .left
        webView.addGestureRecognizer(leftSwipeGestureRecognizer)
        let rightSwipeGestureRecognizer = UISwipeGestureRecognizer(target: viewController, action: #selector(viewController.didSwipe(swipeGestureRecognizer:)))
        rightSwipeGestureRecognizer.direction = .right
        webView.addGestureRecognizer(rightSwipeGestureRecognizer)
        webView.isUserInteractionEnabled = true
        
        viewController.delegate = self
        viewController.viewModel = viewModel
    }
    
    /// UIViewRepresentable

    func makeUIView(context: UIViewRepresentableContext<PeekabooWKWebView>) -> WKWebView {
        webView.navigationDelegate = context.coordinator
        webView.loadHTMLString(viewModel.currentChapterHTMLString, baseURL: nil)
        return webView
    }

    func updateUIView(_ uiView: WKWebView, context: UIViewRepresentableContext<PeekabooWKWebView>) {
        
    }
    
    class Coordinator: NSObject, WKNavigationDelegate {
        private var viewModel: ScrollMapperBibleTextViewModel

        init(_ viewModel: ScrollMapperBibleTextViewModel) {
            self.viewModel = viewModel
        }

        func webView(_ webView: WKWebView, didFinish navigation: WKNavigation!) {
            self.viewModel.webViewDidFinishNavigation()
        }
    }
    
    func makeCoordinator() -> Self.Coordinator {
        Coordinator(viewModel)
    }
}

extension PeekabooWKWebView {
    @discardableResult func evaluateJavaScript(javaScriptString: String, completionHandler: ((Any?, Error?) -> Void)?) -> Self {
        webView.evaluateJavaScript(javaScriptString) { (result, error) in
            if let completionHandler = completionHandler {
                completionHandler(result, error)
            }
        }
        return self

        // e.g.
        // if you had a page that contained <div id="username">@twostraws</div>,
        // you would use "document.getElementById('username').innerText" as javaScriptString
        // and the result would be "@twostraws"
    }
}

extension PeekabooWKWebView { // modifiers
    @discardableResult func onClick(delegate: @escaping (PeekabooWKWebViewMessage) -> ()) -> Self {
        viewController.onClickDelegate = delegate
        return self
    }

    @discardableResult func onSwipe(delegate: @escaping (UISwipeGestureRecognizer.Direction) -> ()) -> Self {
        viewController.onSwipeDelegate = delegate
        return self
    }
}

extension PeekabooWKWebView: PeekabooWKWebViewControllerDelegate {
    func reload(htmlString: String) {
        webView.loadHTMLString(htmlString, baseURL: nil)
    }
}
