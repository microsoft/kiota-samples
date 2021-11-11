<?php

namespace Microsoft\Graph\Users\Item\MailFolders\Item;

use Http\Promise\Promise;
use Microsoft\Graph\Models\Microsoft\Graph\MailFolder;
use Microsoft\Graph\Users\Item\MailFolders\Item\ChildFolders\ChildFoldersRequestBuilder;
use Microsoft\Graph\Users\Item\MailFolders\Item\MessageRules\Item\MessageRuleRequestBuilder;
use Microsoft\Graph\Users\Item\MailFolders\Item\MessageRules\MessageRulesRequestBuilder;
use Microsoft\Graph\Users\Item\MailFolders\Item\Messages\Item\MessageRequestBuilder;
use Microsoft\Graph\Users\Item\MailFolders\Item\Messages\MessagesRequestBuilder;
use Microsoft\Graph\Users\Item\MailFolders\Item\MultiValueExtendedProperties\Item\MultiValueLegacyExtendedPropertyRequestBuilder;
use Microsoft\Graph\Users\Item\MailFolders\Item\MultiValueExtendedProperties\MultiValueExtendedPropertiesRequestBuilder;
use Microsoft\Graph\Users\Item\MailFolders\Item\SingleValueExtendedProperties\Item\SingleValueLegacyExtendedPropertyRequestBuilder;
use Microsoft\Graph\Users\Item\MailFolders\Item\SingleValueExtendedProperties\SingleValueExtendedPropertiesRequestBuilder;
use Microsoft\Kiota\Abstractions\HttpMethod;
use Microsoft\Kiota\Abstractions\MiddlewareOption;
use Microsoft\Kiota\Abstractions\RequestAdapter;
use Microsoft\Kiota\Abstractions\RequestInformation;
use Microsoft\Kiota\Abstractions\ResponseHandler;
use Microsoft\Kiota\Abstractions\Serialization\Parsable;

class MailFolderRequestBuilder 
{
    public function childFolders(): ChildFoldersRequestBuilder {
        return new ChildFoldersRequestBuilder($this->pathParameters, $this->requestAdapter);
    }
    
    public function messageRules(): MessageRulesRequestBuilder {
        return new MessageRulesRequestBuilder($this->pathParameters, $this->requestAdapter);
    }
    
    public function messages(): MessagesRequestBuilder {
        return new MessagesRequestBuilder($this->pathParameters, $this->requestAdapter);
    }
    
    public function multiValueExtendedProperties(): MultiValueExtendedPropertiesRequestBuilder {
        return new MultiValueExtendedPropertiesRequestBuilder($this->pathParameters, $this->requestAdapter);
    }
    
    /** @var array|null $pathParameters Path parameters for the request */
    private array $pathParameters;
    
    /** @var RequestAdapter|null $requestAdapter The request adapter to use to execute the requests. */
    private RequestAdapter $requestAdapter;
    
    public function singleValueExtendedProperties(): SingleValueExtendedPropertiesRequestBuilder {
        return new SingleValueExtendedPropertiesRequestBuilder($this->pathParameters, $this->requestAdapter);
    }
    
    /** @var string|null $urlTemplate Url template to use to build the URL for the current request builder */
    private ?string $urlTemplate;
    
    /**
     * Gets an item from the Microsoft\Graph.users.item.mailFolders.item.childFolders.item collection
     * @param string $id Unique identifier of the item
     * @return MailFolderRequestBuilder
    */
    public function childFoldersById(string $id): MailFolderRequestBuilder {
        $urlTplParams = $this->pathParameters;
        $urlTplParams['mailFolder_id1'] = $id;
        return new MailFolderRequestBuilder($urlTplParams, $this->requestAdapter);
    }

    /**
     * Instantiates a new MailFolderRequestBuilder and sets the default values.
     * @param array $pathParameters Path parameters for the request
     * @param RequestAdapter $requestAdapter The request adapter to use to execute the requests.
    */
    public function __construct(array $pathParameters, RequestAdapter $requestAdapter) {
        $this->urlTemplate = '{+baseurl}/users/{user_id}/mailFolders/{mailFolder_id}{?select,expand}';
        $this->requestAdapter = $requestAdapter;
        $this->pathParameters = $pathParameters;
    }

    /**
     * The user's mail folders. Read-only. Nullable.
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @return RequestInformation
    */
    public function createDeleteRequestInformation(?array $headers, ?array $options): RequestInformation {
        $requestInfo = new RequestInformation();
        $requestInfo->urlTemplate = $this->urlTemplate;
        $requestInfo->pathParameters = $this->pathParameters;
        $requestInfo->httpMethod = HttpMethod::DELETE;
        $requestInfo->setHeadersFromRawObject($headers);
        $requestInfo->addRequestOptions(...$options);
        return $requestInfo;
    }

