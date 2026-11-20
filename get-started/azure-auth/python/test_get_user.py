# Copyright (c) Microsoft Corporation.
# Licensed under the MIT License.

import unittest
from kiota_abstractions.authentication.anonymous_authentication_provider import (
    AnonymousAuthenticationProvider,
)
from kiota_abstractions.method import Method
from kiota_http.httpx_request_adapter import HttpxRequestAdapter
from kiota_serialization_json.json_serialization_writer import JsonSerializationWriter
from client.get_user_api_client import GetUserApiClient
from client.models.user import User


class TestGetUserApiClient(unittest.TestCase):

    def setUp(self):
        self.auth_provider = AnonymousAuthenticationProvider()
        self.request_adapter = HttpxRequestAdapter(self.auth_provider)
        self.client = GetUserApiClient(self.request_adapter)

    def test_client_initialization(self):
        self.assertIsNotNone(self.client)
        self.assertIsNotNone(self.client.me)
        self.assertEqual(
            self.request_adapter.base_url,
            "https://graph.microsoft.com/v1.0",
        )

    def test_me_to_get_request_information(self):
        req_info = self.client.me.to_get_request_information()
        self.assertEqual(req_info.http_method, Method.GET)
        self.assertEqual(req_info.url_template, "{+baseurl}/me")
        self.assertIn("application/json", req_info.headers.get("Accept"))


class TestUserModel(unittest.TestCase):

    def test_user_model_properties_and_serialization(self):
        user = User()
        user.id = "user-123"
        user.display_name = "Jane Doe"
        user.additional_data["custom"] = "value"

        self.assertEqual(user.id, "user-123")
        self.assertEqual(user.display_name, "Jane Doe")
        self.assertEqual(user.additional_data.get("custom"), "value")

        fields = user.get_field_deserializers()
        for expected_field in ["displayName", "id"]:
            self.assertIn(expected_field, fields)

        writer = JsonSerializationWriter()
        user.serialize(writer)
        serialized = writer.get_serialized_content().decode("utf-8")
        self.assertIn('"displayName": "Jane Doe"', serialized)
        self.assertIn('"id": "user-123"', serialized)


if __name__ == "__main__":
    unittest.main()
