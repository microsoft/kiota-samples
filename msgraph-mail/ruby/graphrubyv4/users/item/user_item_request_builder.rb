require 'microsoft_kiota_abstractions'
require_relative '../users'
require_relative './inference_classification/inference_classification_request_builder'
require_relative './item'
require_relative './mail_folders/item/mail_folder_item_request_builder'
require_relative './mail_folders/mail_folders_request_builder'
require_relative './messages/item/message_item_request_builder'
require_relative './messages/messages_request_builder'

module Graphrubyv4::Users::Item
    ## 
    # Builds and executes requests for operations under \users\{user-id}
    class UserItemRequestBuilder
        
        ## 
        # The inferenceClassification property
        def inference_classification()
            return Graphrubyv4::Users::Item::InferenceClassification::InferenceClassificationRequestBuilder.new(@path_parameters, @request_adapter)
        end
        ## 
        # The mailFolders property
        def mail_folders()
            return Graphrubyv4::Users::Item::MailFolders::MailFoldersRequestBuilder.new(@path_parameters, @request_adapter)
        end
        ## 
        # The messages property
        def messages()
            return Graphrubyv4::Users::Item::Messages::MessagesRequestBuilder.new(@path_parameters, @request_adapter)
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
        ## Instantiates a new UserItemRequestBuilder and sets the default values.
        ## @param pathParameters Path parameters for the request
        ## @param requestAdapter The request adapter to use to execute the requests.
        ## @return a void
        ## 
        def initialize(path_parameters, request_adapter) 
            @url_template = "{+baseurl}/users/{user%2Did}"
            @request_adapter = request_adapter
            if path_parameters.is_a? String
                path_parameters = { "request-raw-url" => path_parameters }
            end
            @path_parameters = path_parameters
        end
        ## 
        ## Gets an item from the graphrubyv4.users.item.mailFolders.item collection
        ## @param id Unique identifier of the item
        ## @return a mail_folder_item_request_builder
        ## 
        def mail_folders_by_id(id) 
            url_tpl_params = @path_parameters.clone
            url_tpl_params["mailFolder%2Did"] = id
            return Graphrubyv4::Users::Item::MailFolders::Item::MailFolderItemRequestBuilder.new(url_tpl_params, @request_adapter)
        end
        ## 
        ## Gets an item from the graphrubyv4.users.item.messages.item collection
        ## @param id Unique identifier of the item
        ## @return a message_item_request_builder
        ## 
        def messages_by_id(id) 
            url_tpl_params = @path_parameters.clone
            url_tpl_params["message%2Did"] = id
            return Graphrubyv4::Users::Item::Messages::Item::MessageItemRequestBuilder.new(url_tpl_params, @request_adapter)
        end
    end
end
