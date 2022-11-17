import {MessageRulesRequestBuilderGetQueryParameters} from './messageRulesRequestBuilderGetQueryParameters';
import {RequestOption} from '@microsoft/kiota-abstractions';

<<<<<<< HEAD
export interface MessageRulesRequestBuilderGetRequestConfiguration {
    /** Request headers */
    headers?: Record<string, string> | undefined;
    /** Request options */
    options?: RequestOption[] | undefined;
    /** Request query parameters */
    queryParameters?: MessageRulesRequestBuilderGetQueryParameters | undefined;
=======
/** Configuration for the request such as headers, query parameters, and middleware options. */
export class MessageRulesRequestBuilderGetRequestConfiguration {
    /** Request headers */
    public headers?: Record<string, string> | undefined;
    /** Request options */
    public options?: RequestOption[] | undefined;
    /** Request query parameters */
    public queryParameters?: MessageRulesRequestBuilderGetQueryParameters | undefined;
>>>>>>> main
}
