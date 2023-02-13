/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../prcibc/params";
import { IbcMsg } from "../prcibc/ibc_msg";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "pricechain.cprc.prcibc";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetIbcMsgRequest {
  id: number;
}

export interface QueryGetIbcMsgResponse {
  IbcMsg: IbcMsg | undefined;
}

export interface QueryAllIbcMsgRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllIbcMsgResponse {
  IbcMsg: IbcMsg[];
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

const baseQueryGetIbcMsgRequest: object = { id: 0 };

export const QueryGetIbcMsgRequest = {
  encode(
    message: QueryGetIbcMsgRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetIbcMsgRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetIbcMsgRequest } as QueryGetIbcMsgRequest;
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

  fromJSON(object: any): QueryGetIbcMsgRequest {
    const message = { ...baseQueryGetIbcMsgRequest } as QueryGetIbcMsgRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetIbcMsgRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetIbcMsgRequest>
  ): QueryGetIbcMsgRequest {
    const message = { ...baseQueryGetIbcMsgRequest } as QueryGetIbcMsgRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetIbcMsgResponse: object = {};

export const QueryGetIbcMsgResponse = {
  encode(
    message: QueryGetIbcMsgResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.IbcMsg !== undefined) {
      IbcMsg.encode(message.IbcMsg, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetIbcMsgResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetIbcMsgResponse } as QueryGetIbcMsgResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.IbcMsg = IbcMsg.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetIbcMsgResponse {
    const message = { ...baseQueryGetIbcMsgResponse } as QueryGetIbcMsgResponse;
    if (object.IbcMsg !== undefined && object.IbcMsg !== null) {
      message.IbcMsg = IbcMsg.fromJSON(object.IbcMsg);
    } else {
      message.IbcMsg = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetIbcMsgResponse): unknown {
    const obj: any = {};
    message.IbcMsg !== undefined &&
      (obj.IbcMsg = message.IbcMsg ? IbcMsg.toJSON(message.IbcMsg) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetIbcMsgResponse>
  ): QueryGetIbcMsgResponse {
    const message = { ...baseQueryGetIbcMsgResponse } as QueryGetIbcMsgResponse;
    if (object.IbcMsg !== undefined && object.IbcMsg !== null) {
      message.IbcMsg = IbcMsg.fromPartial(object.IbcMsg);
    } else {
      message.IbcMsg = undefined;
    }
    return message;
  },
};

const baseQueryAllIbcMsgRequest: object = {};

export const QueryAllIbcMsgRequest = {
  encode(
    message: QueryAllIbcMsgRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllIbcMsgRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllIbcMsgRequest } as QueryAllIbcMsgRequest;
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

  fromJSON(object: any): QueryAllIbcMsgRequest {
    const message = { ...baseQueryAllIbcMsgRequest } as QueryAllIbcMsgRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllIbcMsgRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllIbcMsgRequest>
  ): QueryAllIbcMsgRequest {
    const message = { ...baseQueryAllIbcMsgRequest } as QueryAllIbcMsgRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllIbcMsgResponse: object = {};

export const QueryAllIbcMsgResponse = {
  encode(
    message: QueryAllIbcMsgResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.IbcMsg) {
      IbcMsg.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllIbcMsgResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllIbcMsgResponse } as QueryAllIbcMsgResponse;
    message.IbcMsg = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.IbcMsg.push(IbcMsg.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllIbcMsgResponse {
    const message = { ...baseQueryAllIbcMsgResponse } as QueryAllIbcMsgResponse;
    message.IbcMsg = [];
    if (object.IbcMsg !== undefined && object.IbcMsg !== null) {
      for (const e of object.IbcMsg) {
        message.IbcMsg.push(IbcMsg.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllIbcMsgResponse): unknown {
    const obj: any = {};
    if (message.IbcMsg) {
      obj.IbcMsg = message.IbcMsg.map((e) =>
        e ? IbcMsg.toJSON(e) : undefined
      );
    } else {
      obj.IbcMsg = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllIbcMsgResponse>
  ): QueryAllIbcMsgResponse {
    const message = { ...baseQueryAllIbcMsgResponse } as QueryAllIbcMsgResponse;
    message.IbcMsg = [];
    if (object.IbcMsg !== undefined && object.IbcMsg !== null) {
      for (const e of object.IbcMsg) {
        message.IbcMsg.push(IbcMsg.fromPartial(e));
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
  /** Queries a IbcMsg by id. */
  IbcMsg(request: QueryGetIbcMsgRequest): Promise<QueryGetIbcMsgResponse>;
  /** Queries a list of IbcMsg items. */
  IbcMsgAll(request: QueryAllIbcMsgRequest): Promise<QueryAllIbcMsgResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "pricechain.cprc.prcibc.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  IbcMsg(request: QueryGetIbcMsgRequest): Promise<QueryGetIbcMsgResponse> {
    const data = QueryGetIbcMsgRequest.encode(request).finish();
    const promise = this.rpc.request(
      "pricechain.cprc.prcibc.Query",
      "IbcMsg",
      data
    );
    return promise.then((data) =>
      QueryGetIbcMsgResponse.decode(new Reader(data))
    );
  }

  IbcMsgAll(request: QueryAllIbcMsgRequest): Promise<QueryAllIbcMsgResponse> {
    const data = QueryAllIbcMsgRequest.encode(request).finish();
    const promise = this.rpc.request(
      "pricechain.cprc.prcibc.Query",
      "IbcMsgAll",
      data
    );
    return promise.then((data) =>
      QueryAllIbcMsgResponse.decode(new Reader(data))
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
