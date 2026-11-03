# frozen_string_literal: true

require "microsoft_kiota_abstractions"
require_relative "../models/models"
require_relative "./item/post_item_request_builder"

module KiotaPosts
  module Posts
    # Builds and executes requests for operations under #posts
    class PostsRequestBuilder < MicrosoftKiotaAbstractions::BaseRequestBuilder
      # Gets an item from the KiotaPosts.posts.item collection
      # @param post_id key: id of post
      # @return a post_item_request_builder
      def by_post_id(post_id)
        raise StandardError, "post_id cannot be null" if post_id.nil?

        url_tpl_params = @path_parameters.clone
        url_tpl_params["post%2Did"] = post_id
        KiotaPosts::Posts::Item::PostItemRequestBuilder.new(url_tpl_params, @request_adapter)
      end

      # Instantiates a new PostsRequestBuilder and sets the default values.
      # @param path_parameters Path parameters for the request
      # @param request_adapter The request adapter to use to execute the requests.
      # @return a void
      def initialize(path_parameters, request_adapter)
        super(path_parameters, request_adapter, "{+baseurl}/posts{?title*,userId*}")
      end

      # Get posts
      # @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
      # @return a Fiber of post
      def get(request_configuration = nil)
        request_info = to_get_request_information(request_configuration)
        @request_adapter.send_collection_async(request_info, ->(pn) { KiotaPosts::Models::Post.create_from_discriminator_value(pn) }, nil)
      end

      # Create post
      # @param body The request body
      # @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
      # @return a Fiber of post
      def post(body, request_configuration = nil)
        raise StandardError, "body cannot be null" if body.nil?

        request_info = to_post_request_information(body, request_configuration)
        @request_adapter.send_async(request_info, ->(pn) { KiotaPosts::Models::Post.create_from_discriminator_value(pn) }, nil)
      end

      # Get posts
      # @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
      # @return a request_information
      def to_get_request_information(request_configuration = nil)
        request_info = MicrosoftKiotaAbstractions::RequestInformation.new
        unless request_configuration.nil?
          request_info.add_headers_from_raw_object(request_configuration.headers)
          request_info.set_query_string_parameters_from_raw_object(request_configuration.query_parameters)
          request_info.add_request_options(request_configuration.options)
        end
        request_info.url_template = @url_template
        request_info.path_parameters = @path_parameters
        request_info.http_method = :GET
        request_info.headers.try_add("Accept", "application/json")
        request_info
      end

      # Create post
      # @param body The request body
      # @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
      # @return a request_information
      def to_post_request_information(body, request_configuration = nil)
        raise StandardError, "body cannot be null" if body.nil?

        request_info = MicrosoftKiotaAbstractions::RequestInformation.new
        unless request_configuration.nil?
          request_info.add_headers_from_raw_object(request_configuration.headers)
          request_info.add_request_options(request_configuration.options)
        end
        request_info.set_content_from_parsable(@request_adapter, "application/json", body)
        request_info.url_template = @url_template
        request_info.path_parameters = @path_parameters
        request_info.http_method = :POST
        request_info.headers.try_add("Accept", "application/json")
        request_info
      end

      # Returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
      # @param raw_url The raw URL to use for the request builder.
      # @return a posts_request_builder
      def with_url(raw_url)
        raise StandardError, "raw_url cannot be null" if raw_url.nil?

        PostsRequestBuilder.new(raw_url, @request_adapter)
      end

      # Get posts
      class PostsRequestBuilderGetQueryParameters
        # Filter results by title
        attr_accessor :title

        # Filter results by user ID
        attr_accessor :user_id

        # Maps the query parameters names to their encoded names for the URI template parsing.
        # @param original_name The original query parameter name in the class.
        # @return a string
        def get_query_parameter(original_name)
          raise StandardError, "original_name cannot be null" if original_name.nil?

          case original_name
          when "user_id"
            "userId"
          else
            original_name
          end
        end
      end
    end
  end
end
