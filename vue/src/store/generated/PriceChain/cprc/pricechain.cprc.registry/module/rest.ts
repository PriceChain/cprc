/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface ProtobufAny {
  "@type"?: string;
}

export type RegistryMsgCreateRegistryResponse = object;

export type RegistryMsgJoinRegistryCoOperatorResponse = object;

export type RegistryMsgJoinRegistryMemberResponse = object;

export type RegistryMsgModifyRegistryResponse = object;

export type RegistryMsgProposePriceResponse = object;

export type RegistryMsgUnbondRegistryResponse = object;

/**
 * Params defines the parameters for the module.
 */
export interface RegistryParams {
  minimumStakeAmount?: string;
}

export interface RegistryQueryAllRegistryMemberResponse {
  RegistryMember?: RegistryRegistryMember[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface RegistryQueryAllRegistryOwnerResponse {
  RegistryOwner?: RegistryRegistryOwner[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface RegistryQueryAllRegistryResponse {
  Registry?: RegistryRegistry[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface RegistryQueryAllRegistryStakedAmountResponse {
  registryStakedAmount?: RegistryRegistryStakedAmount[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface RegistryQueryAllStakedAmountPerWalletResponse {
  stakedAmountPerWallet?: RegistryStakedAmountPerWallet[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface RegistryQueryGetRegistryMemberResponse {
  RegistryMember?: RegistryRegistryMember;
}

export interface RegistryQueryGetRegistryOwnerResponse {
  RegistryOwner?: RegistryRegistryOwner;
}

export interface RegistryQueryGetRegistryResponse {
  Registry?: RegistryRegistry;
}

export interface RegistryQueryGetRegistryStakedAmountResponse {
  registryStakedAmount?: RegistryRegistryStakedAmount;
}

export interface RegistryQueryGetStakedAmountPerWalletResponse {
  stakedAmountPerWallet?: RegistryStakedAmountPerWallet;
}

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface RegistryQueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: RegistryParams;
}

export interface RegistryRegistry {
  /** @format uint64 */
  id?: string;
  name?: string;
  stakedAmount?: string;
  status?: string;
  description?: string;
  imageUrl?: string;
  priceCount?: string;
  reviewCount?: string;
  timestamp?: string;
  reserved?: string;
  owners?: string[];
}

export interface RegistryRegistryMember {
  /** @format uint64 */
  id?: string;
  registryId?: string;
  stakedAmount?: string;
  address?: string;
  status?: string;
  memberSince?: string;
  memo?: string;
  reserved?: string;
}

export interface RegistryRegistryOwner {
  /** @format uint64 */
  id?: string;
  registryId?: string;
  stakedAmount?: string;
  address?: string;
  status?: string;
  reserved?: string;
}

export interface RegistryRegistryStakedAmount {
  index?: string;
  amount?: string;
}

export interface RegistryStakedAmountPerWallet {
  index?: string;
  wallet?: string;
  stakedAmount?: string;
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
}

/**
* message SomeRequest {
         Foo some_parameter = 1;
         PageRequest pagination = 2;
 }
*/
export interface V1Beta1PageRequest {
  /**
   * key is a value returned in PageResponse.next_key to begin
   * querying the next page most efficiently. Only one of offset or key
   * should be set.
   * @format byte
   */
  key?: string;

  /**
   * offset is a numeric offset that can be used when key is unavailable.
   * It is less efficient than using key. Only one of offset or key should
   * be set.
   * @format uint64
   */
  offset?: string;

  /**
   * limit is the total number of results to be returned in the result page.
   * If left empty it will default to a value to be set by each app.
   * @format uint64
   */
  limit?: string;

  /**
   * count_total is set to true  to indicate that the result set should include
   * a count of the total number of items available for pagination in UIs.
   * count_total is only respected when offset is used. It is ignored when key
   * is set.
   */
  count_total?: boolean;

  /**
   * reverse is set to true if results are to be returned in the descending order.
   *
   * Since: cosmos-sdk 0.43
   */
  reverse?: boolean;
}

/**
* PageResponse is to be embedded in gRPC response messages where the
corresponding request message has used PageRequest.

 message SomeResponse {
         repeated Bar results = 1;
         PageResponse page = 2;
 }
*/
export interface V1Beta1PageResponse {
  /** @format byte */
  next_key?: string;

  /** @format uint64 */
  total?: string;
}

export type QueryParamsType = Record<string | number, any>;
export type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;

export interface FullRequestParams extends Omit<RequestInit, "body"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: keyof Omit<Body, "body" | "bodyUsed">;
  /** request body */
  body?: unknown;
  /** base url */
  baseUrl?: string;
  /** request cancellation token */
  cancelToken?: CancelToken;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> {
  baseUrl?: string;
  baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
  securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
  data: D;
  error: E;
}

type CancelToken = Symbol | string | number;

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public baseUrl: string = "";
  private securityData: SecurityDataType = null as any;
  private securityWorker: null | ApiConfig<SecurityDataType>["securityWorker"] = null;
  private abortControllers = new Map<CancelToken, AbortController>();

  private baseApiParams: RequestParams = {
    credentials: "same-origin",
    headers: {},
    redirect: "follow",
    referrerPolicy: "no-referrer",
  };

  constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
    Object.assign(this, apiConfig);
  }

  public setSecurityData = (data: SecurityDataType) => {
    this.securityData = data;
  };

  private addQueryParam(query: QueryParamsType, key: string) {
    const value = query[key];

    return (
      encodeURIComponent(key) +
      "=" +
      encodeURIComponent(Array.isArray(value) ? value.join(",") : typeof value === "number" ? value : `${value}`)
    );
  }

  protected toQueryString(rawQuery?: QueryParamsType): string {
    const query = rawQuery || {};
    const keys = Object.keys(query).filter((key) => "undefined" !== typeof query[key]);
    return keys
      .map((key) =>
        typeof query[key] === "object" && !Array.isArray(query[key])
          ? this.toQueryString(query[key] as QueryParamsType)
          : this.addQueryParam(query, key),
      )
      .join("&");
  }

  protected addQueryParams(rawQuery?: QueryParamsType): string {
    const queryString = this.toQueryString(rawQuery);
    return queryString ? `?${queryString}` : "";
  }

  private contentFormatters: Record<ContentType, (input: any) => any> = {
    [ContentType.Json]: (input: any) =>
      input !== null && (typeof input === "object" || typeof input === "string") ? JSON.stringify(input) : input,
    [ContentType.FormData]: (input: any) =>
      Object.keys(input || {}).reduce((data, key) => {
        data.append(key, input[key]);
        return data;
      }, new FormData()),
    [ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
  };

  private mergeRequestParams(params1: RequestParams, params2?: RequestParams): RequestParams {
    return {
      ...this.baseApiParams,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.baseApiParams.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createAbortSignal = (cancelToken: CancelToken): AbortSignal | undefined => {
    if (this.abortControllers.has(cancelToken)) {
      const abortController = this.abortControllers.get(cancelToken);
      if (abortController) {
        return abortController.signal;
      }
      return void 0;
    }

    const abortController = new AbortController();
    this.abortControllers.set(cancelToken, abortController);
    return abortController.signal;
  };

  public abortRequest = (cancelToken: CancelToken) => {
    const abortController = this.abortControllers.get(cancelToken);

    if (abortController) {
      abortController.abort();
      this.abortControllers.delete(cancelToken);
    }
  };

  public request = <T = any, E = any>({
    body,
    secure,
    path,
    type,
    query,
    format = "json",
    baseUrl,
    cancelToken,
    ...params
  }: FullRequestParams): Promise<HttpResponse<T, E>> => {
    const secureParams = (secure && this.securityWorker && this.securityWorker(this.securityData)) || {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const queryString = query && this.toQueryString(query);
    const payloadFormatter = this.contentFormatters[type || ContentType.Json];

    return fetch(`${baseUrl || this.baseUrl || ""}${path}${queryString ? `?${queryString}` : ""}`, {
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      signal: cancelToken ? this.createAbortSignal(cancelToken) : void 0,
      body: typeof body === "undefined" || body === null ? null : payloadFormatter(body),
    }).then(async (response) => {
      const r = response as HttpResponse<T, E>;
      r.data = (null as unknown) as T;
      r.error = (null as unknown) as E;

      const data = await response[format]()
        .then((data) => {
          if (r.ok) {
            r.data = data;
          } else {
            r.error = data;
          }
          return r;
        })
        .catch((e) => {
          r.error = e;
          return r;
        });

      if (cancelToken) {
        this.abortControllers.delete(cancelToken);
      }

      if (!response.ok) throw data;
      return data;
    });
  };
}

/**
 * @title registry/genesis.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryParams
   * @summary Parameters queries the parameters of the module.
   * @request GET:/PriceChain/cprc/registry/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<RegistryQueryParamsResponse, RpcStatus>({
      path: `/PriceChain/cprc/registry/params`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryRegistryAll
   * @summary Queries a list of Registry items.
   * @request GET:/PriceChain/cprc/registry/registry
   */
  queryRegistryAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<RegistryQueryAllRegistryResponse, RpcStatus>({
      path: `/PriceChain/cprc/registry/registry`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryRegistry
   * @summary Queries a Registry by id.
   * @request GET:/PriceChain/cprc/registry/registry/{id}
   */
  queryRegistry = (id: string, params: RequestParams = {}) =>
    this.request<RegistryQueryGetRegistryResponse, RpcStatus>({
      path: `/PriceChain/cprc/registry/registry/${id}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryRegistryMemberAll
   * @summary Queries a list of RegistryMember items.
   * @request GET:/PriceChain/cprc/registry/registry_member
   */
  queryRegistryMemberAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<RegistryQueryAllRegistryMemberResponse, RpcStatus>({
      path: `/PriceChain/cprc/registry/registry_member`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryRegistryMember
   * @summary Queries a RegistryMember by id.
   * @request GET:/PriceChain/cprc/registry/registry_member/{id}
   */
  queryRegistryMember = (id: string, params: RequestParams = {}) =>
    this.request<RegistryQueryGetRegistryMemberResponse, RpcStatus>({
      path: `/PriceChain/cprc/registry/registry_member/${id}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryRegistryOwnerAll
   * @summary Queries a list of RegistryOwner items.
   * @request GET:/PriceChain/cprc/registry/registry_owner
   */
  queryRegistryOwnerAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<RegistryQueryAllRegistryOwnerResponse, RpcStatus>({
      path: `/PriceChain/cprc/registry/registry_owner`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryRegistryOwner
   * @summary Queries a RegistryOwner by id.
   * @request GET:/PriceChain/cprc/registry/registry_owner/{id}
   */
  queryRegistryOwner = (id: string, params: RequestParams = {}) =>
    this.request<RegistryQueryGetRegistryOwnerResponse, RpcStatus>({
      path: `/PriceChain/cprc/registry/registry_owner/${id}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryRegistryStakedAmountAll
   * @summary Queries a list of RegistryStakedAmount items.
   * @request GET:/PriceChain/cprc/registry/registry_staked_amount
   */
  queryRegistryStakedAmountAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<RegistryQueryAllRegistryStakedAmountResponse, RpcStatus>({
      path: `/PriceChain/cprc/registry/registry_staked_amount`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryRegistryStakedAmount
   * @summary Queries a RegistryStakedAmount by index.
   * @request GET:/PriceChain/cprc/registry/registry_staked_amount/{index}
   */
  queryRegistryStakedAmount = (index: string, params: RequestParams = {}) =>
    this.request<RegistryQueryGetRegistryStakedAmountResponse, RpcStatus>({
      path: `/PriceChain/cprc/registry/registry_staked_amount/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryStakedAmountPerWalletAll
   * @summary Queries a list of StakedAmountPerWallet items.
   * @request GET:/PriceChain/cprc/registry/staked_amount_per_wallet
   */
  queryStakedAmountPerWalletAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<RegistryQueryAllStakedAmountPerWalletResponse, RpcStatus>({
      path: `/PriceChain/cprc/registry/staked_amount_per_wallet`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryStakedAmountPerWallet
   * @summary Queries a StakedAmountPerWallet by index.
   * @request GET:/PriceChain/cprc/registry/staked_amount_per_wallet/{index}
   */
  queryStakedAmountPerWallet = (index: string, params: RequestParams = {}) =>
    this.request<RegistryQueryGetStakedAmountPerWalletResponse, RpcStatus>({
      path: `/PriceChain/cprc/registry/staked_amount_per_wallet/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });
}
