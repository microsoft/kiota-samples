require './models'
require 'microsoft_kiota_abstractions'

module Graphrubyv4::Models
    class FollowupFlag
        include IAdditionalDataHolder, MicrosoftKiotaAbstractions::Parsable
        ## 
        # Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        @additional_data
        ## 
        # The completedDateTime property
        @completed_date_time
        ## 
        # The dueDateTime property
        @due_date_time
        ## 
        # The flagStatus property
        @flag_status
        ## 
        # The startDateTime property
        @start_date_time
        ## 
        ## Gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        ## @return a i_dictionary
        ## 
        def  additional_data
            return @additional_data
        end
        ## 
        ## Sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        ## @param value Value to set for the AdditionalData property.
        ## @return a void
        ## 
        def  additional_data=(additionalData)
            @additional_data = additionalData
        end
        ## 
        ## Gets the completedDateTime property value. The completedDateTime property
        ## @return a date_time_time_zone
        ## 
        def  completed_date_time
            return @completed_date_time
        end
        ## 
        ## Sets the completedDateTime property value. The completedDateTime property
        ## @param value Value to set for the completedDateTime property.
        ## @return a void
        ## 
        def  completed_date_time=(completedDateTime)
            @completed_date_time = completedDateTime
        end
        ## 
        ## Creates a new instance of the appropriate class based on discriminator value
        ## @param parseNode The parse node to use to read the discriminator value and create the object
        ## @return a followup_flag
        ## 
        def create_from_discriminator_value(parse_node) 
            return nil;
        end
        ## 
        ## Gets the dueDateTime property value. The dueDateTime property
        ## @return a date_time_time_zone
        ## 
        def  due_date_time
            return @due_date_time
        end
        ## 
        ## Sets the dueDateTime property value. The dueDateTime property
        ## @param value Value to set for the dueDateTime property.
        ## @return a void
        ## 
        def  due_date_time=(dueDateTime)
            @due_date_time = dueDateTime
        end
        ## 
        ## Gets the flagStatus property value. The flagStatus property
        ## @return a followup_flag_status
        ## 
        def  flag_status
            return @flag_status
        end
        ## 
        ## Sets the flagStatus property value. The flagStatus property
        ## @param value Value to set for the flagStatus property.
        ## @return a void
        ## 
        def  flag_status=(flagStatus)
            @flag_status = flagStatus
        end
        ## 
        ## The deserialization information for the current model
        ## @return a i_dictionary
        ## 
        def get_field_deserializers() 
            return {
                "completedDateTime" => lambda {|o, n| o.completed_date_time = n.get_object_value(Graphrubyv4::Models::DateTimeTimeZone) },
                "dueDateTime" => lambda {|o, n| o.due_date_time = n.get_object_value(Graphrubyv4::Models::DateTimeTimeZone) },
                "flagStatus" => lambda {|o, n| o.flag_status = n.get_enum_value(Graphrubyv4::Models::FollowupFlagStatus) },
                "startDateTime" => lambda {|o, n| o.start_date_time = n.get_object_value(Graphrubyv4::Models::DateTimeTimeZone) },
            }
        end
        ## 
        ## Serializes information the current object
        ## @param writer Serialization writer to use to serialize this model
        ## @return a void
        ## 
        def serialize(writer) 
            writer.write_object_value("completedDateTime", @completed_date_time)
            writer.write_object_value("dueDateTime", @due_date_time)
            writer.write_enum_value("flagStatus", @flag_status)
            writer.write_object_value("startDateTime", @start_date_time)
            writer.write_additional_data(@additional_data)
        end
        ## 
        ## Gets the startDateTime property value. The startDateTime property
        ## @return a date_time_time_zone
        ## 
        def  start_date_time
            return @start_date_time
        end
        ## 
        ## Sets the startDateTime property value. The startDateTime property
        ## @param value Value to set for the startDateTime property.
        ## @return a void
        ## 
        def  start_date_time=(startDateTime)
            @start_date_time = startDateTime
        end
    end
end
