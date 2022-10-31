/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../registry/params";
import { Registry } from "../registry/registry";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "pricechain.rd_net.registry";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetRegistryRequest {
  id: number;
}

export interface QueryGetRegistryResponse {
  Registry: Registry | undefined;
}

export interface QueryAllRegistryRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllRegistryResponse {
  Registry: Registry[];
  pagination: PageResponse | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetRegistryRequest: object = { id: 0 };

export const QueryGetRegistryRequest = {
  encode(
    message: QueryGetRegistryRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetRegistryRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetRegistryRequest,
    } as QueryGetRegistryRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetRegistryRequest {
    const message = {
      ...baseQueryGetRegistryRequest,
    } as QueryGetRegistryRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetRegistryRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetRegistryRequest>
  ): QueryGetRegistryRequest {
    const message = {
      ...baseQueryGetRegistryRequest,
    } as QueryGetRegistryRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetRegistryResponse: object = {};

export const QueryGetRegistryResponse = {
  encode(
    message: QueryGetRegistryResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.Registry !== undefined) {
      Registry.encode(message.Registry, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetRegistryResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetRegistryResponse,
    } as QueryGetRegistryResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Registry = Registry.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetRegistryResponse {
    const message = {
      ...baseQueryGetRegistryResponse,
    } as QueryGetRegistryResponse;
    if (object.Registry !== undefined && object.Registry !== null) {
      message.Registry = Registry.fromJSON(object.Registry);
    } else {
      message.Registry = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetRegistryResponse): unknown {
    const obj: any = {};
    message.Registry !== undefined &&
      (obj.Registry = message.Registry
        ? Registry.toJSON(message.Registry)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetRegistryResponse>
  ): QueryGetRegistryResponse {
    const message = {
      ...baseQueryGetRegistryResponse,
    } as QueryGetRegistryResponse;
    if (object.Registry !== undefined && object.Registry !== null) {
      message.Registry = Registry.fromPartial(object.Registry);
    } else {
      message.Registry = undefined;
    }
    return message;
  },
};

const baseQueryAllRegistryRequest: object = {};

export const QueryAllRegistryRequest = {
  encode(
    message: QueryAllRegistryRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllRegistryRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllRegistryRequest,
    } as QueryAllRegistryRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllRegistryRequest {
    const message = {
      ...baseQueryAllRegistryRequest,
    } as QueryAllRegistryRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllRegistryRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllRegistryRequest>
  ): QueryAllRegistryRequest {
    const message = {
      ...baseQueryAllRegistryRequest,
    } as QueryAllRegistryRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllRegistryResponse: object = {};

export const QueryAllRegistryResponse = {
  encode(
    message: QueryAllRegistryResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.Registry) {
      Registry.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllRegistryResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllRegistryResponse,
    } as QueryAllRegistryResponse;
    message.Registry = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Registry.push(Registry.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllRegistryResponse {
    const message = {
      ...baseQueryAllRegistryResponse,
    } as QueryAllRegistryResponse;
    message.Registry = [];
    if (object.Registry !== undefined && object.Registry !== null) {
      for (const e of object.Registry) {
        message.Registry.push(Registry.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllRegistryResponse): unknown {
    const obj: any = {};
    if (message.Registry) {
      obj.Registry = message.Registry.map((e) =>
        e ? Registry.toJSON(e) : undefined
      );
    } else {
      obj.Registry = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllRegistryResponse>
  ): QueryAllRegistryResponse {
    const message = {
      ...baseQueryAllRegistryResponse,
    } as QueryAllRegistryResponse;
    message.Registry = [];
    if (object.Registry !== undefined && object.Registry !== null) {
      for (const e of object.Registry) {
        message.Registry.push(Registry.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a Registry by id. */
  Registry(request: QueryGetRegistryRequest): Promise<QueryGetRegistryResponse>;
  /** Queries a list of Registry items. */
  RegistryAll(
    request: QueryAllRegistryRequest
  ): Promise<QueryAllRegistryResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "pricechain.rd_net.registry.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  Registry(
    request: QueryGetRegistryRequest
  ): Promise<QueryGetRegistryResponse> {
    const data = QueryGetRegistryRequest.encode(request).finish();
    const promise = this.rpc.request(
      "pricechain.rd_net.registry.Query",
      "Registry",
      data
    );
    return promise.then((data) =>
      QueryGetRegistryResponse.decode(new Reader(data))
    );
  }

  RegistryAll(
    request: QueryAllRegistryRequest
  ): Promise<QueryAllRegistryResponse> {
    const data = QueryAllRegistryRequest.encode(request).finish();
    const promise = this.rpc.request(
      "pricechain.rd_net.registry.Query",
      "RegistryAll",
      data
    );
    return promise.then((data) =>
      QueryAllRegistryResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
