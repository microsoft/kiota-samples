import {MessageRuleItemRequestBuilderGetQueryParameters} from './messageRuleItemRequestBuilderGetQueryParameters';
import {RequestOption} from '@microsoft/kiota-abstractions';

<<<<<<< HEAD
export interface MessageRuleItemRequestBuilderGetRequestConfiguration {
    /** Request headers */
    headers?: Record<string, string> | undefined;
    /** Request options */
    options?: RequestOption[] | undefined;
    /** Request query parameters */
    queryParameters?: MessageRuleItemRequestBuilderGetQueryParameters | undefined;
=======
/** Configuration for the request such as headers, query parameters, and middleware options. */
export class MessageRuleItemRequestBuilderGetRequestConfiguration {
    /** Request headers */
    public headers?: Record<string, string> | undefined;
    /** Request options */
    public options?: RequestOption[] | undefined;
    /** Request query parameters */
    public queryParameters?: MessageRuleItemRequestBuilderGetQueryParameters | undefined;
>>>>>>> main
}
