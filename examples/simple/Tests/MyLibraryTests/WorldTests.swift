@testable import MyLibrary
import XCTest

final class WorldTests: XCTestCase {
    func test() throws {
        XCTAssertEqual(World().name, "World")
    }
}
