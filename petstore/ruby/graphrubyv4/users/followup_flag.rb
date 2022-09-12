require 'microsoft_kiota_abstractions'
require_relative './users'

module Graphrubyv4::Users
    class FollowupFlag
        include MicrosoftKiotaAbstractions::Parsable
        ## 
        # Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        @additional_data
        @completed_date_time
        @due_date_time
        @flag_status
        @start_date_time
        ## 
        ## Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        ## @return a i_dictionary
        ## 
        def  additional_data
            return @additional_data
        end
        ## 
        ## Gets the completedDateTime property value. 
        ## @return a date_time_time_zone
        ## 
        def  completed_date_time
            return @completed_date_time
        end
        ## 
        ## Gets the dueDateTime property value. 
        ## @return a date_time_time_zone
        ## 
        def  due_date_time
            return @due_date_time
        end
        ## 
        ## Gets the flagStatus property value. 
        ## @return a followup_flag_status
        ## 
        def  flag_status
            return @flag_status
        end
        ## 
        ## Gets the startDateTime property value. 
        ## @return a date_time_time_zone
        ## 
        def  start_date_time
            return @start_date_time
        end
        ## 
        ## The deserialization information for the current model
        ## @return a i_dictionary
        ## 
        def get_field_deserializers() 
            return {
                "completedDateTime" => lambda {|o, n| o.completed_date_time = n.get_object_value(Graphrubyv4::Users::DateTimeTimeZone) },
                "dueDateTime" => lambda {|o, n| o.due_date_time = n.get_object_value(Graphrubyv4::Users::DateTimeTimeZone) },
                "flagStatus" => lambda {|o, n| o.flag_status = n.get_enum_value(Graphrubyv4::Users::FollowupFlagStatus) },
                "startDateTime" => lambda {|o, n| o.start_date_time = n.get_object_value(Graphrubyv4::Users::DateTimeTimeZone) },
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
        ## Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        ## @param value Value to set for the AdditionalData property.
        ## @return a void
        ## 
        def  additional_data=(additionalData)
            @additional_data = additionalData
        end
        ## 
        ## Sets the completedDateTime property value. 
        ## @param value Value to set for the completedDateTime property.
        ## @return a void
        ## 
        def  completed_date_time=(completedDateTime)
            @completed_date_time = completedDateTime
        end
        ## 
        ## Sets the dueDateTime property value. 
        ## @param value Value to set for the dueDateTime property.
        ## @return a void
        ## 
        def  due_date_time=(dueDateTime)
            @due_date_time = dueDateTime
        end
        ## 
        ## Sets the flagStatus property value. 
        ## @param value Value to set for the flagStatus property.
        ## @return a void
        ## 
        def  flag_status=(flagStatus)
            @flag_status = flagStatus
        end
        ## 
        ## Sets the startDateTime property value. 
        ## @param value Value to set for the startDateTime property.
        ## @return a void
        ## 
        def  start_date_time=(startDateTime)
            @start_date_time = startDateTime
        end
    end
end
