import '../client/models/post.dart';
import '../client/posts_client.dart';
import 'package:microsoft_kiota_bundle/microsoft_kiota_bundle.dart';
import 'package:test/test.dart';

void main() {
  group('PostsClient and Request Builders', () {
    late PostsClient client;

    setUp(() {
      final authProvider = AnonymousAuthenticationProvider();
      final requestAdapter = DefaultRequestAdapter(authProvider: authProvider);
      client = PostsClient(requestAdapter);
    });

    test('PostsClient initializes correctly with base url', () {
      expect(client, isNotNull);
      expect(client.posts, isNotNull);
      expect(
        client.requestAdapter.baseUrl,
        'https://jsonplaceholder.typicode.com',
      );
    });

    test('toGetRequestInformation generates expected request', () {
      final reqInfo = client.posts.toGetRequestInformation();
      expect(reqInfo.httpMethod, HttpMethod.get);
      expect(reqInfo.urlTemplate, '{+baseurl}/posts{?title*,userId*}');
      expect(reqInfo.headers['Accept'], contains('application/json'));
    });

    test('toGetRequestInformation applies query parameters correctly', () {
      final reqInfo = client.posts.toGetRequestInformation((config) {
        config.queryParameters.userId = 42;
        config.queryParameters.title = 'Test Title';
      });
      expect(reqInfo.httpMethod, HttpMethod.get);
      final queryParams = reqInfo.queryParameters;
      expect(queryParams['userId'], 42);
      expect(queryParams['title'], 'Test Title');
    });

    test('toPostRequestInformation generates expected request and content', () {
      final post = Post();
      post.userId = 42;
      post.title = 'New Post';
      post.body = 'Post body';

      final reqInfo = client.posts.toPostRequestInformation(post);
      expect(reqInfo.httpMethod, HttpMethod.post);
      expect(reqInfo.headers['Accept'], contains('application/json'));
      expect(reqInfo.content, isNotNull);
    });

    test('byPostId sets path parameter correctly', () {
      final itemReq = client.posts.byPostId(5);
      expect(itemReq, isNotNull);

      final getInfo = itemReq.toGetRequestInformation();
      expect(getInfo.httpMethod, HttpMethod.get);
      expect(getInfo.pathParameters.values, contains(5));

      final deleteInfo = itemReq.toDeleteRequestInformation();
      expect(deleteInfo.httpMethod, HttpMethod.delete);
      expect(deleteInfo.pathParameters.values, contains(5));
    });
  });

  group('Post Model', () {
    test('Post model properties can be set and retrieved', () {
      final post = Post();
      post.id = 10;
      post.userId = 42;
      post.title = 'Hello';
      post.body = 'World';

      expect(post.id, 10);
      expect(post.userId, 42);
      expect(post.title, 'Hello');
      expect(post.body, 'World');

      final fieldDeserializers = post.getFieldDeserializers();
      expect(fieldDeserializers.containsKey('id'), isTrue);
      expect(fieldDeserializers.containsKey('userId'), isTrue);
      expect(fieldDeserializers.containsKey('title'), isTrue);
      expect(fieldDeserializers.containsKey('body'), isTrue);
    });
  });
}
