require 'microsoft_kiota_abstractions'
require_relative './models'

module Graphrubyv4::Models
    class SizeRange
        include MicrosoftKiotaAbstractions::AdditionalDataHolder, MicrosoftKiotaAbstractions::Parsable
        ## 
        # Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        @additional_data
        ## 
        # The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
        @maximum_size
        ## 
        # The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
        @minimum_size
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
        def additional_data=(additionalData)
            @additional_data = additionalData
        end
        ## 
        ## Instantiates a new sizeRange and sets the default values.
        ## @return a void
        ## 
        def initialize() 
            @additional_data = Hash.new
        end
        ## 
        ## Creates a new instance of the appropriate class based on discriminator value
        ## @param parseNode The parse node to use to read the discriminator value and create the object
        ## @return a size_range
        ## 
        def create_from_discriminator_value(parse_node) 
            return nil;
        end
        ## 
        ## The deserialization information for the current model
        ## @return a i_dictionary
        ## 
        def get_field_deserializers() 
            return {
                "maximumSize" => lambda {|n| @maximum_size = n.get_number_value() },
                "minimumSize" => lambda {|n| @minimum_size = n.get_number_value() },
            }
        end
        ## 
        ## Gets the maximumSize property value. The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
        ## @return a integer
        ## 
        def maximum_size
            return @maximum_size
        end
        ## 
        ## Sets the maximumSize property value. The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
        ## @param value Value to set for the maximumSize property.
        ## @return a void
        ## 
        def maximum_size=(maximumSize)
            @maximum_size = maximumSize
        end
        ## 
        ## Gets the minimumSize property value. The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
        ## @return a integer
        ## 
        def minimum_size
            return @minimum_size
        end
        ## 
        ## Sets the minimumSize property value. The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
        ## @param value Value to set for the minimumSize property.
        ## @return a void
        ## 
        def minimum_size=(minimumSize)
            @minimum_size = minimumSize
        end
        ## 
        ## Serializes information the current object
        ## @param writer Serialization writer to use to serialize this model
        ## @return a void
        ## 
        def serialize(writer) 
            writer.write_number_value("maximumSize", @maximum_size)
            writer.write_number_value("minimumSize", @minimum_size)
            writer.write_additional_data(@additional_data)
        end
    end
end
