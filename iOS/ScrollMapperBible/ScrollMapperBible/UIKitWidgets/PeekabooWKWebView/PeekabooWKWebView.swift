//
//  PeekabooWKWebView.swift
//
//  Created by Zhengqian Kuang on 2020-06-26.
//  Copyright © 2019 Zhengqian Kuang. All rights reserved.
//

import SwiftUI
import UIKit
import WebKit

// For your convenience to add this file to your Swift Package, all the types are declared "public" when necessary.

public enum PeekabooWKWebViewEvent: String {
    case wordClickHandler
}

public enum PeekabooWKWebViewLoad {
    case url(_ url: URL)
    case urlString(_ urlString: String)
    case file(_ resourceName: String, _ resourceExt: String)
    case htmlString(_ htmlString: String)
}

// Ref.
//  https://medium.com/john-lewis-software-engineering/ios-wkwebview-communication-using-javascript-and-swift-ee077e0127eb
//  https://www.hackingwithswift.com/quick-start/swiftui/how-to-wrap-a-custom-uiview-for-swiftui
//  https://www.hackingwithswift.com/articles/112/the-ultimate-guide-to-wkwebview

public struct PeekabooWKWebView: UIViewRepresentable {
    private let postMessageHandlers: [PeekabooWKWebViewEvent]?
    private let viewController = PeekabooWKWebViewController()
    
    public init(postMessageHandlers: [PeekabooWKWebViewEvent]? = nil) {
        self.postMessageHandlers = postMessageHandlers
        viewController.createWKWebView(postMessagehandlers: postMessageHandlers)
    }
    
    public func makeUIView(context: UIViewRepresentableContext<PeekabooWKWebView>) -> WKWebView {
        return viewController.peekableWKWebView(postMessageHandlers: postMessageHandlers)
    }
    
    public func updateUIView(_ uiView: WKWebView, context: UIViewRepresentableContext<PeekabooWKWebView>) {
        //
    }

    private class PeekabooWKWebViewController: UIViewController, WKNavigationDelegate, WKUIDelegate, WKScriptMessageHandler {
        private var wkWebView: WKWebView? = nil
        var delegateWrapper = DelegateWrapper()
        
        class DelegateWrapper {
            var onClickWordDelegate: ((String) -> ())? = nil
        }
        
