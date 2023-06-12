# Copyright (c) Microsoft Corporation.
# Licensed under the MIT License.

# <ProgramSnippet>
import asyncio
from kiota_abstractions.authentication.anonymous_authentication_provider import (
    AnonymousAuthenticationProvider)
from kiota_http.httpx_request_adapter import HttpxRequestAdapter
from client.posts_client import PostsClient
from client.models.post import Post

async def main():
    # You may need this if your're using asyncio on windows
    # See: https://stackoverflow.com/questions/63860576
    # asyncio.set_event_loop_policy(asyncio.WindowsSelectorEventLoopPolicy())

    # API requires no authentication, so use the anonymous
    # authentication provider
    auth_provider = AnonymousAuthenticationProvider()
    # Create request adapter using the HTTPX-based implementation
    request_adapter = HttpxRequestAdapter(auth_provider)
    # Create the API client
    client = PostsClient(request_adapter)

    # GET /posts
    all_posts = await client.posts.get()
    print(f"Retrieved {len(all_posts)} posts.")

    # GET /posts/{id}
    specific_post_id = "5"
    specific_post = await client.posts.by_post_id(specific_post_id).get()
    print(f"Retrieved post - ID: {specific_post.id}, " +
          f"Title: {specific_post.title}, " +
          f"Body: {specific_post.body}")

    # POST /posts
    new_post = Post()
    new_post.user_id = 42
    new_post.title = "Testing Kiota-generated API client"
    new_post.body = "Hello world!"

    created_post = await client.posts.post(new_post)
    print(f"Created new post with ID: {created_post.id}")

    # PATCH /posts/{id}
    update = Post()
    # Only update title
    update.title = "Updated title"

    updated_post = await client.posts.by_post_id(specific_post_id).patch(update)
    print(f"Updated post - ID: {updated_post.id}, " +
          f"Title: {updated_post.title}, " +
          f"Body: {updated_post.body}")

    # DELETE /posts/{id}
    await client.posts.by_post_id(specific_post_id).delete()

# Run main
asyncio.run(main())
# </ProgramSnippet>
