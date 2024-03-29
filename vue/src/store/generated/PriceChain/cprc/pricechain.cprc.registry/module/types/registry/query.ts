/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../registry/params";
import { Registry } from "../registry/registry";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { RegistryOwner } from "../registry/registry_owner";
import { RegistryMember } from "../registry/registry_member";
import { RegistryStakedAmount } from "../registry/registry_staked_amount";
import { RegistryStakedAmountPerWallet } from "../registry/registry_staked_amount_per_wallet";
import { PriceData } from "../registry/price_data";

export const protobufPackage = "pricechain.cprc.registry";

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

export interface QueryGetRegistryOwnerRequest {
  id: number;
}

export interface QueryGetRegistryOwnerResponse {
  RegistryOwner: RegistryOwner | undefined;
}

export interface QueryAllRegistryOwnerRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllRegistryOwnerResponse {
  RegistryOwner: RegistryOwner[];
  pagination: PageResponse | undefined;
}

export interface QueryGetRegistryMemberRequest {
  id: number;
}

export interface QueryGetRegistryMemberResponse {
  RegistryMember: RegistryMember | undefined;
}

export interface QueryAllRegistryMemberRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllRegistryMemberResponse {
  RegistryMember: RegistryMember[];
  pagination: PageResponse | undefined;
}

export interface QueryGetRegistryStakedAmountRequest {
  index: string;
}

export interface QueryGetRegistryStakedAmountResponse {
  registryStakedAmount: RegistryStakedAmount | undefined;
}

export interface QueryAllRegistryStakedAmountRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllRegistryStakedAmountResponse {
  registryStakedAmount: RegistryStakedAmount[];
  pagination: PageResponse | undefined;
}

export interface QueryGetRegistryStakedAmountPerWalletRequest {
  index: string;
}

export interface QueryGetRegistryStakedAmountPerWalletResponse {
  registryStakedAmountPerWallet: RegistryStakedAmountPerWallet | undefined;
}

export interface QueryAllRegistryStakedAmountPerWalletRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllRegistryStakedAmountPerWalletResponse {
  registryStakedAmountPerWallet: RegistryStakedAmountPerWallet[];
  pagination: PageResponse | undefined;
}

export interface QueryGetPriceDataRequest {
  index: string;
}

export interface QueryGetPriceDataResponse {
  priceData: PriceData | undefined;
}

