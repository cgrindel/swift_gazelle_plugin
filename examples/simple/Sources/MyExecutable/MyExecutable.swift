import MyLibrary

@main
struct MyExecutable {
    static func main() async {
        let world = World()
        print("Hello, \(world.name)!")
    }
}
