//
//  AppDelegate.swift
//  ScrollMapperBible
//
//  Created by Zhengqian Kuang on 2020-07-01.
//  Copyright © 2020 Kuang. All rights reserved.
//

import UIKit

@UIApplicationMain
class AppDelegate: UIResponder, UIApplicationDelegate {



    func application(_ application: UIApplication, didFinishLaunchingWithOptions launchOptions: [UIApplication.LaunchOptionsKey: Any]?) -> Bool {
        
        // Override point for customization after application launch.
        test()
        
        return true
    }

    // MARK: UISceneSession Lifecycle

    func application(_ application: UIApplication, configurationForConnecting connectingSceneSession: UISceneSession, options: UIScene.ConnectionOptions) -> UISceneConfiguration {
        // Called when a new scene session is being created.
        // Use this method to select a configuration to create the new scene with.
        return UISceneConfiguration(name: "Default Configuration", sessionRole: connectingSceneSession.role)
    }

    func application(_ application: UIApplication, didDiscardSceneSessions sceneSessions: Set<UISceneSession>) {
        // Called when the user discards a scene session.
        // If any sessions were discarded while the application was not running, this will be called shortly after application:didFinishLaunchingWithOptions.
        // Use this method to release any resources that were specific to the discarded scenes, as they will not return.
    }


}

// MARK: - custom

extension AppDelegate {
    func test() {
        ScrollMapperBibleVersion.test()
        ScrollMapperBibleBookInfo.test()
        ScrollMapperBibleCrossReference.test()
        
        print("*** ASV: \(ScrollMapperBibleVersion.BibleVersion.ASV.table() ?? "<nil>")")
        print("*** BBE: \(ScrollMapperBibleVersion.BibleVersion.BBE.table() ?? "<nil>")")
        print("*** DARBY: \(ScrollMapperBibleVersion.BibleVersion.DARBY.table() ?? "<nil>")")
        print("*** KJV: \(ScrollMapperBibleVersion.BibleVersion.KJV.table() ?? "<nil>")")
        print("*** WBT: \(ScrollMapperBibleVersion.BibleVersion.WBT.table() ?? "<nil>")")
        print("*** WEB: \(ScrollMapperBibleVersion.BibleVersion.WEB.table() ?? "<nil>")")
        print("*** YLT: \(ScrollMapperBibleVersion.BibleVersion.YLT.table() ?? "<nil>")")
        
        print("*** Genesis: \(ScrollMapperBibleBookInfo.BibleBook.Genesis.order()), \(ScrollMapperBibleBookInfo.BibleBook.Genesis.titleShort()), \(ScrollMapperBibleBookInfo.BibleBook.Genesis.titleFull()), \(ScrollMapperBibleBookInfo.BibleBook.Genesis.abbreviation()), \(ScrollMapperBibleBookInfo.BibleBook.Genesis.category()), \(ScrollMapperBibleBookInfo.BibleBook.Genesis.otnt()), \(ScrollMapperBibleBookInfo.BibleBook.Genesis.chapters())")
    }
}

