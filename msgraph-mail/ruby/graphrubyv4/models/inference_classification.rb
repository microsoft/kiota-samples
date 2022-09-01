require 'microsoft_kiota_abstractions'
require_relative './entity'
require_relative './models'

module Graphrubyv4::Models
    class InferenceClassification < Entity
        include MicrosoftKiotaAbstractions::Parsable
        ## 
        # A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.
        @overrides
        ## 
        ## Creates a new instance of the appropriate class based on discriminator value
        ## @param parseNode The parse node to use to read the discriminator value and create the object
        ## @return a inference_classification
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
                "overrides" => lambda {|n| @overrides = n.get_collection_of_object_values(Graphrubyv4::Models::InferenceClassificationOverride) },
            })
        end
        ## 
        ## Gets the overrides property value. A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.
        ## @return a inference_classification_override
        ## 
        def  overrides
            return @overrides
        end
        ## 
        ## Sets the overrides property value. A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.
        ## @param value Value to set for the overrides property.
        ## @return a void
        ## 
        def  overrides=(overrides)
            @overrides = overrides
        end
        ## 
        ## Serializes information the current object
        ## @param writer Serialization writer to use to serialize this model
        ## @return a void
        ## 
        def serialize(writer) 
            super
            writer.write_collection_of_object_values("overrides", @overrides)
        end
    end
end
