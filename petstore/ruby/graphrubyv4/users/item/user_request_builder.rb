require 'microsoft_kiota_abstractions'
require_relative '../inference_classification/inference_classification_request_builder'
require_relative '../mail_folders/item/mail_folder_request_builder'
require_relative '../mail_folders/mail_folders_request_builder'
require_relative '../messages/item/message_request_builder'
require_relative '../messages/messages_request_builder'
require_relative '../users'
require_relative './item'

module Graphrubyv4::Users::Item
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
            return Graphrubyv4::Users::InferenceClassification::InferenceClassificationRequestBuilder.new(@current_path + @path_segment , @http_core)
        end
        def mail_folders()
            return Graphrubyv4::Users::MailFolders::MailFoldersRequestBuilder.new(@current_path + @path_segment , @http_core)
        end
        def messages()
            return Graphrubyv4::Users::Messages::MessagesRequestBuilder.new(@current_path + @path_segment , @http_core)
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
        ## 
        ## Gets an item from the graphrubyv4.users.mailFolders collection
        ## @param id Unique identifier of the item
        ## @return a mail_folder_request_builder
        ## 
        def mail_folders_by_id(id) 
            return Graphrubyv4::Users::MailFolders::Item::MailFolderRequestBuilder.new(@current_path + @path_segment  + "/mailFolders/" + id, @http_core)
        end
        ## 
        ## Gets an item from the graphrubyv4.users.messages collection
        ## @param id Unique identifier of the item
        ## @return a message_request_builder
        ## 
        def messages_by_id(id) 
            return Graphrubyv4::Users::Messages::Item::MessageRequestBuilder.new(@current_path + @path_segment  + "/messages/" + id, @http_core)
        end
    end
end
