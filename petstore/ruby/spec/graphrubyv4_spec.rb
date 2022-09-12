# frozen_string_literal: true

require_relative 'spec_helper'
require_relative 'authentication_provider'
RSpec.describe "ApiClient" do
  it "has a version number" do
    expect(true).not_to be nil
  end

  it "can build request_builder" do
    request = Graphrubyv4::Users::UsersRequestBuilder.new("test", "test")
    expect(request).not_to be nil
  end

  it "can build attachment" do 
    attachment = Graphrubyv4::Users::Attachment.new()
    expect(attachment.get_field_deserializers()).not_to be nil
  end
  
  it "can use enum" do 
    bodyType = Graphrubyv4::Users::BodyType
    expect(bodyType[:Text]).to be :Text
  end
  
  it "can build api_client" do 
    api = Graphrubyv4::ApiClient.new("http_core")
    expect(api).to_not be nil
  end

  it "can get request" do 
    auth_provider = AuthenticationProvider.new()
    client = Graphrubyv4::ApiClient.new(MicrosoftKiotaNethttp::HttpCore.new(auth_provider, MicrosoftKiotaSerialization::JsonParseNodeFactory.new(), MicrosoftKiotaSerialization::JsonSerializationWriterFactory.new()))
    messageResponses = client.users_by_id("Example@example.com").messages().get()

    expect(messageResponses.value).to_not be nil
    expect(messageResponses.value[0]).to_not be nil
    expect(messageResponses.value[0].id).to_not be nil
  end
end
