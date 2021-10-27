require '../../users'
require '../item'
require './inference_classification'
require 'microsoft_kiota_abstractions'
require_relative '../../../models/microsoft/graph/inference_classification'
require_relative './overrides/item/inference_classification_override_request_builder'
require_relative './overrides/overrides_request_builder'

module Graphrubyv4::Users::Item::InferenceClassification
    ## 
    # Builds and executes requests for operations under \users\{user-id}\inferenceClassification
    class InferenceClassificationRequestBuilder
        
        def overrides()
            return Graphrubyv4::Users::Item::InferenceClassification::Overrides::OverridesRequestBuilder.new(@path_parameters, @request_adapter)
        end
        ## 
        # Path parameters for the request
        @path_parameters
        ## 
        # The request adapter to use to execute the requests.
        @request_adapter
        ## 
        # Url template to use to build the URL for the current request builder
        @url_template
        ## 
        ## Instantiates a new InferenceClassificationRequestBuilder and sets the default values.
        ## @param pathParameters Path parameters for the request
        ## @param requestAdapter The request adapter to use to execute the requests.
        ## @return a void
        ## 
        def initialize(path_parameters, request_adapter) 
            @url_template = "https://graph.microsoft.com/v1.0/users/{user_id}/inferenceClassification{?select,expand}"
            @request_adapter = request_adapter
            if path_parameters.is_a? String
                path_parameters = { "request-raw-url" => path_parameters }
            end
            @path_parameters = path_parameters
        end
        ## 
        ## Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
        ## @param h Request headers
        ## @param o Request options
        ## @return a request_information
        ## 
        def create_delete_request_information(h=nil, o=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
            request_info.url_template = @url_template
            request_info.path_parameters = @path_parameters
            request_info.http_method = :DELETE
            request_info.set_headers_from_raw_object(h)
            return request_info;
        end
        ## 
        ## Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
        ## @param h Request headers
        ## @param o Request options
        ## @param q Request query parameters
        ## @return a request_information
        ## 
        def create_get_request_information(q=nil, h=nil, o=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
            request_info.url_template = @url_template
            request_info.path_parameters = @path_parameters
            request_info.http_method = :GET
            request_info.set_headers_from_raw_object(h)
            request_info.set_query_string_parameters_from_raw_object(q)
            return request_info;
        end
        ## 
        ## Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
        ## @param body 
        ## @param h Request headers
        ## @param o Request options
        ## @return a request_information
        ## 
        def create_patch_request_information(body, h=nil, o=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
            request_info.url_template = @url_template
            request_info.path_parameters = @path_parameters
            request_info.http_method = :PATCH
            request_info.set_headers_from_raw_object(h)
            request_info.set_content_from_parsable(self.request_adapter, "application/json", body)
            return request_info;
        end
        ## 
        ## Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
        ## @param h Request headers
        ## @param o Request options
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of void
        ## 
        def delete(h=nil, o=nil, response_handler=nil) 
            request_info = self.create_delete_request_information(
                h
            )
            return @http_core.send_async(request_info, nil, response_handler)
        end
        ## 
        ## Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
        ## @param h Request headers
        ## @param o Request options
        ## @param q Request query parameters
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of inference_classification
        ## 
        def get(q=nil, h=nil, o=nil, response_handler=nil) 
            request_info = self.create_get_request_information(
                q, h
            )
            return @http_core.send_async(request_info, Graphrubyv4::Users::Item::InferenceClassification::InferenceClassification, response_handler)
        end
        ## 
        ## Gets an item from the graphrubyv4.users.item.inferenceClassification.overrides.item collection
        ## @param id Unique identifier of the item
        ## @return a inference_classification_override_request_builder
        ## 
        def overrides_by_id(id) 
            url_tpl_params = @path_parameters.clone
            url_tpl_params["inferenceClassificationOverride_id"] = id
            return Graphrubyv4::Users::Item::InferenceClassification::Overrides::Item::InferenceClassificationOverrideRequestBuilder.new(url_tpl_params, @request_adapter)
        end
        ## 
        ## Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
        ## @param body 
        ## @param h Request headers
        ## @param o Request options
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of void
        ## 
        def patch(body, h=nil, o=nil, response_handler=nil) 
            request_info = self.create_patch_request_information(
                body, h
            )
            return @http_core.send_async(request_info, nil, response_handler)
        end
    end
end
