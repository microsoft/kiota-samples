from __future__ import annotations
from kiota_abstractions.serialization import AdditionalDataHolder, Parsable, ParseNode, SerializationWriter
from typing import Any, Callable, Dict, List, Optional, Union

class EmailAddress(AdditionalDataHolder, Parsable):
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

    @property
    def address(self,) -> Optional[str]:
        """
        Gets the address property value. The email address of the person or entity.
        Returns: Optional[str]
        """
        return self._address

    @address.setter
    def address(self,value: Optional[str] = None) -> None:
        """
        Sets the address property value. The email address of the person or entity.
        Args:
            value: Value to set for the address property.
        """
        self._address = value

    def __init__(self,) -> None:
        """
        Instantiates a new emailAddress and sets the default values.
        """
        # The email address of the person or entity.
        self._address: Optional[str] = None

        # The display name of the person or entity.
        self._name: Optional[str] = None

        # Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        self._additional_data: Dict[str, Any] = {}


    @staticmethod
    def create_from_discriminator_value(parse_node: Optional[ParseNode] = None) -> EmailAddress:
        """
        Creates a new instance of the appropriate class based on discriminator value
        Args:
            parseNode: The parse node to use to read the discriminator value and create the object
        Returns: EmailAddress
        """
        if not parse_node:
            raise Exception("parse_node cannot be undefined")
        return EmailAddress()

    def get_field_deserializers(self,) -> Dict[str, Callable[[ParseNode], None]]:
        """
        The deserialization information for the current model
        Returns: Dict[str, Callable[[ParseNode], None]]
        """
        fields = {
            "address": lambda n : setattr(self, 'address', n.get_str_value()),
            "name": lambda n : setattr(self, 'name', n.get_str_value()),
        }
        return fields

    @property
    def name(self,) -> Optional[str]:
        """
        Gets the name property value. The display name of the person or entity.
        Returns: Optional[str]
        """
        return self._name

    @name.setter
    def name(self,value: Optional[str] = None) -> None:
        """
        Sets the name property value. The display name of the person or entity.
        Args:
            value: Value to set for the name property.
        """
        self._name = value

    def serialize(self,writer: SerializationWriter) -> None:
        """
        Serializes information the current object
        Args:
            writer: Serialization writer to use to serialize this model
        """
        if not writer:
            raise Exception("writer cannot be undefined")
        writer.write_str_value("address", self.address)
        writer.write_str_value("name", self.name)
        writer.write_additional_data_value(self.additional_data)


