require './models'
require 'microsoft_kiota_abstractions'
require_relative './graph_c::_models::_entity'

module GraphC::Models
    class Extension < GraphC::Models::Entity
        include MicrosoftKiotaAbstractions::Parsable
        ## 
        ## Creates a new instance of the appropriate class based on discriminator value
        ## @param parseNode The parse node to use to read the discriminator value and create the object
        ## @return a extension
        ## 
        def create_from_discriminator_value(parse_node) 
            return nil;
        end
        ## 
        ## The deserialization information for the current model
        ## @return a i_dictionary
        ## 
        def get_field_deserializers() 
            return super.merge({
            })
        end
        ## 
        ## Serializes information the current object
        ## @param writer Serialization writer to use to serialize this model
        ## @return a void
        ## 
        def serialize(writer) 
            super
        end
    end
end
