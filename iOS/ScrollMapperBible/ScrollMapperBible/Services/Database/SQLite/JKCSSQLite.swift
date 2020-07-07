//
//  JKCSSQLite.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-06.
//  Copyright © 2020 Kuang. All rights reserved.
//

import Foundation
import SQLite3

// https://www.raywenderlich.com/6620276-sqlite-with-swift-tutorial-getting-started

public enum JKCSSQLiteError: Error {
    case noError(message: String = "No error")
    case genericError(message: String = "Generic error")
    case customError(message: String)
    case codedError((code: Int, message: String))
    
    public var code: Int? {
        switch self {
        case .customError, .genericError, .noError:
            return nil
        case .codedError((let code, _)):
            return code
        }
    }
    
    public var message: String {
        switch self {
        case .customError(let msg), .genericError(let msg), .noError(let msg):
            return msg
        case .codedError((_, let msg)):
            return msg
        }
    }
}

public class JKCSSQLite {
    private var resourceName: String? = nil
    private var resourceExt: String? = nil
    var db: OpaquePointer? = nil
    
    public init?(resourceName: String, resourceExt: String) {
        let result = openDatabase(resourceName: resourceName, resourceExt: resourceExt)
        switch result {
        case .success(_):
            break
        default:
            return nil
        }
    }
    
    deinit {
        closeDatabase()
    }
    
    @discardableResult
    private func openDatabase(resourceName: String, resourceExt: String) -> Result<ExpressibleByNilLiteral?, JKCSSQLiteError> {
        guard let dbPath = pathOfResource(name: resourceName, ext: resourceExt) else {
            let message = "Unable to find the path for resource"
            print("useDatabase(name: \(resourceName), ext: \(resourceExt)) failed: \(message)")
            return Result.failure(.customError(message: message))
        }
        var db: OpaquePointer? = nil
        guard sqlite3_open(dbPath, &db) == SQLITE_OK else {
            let message = "Unable to open database at \"\(dbPath)\""
            print("useDatabase(name: \(resourceName), ext: \(resourceExt)) failed: \(message)")
            return Result.failure(.customError(message: message))
        }
        print("useDatabase(name: \(resourceName), ext: \(resourceExt)) success")
        self.db = db
        self.resourceName = resourceName
        self.resourceExt = resourceExt
        return Result.success(nil)
    }
    
    private func closeDatabase() {
        guard let db = db else { return }
        sqlite3_close(db)
    }
    
    private func pathOfResource(name: String, ext: String) -> String? {
        let path = Bundle.main.path(forResource: name, ofType: ext)
        return path
    }
}

public class JKCSSQLiteQuery {
    let sqlite: JKCSSQLite
    var statement: String
    var queryStatement: OpaquePointer?
    
    public init?(resourceName: String, resourceExt: String, statement: String) {
        guard
            let sqlite = JKCSSQLite(resourceName: resourceName, resourceExt: resourceExt),
            let db = sqlite.db
        else {
            return nil
        }
        self.sqlite = sqlite
        self.statement = statement
        if sqlite3_prepare_v2(db, statement, -1, &queryStatement, nil) != SQLITE_OK {
            print("query(\(statement)) failed: \(String(cString: sqlite3_errmsg(db)))")
            return nil
        }
    }
    
    deinit {
        finalize()
    }
    
    public func finalize() {
        if let qStatement = queryStatement {
            sqlite3_finalize(qStatement)
        }
    }
}
