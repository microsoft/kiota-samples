# frozen_string_literal: true

require "microsoft_kiota_abstractions"
require "microsoft_kiota_serialization_form"
require "microsoft_kiota_serialization_json"
require "microsoft_kiota_serialization_multipart"
require "microsoft_kiota_serialization_text"
require_relative "./kiota_posts"
require_relative "./posts/posts_request_builder"

module KiotaPosts
  # The main entry point of the SDK, exposes the configuration and the fluent API.
  class PostsClient < MicrosoftKiotaAbstractions::BaseRequestBuilder
    # The posts property
    def posts
      KiotaPosts::Posts::PostsRequestBuilder.new(@path_parameters, @request_adapter)
    end

    # Instantiates a new PostsClient and sets the default values.
    # @param request_adapter The request adapter to use to execute the requests.
    # @return a void
    def initialize(request_adapter)
      super({}, request_adapter, "{+baseurl}")
      MicrosoftKiotaAbstractions::ApiClientBuilder.register_default_serializer(MicrosoftKiotaSerializationJson::JsonSerializationWriterFactory)
      MicrosoftKiotaAbstractions::ApiClientBuilder.register_default_serializer(MicrosoftKiotaSerializationText::TextSerializationWriterFactory)
      MicrosoftKiotaAbstractions::ApiClientBuilder.register_default_serializer(MicrosoftKiotaSerializationForm::FormSerializationWriterFactory)
      MicrosoftKiotaAbstractions::ApiClientBuilder.register_default_serializer(MicrosoftKiotaSerializationMultipart::MultipartSerializationWriterFactory)
      MicrosoftKiotaAbstractions::ApiClientBuilder.register_default_deserializer(MicrosoftKiotaSerializationJson::JsonParseNodeFactory)
      MicrosoftKiotaAbstractions::ApiClientBuilder.register_default_deserializer(MicrosoftKiotaSerializationText::TextParseNodeFactory)
      MicrosoftKiotaAbstractions::ApiClientBuilder.register_default_deserializer(MicrosoftKiotaSerializationForm::FormParseNodeFactory)
      @request_adapter.set_base_url("https://jsonplaceholder.typicode.com") if @request_adapter.get_base_url.nil? || @request_adapter.get_base_url.empty?
      @path_parameters["baseurl"] = @request_adapter.get_base_url
    end
  end
end