    /**
     * The user's mail folders. Read-only. Nullable.
     * @param GetQueryParameters|null $queryParameters Request query parameters
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @return RequestInformation
    */
    public function createGetRequestInformation(?GetQueryParameters $queryParameters, ?array $headers, ?array $options): RequestInformation {
        $requestInfo = new RequestInformation();
        $requestInfo->urlTemplate = $this->urlTemplate;
        $requestInfo->pathParameters = $this->pathParameters;
        $requestInfo->httpMethod = HttpMethod::GET;
        $requestInfo->setHeadersFromRawObject($headers);
        $requestInfo->setQueryStringParametersFromRawObject($queryParameters);
        $requestInfo->addRequestOptions(...$options);
        return $requestInfo;
    }

    /**
     * The user's mail folders. Read-only. Nullable.
     * @param MailFolder $body 
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @return RequestInformation
    */
    public function createPatchRequestInformation(MailFolder $body, ?array $headers, ?array $options): RequestInformation {
        $requestInfo = new RequestInformation();
        $requestInfo->urlTemplate = $this->urlTemplate;
        $requestInfo->pathParameters = $this->pathParameters;
        $requestInfo->httpMethod = HttpMethod::PATCH;
        $requestInfo->setHeadersFromRawObject($headers);
        $requestInfo->setContentFromParsable($this->requestAdapter, "application/json", $body);
        $requestInfo->addRequestOptions(...$options);
        return $requestInfo;
    }

    /**
     * The user's mail folders. Read-only. Nullable.
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @param ResponseHandler|null $responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return Promise
    */
    public function delete(?array $headers, ?array $options, ?ResponseHandler $responseHandler): Promise {
        $requestInfo = $this->createDeleteRequestInformation($headers, $options);
        return $this->requestAdapter->sendAsync($requestInfo, get_class($body), $responseHandler);
    }

    /**
     * The user's mail folders. Read-only. Nullable.
     * @param GetQueryParameters|null $queryParameters Request query parameters
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @param ResponseHandler|null $responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return Promise
    */
    public function get(?GetQueryParameters $queryParameters, ?array $headers, ?array $options, ?ResponseHandler $responseHandler): Promise {
        $requestInfo = $this->createGetRequestInformation($queryParameters, $headers, $options);
        return $this->requestAdapter->sendAsync($requestInfo, get_class($body), $responseHandler);
    }

    /**
     * Gets an item from the Microsoft\Graph.users.item.mailFolders.item.messageRules.item collection
     * @param string $id Unique identifier of the item
     * @return MessageRuleRequestBuilder
    */
    public function messageRulesById(string $id): MessageRuleRequestBuilder {
        $urlTplParams = $this->pathParameters;
        $urlTplParams['messageRule_id'] = $id;
        return new MessageRuleRequestBuilder($urlTplParams, $this->requestAdapter);
    }

    /**
     * Gets an item from the Microsoft\Graph.users.item.mailFolders.item.messages.item collection
     * @param string $id Unique identifier of the item
     * @return MessageRequestBuilder
    */
    public function messagesById(string $id): MessageRequestBuilder {
        $urlTplParams = $this->pathParameters;
        $urlTplParams['message_id'] = $id;
        return new MessageRequestBuilder($urlTplParams, $this->requestAdapter);
    }

    /**
     * Gets an item from the Microsoft\Graph.users.item.mailFolders.item.multiValueExtendedProperties.item collection
     * @param string $id Unique identifier of the item
     * @return MultiValueLegacyExtendedPropertyRequestBuilder
    */
    public function multiValueExtendedPropertiesById(string $id): MultiValueLegacyExtendedPropertyRequestBuilder {
        $urlTplParams = $this->pathParameters;
        $urlTplParams['multiValueLegacyExtendedProperty_id'] = $id;
        return new MultiValueLegacyExtendedPropertyRequestBuilder($urlTplParams, $this->requestAdapter);
    }

    /**
     * The user's mail folders. Read-only. Nullable.
     * @param MailFolder $body 
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @param ResponseHandler|null $responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return Promise
    */
    public function patch(MailFolder $body, ?array $headers, ?array $options, ?ResponseHandler $responseHandler): Promise {
        $requestInfo = $this->createPatchRequestInformation($body, $headers, $options);
        return $this->requestAdapter->sendAsync($requestInfo, get_class($body), $responseHandler);
    }

    /**
     * Gets an item from the Microsoft\Graph.users.item.mailFolders.item.singleValueExtendedProperties.item collection
     * @param string $id Unique identifier of the item
     * @return SingleValueLegacyExtendedPropertyRequestBuilder
    */
    public function singleValueExtendedPropertiesById(string $id): SingleValueLegacyExtendedPropertyRequestBuilder {
        $urlTplParams = $this->pathParameters;
        $urlTplParams['singleValueLegacyExtendedProperty_id'] = $id;
        return new SingleValueLegacyExtendedPropertyRequestBuilder($urlTplParams, $this->requestAdapter);
    }

}
