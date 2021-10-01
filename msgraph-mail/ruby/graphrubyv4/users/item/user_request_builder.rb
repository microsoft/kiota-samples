require '../users'
require './item'
require 'microsoft_kiota_abstractions'
require_relative './inference_classification/inference_classification_request_builder'
require_relative './mail_folders/item/mail_folder_request_builder'
require_relative './mail_folders/mail_folders_request_builder'
require_relative './messages/item/message_request_builder'
require_relative './messages/messages_request_builder'

module Graphrubyv4::Users::Item
    ## 
    # Builds and executes requests for operations under \users\{user-id}
    class UserRequestBuilder
        
        ## 
        # Current path for the request
        @current_path
        def inference_classification()
            return Graphrubyv4::Users::Item::InferenceClassification::InferenceClassificationRequestBuilder.new(@current_path + @path_segment , @request_adapter, false)
        end
        ## 
        # Whether the current path is a raw URL
        @is_raw_url
        def mail_folders()
            return Graphrubyv4::Users::Item::MailFolders::MailFoldersRequestBuilder.new(@current_path + @path_segment , @request_adapter, false)
        end
        def messages()
            return Graphrubyv4::Users::Item::Messages::MessagesRequestBuilder.new(@current_path + @path_segment , @request_adapter, false)
        end
        ## 
        # Path segment to use to build the URL for the current request builder
        @path_segment
        ## 
        # The http core service to use to execute the requests.
        @request_adapter
        ## 
        ## Instantiates a new UserRequestBuilder and sets the default values.
        ## @param currentPath Current path for the request
        ## @param isRawUrl Whether the current path is a raw URL
        ## @param requestAdapter The http core service to use to execute the requests.
        ## @return a void
        ## 
        def initialize(current_path, request_adapter, is_raw_url=true) 
            @path_segment = ""
            @request_adapter = request_adapter
            @current_path = current_path
            @is_raw_url = is_raw_url
        end
        ## 
        ## Gets an item from the graphrubyv4.users.item.mailFolders.item collection
        ## @param id Unique identifier of the item
        ## @return a mail_folder_request_builder
        ## 
        def mail_folders_by_id(id) 
            return Graphrubyv4::Users::Item::MailFolders::Item::MailFolderRequestBuilder.new(@current_path + @path_segment  + "/mailFolders/" + id, @request_adapter, false)
        end
        ## 
        ## Gets an item from the graphrubyv4.users.item.messages.item collection
        ## @param id Unique identifier of the item
        ## @return a message_request_builder
        ## 
        def messages_by_id(id) 
            return Graphrubyv4::Users::Item::Messages::Item::MessageRequestBuilder.new(@current_path + @path_segment  + "/messages/" + id, @request_adapter, false)
        end
    end
end
