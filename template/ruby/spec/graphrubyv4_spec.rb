# frozen_string_literal: true

require_relative 'spec_helper'
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
end
