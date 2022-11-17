from __future__ import annotations
from kiota_abstractions.serialization import AdditionalDataHolder, Parsable, ParseNode, SerializationWriter
from typing import Any, Callable, Dict, List, Optional, Union

class SizeRange(AdditionalDataHolder, Parsable):
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
        Instantiates a new sizeRange and sets the default values.
        """
        # The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
        self._maximum_size: Optional[int] = None

        # The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
        self._minimum_size: Optional[int] = None

        # Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        self._additional_data: Dict[str, Any] = {}


    @staticmethod
    def create_from_discriminator_value(parse_node: Optional[ParseNode] = None) -> SizeRange:
        """
        Creates a new instance of the appropriate class based on discriminator value
        Args:
            parseNode: The parse node to use to read the discriminator value and create the object
        Returns: SizeRange
        """
        if not parse_node:
            raise Exception("parse_node cannot be undefined")
        return SizeRange()

    def get_field_deserializers(self,) -> Dict[str, Callable[[ParseNode], None]]:
        """
        The deserialization information for the current model
        Returns: Dict[str, Callable[[ParseNode], None]]
        """
        fields = {
            "maximum_size": lambda n : setattr(self, 'maximum_size', n.get_int_value()),
            "minimum_size": lambda n : setattr(self, 'minimum_size', n.get_int_value()),
        }
        return fields

    @property
    def maximum_size(self,) -> Optional[int]:
        """
        Gets the maximumSize property value. The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
        Returns: Optional[int]
        """
        return self._maximum_size

    @maximum_size.setter
    def maximum_size(self,value: Optional[int] = None) -> None:
        """
        Sets the maximumSize property value. The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
        Args:
            value: Value to set for the maximumSize property.
        """
        self._maximum_size = value

    @property
    def minimum_size(self,) -> Optional[int]:
        """
        Gets the minimumSize property value. The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
        Returns: Optional[int]
        """
        return self._minimum_size

    @minimum_size.setter
    def minimum_size(self,value: Optional[int] = None) -> None:
        """
        Sets the minimumSize property value. The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
        Args:
            value: Value to set for the minimumSize property.
        """
        self._minimum_size = value

    def serialize(self,writer: SerializationWriter) -> None:
        """
        Serializes information the current object
        Args:
            writer: Serialization writer to use to serialize this model
        """
        if not writer:
            raise Exception("writer cannot be undefined")
        writer.write_int_value("maximumSize", self.maximum_size)
        writer.write_int_value("minimumSize", self.minimum_size)
        writer.write_additional_data_value(self.additional_data)


