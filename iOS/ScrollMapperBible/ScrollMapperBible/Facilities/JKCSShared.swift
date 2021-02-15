//
//  JKCSShared.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-07.
//  Copyright © 2020 Kuang. All rights reserved.
//

import Foundation

public enum JKCSError: Error {
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

