//
//  ScrollMapperBibleJumpToView.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-14.
//  Copyright © 2020 Kuang. All rights reserved.
//

import SwiftUI

struct ScrollMapperBibleJumpToView: View {
    @ObservedObject var viewModel = ScrollMapperBibleJumpToViewModel()
    @Environment(\.presentationMode) var presentationMode

    var body: some View {
        GeometryReader { (geometryProxy)  in
            VStack {
                Picker(selection: self.$viewModel.selectedScopeIndex, label: Text("Scope")) {
                    ForEach(self.viewModel.scopes) { scope in
                        Text("\(scope.title)")
                    }
                }
                .pickerStyle(SegmentedPickerStyle())
                
                List {
                    ForEach(self.viewModel.listData) { section in
                        Section(header: Text(section.title.uppercased())) {
                            ForEach(section.books) { item in
                                VStack(alignment: .leading) {
                                    Text(item.book.title_short).font(.system(size: 15))
                                    Text(item.book.title_full).foregroundColor(Color(.systemGray)).font(.system(size: 11))
                                    self.chaptersView(item: item, geometryProxy: geometryProxy)
                                }
                                .onTapGesture {
                                    self.viewModel.expand(book: item.book)
                                }
                            }
                        }
                    }
                }
            }
            .navigationBarTitle("Jump To")
        }
    }
    
    private func chaptersView(item: ScrollMapperBibleJumpToViewModel.Item, geometryProxy: GeometryProxy) -> AnyView? {
        guard
            let expandedBook = viewModel.expandedBook,
            expandedBook.title_short == item.book.title_short
        else {
            return nil
        }
        
        let itemWidth: CGFloat = 32
        let horizontalSpace: CGFloat = 8
        let verticalSpace: CGFloat = 16
        let leadingSpace: CGFloat = 8
        let trailingSpace: CGFloat = 8
        let numberOfColumns: Int = Int((geometryProxy.size.width - (leadingSpace + trailingSpace)) / (itemWidth + horizontalSpace))
        if numberOfColumns < 1 { return nil }
        let numberOfRows = (item.book.chapters + (numberOfColumns - 1)) / numberOfColumns
        
        let gridView = HStack {
            Spacer().frame(width: leadingSpace)
            VStack {
                ForEach(0..<numberOfRows, id: \.self) { i in
                    VStack{
                        Spacer().frame(height: verticalSpace / 2)
                        HStack {
                            HStack {
                                ForEach(0..<numberOfColumns, id: \.self) { j in
                                    HStack {
                                        Spacer().frame(width: horizontalSpace / 2)
                                        if i * numberOfColumns + j < item.book.chapters {
                                            Text("\(i * numberOfColumns + j + 1)").frame(width: itemWidth)
                                                .onTapGesture {
                                                    self.viewModel.jumpTo(book: item.book, chapter: i * numberOfColumns + j + 1)
                                                    self.presentationMode.wrappedValue.dismiss()
                                            }
                                        }
                                        else {
                                            Text("").frame(width: itemWidth)
                                        }
                                        Spacer().frame(width: horizontalSpace / 2)
                                    }
                                }
                            }
                            Spacer()
                        }
                        Spacer().frame(height: verticalSpace / 2)
                    }
                }
            }
            Spacer().frame(width: trailingSpace)
        }
        
        return AnyView(gridView)
    }
    
}

struct JumpToView_Previews: PreviewProvider {
    static var previews: some View {
        ScrollMapperBibleJumpToView()
    }
}
