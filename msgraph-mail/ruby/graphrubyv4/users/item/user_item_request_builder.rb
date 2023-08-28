require 'microsoft_kiota_abstractions'
require_relative '../../graphrubyv4'
require_relative '../users'
require_relative './inference_classification/inference_classification_request_builder'
require_relative './item'
require_relative './mail_folders/mail_folders_request_builder'
require_relative './messages/messages_request_builder'

module Graphrubyv4
    module Users
        module Item
            ## 
            # Builds and executes requests for operations under \users\{user-id}
            class UserItemRequestBuilder < MicrosoftKiotaAbstractions::BaseRequestBuilder
                
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
                ## Instantiates a new UserItemRequestBuilder and sets the default values.
                ## @param path_parameters Path parameters for the request
                ## @param request_adapter The request adapter to use to execute the requests.
                ## @return a void
                ## 
                def initialize(path_parameters, request_adapter)
                    super(path_parameters, request_adapter, "{+baseurl}/users/{user%2Did}")
                end
            end
        end
    end
end
