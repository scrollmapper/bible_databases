//
//  ScrollMapperBibleCrossReferenceView.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-16.
//  Copyright © 2020 Kuang. All rights reserved.
//

import SwiftUI

struct ScrollMapperBibleCrossReferenceView: View {
    @ObservedObject var viewModel: ScrollMapperBibleCrossReferenceViewModel
    private var vid: Int
    
    init(vid: Int) {
        self.vid = vid
        viewModel = ScrollMapperBibleCrossReferenceViewModel(vid: vid)
    }

    var body: some View {
        GeometryReader { (geometryProxy) in
            VStack {
                self.viewHeader(geometryProxy: geometryProxy)
                List {
                    ForEach(self.viewModel.listData) { section in
                        Section(header: self.sectionHeader(section)) {
                            ForEach(section.items) { (item) in
                                Text("\(item.verse.t)")
                            }
                        }
                    }
                }
            }
        }
    }
    
    private func viewHeader(geometryProxy: GeometryProxy) -> some View {
        let b = vid / 1_000_000
        let c = vid / 1_000 % 1_000
        let v = vid % 1_000
        let book = ScrollMapperBibleBookInfo.BibleBook(rawValue: b)?.bookInfo()?.abbreviation ?? ""
        return Text("\(book) \(c):\(v)").frame(width: geometryProxy.size.width).background(Color(.systemBackground))
    }
    
    private func sectionHeader(_ section: ScrollMapperBibleCrossReferenceViewModel.Section) -> some View {
        let sb = ScrollMapperBibleBookInfo.BookInfo(order: section.sv / 1_000_000)?.abbreviation ?? ""
        let sc = section.sv / 1_000 % 1_000
        let sv = section.sv % 1_000
        let eb = ((section.ev == 0) ? sb : (ScrollMapperBibleBookInfo.BookInfo(order: section.ev / 1_000_000)?.abbreviation ?? ""))
        let ec = ((section.ev == 0) ? sc : (section.ev / 1_000 % 1_000))
        let ev = ((section.ev == 0) ? sv : (section.ev % 1_000))
        var header = "\(sb) \(sc):\(sv)"
        if section.ev != 0 {
            header += "-"
            if eb != sb {
                header += "\(eb) "
            }
            if (eb != sb) || (ec != sc) {
                header += "\(ec):"
            }
            header += "\(ev)"
        }
        
        return Text("\(header)")
    }
}

struct ScrollMapperBibleCrossReferenceView_Previews: PreviewProvider {
    static var previews: some View {
        ScrollMapperBibleCrossReferenceView(vid: 0)
    }
}
