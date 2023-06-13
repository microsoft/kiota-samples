require 'microsoft_kiota_abstractions'
require_relative '../get_user'
require_relative '../models/user'
require_relative './me'

module GetUser
    module Me
        ## 
        # Builds and executes requests for operations under \me
        class MeRequestBuilder < MicrosoftKiotaAbstractions::BaseRequestBuilder
            
            ## 
            ## Instantiates a new MeRequestBuilder and sets the default values.
            ## @param path_parameters Path parameters for the request
            ## @param request_adapter The request adapter to use to execute the requests.
            ## @return a void
            ## 
            def initialize(path_parameters, request_adapter)
                super(path_parameters, request_adapter, "{+baseurl}/me")
            end
            def get(request_configuration=nil)
                request_info = self.to_get_request_information(
                    request_configuration
                )
                return @request_adapter.send_async(request_info, lambda {|pn| GetUser::Models::User.create_from_discriminator_value(pn) }, nil)
            end
            def to_get_request_information(request_configuration=nil)
                request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
                request_info.url_template = @url_template
                request_info.path_parameters = @path_parameters
                request_info.http_method = :GET
                request_info.headers.add('Accept', 'application/json')
                unless request_configuration.nil?
                    request_info.add_headers_from_raw_object(request_configuration.headers)
                    request_info.add_request_options(request_configuration.options)
                end
                return request_info
            end
        end
    end
end
