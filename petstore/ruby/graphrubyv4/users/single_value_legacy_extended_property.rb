require 'microsoft_kiota_abstractions'
require_relative './entity'
require_relative './users'

module Graphrubyv4::Users
    class SingleValueLegacyExtendedProperty < Graphrubyv4::Users::Entity
        include MicrosoftKiotaAbstractions::Parsable
        ## 
        # A property value.
        @value
        ## 
        ## Gets the value property value. A property value.
        ## @return a string
        ## 
        def  value
            return @value
        end
        ## 
        ## The deserialization information for the current model
        ## @return a i_dictionary
        ## 
        def get_field_deserializers() 
            return super.merge({
                "value" => lambda {|o, n| o.value = n.get_string_value() },
            })
        end
        ## 
        ## Serializes information the current object
        ## @param writer Serialization writer to use to serialize this model
        ## @return a void
        ## 
        def serialize(writer) 
            super
            writer.write_string_value("value", @value)
        end
        ## 
        ## Sets the value property value. A property value.
        ## @param value Value to set for the value property.
        ## @return a void
        ## 
        def  value=(value)
            @value = value
        end
    end
end
