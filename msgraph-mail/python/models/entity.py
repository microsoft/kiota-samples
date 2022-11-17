from __future__ import annotations
from kiota_abstractions.serialization import AdditionalDataHolder, Parsable, ParseNode, SerializationWriter
from typing import Any, Callable, Dict, List, Optional, Union

class Entity(AdditionalDataHolder, Parsable):
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
        Instantiates a new entity and sets the default values.
        """
        # The unique idenfier for an entity. Read-only.
        self._id: Optional[str] = None

        # Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        self._additional_data: Dict[str, Any] = {}


    @staticmethod
    def create_from_discriminator_value(parse_node: Optional[ParseNode] = None) -> Entity:
        """
        Creates a new instance of the appropriate class based on discriminator value
        Args:
            parseNode: The parse node to use to read the discriminator value and create the object
        Returns: Entity
        """
        if not parse_node:
            raise Exception("parse_node cannot be undefined")
        return Entity()

    def get_field_deserializers(self,) -> Dict[str, Callable[[ParseNode], None]]:
        """
        The deserialization information for the current model
        Returns: Dict[str, Callable[[ParseNode], None]]
        """
        fields = {
            "id": lambda n : setattr(self, 'id', n.get_str_value()),
        }
        return fields

    @property
    def id(self,) -> Optional[str]:
        """
        Gets the id property value. The unique idenfier for an entity. Read-only.
        Returns: Optional[str]
        """
        return self._id

    @id.setter
    def id(self,value: Optional[str] = None) -> None:
        """
        Sets the id property value. The unique idenfier for an entity. Read-only.
        Args:
            value: Value to set for the id property.
        """
        self._id = value

    def serialize(self,writer: SerializationWriter) -> None:
        """
        Serializes information the current object
        Args:
            writer: Serialization writer to use to serialize this model
        """
        if not writer:
            raise Exception("writer cannot be undefined")
        writer.write_str_value("id", self.id)
        writer.write_additional_data_value(self.additional_data)


