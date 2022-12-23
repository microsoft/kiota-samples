require 'microsoft_kiota_abstractions'
require_relative '../models'
require_relative './o_data_errors'

module Graphrubyv4::Models::ODataErrors
    class MainError
        include MicrosoftKiotaAbstractions::AdditionalDataHolder, MicrosoftKiotaAbstractions::Parsable
        ## 
        # Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        @additional_data
        ## 
        # The code property
        @code
        ## 
        # The details property
        @details
        ## 
        # The structure of this object is service-specific
        @innererror
        ## 
        # The message property
        @message
        ## 
        # The target property
        @target
        ## 
        ## Gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        ## @return a i_dictionary
        ## 
        def additional_data
            return @additional_data
        end
        ## 
        ## Sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        ## @param value Value to set for the AdditionalData property.
        ## @return a void
        ## 
        def additional_data=(value)
            @additional_data = value
        end
        ## 
        ## Gets the code property value. The code property
        ## @return a string
        ## 
        def code
            return @code
        end
        ## 
        ## Sets the code property value. The code property
        ## @param value Value to set for the code property.
        ## @return a void
        ## 
        def code=(value)
            @code = value
        end
        ## 
        ## Instantiates a new MainError and sets the default values.
        ## @return a void
        ## 
        def initialize()
            @additional_data = Hash.new
        end
        ## 
        ## Creates a new instance of the appropriate class based on discriminator value
        ## @param parseNode The parse node to use to read the discriminator value and create the object
        ## @return a main_error
        ## 
        def self.create_from_discriminator_value(parse_node)
            raise StandardError, 'parse_node cannot be null' if parse_node.nil?
            return MainError.new
        end
        ## 
        ## Gets the details property value. The details property
        ## @return a error_details
        ## 
        def details
            return @details
        end
        ## 
        ## Sets the details property value. The details property
        ## @param value Value to set for the details property.
        ## @return a void
        ## 
        def details=(value)
            @details = value
        end
        ## 
        ## The deserialization information for the current model
        ## @return a i_dictionary
        ## 
        def get_field_deserializers()
            return {
                "code" => lambda {|n| @code = n.get_string_value() },
                "details" => lambda {|n| @details = n.get_collection_of_object_values(lambda {|pn| Graphrubyv4::Models::ODataErrors::ErrorDetails.create_from_discriminator_value(pn) }) },
                "innererror" => lambda {|n| @innererror = n.get_object_value(lambda {|pn| Graphrubyv4::Models::ODataErrors::InnerError.create_from_discriminator_value(pn) }) },
                "message" => lambda {|n| @message = n.get_string_value() },
                "target" => lambda {|n| @target = n.get_string_value() },
            }
        end
        ## 
        ## Gets the innererror property value. The structure of this object is service-specific
        ## @return a inner_error
        ## 
        def innererror
            return @innererror
        end
        ## 
        ## Sets the innererror property value. The structure of this object is service-specific
        ## @param value Value to set for the innererror property.
        ## @return a void
        ## 
        def innererror=(value)
            @innererror = value
        end
        ## 
        ## Gets the message property value. The message property
        ## @return a string
        ## 
        def message
            return @message
        end
        ## 
        ## Sets the message property value. The message property
        ## @param value Value to set for the message property.
        ## @return a void
        ## 
        def message=(value)
            @message = value
        end
        ## 
        ## Serializes information the current object
        ## @param writer Serialization writer to use to serialize this model
        ## @return a void
        ## 
        def serialize(writer)
            raise StandardError, 'writer cannot be null' if writer.nil?
            writer.write_string_value("code", @code)
            writer.write_collection_of_object_values("details", @details)
            writer.write_object_value("innererror", @innererror)
            writer.write_string_value("message", @message)
            writer.write_string_value("target", @target)
            writer.write_additional_data(@additional_data)
        end
        ## 
        ## Gets the target property value. The target property
        ## @return a string
        ## 
        def target
            return @target
        end
        ## 
        ## Sets the target property value. The target property
        ## @param value Value to set for the target property.
        ## @return a void
        ## 
        def target=(value)
            @target = value
        end
    end
end
