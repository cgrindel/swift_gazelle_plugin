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

import GRPCCore
import GRPCInProcessTransport
import XCTest
import EchoServiceClient
import EchoServiceServer

struct TestEchoProvider: EchoServiceServer.EchoService_Echo.SimpleServiceProtocol {
  func echo(
    request: EchoServiceServer.EchoService_EchoRequest,
    context: ServerContext
  ) async throws -> EchoServiceServer.EchoService_EchoResponse {
    return EchoServiceServer.EchoService_EchoResponse.with {
      $0.contents = request.contents
    }
  }
}

final class ClientUnitTest: XCTestCase {

  func testGetWithRealClientAndServer() async throws {
    let inProcess = InProcessTransport()

    try await withThrowingTaskGroup(of: Void.self) { group in
      let server = GRPCServer(
        transport: inProcess.server,
        services: [TestEchoProvider()]
      )

      group.addTask {
        try await server.serve()
      }

      try await withGRPCClient(
        transport: inProcess.client
      ) { client in
        let echo = EchoServiceClient.EchoService_Echo.Client(wrapping: client)
        let request = EchoServiceClient.EchoService_EchoRequest.with {
          $0.contents = "Hello"
        }
        let response = try await echo.echo(request)
        XCTAssertEqual(response.contents, "Hello")
      }

      server.beginGracefulShutdown()
    }
  }
}