export interface QueryAllPriceDataRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllPriceDataResponse {
  priceData: PriceData[];
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

const baseQueryGetRegistryOwnerRequest: object = { id: 0 };

export const QueryGetRegistryOwnerRequest = {
  encode(
    message: QueryGetRegistryOwnerRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetRegistryOwnerRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetRegistryOwnerRequest,
    } as QueryGetRegistryOwnerRequest;
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

  fromJSON(object: any): QueryGetRegistryOwnerRequest {
    const message = {
      ...baseQueryGetRegistryOwnerRequest,
    } as QueryGetRegistryOwnerRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetRegistryOwnerRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetRegistryOwnerRequest>
  ): QueryGetRegistryOwnerRequest {
    const message = {
      ...baseQueryGetRegistryOwnerRequest,
    } as QueryGetRegistryOwnerRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetRegistryOwnerResponse: object = {};

export const QueryGetRegistryOwnerResponse = {
  encode(
    message: QueryGetRegistryOwnerResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.RegistryOwner !== undefined) {
      RegistryOwner.encode(
        message.RegistryOwner,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetRegistryOwnerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetRegistryOwnerResponse,
    } as QueryGetRegistryOwnerResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.RegistryOwner = RegistryOwner.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetRegistryOwnerResponse {
    const message = {
      ...baseQueryGetRegistryOwnerResponse,
    } as QueryGetRegistryOwnerResponse;
    if (object.RegistryOwner !== undefined && object.RegistryOwner !== null) {
      message.RegistryOwner = RegistryOwner.fromJSON(object.RegistryOwner);
    } else {
      message.RegistryOwner = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetRegistryOwnerResponse): unknown {
    const obj: any = {};
    message.RegistryOwner !== undefined &&
      (obj.RegistryOwner = message.RegistryOwner
        ? RegistryOwner.toJSON(message.RegistryOwner)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetRegistryOwnerResponse>
  ): QueryGetRegistryOwnerResponse {
    const message = {
      ...baseQueryGetRegistryOwnerResponse,
    } as QueryGetRegistryOwnerResponse;
    if (object.RegistryOwner !== undefined && object.RegistryOwner !== null) {
      message.RegistryOwner = RegistryOwner.fromPartial(object.RegistryOwner);
    } else {
      message.RegistryOwner = undefined;
    }
    return message;
  },
};

const baseQueryAllRegistryOwnerRequest: object = {};

export const QueryAllRegistryOwnerRequest = {
  encode(
    message: QueryAllRegistryOwnerRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllRegistryOwnerRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllRegistryOwnerRequest,
    } as QueryAllRegistryOwnerRequest;
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

  fromJSON(object: any): QueryAllRegistryOwnerRequest {
    const message = {
      ...baseQueryAllRegistryOwnerRequest,
    } as QueryAllRegistryOwnerRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllRegistryOwnerRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllRegistryOwnerRequest>
  ): QueryAllRegistryOwnerRequest {
    const message = {
      ...baseQueryAllRegistryOwnerRequest,
    } as QueryAllRegistryOwnerRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllRegistryOwnerResponse: object = {};

export const QueryAllRegistryOwnerResponse = {
  encode(
    message: QueryAllRegistryOwnerResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.RegistryOwner) {
      RegistryOwner.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryAllRegistryOwnerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllRegistryOwnerResponse,
    } as QueryAllRegistryOwnerResponse;
    message.RegistryOwner = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.RegistryOwner.push(
            RegistryOwner.decode(reader, reader.uint32())
          );
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

  fromJSON(object: any): QueryAllRegistryOwnerResponse {
    const message = {
      ...baseQueryAllRegistryOwnerResponse,
    } as QueryAllRegistryOwnerResponse;
    message.RegistryOwner = [];
    if (object.RegistryOwner !== undefined && object.RegistryOwner !== null) {
      for (const e of object.RegistryOwner) {
        message.RegistryOwner.push(RegistryOwner.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllRegistryOwnerResponse): unknown {
    const obj: any = {};
    if (message.RegistryOwner) {
      obj.RegistryOwner = message.RegistryOwner.map((e) =>
        e ? RegistryOwner.toJSON(e) : undefined
      );
    } else {
      obj.RegistryOwner = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllRegistryOwnerResponse>
  ): QueryAllRegistryOwnerResponse {
    const message = {
      ...baseQueryAllRegistryOwnerResponse,
    } as QueryAllRegistryOwnerResponse;
    message.RegistryOwner = [];
    if (object.RegistryOwner !== undefined && object.RegistryOwner !== null) {
      for (const e of object.RegistryOwner) {
        message.RegistryOwner.push(RegistryOwner.fromPartial(e));
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

const baseQueryGetRegistryMemberRequest: object = { id: 0 };

export const QueryGetRegistryMemberRequest = {
  encode(
    message: QueryGetRegistryMemberRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetRegistryMemberRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetRegistryMemberRequest,
    } as QueryGetRegistryMemberRequest;
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

  fromJSON(object: any): QueryGetRegistryMemberRequest {
    const message = {
      ...baseQueryGetRegistryMemberRequest,
    } as QueryGetRegistryMemberRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetRegistryMemberRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetRegistryMemberRequest>
  ): QueryGetRegistryMemberRequest {
    const message = {
      ...baseQueryGetRegistryMemberRequest,
    } as QueryGetRegistryMemberRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetRegistryMemberResponse: object = {};

export const QueryGetRegistryMemberResponse = {
  encode(
    message: QueryGetRegistryMemberResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.RegistryMember !== undefined) {
      RegistryMember.encode(
        message.RegistryMember,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetRegistryMemberResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetRegistryMemberResponse,
    } as QueryGetRegistryMemberResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.RegistryMember = RegistryMember.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetRegistryMemberResponse {
    const message = {
      ...baseQueryGetRegistryMemberResponse,
    } as QueryGetRegistryMemberResponse;
    if (object.RegistryMember !== undefined && object.RegistryMember !== null) {
      message.RegistryMember = RegistryMember.fromJSON(object.RegistryMember);
    } else {
      message.RegistryMember = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetRegistryMemberResponse): unknown {
    const obj: any = {};
    message.RegistryMember !== undefined &&
      (obj.RegistryMember = message.RegistryMember
        ? RegistryMember.toJSON(message.RegistryMember)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetRegistryMemberResponse>
  ): QueryGetRegistryMemberResponse {
    const message = {
      ...baseQueryGetRegistryMemberResponse,
    } as QueryGetRegistryMemberResponse;
    if (object.RegistryMember !== undefined && object.RegistryMember !== null) {
      message.RegistryMember = RegistryMember.fromPartial(
        object.RegistryMember
      );
    } else {
      message.RegistryMember = undefined;
    }
    return message;
  },
};

const baseQueryAllRegistryMemberRequest: object = {};

export const QueryAllRegistryMemberRequest = {
  encode(
    message: QueryAllRegistryMemberRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllRegistryMemberRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllRegistryMemberRequest,
    } as QueryAllRegistryMemberRequest;
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

  fromJSON(object: any): QueryAllRegistryMemberRequest {
    const message = {
      ...baseQueryAllRegistryMemberRequest,
    } as QueryAllRegistryMemberRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllRegistryMemberRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllRegistryMemberRequest>
  ): QueryAllRegistryMemberRequest {
    const message = {
      ...baseQueryAllRegistryMemberRequest,
    } as QueryAllRegistryMemberRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllRegistryMemberResponse: object = {};

export const QueryAllRegistryMemberResponse = {
  encode(
    message: QueryAllRegistryMemberResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.RegistryMember) {
      RegistryMember.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryAllRegistryMemberResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllRegistryMemberResponse,
    } as QueryAllRegistryMemberResponse;
    message.RegistryMember = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.RegistryMember.push(
            RegistryMember.decode(reader, reader.uint32())
          );
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

  fromJSON(object: any): QueryAllRegistryMemberResponse {
    const message = {
      ...baseQueryAllRegistryMemberResponse,
    } as QueryAllRegistryMemberResponse;
    message.RegistryMember = [];
    if (object.RegistryMember !== undefined && object.RegistryMember !== null) {
      for (const e of object.RegistryMember) {
        message.RegistryMember.push(RegistryMember.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllRegistryMemberResponse): unknown {
    const obj: any = {};
    if (message.RegistryMember) {
      obj.RegistryMember = message.RegistryMember.map((e) =>
        e ? RegistryMember.toJSON(e) : undefined
      );
    } else {
      obj.RegistryMember = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllRegistryMemberResponse>
  ): QueryAllRegistryMemberResponse {
    const message = {
      ...baseQueryAllRegistryMemberResponse,
    } as QueryAllRegistryMemberResponse;
    message.RegistryMember = [];
    if (object.RegistryMember !== undefined && object.RegistryMember !== null) {
      for (const e of object.RegistryMember) {
        message.RegistryMember.push(RegistryMember.fromPartial(e));
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

const baseQueryGetRegistryStakedAmountRequest: object = { index: "" };

export const QueryGetRegistryStakedAmountRequest = {
  encode(
    message: QueryGetRegistryStakedAmountRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetRegistryStakedAmountRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetRegistryStakedAmountRequest,
    } as QueryGetRegistryStakedAmountRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetRegistryStakedAmountRequest {
    const message = {
      ...baseQueryGetRegistryStakedAmountRequest,
    } as QueryGetRegistryStakedAmountRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetRegistryStakedAmountRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetRegistryStakedAmountRequest>
  ): QueryGetRegistryStakedAmountRequest {
    const message = {
      ...baseQueryGetRegistryStakedAmountRequest,
    } as QueryGetRegistryStakedAmountRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetRegistryStakedAmountResponse: object = {};

export const QueryGetRegistryStakedAmountResponse = {
  encode(
    message: QueryGetRegistryStakedAmountResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.registryStakedAmount !== undefined) {
      RegistryStakedAmount.encode(
        message.registryStakedAmount,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetRegistryStakedAmountResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetRegistryStakedAmountResponse,
    } as QueryGetRegistryStakedAmountResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.registryStakedAmount = RegistryStakedAmount.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetRegistryStakedAmountResponse {
    const message = {
      ...baseQueryGetRegistryStakedAmountResponse,
    } as QueryGetRegistryStakedAmountResponse;
    if (
      object.registryStakedAmount !== undefined &&
      object.registryStakedAmount !== null
    ) {
      message.registryStakedAmount = RegistryStakedAmount.fromJSON(
        object.registryStakedAmount
      );
    } else {
      message.registryStakedAmount = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetRegistryStakedAmountResponse): unknown {
    const obj: any = {};
    message.registryStakedAmount !== undefined &&
      (obj.registryStakedAmount = message.registryStakedAmount
        ? RegistryStakedAmount.toJSON(message.registryStakedAmount)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetRegistryStakedAmountResponse>
  ): QueryGetRegistryStakedAmountResponse {
    const message = {
      ...baseQueryGetRegistryStakedAmountResponse,
    } as QueryGetRegistryStakedAmountResponse;
    if (
      object.registryStakedAmount !== undefined &&
      object.registryStakedAmount !== null
    ) {
      message.registryStakedAmount = RegistryStakedAmount.fromPartial(
        object.registryStakedAmount
      );
    } else {
      message.registryStakedAmount = undefined;
    }
    return message;
  },
};

const baseQueryAllRegistryStakedAmountRequest: object = {};

export const QueryAllRegistryStakedAmountRequest = {
  encode(
    message: QueryAllRegistryStakedAmountRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllRegistryStakedAmountRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllRegistryStakedAmountRequest,
    } as QueryAllRegistryStakedAmountRequest;
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

  fromJSON(object: any): QueryAllRegistryStakedAmountRequest {
    const message = {
      ...baseQueryAllRegistryStakedAmountRequest,
    } as QueryAllRegistryStakedAmountRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllRegistryStakedAmountRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllRegistryStakedAmountRequest>
  ): QueryAllRegistryStakedAmountRequest {
    const message = {
      ...baseQueryAllRegistryStakedAmountRequest,
    } as QueryAllRegistryStakedAmountRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllRegistryStakedAmountResponse: object = {};

export const QueryAllRegistryStakedAmountResponse = {
  encode(
    message: QueryAllRegistryStakedAmountResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.registryStakedAmount) {
      RegistryStakedAmount.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryAllRegistryStakedAmountResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllRegistryStakedAmountResponse,
    } as QueryAllRegistryStakedAmountResponse;
    message.registryStakedAmount = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.registryStakedAmount.push(
            RegistryStakedAmount.decode(reader, reader.uint32())
          );
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

  fromJSON(object: any): QueryAllRegistryStakedAmountResponse {
    const message = {
      ...baseQueryAllRegistryStakedAmountResponse,
    } as QueryAllRegistryStakedAmountResponse;
    message.registryStakedAmount = [];
    if (
      object.registryStakedAmount !== undefined &&
      object.registryStakedAmount !== null
    ) {
      for (const e of object.registryStakedAmount) {
        message.registryStakedAmount.push(RegistryStakedAmount.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllRegistryStakedAmountResponse): unknown {
    const obj: any = {};
    if (message.registryStakedAmount) {
      obj.registryStakedAmount = message.registryStakedAmount.map((e) =>
        e ? RegistryStakedAmount.toJSON(e) : undefined
      );
    } else {
      obj.registryStakedAmount = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllRegistryStakedAmountResponse>
  ): QueryAllRegistryStakedAmountResponse {
    const message = {
      ...baseQueryAllRegistryStakedAmountResponse,
    } as QueryAllRegistryStakedAmountResponse;
    message.registryStakedAmount = [];
    if (
      object.registryStakedAmount !== undefined &&
      object.registryStakedAmount !== null
    ) {
      for (const e of object.registryStakedAmount) {
        message.registryStakedAmount.push(RegistryStakedAmount.fromPartial(e));
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

const baseQueryGetRegistryStakedAmountPerWalletRequest: object = { index: "" };

export const QueryGetRegistryStakedAmountPerWalletRequest = {
  encode(
    message: QueryGetRegistryStakedAmountPerWalletRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetRegistryStakedAmountPerWalletRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetRegistryStakedAmountPerWalletRequest,
    } as QueryGetRegistryStakedAmountPerWalletRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetRegistryStakedAmountPerWalletRequest {
    const message = {
      ...baseQueryGetRegistryStakedAmountPerWalletRequest,
    } as QueryGetRegistryStakedAmountPerWalletRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetRegistryStakedAmountPerWalletRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetRegistryStakedAmountPerWalletRequest>
  ): QueryGetRegistryStakedAmountPerWalletRequest {
    const message = {
      ...baseQueryGetRegistryStakedAmountPerWalletRequest,
    } as QueryGetRegistryStakedAmountPerWalletRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetRegistryStakedAmountPerWalletResponse: object = {};

export const QueryGetRegistryStakedAmountPerWalletResponse = {
  encode(
    message: QueryGetRegistryStakedAmountPerWalletResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.registryStakedAmountPerWallet !== undefined) {
      RegistryStakedAmountPerWallet.encode(
        message.registryStakedAmountPerWallet,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetRegistryStakedAmountPerWalletResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetRegistryStakedAmountPerWalletResponse,
    } as QueryGetRegistryStakedAmountPerWalletResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.registryStakedAmountPerWallet = RegistryStakedAmountPerWallet.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetRegistryStakedAmountPerWalletResponse {
    const message = {
      ...baseQueryGetRegistryStakedAmountPerWalletResponse,
    } as QueryGetRegistryStakedAmountPerWalletResponse;
    if (
      object.registryStakedAmountPerWallet !== undefined &&
      object.registryStakedAmountPerWallet !== null
    ) {
      message.registryStakedAmountPerWallet = RegistryStakedAmountPerWallet.fromJSON(
        object.registryStakedAmountPerWallet
      );
    } else {
      message.registryStakedAmountPerWallet = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetRegistryStakedAmountPerWalletResponse): unknown {
    const obj: any = {};
    message.registryStakedAmountPerWallet !== undefined &&
      (obj.registryStakedAmountPerWallet = message.registryStakedAmountPerWallet
        ? RegistryStakedAmountPerWallet.toJSON(
            message.registryStakedAmountPerWallet
          )
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetRegistryStakedAmountPerWalletResponse>
  ): QueryGetRegistryStakedAmountPerWalletResponse {
    const message = {
      ...baseQueryGetRegistryStakedAmountPerWalletResponse,
    } as QueryGetRegistryStakedAmountPerWalletResponse;
    if (
      object.registryStakedAmountPerWallet !== undefined &&
      object.registryStakedAmountPerWallet !== null
    ) {
      message.registryStakedAmountPerWallet = RegistryStakedAmountPerWallet.fromPartial(
        object.registryStakedAmountPerWallet
      );
    } else {
      message.registryStakedAmountPerWallet = undefined;
    }
    return message;
  },
};

const baseQueryAllRegistryStakedAmountPerWalletRequest: object = {};

export const QueryAllRegistryStakedAmountPerWalletRequest = {
  encode(
    message: QueryAllRegistryStakedAmountPerWalletRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllRegistryStakedAmountPerWalletRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllRegistryStakedAmountPerWalletRequest,
    } as QueryAllRegistryStakedAmountPerWalletRequest;
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

  fromJSON(object: any): QueryAllRegistryStakedAmountPerWalletRequest {
    const message = {
      ...baseQueryAllRegistryStakedAmountPerWalletRequest,
    } as QueryAllRegistryStakedAmountPerWalletRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllRegistryStakedAmountPerWalletRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllRegistryStakedAmountPerWalletRequest>
  ): QueryAllRegistryStakedAmountPerWalletRequest {
    const message = {
      ...baseQueryAllRegistryStakedAmountPerWalletRequest,
    } as QueryAllRegistryStakedAmountPerWalletRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllRegistryStakedAmountPerWalletResponse: object = {};

export const QueryAllRegistryStakedAmountPerWalletResponse = {
  encode(
    message: QueryAllRegistryStakedAmountPerWalletResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.registryStakedAmountPerWallet) {
      RegistryStakedAmountPerWallet.encode(
        v!,
        writer.uint32(10).fork()
      ).ldelim();
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
  ): QueryAllRegistryStakedAmountPerWalletResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllRegistryStakedAmountPerWalletResponse,
    } as QueryAllRegistryStakedAmountPerWalletResponse;
    message.registryStakedAmountPerWallet = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.registryStakedAmountPerWallet.push(
            RegistryStakedAmountPerWallet.decode(reader, reader.uint32())
          );
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

  fromJSON(object: any): QueryAllRegistryStakedAmountPerWalletResponse {
    const message = {
      ...baseQueryAllRegistryStakedAmountPerWalletResponse,
    } as QueryAllRegistryStakedAmountPerWalletResponse;
    message.registryStakedAmountPerWallet = [];
    if (
      object.registryStakedAmountPerWallet !== undefined &&
      object.registryStakedAmountPerWallet !== null
    ) {
      for (const e of object.registryStakedAmountPerWallet) {
        message.registryStakedAmountPerWallet.push(
          RegistryStakedAmountPerWallet.fromJSON(e)
        );
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllRegistryStakedAmountPerWalletResponse): unknown {
    const obj: any = {};
    if (message.registryStakedAmountPerWallet) {
      obj.registryStakedAmountPerWallet = message.registryStakedAmountPerWallet.map(
        (e) => (e ? RegistryStakedAmountPerWallet.toJSON(e) : undefined)
      );
    } else {
      obj.registryStakedAmountPerWallet = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllRegistryStakedAmountPerWalletResponse>
  ): QueryAllRegistryStakedAmountPerWalletResponse {
    const message = {
      ...baseQueryAllRegistryStakedAmountPerWalletResponse,
    } as QueryAllRegistryStakedAmountPerWalletResponse;
    message.registryStakedAmountPerWallet = [];
    if (
      object.registryStakedAmountPerWallet !== undefined &&
      object.registryStakedAmountPerWallet !== null
    ) {
      for (const e of object.registryStakedAmountPerWallet) {
        message.registryStakedAmountPerWallet.push(
          RegistryStakedAmountPerWallet.fromPartial(e)
        );
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

const baseQueryGetPriceDataRequest: object = { index: "" };

export const QueryGetPriceDataRequest = {
  encode(
    message: QueryGetPriceDataRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetPriceDataRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetPriceDataRequest,
    } as QueryGetPriceDataRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetPriceDataRequest {
    const message = {
      ...baseQueryGetPriceDataRequest,
    } as QueryGetPriceDataRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetPriceDataRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetPriceDataRequest>
  ): QueryGetPriceDataRequest {
    const message = {
      ...baseQueryGetPriceDataRequest,
    } as QueryGetPriceDataRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetPriceDataResponse: object = {};

export const QueryGetPriceDataResponse = {
  encode(
    message: QueryGetPriceDataResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.priceData !== undefined) {
      PriceData.encode(message.priceData, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetPriceDataResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetPriceDataResponse,
    } as QueryGetPriceDataResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.priceData = PriceData.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetPriceDataResponse {
    const message = {
      ...baseQueryGetPriceDataResponse,
    } as QueryGetPriceDataResponse;
    if (object.priceData !== undefined && object.priceData !== null) {
      message.priceData = PriceData.fromJSON(object.priceData);
    } else {
      message.priceData = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetPriceDataResponse): unknown {
    const obj: any = {};
    message.priceData !== undefined &&
      (obj.priceData = message.priceData
        ? PriceData.toJSON(message.priceData)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetPriceDataResponse>
  ): QueryGetPriceDataResponse {
    const message = {
      ...baseQueryGetPriceDataResponse,
    } as QueryGetPriceDataResponse;
    if (object.priceData !== undefined && object.priceData !== null) {
      message.priceData = PriceData.fromPartial(object.priceData);
    } else {
      message.priceData = undefined;
    }
    return message;
  },
};

const baseQueryAllPriceDataRequest: object = {};

export const QueryAllPriceDataRequest = {
  encode(
    message: QueryAllPriceDataRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllPriceDataRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllPriceDataRequest,
    } as QueryAllPriceDataRequest;
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

  fromJSON(object: any): QueryAllPriceDataRequest {
    const message = {
      ...baseQueryAllPriceDataRequest,
    } as QueryAllPriceDataRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllPriceDataRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllPriceDataRequest>
  ): QueryAllPriceDataRequest {
    const message = {
      ...baseQueryAllPriceDataRequest,
    } as QueryAllPriceDataRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllPriceDataResponse: object = {};

export const QueryAllPriceDataResponse = {
  encode(
    message: QueryAllPriceDataResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.priceData) {
      PriceData.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryAllPriceDataResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllPriceDataResponse,
    } as QueryAllPriceDataResponse;
    message.priceData = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.priceData.push(PriceData.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllPriceDataResponse {
    const message = {
      ...baseQueryAllPriceDataResponse,
    } as QueryAllPriceDataResponse;
    message.priceData = [];
    if (object.priceData !== undefined && object.priceData !== null) {
      for (const e of object.priceData) {
        message.priceData.push(PriceData.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllPriceDataResponse): unknown {
    const obj: any = {};
    if (message.priceData) {
      obj.priceData = message.priceData.map((e) =>
        e ? PriceData.toJSON(e) : undefined
      );
    } else {
      obj.priceData = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllPriceDataResponse>
  ): QueryAllPriceDataResponse {
    const message = {
      ...baseQueryAllPriceDataResponse,
    } as QueryAllPriceDataResponse;
    message.priceData = [];
    if (object.priceData !== undefined && object.priceData !== null) {
      for (const e of object.priceData) {
        message.priceData.push(PriceData.fromPartial(e));
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
  /** Queries a RegistryOwner by id. */
  RegistryOwner(
    request: QueryGetRegistryOwnerRequest
  ): Promise<QueryGetRegistryOwnerResponse>;
  /** Queries a list of RegistryOwner items. */
  RegistryOwnerAll(
    request: QueryAllRegistryOwnerRequest
  ): Promise<QueryAllRegistryOwnerResponse>;
  /** Queries a RegistryMember by id. */
  RegistryMember(
    request: QueryGetRegistryMemberRequest
  ): Promise<QueryGetRegistryMemberResponse>;
  /** Queries a list of RegistryMember items. */
  RegistryMemberAll(
    request: QueryAllRegistryMemberRequest
  ): Promise<QueryAllRegistryMemberResponse>;
  /** Queries a RegistryStakedAmount by index. */
  RegistryStakedAmount(
    request: QueryGetRegistryStakedAmountRequest
  ): Promise<QueryGetRegistryStakedAmountResponse>;
  /** Queries a list of RegistryStakedAmount items. */
  RegistryStakedAmountAll(
    request: QueryAllRegistryStakedAmountRequest
  ): Promise<QueryAllRegistryStakedAmountResponse>;
  /** Queries a StakedAmountPerWallet by index. */
  RegistryStakedAmountPerWallet(
    request: QueryGetRegistryStakedAmountPerWalletRequest
  ): Promise<QueryGetRegistryStakedAmountPerWalletResponse>;
  /** Queries a list of StakedAmountPerWallet items. */
  RegistryStakedAmountPerWalletAll(
    request: QueryAllRegistryStakedAmountPerWalletRequest
  ): Promise<QueryAllRegistryStakedAmountPerWalletResponse>;
  /** Queries a PriceData by index. */
  PriceData(
    request: QueryGetPriceDataRequest
  ): Promise<QueryGetPriceDataResponse>;
  /** Queries a list of PriceData items. */
  PriceDataAll(
    request: QueryAllPriceDataRequest
  ): Promise<QueryAllPriceDataResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "pricechain.cprc.registry.Query",
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
      "pricechain.cprc.registry.Query",
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
      "pricechain.cprc.registry.Query",
      "RegistryAll",
      data
    );
    return promise.then((data) =>
      QueryAllRegistryResponse.decode(new Reader(data))
    );
  }

  RegistryOwner(
    request: QueryGetRegistryOwnerRequest
  ): Promise<QueryGetRegistryOwnerResponse> {
    const data = QueryGetRegistryOwnerRequest.encode(request).finish();
    const promise = this.rpc.request(
      "pricechain.cprc.registry.Query",
      "RegistryOwner",
      data
    );
    return promise.then((data) =>
      QueryGetRegistryOwnerResponse.decode(new Reader(data))
    );
  }

  RegistryOwnerAll(
    request: QueryAllRegistryOwnerRequest
  ): Promise<QueryAllRegistryOwnerResponse> {
    const data = QueryAllRegistryOwnerRequest.encode(request).finish();
    const promise = this.rpc.request(
      "pricechain.cprc.registry.Query",
      "RegistryOwnerAll",
      data
    );
    return promise.then((data) =>
      QueryAllRegistryOwnerResponse.decode(new Reader(data))
    );
  }

  RegistryMember(
    request: QueryGetRegistryMemberRequest
  ): Promise<QueryGetRegistryMemberResponse> {
    const data = QueryGetRegistryMemberRequest.encode(request).finish();
    const promise = this.rpc.request(
      "pricechain.cprc.registry.Query",
      "RegistryMember",
      data
    );
    return promise.then((data) =>
      QueryGetRegistryMemberResponse.decode(new Reader(data))
    );
  }

  RegistryMemberAll(
    request: QueryAllRegistryMemberRequest
  ): Promise<QueryAllRegistryMemberResponse> {
    const data = QueryAllRegistryMemberRequest.encode(request).finish();
    const promise = this.rpc.request(
      "pricechain.cprc.registry.Query",
      "RegistryMemberAll",
      data
    );
    return promise.then((data) =>
      QueryAllRegistryMemberResponse.decode(new Reader(data))
    );
  }

  RegistryStakedAmount(
    request: QueryGetRegistryStakedAmountRequest
  ): Promise<QueryGetRegistryStakedAmountResponse> {
    const data = QueryGetRegistryStakedAmountRequest.encode(request).finish();
    const promise = this.rpc.request(
      "pricechain.cprc.registry.Query",
      "RegistryStakedAmount",
      data
    );
    return promise.then((data) =>
      QueryGetRegistryStakedAmountResponse.decode(new Reader(data))
    );
  }

  RegistryStakedAmountAll(
    request: QueryAllRegistryStakedAmountRequest
  ): Promise<QueryAllRegistryStakedAmountResponse> {
    const data = QueryAllRegistryStakedAmountRequest.encode(request).finish();
    const promise = this.rpc.request(
      "pricechain.cprc.registry.Query",
      "RegistryStakedAmountAll",
      data
    );
    return promise.then((data) =>
      QueryAllRegistryStakedAmountResponse.decode(new Reader(data))
    );
  }

  RegistryStakedAmountPerWallet(
    request: QueryGetRegistryStakedAmountPerWalletRequest
  ): Promise<QueryGetRegistryStakedAmountPerWalletResponse> {
    const data = QueryGetRegistryStakedAmountPerWalletRequest.encode(
      request
    ).finish();
    const promise = this.rpc.request(
      "pricechain.cprc.registry.Query",
      "RegistryStakedAmountPerWallet",
      data
    );
    return promise.then((data) =>
      QueryGetRegistryStakedAmountPerWalletResponse.decode(new Reader(data))
    );
  }

  RegistryStakedAmountPerWalletAll(
    request: QueryAllRegistryStakedAmountPerWalletRequest
  ): Promise<QueryAllRegistryStakedAmountPerWalletResponse> {
    const data = QueryAllRegistryStakedAmountPerWalletRequest.encode(
      request
    ).finish();
    const promise = this.rpc.request(
      "pricechain.cprc.registry.Query",
      "RegistryStakedAmountPerWalletAll",
      data
    );
    return promise.then((data) =>
      QueryAllRegistryStakedAmountPerWalletResponse.decode(new Reader(data))
    );
  }

  PriceData(
    request: QueryGetPriceDataRequest
  ): Promise<QueryGetPriceDataResponse> {
    const data = QueryGetPriceDataRequest.encode(request).finish();
    const promise = this.rpc.request(
      "pricechain.cprc.registry.Query",
      "PriceData",
      data
    );
    return promise.then((data) =>
      QueryGetPriceDataResponse.decode(new Reader(data))
    );
  }

  PriceDataAll(
    request: QueryAllPriceDataRequest
  ): Promise<QueryAllPriceDataResponse> {
    const data = QueryAllPriceDataRequest.encode(request).finish();
    const promise = this.rpc.request(
      "pricechain.cprc.registry.Query",
      "PriceDataAll",
      data
    );
    return promise.then((data) =>
      QueryAllPriceDataResponse.decode(new Reader(data))
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
