from __future__ import annotations
from kiota_abstractions.serialization import AdditionalDataHolder, Parsable, ParseNode, SerializationWriter
from kiota_abstractions.utils import lazy_import
from typing import Any, Callable, Dict, List, Optional, Union

message = lazy_import('graph_pythonv1.models.message')

class MessageCollectionResponse(AdditionalDataHolder, Parsable):
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
        Instantiates a new MessageCollectionResponse and sets the default values.
        """
        # Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        self._additional_data: Dict[str, Any] = {}

        # The OdataNextLink property
        self._odata_next_link: Optional[str] = None
        # The value property
        self._value: Optional[List[message.Message]] = None
    
    @staticmethod
    def create_from_discriminator_value(parse_node: Optional[ParseNode] = None) -> MessageCollectionResponse:
        """
        Creates a new instance of the appropriate class based on discriminator value
        Args:
            parseNode: The parse node to use to read the discriminator value and create the object
        Returns: MessageCollectionResponse
        """
        if parse_node is None:
            raise Exception("parse_node cannot be undefined")
        return MessageCollectionResponse()
    
    def get_field_deserializers(self,) -> Dict[str, Callable[[ParseNode], None]]:
        """
        The deserialization information for the current model
        Returns: Dict[str, Callable[[ParseNode], None]]
        """
        fields = {
            "@odata.nextLink": lambda n : setattr(self, 'odata_next_link', n.get_str_value()),
            "value": lambda n : setattr(self, 'value', n.get_collection_of_object_values(message.Message)),
        }
        return fields
    
    @property
    def odata_next_link(self,) -> Optional[str]:
        """
        Gets the @odata.nextLink property value. The OdataNextLink property
        Returns: Optional[str]
        """
        return self._odata_next_link
    
    @odata_next_link.setter
    def odata_next_link(self,value: Optional[str] = None) -> None:
        """
        Sets the @odata.nextLink property value. The OdataNextLink property
        Args:
            value: Value to set for the OdataNextLink property.
        """
        self._odata_next_link = value
    
    def serialize(self,writer: SerializationWriter) -> None:
        """
        Serializes information the current object
        Args:
            writer: Serialization writer to use to serialize this model
        """
        if writer is None:
            raise Exception("writer cannot be undefined")
        writer.write_str_value("@odata.nextLink", self.odata_next_link)
        writer.write_collection_of_object_values("value", self.value)
        writer.write_additional_data_value(self.additional_data)
    
    @property
    def value(self,) -> Optional[List[message.Message]]:
        """
        Gets the value property value. The value property
        Returns: Optional[List[message.Message]]
        """
        return self._value
    
    @value.setter
    def value(self,value: Optional[List[message.Message]] = None) -> None:
        """
        Sets the value property value. The value property
        Args:
            value: Value to set for the value property.
        """
        self._value = value
    

