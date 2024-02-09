require 'microsoft_kiota_abstractions'
require_relative '../graphrubyv4'
require_relative './models'

module Graphrubyv4
    module Models
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
            ## Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
            ## @return a i_dictionary
            ## 
            def additional_data
                return @additional_data
            end
            ## 
            ## Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
            ## @param value Value to set for the AdditionalData property.
            ## @return a void
            ## 
            def additional_data=(value)
                @additional_data = value
            end
            ## 
            ## Instantiates a new SizeRange and sets the default values.
            ## @return a void
            ## 
            def initialize()
                @additional_data = Hash.new
            end
            ## 
            ## Creates a new instance of the appropriate class based on discriminator value
            ## @param parse_node The parse node to use to read the discriminator value and create the object
            ## @return a size_range
            ## 
            def self.create_from_discriminator_value(parse_node)
                raise StandardError, 'parse_node cannot be null' if parse_node.nil?
                return SizeRange.new
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
            def maximum_size=(value)
                @maximum_size = value
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
            def minimum_size=(value)
                @minimum_size = value
            end
            ## 
            ## Serializes information the current object
            ## @param writer Serialization writer to use to serialize this model
            ## @return a void
            ## 
            def serialize(writer)
                raise StandardError, 'writer cannot be null' if writer.nil?
                writer.write_number_value("maximumSize", @maximum_size)
                writer.write_number_value("minimumSize", @minimum_size)
                writer.write_additional_data(@additional_data)
            end
        end
    end
end
