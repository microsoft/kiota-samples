import 'package:cli/cli.dart';
import '../client/models/post.dart';
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

    test('toGetRequestInformation generates expected request', () {
      var reqInfo = client.posts.toGetRequestInformation();
      expect(reqInfo.httpMethod, HttpMethod.get);
      expect(reqInfo.urlTemplate, '{+baseurl}/posts{?title*,userId*}');
      expect(reqInfo.headers['Accept'], contains('application/json'));
    });

    test('toGetRequestInformation applies query parameters correctly', () {
      var reqInfo = client.posts.toGetRequestInformation((config) {
        config.queryParameters.userId = 42;
        config.queryParameters.title = 'Test Title';
      });
      expect(reqInfo.httpMethod, HttpMethod.get);
      var queryParams = reqInfo.queryParameters;
      expect(queryParams['userId'], 42);
      expect(queryParams['title'], 'Test Title');
    });

    test('toPostRequestInformation generates expected request and content', () {
      var post = Post();
      post.userId = 42;
      post.title = 'New Post';
      post.body = 'Post body';

      var reqInfo = client.posts.toPostRequestInformation(post);
      expect(reqInfo.httpMethod, HttpMethod.post);
      expect(reqInfo.headers['Accept'], contains('application/json'));
      expect(reqInfo.content, isNotNull);
    });

    test('byPostId sets path parameter correctly', () {
      var itemReq = client.posts.byPostId(5);
      expect(itemReq, isNotNull);

      var getInfo = itemReq.toGetRequestInformation();
      expect(getInfo.httpMethod, HttpMethod.get);
      expect(getInfo.pathParameters['post%2Did'], 5);

      var deleteInfo = itemReq.toDeleteRequestInformation();
      expect(deleteInfo.httpMethod, HttpMethod.delete);
      expect(deleteInfo.pathParameters['post%2Did'], 5);
    });
  });
}
