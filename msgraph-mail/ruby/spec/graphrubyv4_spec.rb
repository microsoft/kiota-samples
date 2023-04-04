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
    attachment = Graphrubyv4::Models::Attachment.new()
    expect(attachment.get_field_deserializers()).not_to be nil
  end
  
  it "can use enum" do 
    bodyType = Graphrubyv4::Models::BodyType
    expect(bodyType[:Text]).to be :Text
  end
  
  it "can build api_client" do 
    auth_provider = AuthenticationProvider.new()
    api = Graphrubyv4::ApiClient.new(MicrosoftKiotaFaraday::FaradayRequestAdapter.new(auth_provider))
    expect(api).to_not be nil
  end

  it "can get request" do 
    auth_provider = AuthenticationProvider.new()
    client = Graphrubyv4::ApiClient.new(MicrosoftKiotaFaraday::FaradayRequestAdapter.new(auth_provider))
    begin
      messageRequestConfiguration = MicrosoftKiotaAbstractions::RequestConfiguration.new
      messageRequestConfiguration.query_parameters = Graphrubyv4::Users::Item::Messages::MessagesRequestBuilder::MessagesRequestBuilderGetQueryParameters.new
      messageRequestConfiguration.query_parameters.select = "id,subject,toRecipients"
      messageResponses = client.users_by_id("vincent@biret365.onmicrosoft.com").messages().get(messageRequestConfiguration).resume

      expect(messageResponses.value).to_not be nil
      expect(messageResponses.value.first).to_not be nil
      expect(messageResponses.value.first.id).to_not be nil
      expect(messageResponses.value.first.to_recipients.first.email_address.address).to_not be nil
    rescue => any_error
      expect (any_error.error.message).to be nil
    end
  end
end
