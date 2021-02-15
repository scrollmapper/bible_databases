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
    private var uiViewModel = JKCSUIViewModel()
    
    @Published var translation: ScrollMapperBibleVersion.BibleVersion = .KJV {
        didSet {
            if translation != oldValue {
                createCurrentChapterPage()
            }
        }
    }
    var translationSubscriber: AnyCancellable? = nil
    
    var alert: AlertType = .unknown {
        didSet {
            switch alert {
            case .verseNumber( _, _, _):
                showAlert = true
            default:
                break
            }
        }
    }
    @Published var showAlert = false
    
    @Published var crossReferenceVid: Int = 0
    
    private var currentChapter: ScrollMapperBibleChapter.BibleChapter = ScrollMapperBibleChapter.BibleChapter(b: 1, c: 1)! {
        didSet {
            if currentChapter != oldValue {
                createCurrentChapterPage()
                scrollMapperBiblePublishers.publishCurrentChapter(cid: currentChapter.cid)
            }
        }
    }
    var currentChapterSubscriber: AnyCancellable? = nil
    
    private var darkMode: Bool = false {
        didSet {
            if darkMode != oldValue {
                createCurrentChapterPage()
            }
        }
    }
    
    var bibleText: [ScrollMapperBibleText.BibleText] = []
    var currentChapterHTMLString: String = ""
    let currentChapterUpdatedSubject = CurrentValueSubject<String, Never>("")
    let currentChapterUpdatedPublisher: AnyPublisher<String, Never>
    
    init() {
        currentChapterUpdatedPublisher = currentChapterUpdatedSubject.eraseToAnyPublisher()
        subscribe()
        createCurrentChapterPage()
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
        
        uiViewModel.onInterfaceOrientationChange { (_) in
            self.createCurrentChapterPage()
        }
    }
    
    func unsubscribe() {
        translationSubscriber?.cancel()
        translationSubscriber = nil
        currentChapterSubscriber?.cancel()
        currentChapterSubscriber = nil
    }
    
    private func createCurrentChapterPage() {
        bibleText.removeAll()
        
        // https://useyourloaf.com/blog/supporting-dark-mode-in-wkwebview/
        
        let textFontSize = (isPhone && uiViewModel.isPortrait) ? 60 : 24
        let verseNumberFontSize = (isPhone && uiViewModel.isPortrait) ? 36 : 14

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
        htmlString += "  <h1 style=\"text-align:center; font-size:48pt; color:\(darkMode ? "white" : "black")\">\(currentChapter.bibleBook.bookInfo()?.title_short ?? "") \(currentChapter.c)</h1>\n"
        htmlString += "  <body>\n"
        var bodyContent = ""
        if let result = ScrollMapperBibleVerseWithCrossReference(version: translation, book: currentChapter.bibleBook, chapter: currentChapter.c)?.result {
            for verse in result {
                bibleText.append(ScrollMapperBibleText.BibleText(id: verse.id, b: verse.b, c: verse.c, v: verse.v, t: verse.t))
                bodyContent += "\(" ".withHTMLTags(fontSize: textFontSize))\("\(verse.v)".withHTMLTags(fontSize: verseNumberFontSize, color: ((verse.cr.count > 0) ? (darkMode ? "#FFFF00" : "#0000FF") : ""), sup: true))\(" ".withHTMLTags(fontSize: textFontSize))\(verse.t.withHTMLTags(fontSize: textFontSize, color: darkMode ? "white" : "black"))"
            }
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

extension ScrollMapperBibleTextViewModel {
    enum AlertType {
        case unknown
        case verseNumber(_ book: ScrollMapperBibleBookInfo.BibleBook, _ chapter: Int, _ number: Int)
        case word(_ word: String)
    }
    
    func checkClickedMessage(message: Any?) {
        guard
            let message = message as? [String : Any],
            let clicked = message["clicked"] as? [String : Any],
            let word = clicked["word"] as? String,
            let sentence = clicked["sentence"] as? String
        else {
            return
        }
        if CharacterSet.decimalDigits.isSuperset(of: CharacterSet(charactersIn: word)) { // number
            for text in bibleText {
                let verse = "\(text.v) \(text.t)"
                let length = min(sentence.count, verse.count)
                if sentence.prefix(length) == verse.prefix(length),
                    let number = Int(word) {
                    alert = .verseNumber(currentChapter.bibleBook, currentChapter.c, number)
                    return
                }
            }
        }
        else { // word
            alert = .word(word)
        }
    }
    
    func retrieveCrossReference(book: ScrollMapperBibleBookInfo.BibleBook, chapter: Int, verse: Int) {
        crossReferenceVid = book.rawValue * 1_000_000 + chapter * 1_000 + verse
    }
    
    func lookUp(_ word: String) {
        print("*** lookUp: \(word)")
    }
    
    func dismissCrossReference() {
        crossReferenceVid = 0
    }
}
