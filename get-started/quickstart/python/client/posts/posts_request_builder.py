from __future__ import annotations
from dataclasses import dataclass, field
from kiota_abstractions.base_request_builder import BaseRequestBuilder
from kiota_abstractions.base_request_configuration import RequestConfiguration
from kiota_abstractions.get_path_parameters import get_path_parameters
from kiota_abstractions.method import Method
from kiota_abstractions.request_adapter import RequestAdapter
from kiota_abstractions.request_information import RequestInformation
from kiota_abstractions.request_option import RequestOption
from kiota_abstractions.serialization import Parsable, ParsableFactory
from typing import Any, Callable, Dict, List, Optional, TYPE_CHECKING, Union

if TYPE_CHECKING:
    from ..models.post import Post
    from .item.post_item_request_builder import PostItemRequestBuilder

class PostsRequestBuilder(BaseRequestBuilder):
    """
    Builds and executes requests for operations under /posts
    """
    def __init__(self,request_adapter: RequestAdapter, path_parameters: Union[str, Dict[str, Any]]) -> None:
        """
        Instantiates a new PostsRequestBuilder and sets the default values.
        param path_parameters: The raw url or the url-template parameters for the request.
        param request_adapter: The request adapter to use to execute the requests.
        Returns: None
        """
        super().__init__(request_adapter, "{+baseurl}/posts{?title*,userId*}", path_parameters)
    
    def by_post_id(self,post_id: int) -> PostItemRequestBuilder:
        """
        Gets an item from the client.posts.item collection
        param post_id: key: id of post
        Returns: PostItemRequestBuilder
        """
        if not post_id:
            raise TypeError("post_id cannot be null.")
        from .item.post_item_request_builder import PostItemRequestBuilder

        url_tpl_params = get_path_parameters(self.path_parameters)
        url_tpl_params["post%2Did"] = post_id
        return PostItemRequestBuilder(self.request_adapter, url_tpl_params)
    
    async def get(self,request_configuration: Optional[RequestConfiguration] = None) -> Optional[List[Post]]:
        """
        Get posts
        param request_configuration: Configuration for the request such as headers, query parameters, and middleware options.
        Returns: Optional[List[Post]]
        """
        request_info = self.to_get_request_information(
            request_configuration
        )
        if not self.request_adapter:
            raise Exception("Http core is null") 
        from ..models.post import Post

        return await self.request_adapter.send_collection_async(request_info, Post, None)
    
    async def post(self,body: Optional[Post] = None, request_configuration: Optional[RequestConfiguration] = None) -> Optional[Post]:
        """
        Create post
        param body: The request body
        param request_configuration: Configuration for the request such as headers, query parameters, and middleware options.
        Returns: Optional[Post]
        """
        if not body:
            raise TypeError("body cannot be null.")
        request_info = self.to_post_request_information(
            body, request_configuration
        )
        if not self.request_adapter:
            raise Exception("Http core is null") 
        from ..models.post import Post

        return await self.request_adapter.send_async(request_info, Post, None)
    
    def to_get_request_information(self,request_configuration: Optional[RequestConfiguration] = None) -> RequestInformation:
        """
        Get posts
        param request_configuration: Configuration for the request such as headers, query parameters, and middleware options.
        Returns: RequestInformation
        """
        request_info = RequestInformation(Method.GET, self.url_template, self.path_parameters)
        request_info.configure(request_configuration)
        request_info.headers.try_add("Accept", "application/json")
        return request_info
    
    def to_post_request_information(self,body: Optional[Post] = None, request_configuration: Optional[RequestConfiguration] = None) -> RequestInformation:
        """
        Create post
        param body: The request body
        param request_configuration: Configuration for the request such as headers, query parameters, and middleware options.
        Returns: RequestInformation
        """
        if not body:
            raise TypeError("body cannot be null.")
        request_info = RequestInformation(Method.POST, '{+baseurl}/posts', self.path_parameters)
        request_info.configure(request_configuration)
        request_info.headers.try_add("Accept", "application/json")
        request_info.set_content_from_parsable(self.request_adapter, "application/json", body)
        return request_info
    
    def with_url(self,raw_url: Optional[str] = None) -> PostsRequestBuilder:
        """
        Returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
        param raw_url: The raw URL to use for the request builder.
        Returns: PostsRequestBuilder
        """
        if not raw_url:
            raise TypeError("raw_url cannot be null.")
        return PostsRequestBuilder(self.request_adapter, raw_url)
    
    @dataclass
    class PostsRequestBuilderGetQueryParameters():
        """
        Get posts
        """
        def get_query_parameter(self,original_name: Optional[str] = None) -> str:
            """
            Maps the query parameters names to their encoded names for the URI template parsing.
            param original_name: The original query parameter name in the class.
            Returns: str
            """
            if not original_name:
                raise TypeError("original_name cannot be null.")
            if original_name == "user_id":
                return "userId"
            if original_name == "title":
                return "title"
            return original_name
        
        # Filter results by title
        title: Optional[str] = None

        # Filter results by user ID
        user_id: Optional[int] = None

    

