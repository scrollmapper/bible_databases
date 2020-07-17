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
    @Published var translation: ScrollMapperBibleVersion.BibleVersion = .KJV {
        didSet {
            if translation != oldValue {
                generateCurrentChapterHTMLString()
            }
        }
    }
    var translationSubscriber: AnyCancellable? = nil
    
    private var currentChapter: ScrollMapperBibleChapter.BibleChapter = ScrollMapperBibleChapter.BibleChapter(b: 1, c: 1)! {
        didSet {
            if currentChapter != oldValue {
                generateCurrentChapterHTMLString()
                scrollMapperBiblePublishers.publishCurrentChapter(cid: currentChapter.cid)
            }
        }
    }
    var currentChapterSubscriber: AnyCancellable? = nil
    
    private var darkMode: Bool = false {
        didSet {
            if darkMode != oldValue {
                generateCurrentChapterHTMLString()
            }
        }
    }
    
    var currentChapterHTMLString: String = ""
    let currentChapterUpdatedSubject = CurrentValueSubject<String, Never>("")
    let currentChapterUpdatedPublisher: AnyPublisher<String, Never>
    
    init() {
        currentChapterUpdatedPublisher = currentChapterUpdatedSubject.eraseToAnyPublisher()
        subscribe()
        generateCurrentChapterHTMLString()
    }
    
    deinit {
        unsubscribe()
    }
    
    func subscribe() {
        translationSubscriber = scrollMapperBiblePublishers.tranlationPublisher.sink(receiveValue: { (translation) in
            self.translation = ScrollMapperBibleVersion.BibleVersion(rawValue: translation) ?? ScrollMapperBibleVersion.BibleVersion.KJV
        })
        currentChapterSubscriber = scrollMapperBiblePublishers.currentChapterCidPublisher.sink(receiveValue: { (cid) in
            self.currentChapter = ScrollMapperBibleChapter.BibleChapter(cid: cid) ?? ScrollMapperBibleChapter.BibleChapter(b: 1, c: 1)!
        })
    }
    
    func unsubscribe() {
        translationSubscriber?.cancel()
        translationSubscriber = nil
        currentChapterSubscriber?.cancel()
        currentChapterSubscriber = nil
    }
    
    private func generateCurrentChapterHTMLString() {
        // https://useyourloaf.com/blog/supporting-dark-mode-in-wkwebview/
        
        let textFontSize = 60
        let verseNumberFontSize = 36

        var htmlString = ""
        htmlString += "<!DOCTYPE html>\n"
        htmlString += "<html>\n"
        htmlString += "  <head>\n"
        htmlString += "    <script src=\"https://ajax.googleapis.com/ajax/libs/jquery/3.4.1/jquery.min.js\"></script>\n"
        htmlString += "    <script>\n"
        htmlString += "      $(document).ready(function() {\n"
        htmlString += "        let p_clickable = $(\".p_clickable\");\n"
        htmlString += "        p_clickable.click(function(e) {\n"
        htmlString += "          let sel_obj = window.getSelection();\n"
        htmlString += "          var sel_obj_word = sel_obj;\n"
        htmlString += "          sel_obj_word.modify(\"move\", \"forward\", \"character\");\n"
        htmlString += "          sel_obj_word.modify(\"extend\", \"backward\", \"word\");\n"
        htmlString += "          sel_obj_word.modify(\"move\", \"backward\", \"character\");\n"
        htmlString += "          sel_obj_word.modify(\"extend\", \"forward\", \"word\");\n"
        htmlString += "          let sel_word = sel_obj_word.toString().trim();\n"
        htmlString += "          var sel_obj_sentence = sel_obj;\n"
        htmlString += "          sel_obj_sentence.modify(\"move\", \"forward\", \"character\");\n"
        htmlString += "          sel_obj_sentence.modify(\"extend\", \"backward\", \"word\");\n"
        htmlString += "          sel_obj_sentence.modify(\"move\", \"backward\", \"character\");\n"
        htmlString += "          sel_obj_sentence.modify(\"extend\", \"forward\", \"sentence\");\n"
        htmlString += "          let sel_sentence = sel_obj_sentence.toString().trim();\n"
        htmlString += "          let message = {clicked: {word: sel_word, sentence: sel_sentence}};\n"
        htmlString += "          postMessage(message);\n"
        htmlString += "          window.getSelection().removeAllRanges();\n"
        htmlString += "        });\n"
        htmlString += "      });\n"
        htmlString += "    </script>\n"
        htmlString += "    <script>\n"
        htmlString += "      postMessage = function(message) {\n"
        htmlString += "        window.webkit.messageHandlers.wordClickHandler.postMessage(message);\n"
        htmlString += "      }\n"
        htmlString += "    </script>\n"
        htmlString += "  </head>\n"
        htmlString += "  <h1 style=\"text-align:center; font-size:48pt\">\(currentChapter.bibleBook.bookInfo()?.title_short ?? "") \(currentChapter.c)</h1>\n"
        htmlString += "  <body>\n"
        var bodyContent = ""
        _ = ScrollMapperBibleVerseWithCrossReference(version: translation, book: currentChapter.bibleBook, chapter: currentChapter.c)?.result.map {
            bodyContent += "\(" ".withHTMLTags(fontSize: textFontSize))\("\($0.v)".withHTMLTags(fontSize: verseNumberFontSize, color: (($0.cr.count > 0) ? (darkMode ? "#FFFF00" : "#0000FF") : ""), sup: true))\(" ".withHTMLTags(fontSize: textFontSize))\($0.t.withHTMLTags(fontSize: textFontSize, color: darkMode ? "white" : "black"))"
        }
        htmlString += "    <p class=\"p_clickable\">\(bodyContent)</p>\n"
        htmlString += "  </body>\n"
        htmlString += "</html>\n"
        
        currentChapterHTMLString = htmlString
        currentChapterUpdatedSubject.send(currentChapterHTMLString)
    }
    
    func gotoPreviousChapter() {
        if let previousChapter = <~currentChapter {
            currentChapter = previousChapter
            
        }
    }
    
    func gotoNextChapter() {
        if let nextChapter = currentChapter~> {
            currentChapter = nextChapter
        }
    }
    
    func webViewDidFinishNavigation() {
        
    }
    
    func colorSchemeDidChange(darkMode: Bool) {
        self.darkMode = darkMode
    }
}
