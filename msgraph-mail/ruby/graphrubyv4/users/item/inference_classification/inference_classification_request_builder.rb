require '../../users'
require '../item'
require './inference_classification'
require 'microsoft_kiota_abstractions'
require_relative '../../../models/inference_classification'
require_relative './overrides/item/inference_classification_override_item_request_builder'
require_relative './overrides/overrides_request_builder'

module GraphC::Users::Item::InferenceClassification
    ## 
    # Builds and executes requests for operations under \users\{user-id}\inferenceClassification
    class InferenceClassificationRequestBuilder
        
        ## 
        # The overrides property
        def overrides()
            return GraphC::Users::Item::InferenceClassification::Overrides::OverridesRequestBuilder.new(@path_parameters, @request_adapter)
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
            @url_template = "{+baseurl}/users/{user%2Did}/inferenceClassification{?%24select}"
            @request_adapter = request_adapter
            if path_parameters.is_a? String
                path_parameters = { "request-raw-url" => path_parameters }
            end
            @path_parameters = path_parameters
        end
        ## 
        ## Delete navigation property inferenceClassification for users
        ## @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
        ## @return a request_information
        ## 
        def create_delete_request_information(request_configuration=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
            request_info.url_template = @url_template
            request_info.path_parameters = @path_parameters
            request_info.http_method = :DELETE
            request_info.set_headers_from_raw_object(requestConfiguration.headers)
            return request_info;
        end
        ## 
        ## Get inferenceClassification from users
        ## @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
        ## @return a request_information
        ## 
        def create_get_request_information(request_configuration=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
            request_info.url_template = @url_template
            request_info.path_parameters = @path_parameters
            request_info.http_method = :GET
            request_info.set_headers_from_raw_object(requestConfiguration.headers)
            request_info.set_query_string_parameters_from_raw_object(requestConfiguration.query_parameters)
            return request_info;
        end
        ## 
        ## Update the navigation property inferenceClassification in users
        ## @param body 
        ## @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
        ## @return a request_information
        ## 
        def create_patch_request_information(body, request_configuration=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
            request_info.url_template = @url_template
            request_info.path_parameters = @path_parameters
            request_info.http_method = :PATCH
            request_info.set_headers_from_raw_object(requestConfiguration.headers)
            request_info.set_content_from_parsable(self.request_adapter, "application/json", body)
            return request_info;
        end
        ## 
        ## Delete navigation property inferenceClassification for users
        ## @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of void
        ## 
        def delete(request_configuration=nil, response_handler=nil) 
            request_info = self.create_delete_request_information(
                request_configuration
            )
            return @http_core.send_async(request_info, nil, response_handler)
        end
        ## 
        ## Get inferenceClassification from users
        ## @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of inference_classification
        ## 
        def get(request_configuration=nil, response_handler=nil) 
            request_info = self.create_get_request_information(
                request_configuration
            )
            return @http_core.send_async(request_info, GraphC::Users::Item::InferenceClassification::InferenceClassification, response_handler)
        end
        ## 
        ## Gets an item from the GraphC.users.item.inferenceClassification.overrides.item collection
        ## @param id Unique identifier of the item
        ## @return a inference_classification_override_item_request_builder
        ## 
        def overrides_by_id(id) 
            url_tpl_params = @path_parameters.clone
            url_tpl_params["inferenceClassificationOverride%2Did"] = id
            return GraphC::Users::Item::InferenceClassification::Overrides::Item::InferenceClassificationOverrideItemRequestBuilder.new(url_tpl_params, @request_adapter)
        end
        ## 
        ## Update the navigation property inferenceClassification in users
        ## @param body 
        ## @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of void
        ## 
        def patch(body, request_configuration=nil, response_handler=nil) 
            request_info = self.create_patch_request_information(
                body, request_configuration
            )
            return @http_core.send_async(request_info, nil, response_handler)
        end
        require '../../users'
        require '../item'
        require './inference_classification'

        ## 
        # Configuration for the request such as headers, query parameters, and middleware options.
        class InferenceClassificationRequestBuilderDeleteRequestConfiguration
            
            ## 
            # Request headers
            @headers
            ## 
            # Request options
            @options
        end
        require '../../users'
        require '../item'
        require './inference_classification'

        ## 
        # Get inferenceClassification from users
        class InferenceClassificationRequestBuilderGetQueryParameters
            
            ## 
            # Select properties to be returned
            @select
        end
        require '../../users'
        require '../item'
        require './inference_classification'

        ## 
        # Configuration for the request such as headers, query parameters, and middleware options.
        class InferenceClassificationRequestBuilderGetRequestConfiguration
            
            ## 
            # Request headers
            @headers
            ## 
            # Request options
            @options
            ## 
            # Request query parameters
            @query_parameters
        end
        require '../../users'
        require '../item'
        require './inference_classification'

        ## 
        # Configuration for the request such as headers, query parameters, and middleware options.
        class InferenceClassificationRequestBuilderPatchRequestConfiguration
            
            ## 
            # Request headers
            @headers
            ## 
            # Request options
            @options
        end
    end
end
