# Copyright (c) Microsoft Corporation.
# Licensed under the MIT License.

# <ProgramSnippet>
# frozen_string_literal: true
require 'microsoft_kiota_serialization_json'
require 'microsoft_kiota_abstractions'
require 'microsoft_kiota_authentication_oauth'
require 'microsoft_kiota_faraday'
require_relative './client/get_user_api_client'

client_id = 'CLIENT_ID'
client_secret = 'CLIENT_SECRET'
auth_code = 'AUTH_CODE'

tenant_id = 'common'
redirect_uri = 'http://localhost'

# The auth provider will only authorize requests to
# the allowed hosts, in this case Microsoft Graph
allowed_hosts = ['graph.microsoft.com']
graph_scopes = ['User.Read']
token_request_context = MicrosoftKiotaAuthenticationOAuth::AuthorizationCodeContext
  .new(tenant_id, client_id, client_secret, redirect_uri, auth_code)

auth_provider = MicrosoftKiotaAuthenticationOAuth::OAuthAuthenticationProvider
  .new(token_request_context, allowed_hosts, graph_scopes)

request_adapter = MicrosoftKiotaFaraday::FaradayRequestAdapter
  .new(auth_provider,
       MicrosoftKiotaSerializationJson::JsonParseNodeFactory.new,
       MicrosoftKiotaSerializationJson::JsonSerializationWriterFactory.new)

client = GetUser::GetUserApiClient.new(request_adapter)

me = client.me().get().resume

puts "Hi! My name is #{me.display_name}, and my ID is #{me.id}."
# </ProgramSnippet>
