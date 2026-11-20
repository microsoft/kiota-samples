# Copyright (c) Microsoft Corporation.
# Licensed under the MIT License.

import unittest
from kiota_abstractions.authentication.anonymous_authentication_provider import (
    AnonymousAuthenticationProvider,
)
from kiota_abstractions.base_request_configuration import RequestConfiguration
from kiota_abstractions.method import Method
from kiota_http.httpx_request_adapter import HttpxRequestAdapter
from kiota_serialization_json.json_serialization_writer import JsonSerializationWriter
from client.models.post import Post
from client.posts.posts_request_builder import PostsRequestBuilder
from client.posts_client import PostsClient


class TestPostsClient(unittest.TestCase):

    def setUp(self):
        self.auth_provider = AnonymousAuthenticationProvider()
        self.request_adapter = HttpxRequestAdapter(self.auth_provider)
        self.client = PostsClient(self.request_adapter)

    def test_client_initialization(self):
        self.assertIsNotNone(self.client)
        self.assertIsNotNone(self.client.posts)
        self.assertEqual(
            self.request_adapter.base_url,
            "https://jsonplaceholder.typicode.com",
        )

    def test_to_get_request_information(self):
        req_info = self.client.posts.to_get_request_information()
        self.assertEqual(req_info.http_method, Method.GET)
        self.assertEqual(req_info.url_template, "{+baseurl}/posts{?title*,userId*}")
        self.assertIn("application/json", req_info.headers.get("Accept"))

    def test_to_get_request_information_with_query_parameters(self):
        query_params = PostsRequestBuilder.PostsRequestBuilderGetQueryParameters(
            user_id=42,
            title="Testing Kiota",
        )
        config = RequestConfiguration(query_parameters=query_params)
        req_info = self.client.posts.to_get_request_information(config)
        self.assertEqual(req_info.http_method, Method.GET)
        self.assertEqual(req_info.query_parameters.get("userId"), 42)
        self.assertEqual(req_info.query_parameters.get("title"), "Testing Kiota")

    def test_by_post_id(self):
        post_item = self.client.posts.by_post_id(5)
        self.assertIsNotNone(post_item)
        self.assertEqual(post_item.path_parameters.get("post%2Did"), 5)

        with self.assertRaises(TypeError):
            self.client.posts.by_post_id(None)

    def test_item_to_get_request_information(self):
        post_item = self.client.posts.by_post_id(5)
        req_info = post_item.to_get_request_information()
        self.assertEqual(req_info.http_method, Method.GET)
        self.assertEqual(req_info.url_template, "{+baseurl}/posts/{post%2Did}")
        self.assertIn("application/json", req_info.headers.get("Accept"))

    def test_to_post_request_information(self):
        new_post = Post()
        new_post.user_id = 42
        new_post.title = "Testing Kiota"
        new_post.body = "Hello world!"

        req_info = self.client.posts.to_post_request_information(new_post)
        self.assertEqual(req_info.http_method, Method.POST)
        self.assertEqual(req_info.url_template, "{+baseurl}/posts{?title*,userId*}")

        with self.assertRaises(TypeError):
            self.client.posts.to_post_request_information(None)

    def test_item_to_patch_request_information(self):
        update = Post()
        update.title = "Updated title"

        post_item = self.client.posts.by_post_id(5)
        req_info = post_item.to_patch_request_information(update)
        self.assertEqual(req_info.http_method, Method.PATCH)
        self.assertEqual(req_info.url_template, "{+baseurl}/posts/{post%2Did}")

        with self.assertRaises(TypeError):
            post_item.to_patch_request_information(None)

    def test_item_to_delete_request_information(self):
        post_item = self.client.posts.by_post_id(5)
        req_info = post_item.to_delete_request_information()
        self.assertEqual(req_info.http_method, Method.DELETE)
        self.assertEqual(req_info.url_template, "{+baseurl}/posts/{post%2Did}")


class TestPostModel(unittest.TestCase):

    def test_post_model_properties_and_serialization(self):
        post = Post()
        post.user_id = 42
        post.id = 5
        post.title = "Testing Kiota"
        post.body = "Hello world!"
        post.additional_data["custom"] = "value"

        self.assertEqual(post.user_id, 42)
        self.assertEqual(post.id, 5)
        self.assertEqual(post.title, "Testing Kiota")
        self.assertEqual(post.body, "Hello world!")
        self.assertEqual(post.additional_data.get("custom"), "value")

        fields = post.get_field_deserializers()
        for expected_field in ["body", "id", "title", "userId"]:
            self.assertIn(expected_field, fields)

        writer = JsonSerializationWriter()
        post.serialize(writer)
        serialized = writer.get_serialized_content().decode("utf-8")
        self.assertIn('"userId": 42', serialized)
        self.assertIn('"id": 5', serialized)
        self.assertIn('"title": "Testing Kiota"', serialized)
        self.assertIn('"body": "Hello world!"', serialized)


if __name__ == "__main__":
    unittest.main()
