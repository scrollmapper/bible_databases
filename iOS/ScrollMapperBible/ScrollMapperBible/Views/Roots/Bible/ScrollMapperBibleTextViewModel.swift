//
//  ScrollMapperBibleTextViewModel.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-07.
//  Copyright © 2020 Kuang. All rights reserved.
//

import Foundation
import Combine

let scrollMapperBibleSceneStorageKeyCurrentChapterCid = "ED1BD62A-8E2F-4626-9D17-22E78298FE69"

class ScrollMapperBibleTextViewModel: ScrollMapperBibleViewModelBase {
    private var currentChapter: ScrollMapperBibleChapter.BibleChapter {
        didSet {
            if currentChapter != oldValue {
                UserDefaults.standard.set(currentChapter.cid, forKey: scrollMapperBibleSceneStorageKeyCurrentChapterCid)
            }
        }
    }
    
    override init() {
        var currentChapterCid = UserDefaults.standard.integer(forKey: scrollMapperBibleSceneStorageKeyCurrentChapterCid)
        if currentChapterCid == 0 {
            currentChapterCid = 1001
        }
        currentChapter = ScrollMapperBibleChapter.BibleChapter(cid: currentChapterCid)!
        
        super.init()
        
        generateCurrentChapterHTMLString()
    }
    
    override func translationDidChange() {
        super.translationDidChange()
        
        generateCurrentChapterHTMLString()
    }
    
    @Published var currentChapterHTMLString: String = ""
    
    private func generateCurrentChapterHTMLString() {
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
        htmlString += "  <body>\n"
        var bodyContent = ""
        _ = ScrollMapperBibleVerseWithCrossReference(version: translation, book: currentChapter.bibleBook, chapter: currentChapter.c)?.result.map {
            bodyContent += "\(" ".withHTMLTags(fontSize: textFontSize))\("\($0.v)".withHTMLTags(fontSize: verseNumberFontSize, color: (($0.cr.count > 0) ? "blue" : ""), sup: true))\(" ".withHTMLTags(fontSize: textFontSize))\($0.t.withHTMLTags(fontSize: textFontSize))"
        }
        htmlString += "    <p class=\"p_clickable\">\(bodyContent)</p>\n"
        htmlString += "  </body>\n"
        htmlString += "</html>\n"
        
        currentChapterHTMLString = htmlString
    }
}
