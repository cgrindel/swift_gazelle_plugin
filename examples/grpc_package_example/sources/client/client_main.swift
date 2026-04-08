// Copyright 2019 The Bazel Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import Foundation
import GRPCCore
import GRPCNIOTransportHTTP2
import SwiftProtobuf
import EchoServiceClient

@main
struct ClientMain {
  static func main() async throws {
    try await withGRPCClient(
      transport: .http2NIOPosix(
        target: .ipv4(host: "localhost", port: 9000),
        transportSecurity: .plaintext
      )
    ) { client in
      let echo = EchoService_Echo.Client(wrapping: client)

      let request = EchoService_EchoRequest.with {
        $0.contents = "Hello, world!"
        let timestamp = Google_Protobuf_Timestamp(date: Date())
        $0.extra = try! Google_Protobuf_Any(message: timestamp)
      }

      do {
        let response = try await echo.echo(request)
        print(response.contents)
      } catch {
        print("Echo failed: \(error)")
      }
    }
  }
}
