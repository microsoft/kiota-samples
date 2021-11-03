<?php

namespace Microsoft\Graph\Users\Item\MailFolders\Item\MessageRules;

use Microsoft\Graph\Models\Microsoft\Graph\MessageRule;
use Microsoft\Graph\Users\Item\MailFolders\Item\MessageRules\Item\MessageRuleRequestBuilder;
use Microsoft\Kiota\Abstractions\HttpMethod;
use Microsoft\Kiota\Abstractions\MiddlewareOption;
use Microsoft\Kiota\Abstractions\RequestAdapter;
use Microsoft\Kiota\Abstractions\RequestInformation;
use Microsoft\Kiota\Abstractions\ResponseHandler;
use Microsoft\Kiota\Abstractions\Serialization\Parsable;

class MessageRulesRequestBuilder 
{
    /** @var array $pathParameters Path parameters for the request */
    private array $pathParameters;
    
    /** @var IRequestAdapter $requestAdapter The request adapter to use to execute the requests. */
    private RequestAdapter $requestAdapter;
    
    /** @var string $urlTemplate Url template to use to build the URL for the current request builder */
    private string $urlTemplate;
    
    /**
     * Instantiates a new MessageRulesRequestBuilder and sets the default values.
     * @param array $pathParameters Path parameters for the request
     * @param RequestAdapter $requestAdapter The request adapter to use to execute the requests.
    */
    public function __construct(array $pathParameters, RequestAdapter $requestAdapter) {
        $this->urlTemplate = 'https://graph.microsoft.com/v1.0/users/{user_id}/mailFolders/{mailFolder_id}/messageRules{?top,skip,search,filter,count,orderby,select,expand}';
        $this->requestAdapter = $requestAdapter;
        $this->pathParameters = $pathParameters;
    }

    /**
     * The collection of rules that apply to the user's Inbox folder.
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
     * The collection of rules that apply to the user's Inbox folder.
     * @param MessageRule $body 
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @return RequestInformation
    */
    public function createPostRequestInformation(MessageRule $body, ?array $headers, ?array $options): RequestInformation {
        $requestInfo = new RequestInformation();
        $requestInfo->urlTemplate = $this->urlTemplate;
        $requestInfo->pathParameters = $this->pathParameters;
        $requestInfo->httpMethod = HttpMethod::POST;
        $requestInfo->setHeadersFromRawObject($headers);
        $requestInfo->setContentFromParsable($this->requestAdapter, "application/json", $body);
        $requestInfo->addRequestOptions(...$options);
        return $requestInfo;
    }

    /**
     * The collection of rules that apply to the user's Inbox folder.
     * @param GetQueryParameters|null $queryParameters Request query parameters
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @param ResponseHandler|null $responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return MessageRulesResponse|null
    */
    public function get(?GetQueryParameters $queryParameters, ?array $headers, ?array $options, ?ResponseHandler $responseHandler): ?MessageRulesResponse {
    }

    /**
     * The collection of rules that apply to the user's Inbox folder.
     * @param MessageRule $body 
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @param ResponseHandler|null $responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return MessageRule|null
    */
    public function post(MessageRule $body, ?array $headers, ?array $options, ?ResponseHandler $responseHandler): ?MessageRule {
    }

}
