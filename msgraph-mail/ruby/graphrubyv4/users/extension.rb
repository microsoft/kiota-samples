require 'microsoft_kiota_abstractions'
require_relative './entity'
require_relative './users'

module Graphrubyv4::Users
    class Extension < Graphrubyv4::Users::Entity
        include MicrosoftKiotaAbstractions::Parsable
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
