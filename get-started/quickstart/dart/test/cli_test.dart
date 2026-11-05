import 'package:cli/cli.dart';
import '../client/posts_client.dart';
import 'package:microsoft_kiota_bundle/microsoft_kiota_bundle.dart';
import 'package:test/test.dart';

void main() {
  test('calculate', () {
    expect(calculate(), 42);
  });

  group('PostsClient and Request Builders', () {
    late PostsClient client;

    setUp(() {
      var authProvider = AnonymousAuthenticationProvider();
      var requestAdapter = DefaultRequestAdapter(authProvider: authProvider);
      client = PostsClient(requestAdapter);
    });

    test('PostsClient initializes correctly with base url', () {
      expect(client, isNotNull);
      expect(client.posts, isNotNull);
      expect(client.requestAdapter.baseUrl, 'https://jsonplaceholder.typicode.com');
    });
  });
}
