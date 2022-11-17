# frozen_string_literal: true
require 'microsoft_kiota_serialization'
require 'microsoft_kiota_abstractions'
require 'microsoft_kiota_authentication_oauth'
require 'microsoft_kiota_nethttplibrary'
require_relative './client/api_client'

tenant_id = 'TENANT_ID'
client_id = 'CLIENT_ID'
client_secret = 'CLIENT_SECRET'
auth_code = 'AUTH_CODE'
redirect_uri = 'REDIRECT_URI'

# The auth provider will only authorize requests to
# the allowed hosts, in this case Microsoft Graph
allowed_hosts = ['graph.microsoft.com']
graph_scopes = ['User.Read']
token_request_context = MicrosoftKiotaAuthenticationOAuth::AuthorizationCodeContext.new(tenant_id, client_id, client_secret, redirect_uri, auth_code)

auth_provider = MicrosoftKiotaAuthenticationOAuth::OAuthAuthenticationProvider.new(token_request_context, allowed_hosts, graph_scopes)

request_adapter = MicrosoftKiotaNethttplibrary::NetHttpRequestAdapter.new(auth_provider, MicrosoftKiotaSerialization::JsonParseNodeFactory.new, MicrosoftKiotaSerialization::JsonSerializationWriterFactory.new)

client = Graph::ApiClient.new(request_adapter)

me = client.me.get

puts "Hi! My name is #{me.display_name}, and my ID is #{me.id}."
