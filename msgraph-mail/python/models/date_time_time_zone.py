from __future__ import annotations
from kiota_abstractions.serialization import AdditionalDataHolder, Parsable, ParseNode, SerializationWriter
from typing import Any, Callable, Dict, List, Optional, Union

class DateTimeTimeZone(AdditionalDataHolder, Parsable):
    @property
    def additional_data(self,) -> Dict[str, Any]:
        """
        Gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        Returns: Dict[str, Any]
        """
        return self._additional_data

    @additional_data.setter
    def additional_data(self,value: Dict[str, Any]) -> None:
        """
        Sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        Args:
            value: Value to set for the AdditionalData property.
        """
        self._additional_data = value

    def __init__(self,) -> None:
        """
        Instantiates a new dateTimeTimeZone and sets the default values.
        """
        # A single point of time in a combined date and time representation ({date}T{time}; for example, 2017-08-29T04:00:00.0000000).
        self._date_time: Optional[str] = None

        # Represents a time zone, for example, 'Pacific Standard Time'. See below for more possible values.
        self._time_zone: Optional[str] = None

        # Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        self._additional_data: Dict[str, Any] = {}


    @staticmethod
    def create_from_discriminator_value(parse_node: Optional[ParseNode] = None) -> DateTimeTimeZone:
        """
        Creates a new instance of the appropriate class based on discriminator value
        Args:
            parseNode: The parse node to use to read the discriminator value and create the object
        Returns: DateTimeTimeZone
        """
        if not parse_node:
            raise Exception("parse_node cannot be undefined")
        return DateTimeTimeZone()

    @property
    def date_time(self,) -> Optional[str]:
        """
        Gets the dateTime property value. A single point of time in a combined date and time representation ({date}T{time}; for example, 2017-08-29T04:00:00.0000000).
        Returns: Optional[str]
        """
        return self._date_time

    @date_time.setter
    def date_time(self,value: Optional[str] = None) -> None:
        """
        Sets the dateTime property value. A single point of time in a combined date and time representation ({date}T{time}; for example, 2017-08-29T04:00:00.0000000).
        Args:
            value: Value to set for the dateTime property.
        """
        self._date_time = value

    def get_field_deserializers(self,) -> Dict[str, Callable[[ParseNode], None]]:
        """
        The deserialization information for the current model
        Returns: Dict[str, Callable[[ParseNode], None]]
        """
        fields = {
            "date_time": lambda n : setattr(self, 'date_time', n.get_str_value()),
            "time_zone": lambda n : setattr(self, 'time_zone', n.get_str_value()),
        }
        return fields

    def serialize(self,writer: SerializationWriter) -> None:
        """
        Serializes information the current object
        Args:
            writer: Serialization writer to use to serialize this model
        """
        if not writer:
            raise Exception("writer cannot be undefined")
        writer.write_str_value("dateTime", self.date_time)
        writer.write_str_value("timeZone", self.time_zone)
        writer.write_additional_data_value(self.additional_data)

    @property
    def time_zone(self,) -> Optional[str]:
        """
        Gets the timeZone property value. Represents a time zone, for example, 'Pacific Standard Time'. See below for more possible values.
        Returns: Optional[str]
        """
        return self._time_zone

    @time_zone.setter
    def time_zone(self,value: Optional[str] = None) -> None:
        """
        Sets the timeZone property value. Represents a time zone, for example, 'Pacific Standard Time'. See below for more possible values.
        Args:
            value: Value to set for the timeZone property.
        """
        self._time_zone = value


