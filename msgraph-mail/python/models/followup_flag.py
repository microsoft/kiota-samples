from __future__ import annotations
from collections.abc import Callable
from dataclasses import dataclass, field
from kiota_abstractions.serialization import AdditionalDataHolder, Parsable, ParseNode, SerializationWriter
from typing import Any, Optional, TYPE_CHECKING, Union

if TYPE_CHECKING:
    from .date_time_time_zone import DateTimeTimeZone
    from .followup_flag_status import FollowupFlagStatus

@dataclass
class FollowupFlag(AdditionalDataHolder, Parsable):
    # Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additional_data: dict[str, Any] = field(default_factory=dict)

    # The completedDateTime property
    completed_date_time: Optional[DateTimeTimeZone] = None
    # The dueDateTime property
    due_date_time: Optional[DateTimeTimeZone] = None
    # The flagStatus property
    flag_status: Optional[FollowupFlagStatus] = None
    # The startDateTime property
    start_date_time: Optional[DateTimeTimeZone] = None
    
    @staticmethod
    def create_from_discriminator_value(parse_node: ParseNode) -> FollowupFlag:
        """
        Creates a new instance of the appropriate class based on discriminator value
        param parse_node: The parse node to use to read the discriminator value and create the object
        Returns: FollowupFlag
        """
        if parse_node is None:
            raise TypeError("parse_node cannot be null.")
        return FollowupFlag()
    
    def get_field_deserializers(self,) -> dict[str, Callable[[ParseNode], None]]:
        """
        The deserialization information for the current model
        Returns: dict[str, Callable[[ParseNode], None]]
        """
        from .date_time_time_zone import DateTimeTimeZone
        from .followup_flag_status import FollowupFlagStatus

        from .date_time_time_zone import DateTimeTimeZone
        from .followup_flag_status import FollowupFlagStatus

        fields: dict[str, Callable[[Any], None]] = {
            "completedDateTime": lambda n : setattr(self, 'completed_date_time', n.get_object_value(DateTimeTimeZone)),
            "dueDateTime": lambda n : setattr(self, 'due_date_time', n.get_object_value(DateTimeTimeZone)),
            "flagStatus": lambda n : setattr(self, 'flag_status', n.get_enum_value(FollowupFlagStatus)),
            "startDateTime": lambda n : setattr(self, 'start_date_time', n.get_object_value(DateTimeTimeZone)),
        }
        return fields
    
    def serialize(self,writer: SerializationWriter) -> None:
        """
        Serializes information the current object
        param writer: Serialization writer to use to serialize this model
        Returns: None
        """
        if writer is None:
            raise TypeError("writer cannot be null.")
        writer.write_object_value("completedDateTime", self.completed_date_time)
        writer.write_object_value("dueDateTime", self.due_date_time)
        writer.write_enum_value("flagStatus", self.flag_status)
        writer.write_object_value("startDateTime", self.start_date_time)
        writer.write_additional_data_value(self.additional_data)
    

