import {MessageRulesRequestBuilderGetQueryParameters} from './messageRulesRequestBuilderGetQueryParameters';
import {RequestOption} from '@microsoft/kiota-abstractions';

export interface MessageRulesRequestBuilderGetRequestConfiguration {
    /** Request headers */
    headers?: Record<string, string> | undefined;
    /** Request options */
    options?: RequestOption[] | undefined;
    /** Request query parameters */
    queryParameters?: MessageRulesRequestBuilderGetQueryParameters | undefined;
}
