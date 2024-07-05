<?php

namespace Microsoft\Graph\Users\Item\MailFolders\Item\ChildFolders;

use Exception;
use Http\Promise\Promise;
use Microsoft\Graph\Models\MailFolder;
use Microsoft\Graph\Models\MailFolderCollectionResponse;
use Microsoft\Graph\Models\ODataErrors\ODataError;
use Microsoft\Graph\Users\Item\MailFolders\Item\ChildFolders\Count\CountRequestBuilder;
use Microsoft\Graph\Users\Item\MailFolders\Item\ChildFolders\Item\MailFolderItemRequestBuilder;
use Microsoft\Kiota\Abstractions\BaseRequestBuilder;
use Microsoft\Kiota\Abstractions\HttpMethod;
use Microsoft\Kiota\Abstractions\RequestAdapter;
use Microsoft\Kiota\Abstractions\RequestInformation;

/**
 * Builds and executes requests for operations under /users/{user-id}/mailFolders/{mailFolder-id}/childFolders
*/
class ChildFoldersRequestBuilder extends BaseRequestBuilder 
{
    /**
     * The Count property
    */
    public function count(): CountRequestBuilder {
        return new CountRequestBuilder($this->pathParameters, $this->requestAdapter);
    }
    
    /**
     * Gets an item from the Microsoft/Graph.users.item.mailFolders.item.childFolders.item collection
     * @param string $mailFolderId1 The unique identifier of mailFolder
     * @return MailFolderItemRequestBuilder
    */
    public function byMailFolderId1(string $mailFolderId1): MailFolderItemRequestBuilder {
        $urlTplParams = $this->pathParameters;
        $urlTplParams['mailFolder%2Did1'] = $mailFolderId1;
        return new MailFolderItemRequestBuilder($urlTplParams, $this->requestAdapter);
    }

    /**
     * Instantiates a new ChildFoldersRequestBuilder and sets the default values.
     * @param array<string, mixed>|string $pathParametersOrRawUrl Path parameters for the request or a String representing the raw URL.
     * @param RequestAdapter $requestAdapter The request adapter to use to execute the requests.
    */
    public function __construct($pathParametersOrRawUrl, RequestAdapter $requestAdapter) {
        parent::__construct($requestAdapter, [], '{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders{?%24count,%24expand,%24filter,%24orderby,%24select,%24skip,%24top,includeHiddenFolders}');
        if (is_array($pathParametersOrRawUrl)) {
            $this->pathParameters = $pathParametersOrRawUrl;
        } else {
            $this->pathParameters = ['request-raw-url' => $pathParametersOrRawUrl];
        }
    }

    /**
     * The collection of child folders in the mailFolder.
     * @param ChildFoldersRequestBuilderGetRequestConfiguration|null $requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @return Promise<MailFolderCollectionResponse|null>
     * @throws Exception
    */
    public function get(?ChildFoldersRequestBuilderGetRequestConfiguration $requestConfiguration = null): Promise {
        $requestInfo = $this->toGetRequestInformation($requestConfiguration);
        $errorMappings = [
                'XXX' => [ODataError::class, 'createFromDiscriminatorValue'],
        ];
        return $this->requestAdapter->sendAsync($requestInfo, [MailFolderCollectionResponse::class, 'createFromDiscriminatorValue'], $errorMappings);
    }

    /**
     * Create new navigation property to childFolders for users
     * @param MailFolder $body The request body
     * @param ChildFoldersRequestBuilderPostRequestConfiguration|null $requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @return Promise<MailFolder|null>
     * @throws Exception
    */
    public function post(MailFolder $body, ?ChildFoldersRequestBuilderPostRequestConfiguration $requestConfiguration = null): Promise {
        $requestInfo = $this->toPostRequestInformation($body, $requestConfiguration);
        $errorMappings = [
                'XXX' => [ODataError::class, 'createFromDiscriminatorValue'],
        ];
        return $this->requestAdapter->sendAsync($requestInfo, [MailFolder::class, 'createFromDiscriminatorValue'], $errorMappings);
    }

    /**
     * The collection of child folders in the mailFolder.
     * @param ChildFoldersRequestBuilderGetRequestConfiguration|null $requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @return RequestInformation
    */
    public function toGetRequestInformation(?ChildFoldersRequestBuilderGetRequestConfiguration $requestConfiguration = null): RequestInformation {
        $requestInfo = new RequestInformation();
        $requestInfo->urlTemplate = $this->urlTemplate;
        $requestInfo->pathParameters = $this->pathParameters;
        $requestInfo->httpMethod = HttpMethod::GET;
        if ($requestConfiguration !== null) {
            $requestInfo->addHeaders($requestConfiguration->headers);
            if ($requestConfiguration->queryParameters !== null) {
                $requestInfo->setQueryParameters($requestConfiguration->queryParameters);
            }
            $requestInfo->addRequestOptions(...$requestConfiguration->options);
        }
        $requestInfo->tryAddHeader('Accept', "application/json");
        return $requestInfo;
    }

    /**
     * Create new navigation property to childFolders for users
     * @param MailFolder $body The request body
     * @param ChildFoldersRequestBuilderPostRequestConfiguration|null $requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @return RequestInformation
    */
    public function toPostRequestInformation(MailFolder $body, ?ChildFoldersRequestBuilderPostRequestConfiguration $requestConfiguration = null): RequestInformation {
        $requestInfo = new RequestInformation();
        $requestInfo->urlTemplate = $this->urlTemplate;
        $requestInfo->pathParameters = $this->pathParameters;
        $requestInfo->httpMethod = HttpMethod::POST;
        if ($requestConfiguration !== null) {
            $requestInfo->addHeaders($requestConfiguration->headers);
            $requestInfo->addRequestOptions(...$requestConfiguration->options);
        }
        $requestInfo->tryAddHeader('Accept', "application/json");
        $requestInfo->setContentFromParsable($this->requestAdapter, "application/json", $body);
        return $requestInfo;
    }

    /**
     * Returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
     * @param string $rawUrl The raw URL to use for the request builder.
     * @return ChildFoldersRequestBuilder
    */
    public function withUrl(string $rawUrl): ChildFoldersRequestBuilder {
        return new ChildFoldersRequestBuilder($rawUrl, $this->requestAdapter);
    }

}
