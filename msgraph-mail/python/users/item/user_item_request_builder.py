from __future__ import annotations
from kiota_abstractions.get_path_parameters import get_path_parameters
from kiota_abstractions.request_adapter import RequestAdapter
from kiota_abstractions.utils import lazy_import
from typing import Any, Callable, Dict, List, Optional, Union

inference_classification_request_builder = lazy_import('graph_pythonv1.users.item.inference_classification.inference_classification_request_builder')
mail_folders_request_builder = lazy_import('graph_pythonv1.users.item.mail_folders.mail_folders_request_builder')
mail_folder_item_request_builder = lazy_import('graph_pythonv1.users.item.mail_folders.item.mail_folder_item_request_builder')
messages_request_builder = lazy_import('graph_pythonv1.users.item.messages.messages_request_builder')
message_item_request_builder = lazy_import('graph_pythonv1.users.item.messages.item.message_item_request_builder')

class UserItemRequestBuilder():
    """
    Builds and executes requests for operations under /users/{user-id}
    """
    @property
    def inference_classification(self) -> inference_classification_request_builder.InferenceClassificationRequestBuilder:
        """
        The inferenceClassification property
        """
        return inference_classification_request_builder.InferenceClassificationRequestBuilder(self.request_adapter, self.path_parameters)
    
    @property
    def mail_folders(self) -> mail_folders_request_builder.MailFoldersRequestBuilder:
        """
        The mailFolders property
        """
        return mail_folders_request_builder.MailFoldersRequestBuilder(self.request_adapter, self.path_parameters)
    
    @property
    def messages(self) -> messages_request_builder.MessagesRequestBuilder:
        """
        The messages property
        """
        return messages_request_builder.MessagesRequestBuilder(self.request_adapter, self.path_parameters)
    
    def __init__(self,request_adapter: RequestAdapter, path_parameters: Optional[Union[Dict[str, Any], str]] = None) -> None:
        """
        Instantiates a new UserItemRequestBuilder and sets the default values.
        Args:
            pathParameters: The raw url or the Url template parameters for the request.
            requestAdapter: The request adapter to use to execute the requests.
        """
        if path_parameters is None:
            raise Exception("path_parameters cannot be undefined")
        if request_adapter is None:
            raise Exception("request_adapter cannot be undefined")
        # Url template to use to build the URL for the current request builder
        self.url_template: str = "{+baseurl}/users/{user%2Did}"

        url_tpl_params = get_path_parameters(path_parameters)
        self.path_parameters = url_tpl_params
        self.request_adapter = request_adapter
    
    def mail_folders_by_id(self,id: str) -> mail_folder_item_request_builder.MailFolderItemRequestBuilder:
        """
        Gets an item from the GraphPythonv1.users.item.mailFolders.item collection
        Args:
            id: Unique identifier of the item
        Returns: mail_folder_item_request_builder.MailFolderItemRequestBuilder
        """
        if id is None:
            raise Exception("id cannot be undefined")
        url_tpl_params = get_path_parameters(self.path_parameters)
        url_tpl_params["mailFolder%2Did"] = id
        return mail_folder_item_request_builder.MailFolderItemRequestBuilder(self.request_adapter, url_tpl_params)
    
    def messages_by_id(self,id: str) -> message_item_request_builder.MessageItemRequestBuilder:
        """
        Gets an item from the GraphPythonv1.users.item.messages.item collection
        Args:
            id: Unique identifier of the item
        Returns: message_item_request_builder.MessageItemRequestBuilder
        """
        if id is None:
            raise Exception("id cannot be undefined")
        url_tpl_params = get_path_parameters(self.path_parameters)
        url_tpl_params["message%2Did"] = id
        return message_item_request_builder.MessageItemRequestBuilder(self.request_adapter, url_tpl_params)
    

