# Copyright (c) Microsoft Corporation.
# Licensed under the MIT License.

# <ProgramSnippet>
# frozen_string_literal: true
require 'microsoft_kiota_abstractions'
require 'microsoft_kiota_faraday'
require 'microsoft_kiota_serialization_json'
require_relative './client/posts_client'

# API requires no authentication, so use the anonymous
# authentication provider
auth_provider = MicrosoftKiotaAbstractions::AnonymousAuthenticationProvider.new
# Create request adapter using the Faraday-based implementation
request_adapter = MicrosoftKiotaFaraday::FaradayRequestAdapter
  .new(auth_provider,
       MicrosoftKiotaSerializationJson::JsonParseNodeFactory.new,
       MicrosoftKiotaSerializationJson::JsonSerializationWriterFactory.new)
# Create the API client
client = KiotaPosts::PostsClient.new(request_adapter)

# GET /posts
all_posts = client.posts.get.resume
puts "Retrieved #{all_posts.length} posts."

# GET /posts/{id}
specific_post_id = 5
specific_post = client.posts.by_post_id(specific_post_id).get.resume
puts "Retrieved post - ID: #{specific_post.id}, " \
     "Title: #{specific_post.title}, " \
     "Body: #{specific_post.body}"

# POST /posts
new_post = KiotaPosts::Models::Post.new
new_post.user_id = 42
new_post.title = 'Testing Kiota-generated API client'
new_post.body = 'Hello world!'

created_post = client.posts.post(new_post).resume
puts "Created new post with ID: #{created_post.id}"

# PATCH /posts/{id}
update = KiotaPosts::Models::Post.new
# Only update title
update.title = 'Updated title'

updated_post = client.posts.by_post_id(specific_post_id).patch(update).resume
puts "Updated post - ID: #{updated_post.id}, " \
     "Title: #{updated_post.title}, " \
     "Body: #{updated_post.body}"

# DELETE /posts/{id}
client.posts.by_post_id(specific_post_id).delete.resume
# </ProgramSnippet>
