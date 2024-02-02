from __future__ import annotations
from kiota_abstractions.base_request_builder import BaseRequestBuilder
from kiota_abstractions.get_path_parameters import get_path_parameters
from kiota_abstractions.request_adapter import RequestAdapter
from typing import Any, Callable, Dict, List, Optional, TYPE_CHECKING, Union

if TYPE_CHECKING:
    from .inference_classification.inference_classification_request_builder import InferenceClassificationRequestBuilder
    from .mail_folders.mail_folders_request_builder import MailFoldersRequestBuilder
    from .messages.messages_request_builder import MessagesRequestBuilder

class UserItemRequestBuilder(BaseRequestBuilder):
    """
    Builds and executes requests for operations under /users/{user-id}
    """
    def __init__(self,request_adapter: RequestAdapter, path_parameters: Union[str, Dict[str, Any]]) -> None:
        """
        Instantiates a new UserItemRequestBuilder and sets the default values.
        param path_parameters: The raw url or the url-template parameters for the request.
        param request_adapter: The request adapter to use to execute the requests.
        Returns: None
        """
        super().__init__(request_adapter, "{+baseurl}/users/{user%2Did}", path_parameters)
    
    @property
    def inference_classification(self) -> InferenceClassificationRequestBuilder:
        """
        The inferenceClassification property
        """
        from .inference_classification.inference_classification_request_builder import InferenceClassificationRequestBuilder

        return InferenceClassificationRequestBuilder(self.request_adapter, self.path_parameters)
    
    @property
    def mail_folders(self) -> MailFoldersRequestBuilder:
        """
        The mailFolders property
        """
        from .mail_folders.mail_folders_request_builder import MailFoldersRequestBuilder

        return MailFoldersRequestBuilder(self.request_adapter, self.path_parameters)
    
    @property
    def messages(self) -> MessagesRequestBuilder:
        """
        The messages property
        """
        from .messages.messages_request_builder import MessagesRequestBuilder

        return MessagesRequestBuilder(self.request_adapter, self.path_parameters)
    

