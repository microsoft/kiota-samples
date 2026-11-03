# frozen_string_literal: true

require "microsoft_kiota_abstractions"
require "stringio"
require_relative "../../models/models"

module KiotaPosts
  module Posts
    module Item
      # Builds and executes requests for operations under #posts#{post-id}
      class PostItemRequestBuilder < MicrosoftKiotaAbstractions::BaseRequestBuilder
        # Instantiates a new PostItemRequestBuilder and sets the default values.
        # @param path_parameters Path parameters for the request
        # @param request_adapter The request adapter to use to execute the requests.
        # @return a void
        def initialize(path_parameters, request_adapter)
          super(path_parameters, request_adapter, "{+baseurl}/posts/{post%2Did}")
        end

        # Delete post
        # @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
        # @return a Fiber of string_i_o
        def delete(request_configuration = nil)
          request_info = to_delete_request_information(request_configuration)
          @request_adapter.send_primitive_async(request_info, StringIO, nil)
        end

        # Get post by ID
        # @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
        # @return a Fiber of post
        def get(request_configuration = nil)
          request_info = to_get_request_information(request_configuration)
          @request_adapter.send_async(request_info, ->(pn) { KiotaPosts::Models::Post.create_from_discriminator_value(pn) }, nil)
        end

        # Update post
        # @param body The request body
        # @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
        # @return a Fiber of post
        def patch(body, request_configuration = nil)
          raise StandardError, "body cannot be null" if body.nil?

          request_info = to_patch_request_information(body, request_configuration)
          @request_adapter.send_async(request_info, ->(pn) { KiotaPosts::Models::Post.create_from_discriminator_value(pn) }, nil)
        end

        # Delete post
        # @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
        # @return a request_information
        def to_delete_request_information(request_configuration = nil)
          request_info = MicrosoftKiotaAbstractions::RequestInformation.new
          unless request_configuration.nil?
            request_info.add_headers_from_raw_object(request_configuration.headers)
            request_info.add_request_options(request_configuration.options)
          end
          request_info.url_template = @url_template
          request_info.path_parameters = @path_parameters
          request_info.http_method = :DELETE
          request_info
        end

        # Get post by ID
        # @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
        # @return a request_information
        def to_get_request_information(request_configuration = nil)
          request_info = MicrosoftKiotaAbstractions::RequestInformation.new
          unless request_configuration.nil?
            request_info.add_headers_from_raw_object(request_configuration.headers)
            request_info.add_request_options(request_configuration.options)
          end
          request_info.url_template = @url_template
          request_info.path_parameters = @path_parameters
          request_info.http_method = :GET
          request_info.headers.try_add("Accept", "application/json")
          request_info
        end

        # Update post
        # @param body The request body
        # @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
        # @return a request_information
        def to_patch_request_information(body, request_configuration = nil)
          raise StandardError, "body cannot be null" if body.nil?

          request_info = MicrosoftKiotaAbstractions::RequestInformation.new
          unless request_configuration.nil?
            request_info.add_headers_from_raw_object(request_configuration.headers)
            request_info.add_request_options(request_configuration.options)
          end
          request_info.set_content_from_parsable(@request_adapter, "application/json", body)
          request_info.url_template = @url_template
          request_info.path_parameters = @path_parameters
          request_info.http_method = :PATCH
          request_info.headers.try_add("Accept", "application/json")
          request_info
        end

        # Returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
        # @param raw_url The raw URL to use for the request builder.
        # @return a post_item_request_builder
        def with_url(raw_url)
          raise StandardError, "raw_url cannot be null" if raw_url.nil?

          PostItemRequestBuilder.new(raw_url, @request_adapter)
        end
      end
    end
  end
end
