require 'microsoft_kiota_abstractions'
require_relative './users'

module Graphrubyv4::Users
    class DateTimeTimeZone
        include MicrosoftKiotaAbstractions::Parsable
        ## 
        # Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        @additional_data
        ## 
        # A single point of time in a combined date and time representation ({date}T{time}). For example, '2019-04-16T09:00:00'.
        @date_time
        ## 
        # Represents a time zone, for example, 'Pacific Standard Time'. See below for possible values.
        @time_zone
        ## 
        ## Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        ## @return a i_dictionary
        ## 
        def  additional_data
            return @additional_data
        end
        ## 
        ## Gets the dateTime property value. A single point of time in a combined date and time representation ({date}T{time}). For example, '2019-04-16T09:00:00'.
        ## @return a string
        ## 
        def  date_time
            return @date_time
        end
        ## 
        ## Gets the timeZone property value. Represents a time zone, for example, 'Pacific Standard Time'. See below for possible values.
        ## @return a string
        ## 
        def  time_zone
            return @time_zone
        end
        ## 
        ## The deserialization information for the current model
        ## @return a i_dictionary
        ## 
        def get_field_deserializers() 
            return {
                "dateTime" => lambda {|o, n| o.date_time = n.get_string_value() },
                "timeZone" => lambda {|o, n| o.time_zone = n.get_string_value() },
            }
        end
        ## 
        ## Serializes information the current object
        ## @param writer Serialization writer to use to serialize this model
        ## @return a void
        ## 
        def serialize(writer) 
            writer.write_string_value("dateTime", @date_time)
            writer.write_string_value("timeZone", @time_zone)
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
        ## Sets the dateTime property value. A single point of time in a combined date and time representation ({date}T{time}). For example, '2019-04-16T09:00:00'.
        ## @param value Value to set for the dateTime property.
        ## @return a void
        ## 
        def  date_time=(dateTime)
            @date_time = dateTime
        end
        ## 
        ## Sets the timeZone property value. Represents a time zone, for example, 'Pacific Standard Time'. See below for possible values.
        ## @param value Value to set for the timeZone property.
        ## @return a void
        ## 
        def  time_zone=(timeZone)
            @time_zone = timeZone
        end
    end
end