        @discardableResult func createWKWebView(postMessagehandlers: [PeekabooWKWebViewEvent]? = nil) -> WKWebView {
            let config = WKWebViewConfiguration()
            // config.preferences.javaScriptEnabled = true
            // config.dataDetectorTypes = [.address, .calendarEvent, .flightNumber, .link, .lookupSuggestion, .phoneNumber, .trackingNumber]
            // wkWebView = WKWebView(frame: .zero, configuration: config)
            if let postMessagehandlers = postMessagehandlers {
                for postMessagehandler in postMessagehandlers {
                    config.userContentController.add(self, name: postMessagehandler.rawValue)
                }
            }
            
            wkWebView = WKWebView(frame: .zero, configuration: config)
            wkWebView?.navigationDelegate = self
            wkWebView?.uiDelegate = self
            
            wkWebView?.addObserver(self, forKeyPath: #keyPath(WKWebView.estimatedProgress), options: .new, context: nil)
            wkWebView?.addObserver(self, forKeyPath: #keyPath(WKWebView.title), options: .new, context: nil)
            
            return wkWebView!
        }
        
        func peekableWKWebView(postMessageHandlers: [PeekabooWKWebViewEvent]? = nil) -> WKWebView {
            if wkWebView == nil {
                createWKWebView(postMessagehandlers: postMessageHandlers)
            }
            return wkWebView!
        }
        
        func webView(_ webView: WKWebView, decidePolicyFor navigationAction: WKNavigationAction, decisionHandler: @escaping (WKNavigationActionPolicy) -> Void) {
            // if let url = navigationAction.request.url {
                // check url and
                //   decisionHandler(.cancel)
                // to cancel
            // }
            decisionHandler(.allow)
        }
        
        func userContentController(_ userContentController: WKUserContentController, didReceive message: WKScriptMessage) {
            if message.name == PeekabooWKWebViewEvent.wordClickHandler.rawValue {
                guard
                    let body = message.body as? [String: Any],
                    let param1 = body["param1"] as? String
                else {
                    return
                }
                if let delegate = delegateWrapper.onClickWordDelegate {
                    delegate(param1)
                }
            }
        }
        
//        func webView(_ webView: WKWebView, runJavaScriptAlertPanelWithMessage message: String, initiatedByFrame frame: WKFrameInfo, completionHandler: @escaping () -> Void) {
//            // let ac = UIAlertController(title: "Hey, listen!", message: message, preferredStyle: .alert)
//            // ac.addAction(UIAlertAction(title: "OK", style: .default, handler: nil))
//            // present(ac, animated: true)
//            completionHandler()
//        }
//
//        func webView(_ webView: WKWebView, runJavaScriptConfirmPanelWithMessage message: String, initiatedByFrame frame: WKFrameInfo, completionHandler: @escaping (Bool) -> Void) {
//            // show confirm panel
//            completionHandler(true or false according to the confirm panel)
//        }
//
//        func webView(_ webView: WKWebView, runJavaScriptTextInputPanelWithPrompt prompt: String, defaultText: String?, initiatedByFrame frame: WKFrameInfo, completionHandler: @escaping (String?) -> Void) {
//            // show text input panel
//            completionHandler(text according to the input panel)
//        }
        
        override func observeValue(forKeyPath keyPath: String?, of object: Any?, change: [NSKeyValueChangeKey : Any]?, context: UnsafeMutableRawPointer?) {
            if keyPath == "estimatedProgress" {
                // print(Float(wkWebView!.estimatedProgress))
            }
            else if keyPath == "title" {
                // if let title = wkWebView?.title {
                    //
                // }
            }
        }
    }
}

public extension PeekabooWKWebView {
    @discardableResult func load(source: PeekabooWKWebViewLoad) -> Self {
        switch source {
        case .url(let url):
            let request = URLRequest(url: url)
            viewController.peekableWKWebView().load(request)
        case .urlString(let urlString):
            if let url = URL(string: urlString) {
                let request = URLRequest(url: url)
                viewController.peekableWKWebView().load(request)
            }
        case .file(let resourceName, let resourceExt):
            if let url = Bundle.main.url(forResource: resourceName, withExtension: resourceExt) {
                // url.deletingLastPathComponent() part tells WebKit it can read from the directory that contains help.html – that’s a good place to put any assets such as images, JavaScript, or CSS.
                viewController.peekableWKWebView().loadFileURL(url, allowingReadAccessTo: url.deletingPathExtension())
            }
        case .htmlString(let htmlString):
            viewController.peekableWKWebView().loadHTMLString(htmlString, baseURL: nil)
        }
        return self
    }
    
    @discardableResult func evaluateJavaScript(javaScriptString: String, completionHandler: ((Any?, Error?) -> Void)?) -> Self {
        viewController.peekableWKWebView().evaluateJavaScript(javaScriptString) { (result, error) in
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
    
    func backList() -> [WKBackForwardListItem] {
        return viewController.peekableWKWebView().backForwardList.backList
    }
    
    func forwardList() -> [WKBackForwardListItem] {
        return viewController.peekableWKWebView().backForwardList.forwardList
    }
    
    @discardableResult func handleCookiesExample() -> Self {
        // As an example, this code loops over all cookies, and when it finds one called “authentication” deletes it – all other cookies are just printed out
        viewController.peekableWKWebView().configuration.websiteDataStore.httpCookieStore.getAllCookies { cookies in
            for cookie in cookies {
                if cookie.name == "authentication" {
                    self.viewController.peekableWKWebView().configuration.websiteDataStore.httpCookieStore.delete(cookie)
                } else {
                    print("\(cookie.name) is set to \(cookie.value)")
                }
            }
        }
        return self
    }
    
    @discardableResult func takeSnapshot(rect: CGRect, completionHandler: @escaping (UIImage?, Error?) -> Void) -> Self {
        let config = WKSnapshotConfiguration()
        config.rect = rect
        viewController.peekableWKWebView().takeSnapshot(with: config) { (image, error) in
            completionHandler(image, error)
        }
        return self
    }
}

public extension PeekabooWKWebView {
    @discardableResult func onClickWord(delegate: @escaping (String) -> ()) -> Self {
        viewController.delegateWrapper.onClickWordDelegate = delegate
        return self
    }
}
