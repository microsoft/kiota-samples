// <ProgramSnippet>
import '../client/posts_client.dart';
import '../client/models/post.dart';
import 'package:microsoft_kiota_bundle/microsoft_kiota_bundle.dart';

void main(List<String> arguments) async {
  var authenticationProvider = AnonymousAuthenticationProvider();
  var requestAdapter =
      DefaultRequestAdapter(authProvider: authenticationProvider);
  var client = PostsClient(requestAdapter);

  // GET /posts
  var posts = await client.posts.getAsync();
  print('Posts: ${posts?.length ?? 0}');

  // GET /posts/{id}
  var specificPostId = 5;
  var specificPost = await client.posts.byPostId(specificPostId).getAsync();
  print(
      'Retrieved post - ID: ${specificPost?.id}, Title: ${specificPost?.title}, Body: ${specificPost?.body}');

  // POST /posts
  var newPost = Post();
  newPost.body = 'Hello world';
  newPost.title = 'Testing Kiota-generated API client';
  newPost.userId = 42;
  var createdPost = await client.posts.postAsync(newPost);
  print('Created new post with ID: ${createdPost?.id}');

  // PATCH /posts/{id}
  var updatePost = Post();
  updatePost.title = 'Updated title';
  var updatedPost =
      await client.posts.byPostId(specificPostId).patchAsync(updatePost);
  print(
      'Updated post - ID: ${updatedPost?.id}, Title: ${updatedPost?.title}, Body: ${updatedPost?.body}');

  // DELETE /posts/{id}
  await client.posts.byPostId(specificPostId).deleteAsync();
}
// </ProgramSnippet>
