require 'microsoft_kiota_abstractions'
require_relative '../entity'
require_relative '../users'
require_relative './inference_classification'

module Graphrubyv4::Users::InferenceClassification
    class InferenceClassification < Graphrubyv4::Users::Entity
        include MicrosoftKiotaAbstractions::Parsable
        ## 
        # A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.
        @overrides
        ## 
        ## Gets the overrides property value. A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.
        ## @return a inference_classification_override
        ## 
        def  overrides
            return @overrides
        end
        ## 
        ## The deserialization information for the current model
        ## @return a i_dictionary
        ## 
        def get_field_deserializers() 
            return super.merge({
                "overrides" => lambda {|o, n| o.overrides = n.get_collection_of_object_values(Graphrubyv4::Users::InferenceClassification::InferenceClassificationOverride) },
            })
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
        ## 
        ## Sets the overrides property value. A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.
        ## @param value Value to set for the overrides property.
        ## @return a void
        ## 
        def  overrides=(overrides)
            @overrides = overrides
        end
    end
end
