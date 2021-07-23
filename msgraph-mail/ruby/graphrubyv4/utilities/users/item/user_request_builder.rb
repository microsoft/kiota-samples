require 'microsoft_kiota_abstractions'
require_relative '../inference_classification/inference_classification_request_builder'
require_relative '../mail_folders/mail_folders_request_builder'
require_relative '../messages/messages_request_builder'

module Graphrubyv4::Utilities::Users::Item
    ## 
    # Builds and executes requests for operations under \users\{user-id}
    class UserRequestBuilder
        
        ## 
        # Current path for the request
        @current_path
        ## 
        # The http core service to use to execute the requests.
        @http_core
        def inference_classification()
            InferenceClassificationRequestBuilder.new(@current_path + @path_segment , @http_core)
        end
        def mail_folders()
            MailFoldersRequestBuilder.new(@current_path + @path_segment , @http_core)
        end
        def messages()
            MessagesRequestBuilder.new(@current_path + @path_segment , @http_core)
        end
        ## 
        # Path segment to use to build the URL for the current request builder
        @path_segment
        ## 
        ## Instantiates a new UserRequestBuilder and sets the default values.
        ## @param currentPath Current path for the request
        ## @param httpCore The http core service to use to execute the requests.
        ## @return a void
        ## 
        def initialize(current_path, http_core) 
            @path_segment = ""
            @http_core = http_core
            @current_path = current_path
        end
    end
end
