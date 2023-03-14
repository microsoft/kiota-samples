require 'microsoft_kiota_abstractions'
require_relative '../../../../../graphrubyv4'
require_relative '../../../../../models/message_rule'
require_relative '../../../../../models/message_rule_collection_response'
require_relative '../../../../users'
require_relative '../../../item'
require_relative '../../mail_folders'
require_relative '../item'
require_relative './message_rules'

module Graphrubyv4
    module Users
        module Item
            module MailFolders
                module Item
                    module MessageRules
                        ## 
                        # Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\messageRules
                        class MessageRulesRequestBuilder < MicrosoftKiotaAbstractions::BaseRequestBuilder
                            
                            ## 
                            ## Instantiates a new MessageRulesRequestBuilder and sets the default values.
                            ## @param path_parameters Path parameters for the request
                            ## @param request_adapter The request adapter to use to execute the requests.
                            ## @return a void
                            ## 
                            def initialize(path_parameters, request_adapter)
                                super(path_parameters, request_adapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/messageRules{?%24top,%24skip,%24filter,%24count,%24orderby,%24select}")
                            end
                            ## 
                            ## Get all the messageRule objects defined for the user's inbox.
                            ## @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
                            ## @return a Fiber of message_rule_collection_response
                            ## 
                            def get(request_configuration=nil)
                                request_info = self.to_get_request_information(
                                    request_configuration
                                )
                                return @request_adapter.send_async(request_info, lambda {|pn| Graphrubyv4::Models::MessageRuleCollectionResponse.create_from_discriminator_value(pn) }, nil)
                            end
                            ## 
                            ## Create a messageRule object by specifying a set of conditions and actions.  Outlook carries out those actions if an incoming message in the user's Inbox meets the specified conditions.
                            ## @param body The request body
                            ## @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
                            ## @return a Fiber of message_rule
                            ## 
                            def post(body, request_configuration=nil)
                                raise StandardError, 'body cannot be null' if body.nil?
                                request_info = self.to_post_request_information(
                                    body, request_configuration
                                )
                                return @request_adapter.send_async(request_info, lambda {|pn| Graphrubyv4::Models::MessageRule.create_from_discriminator_value(pn) }, nil)
                            end
                            ## 
                            ## Get all the messageRule objects defined for the user's inbox.
                            ## @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
                            ## @return a request_information
                            ## 
                            def to_get_request_information(request_configuration=nil)
                                request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
                                request_info.url_template = @url_template
                                request_info.path_parameters = @path_parameters
                                request_info.http_method = :GET
                                request_info.headers.add('Accept', 'application/json')
                                unless request_configuration.nil?
                                    request_info.add_headers_from_raw_object(request_configuration.headers)
                                    request_info.set_query_string_parameters_from_raw_object(request_configuration.query_parameters)
                                    request_info.add_request_options(request_configuration.options)
                                end
                                return request_info
                            end
                            ## 
                            ## Create a messageRule object by specifying a set of conditions and actions.  Outlook carries out those actions if an incoming message in the user's Inbox meets the specified conditions.
                            ## @param body The request body
                            ## @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
                            ## @return a request_information
                            ## 
                            def to_post_request_information(body, request_configuration=nil)
                                raise StandardError, 'body cannot be null' if body.nil?
                                request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
                                request_info.url_template = @url_template
                                request_info.path_parameters = @path_parameters
                                request_info.http_method = :POST
                                request_info.headers.add('Accept', 'application/json')
                                unless request_configuration.nil?
                                    request_info.add_headers_from_raw_object(request_configuration.headers)
                                    request_info.add_request_options(request_configuration.options)
                                end
                                request_info.set_content_from_parsable(self.request_adapter, "application/json", body)
                                return request_info
                            end

                            ## 
                            # Get all the messageRule objects defined for the user's inbox.
                            class MessageRulesRequestBuilderGetQueryParameters
                                
                                ## 
                                # Include count of items
                                attr_accessor :count
                                ## 
                                # Filter items by property values
                                attr_accessor :filter
                                ## 
                                # Order items by property values
                                attr_accessor :orderby
                                ## 
                                # Select properties to be returned
                                attr_accessor :select
                                ## 
                                # Skip the first n items
                                attr_accessor :skip
                                ## 
                                # Show only the first n items
                                attr_accessor :top
                                ## 
                                ## Maps the query parameters names to their encoded names for the URI template parsing.
                                ## @param original_name The original query parameter name in the class.
                                ## @return a string
                                ## 
                                def get_query_parameter(original_name)
                                    raise StandardError, 'original_name cannot be null' if original_name.nil?
                                    case original_name
                                        when "count"
                                            return "%24count"
                                        when "filter"
                                            return "%24filter"
                                        when "orderby"
                                            return "%24orderby"
                                        when "select"
                                            return "%24select"
                                        when "skip"
                                            return "%24skip"
                                        when "top"
                                            return "%24top"
                                        else
                                            return original_name
                                    end
                                end
                            end

                            ## 
                            # Configuration for the request such as headers, query parameters, and middleware options.
                            class MessageRulesRequestBuilderGetRequestConfiguration
                                
                                ## 
                                # Request headers
                                attr_accessor :headers
                                ## 
                                # Request options
                                attr_accessor :options
                                ## 
                                # Request query parameters
                                attr_accessor :query_parameters
                            end

                            ## 
                            # Configuration for the request such as headers, query parameters, and middleware options.
                            class MessageRulesRequestBuilderPostRequestConfiguration
                                
                                ## 
                                # Request headers
                                attr_accessor :headers
                                ## 
                                # Request options
                                attr_accessor :options
                            end
                        end
                    end
                end
            end
        end
    end
end
