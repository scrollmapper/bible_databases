//
//  ScrollMapperBibleVersion.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-06.
//  Copyright © 2020 Kuang. All rights reserved.
//

import Foundation
import SQLite3

public class ScrollMapperBibleVersion: ScrollMapperBibleModelBase {
    public enum BibleVersion: String {
        static let versions: [BibleVersionKey] = {
            if let versions = ScrollMapperBibleVersion()?.result {
                return versions
            }
            return []
        }()
        
        case ASV, BBE, DARBY, KJV, WBT, WEB, YLT
        
        public func abbreviation() -> String {
            return self.rawValue
        }
        
        public func table() -> String? {
            let version = BibleVersion.versions.first { $0.abbreviation == self.abbreviation() }
            return version?.table
        }
        
        public func next() -> BibleVersion {
            switch self {
            case .ASV:
                return .BBE
            case .BBE:
                return .DARBY
            case .DARBY:
                return .KJV
            case .KJV:
                return .WBT
            case .WBT:
                return .WEB
            case .WEB:
                return .YLT
            case .YLT:
                return .ASV
            }
        }
    }
    
    public typealias StructType = BibleVersionKey
    
    public struct BibleVersionKey {
        var id: Int = 0
        var table: String = ""
        var abbreviation: String = ""
        var language: String = ""
        var version: String = ""
        var info_text: String = ""
        var info_url: String = ""
        var publisher: String = ""
        var copyright: String = ""
        var copyright_info: String = ""
    }
    
    public lazy var result: [StructType] = {
        return getResult()
    }()
    
    public required init?(statement: String = "SELECT * FROM bible_version_key ORDER BY id") {
        super.init(statement: statement)
    }
    
    private func getResult() -> [StructType] {
        guard let queryStatement = queryStatement else {
            return []
        }
        var result: [StructType] = []
        while sqlite3_step(queryStatement) == SQLITE_ROW {
            let id = sqlite3_column_int(queryStatement, 0)
            guard
                let table = sqlite3_column_text(queryStatement, 1),
                let abbreviation = sqlite3_column_text(queryStatement, 2),
                let language = sqlite3_column_text(queryStatement, 3),
                let version = sqlite3_column_text(queryStatement, 4),
                let info_text = sqlite3_column_text(queryStatement, 5),
                let info_url = sqlite3_column_text(queryStatement, 6),
                let publisher = sqlite3_column_text(queryStatement, 7),
                let copyright = sqlite3_column_text(queryStatement, 8),
                let copyright_info = sqlite3_column_text(queryStatement, 9)
            else {
                continue
            }
            let row = StructType(id: Int(id), table: String(cString: table), abbreviation: String(cString: abbreviation), language: String(cString: language), version: String(cString: version), info_text: String(cString: info_text), info_url: String(cString: info_url), publisher: String(cString: publisher), copyright: String(cString: copyright), copyright_info: String(cString: copyright_info))
            result.append(row)
        }
        return result
    }
    
    public static func test() {
        print("testScrollMapperBibleVersion")
        _ = self.init()?.result.map {
            print("id: \($0.id), table: \($0.table), abbreviation: \($0.abbreviation), language: \($0.language), version: \($0.version), info_text: \($0.info_text), info_url: \($0.info_url), publisher: \($0.publisher), copyright: \($0.copyright), copyright_info: \($0.copyright_info)")
        }
    }
}
