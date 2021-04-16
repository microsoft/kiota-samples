import {HttpCore, HttpMethod, RequestInfo, ResponseHandler} from '@microsoft/kiota-abstractions';
import {Message} from '../../../message';
import {AttachmentsRequestBuilder} from '../attachments/attachmentsRequestBuilder';
import {AttachmentRequestBuilder} from '../attachments/item/attachmentRequestBuilder';
import {ContentRequestBuilder} from '../content/contentRequestBuilder';
import {ExtensionsRequestBuilder} from '../extensions/extensionsRequestBuilder';
import {ExtensionRequestBuilder} from '../extensions/item/extensionRequestBuilder';
import {MultiValueLegacyExtendedPropertyRequestBuilder} from '../multiValueExtendedProperties/item/multiValueLegacyExtendedPropertyRequestBuilder';
import {MultiValueExtendedPropertiesRequestBuilder} from '../multiValueExtendedProperties/multiValueExtendedPropertiesRequestBuilder';
import {SingleValueLegacyExtendedPropertyRequestBuilder} from '../singleValueExtendedProperties/item/singleValueLegacyExtendedPropertyRequestBuilder';
import {SingleValueExtendedPropertiesRequestBuilder} from '../singleValueExtendedProperties/singleValueExtendedPropertiesRequestBuilder';

export class MessageRequestBuilder {
    public get content(): ContentRequestBuilder {
        const builder = new ContentRequestBuilder();
        builder.currentPath = (this.currentPath && this.currentPath) + this.pathSegment;
        return builder;
    }
    public get attachments(): AttachmentsRequestBuilder {
        const builder = new AttachmentsRequestBuilder();
        builder.currentPath = (this.currentPath && this.currentPath) + this.pathSegment;
        return builder;
    }
    public get extensions(): ExtensionsRequestBuilder {
        const builder = new ExtensionsRequestBuilder();
        builder.currentPath = (this.currentPath && this.currentPath) + this.pathSegment;
        return builder;
    }
    public get multiValueExtendedProperties(): MultiValueExtendedPropertiesRequestBuilder {
        const builder = new MultiValueExtendedPropertiesRequestBuilder();
        builder.currentPath = (this.currentPath && this.currentPath) + this.pathSegment;
        return builder;
    }
    public get singleValueExtendedProperties(): SingleValueExtendedPropertiesRequestBuilder {
        const builder = new SingleValueExtendedPropertiesRequestBuilder();
        builder.currentPath = (this.currentPath && this.currentPath) + this.pathSegment;
        return builder;
    }
    public readonly get = async (q?: {
                    select?: string[],
                    expand?: string[]
                    } | undefined, h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<Message | undefined> => {
        const requestInfo = {
            URI: this.currentPath ? new URL(this.currentPath): null,
            headers: h,
            httpMethod: HttpMethod.GET,
            queryParameters: q,
        } as RequestInfo;
        return await this.httpCore?.sendAsync<Message>(requestInfo, responseHandler);
    }
    public readonly patch = async (h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<object | undefined> => {
        const requestInfo = {
            URI: this.currentPath ? new URL(this.currentPath): null,
            headers: h,
            httpMethod: HttpMethod.PATCH,
        } as RequestInfo;
        return await this.httpCore?.sendAsync<object>(requestInfo, responseHandler);
    }
    public readonly delete = async (h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<object | undefined> => {
        const requestInfo = {
            URI: this.currentPath ? new URL(this.currentPath): null,
            headers: h,
            httpMethod: HttpMethod.DELETE,
        } as RequestInfo;
        return await this.httpCore?.sendAsync<object>(requestInfo, responseHandler);
    }
    private readonly pathSegment: string = "";
    public currentPath?: string | undefined;
    public httpCore?: HttpCore | undefined;
    public readonly attachmentsById = (id: String) : AttachmentRequestBuilder => {
        const builder = new AttachmentRequestBuilder();
        builder.currentPath = (this.currentPath && this.currentPath) + this.pathSegment + "/attachments/" + id;
        return builder;
    }
    public readonly extensionsById = (id: String) : ExtensionRequestBuilder => {
        const builder = new ExtensionRequestBuilder();
        builder.currentPath = (this.currentPath && this.currentPath) + this.pathSegment + "/extensions/" + id;
        return builder;
    }
    public readonly multiValueExtendedPropertiesById = (id: String) : MultiValueLegacyExtendedPropertyRequestBuilder => {
        const builder = new MultiValueLegacyExtendedPropertyRequestBuilder();
        builder.currentPath = (this.currentPath && this.currentPath) + this.pathSegment + "/multiValueExtendedProperties/" + id;
        return builder;
    }
    public readonly singleValueExtendedPropertiesById = (id: String) : SingleValueLegacyExtendedPropertyRequestBuilder => {
        const builder = new SingleValueLegacyExtendedPropertyRequestBuilder();
        builder.currentPath = (this.currentPath && this.currentPath) + this.pathSegment + "/singleValueExtendedProperties/" + id;
        return builder;
    }
}
