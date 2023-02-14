/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "pricechain.cprc.registry";

export interface MsgCreateRegistry {
  creator: string;
  name: string;
  stakeAmount: string;
  description: string;
  imageUrl: string;
}

export interface MsgCreateRegistryResponse {}

export interface MsgJoinRegistryCoOperator {
  creator: string;
  registryId: string;
  stakeAmount: string;
}

export interface MsgJoinRegistryCoOperatorResponse {}

export interface MsgJoinRegistryMember {
  creator: string;
  registryId: string;
  stakeAmount: string;
}

export interface MsgJoinRegistryMemberResponse {}

export interface MsgUnbondRegistry {
  creator: string;
  registryId: string;
  unbondAmount: string;
  reason: string;
}

export interface MsgUnbondRegistryResponse {}

export interface MsgModifyRegistry {
  creator: string;
  registryId: string;
  stakeAmount: string;
  name: string;
  description: string;
  imageUrl: string;
  reason: string;
}

export interface MsgModifyRegistryResponse {}

export interface MsgProposePrice {
  creator: string;
  registryId: string;
  storeName: string;
  storeAddr: string;
  purchaseTime: string;
  prodName: string;
  price: string;
  receiptCode: string;
  reserved: string;
}

export interface MsgProposePriceResponse {}

const baseMsgCreateRegistry: object = {
  creator: "",
  name: "",
  stakeAmount: "",
  description: "",
  imageUrl: "",
};

export const MsgCreateRegistry = {
  encode(message: MsgCreateRegistry, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.stakeAmount !== "") {
      writer.uint32(26).string(message.stakeAmount);
    }
    if (message.description !== "") {
      writer.uint32(34).string(message.description);
    }
    if (message.imageUrl !== "") {
      writer.uint32(42).string(message.imageUrl);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateRegistry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateRegistry } as MsgCreateRegistry;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.stakeAmount = reader.string();
          break;
        case 4:
          message.description = reader.string();
          break;
        case 5:
          message.imageUrl = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateRegistry {
    const message = { ...baseMsgCreateRegistry } as MsgCreateRegistry;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.stakeAmount !== undefined && object.stakeAmount !== null) {
      message.stakeAmount = String(object.stakeAmount);
    } else {
      message.stakeAmount = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.imageUrl !== undefined && object.imageUrl !== null) {
      message.imageUrl = String(object.imageUrl);
    } else {
      message.imageUrl = "";
    }
    return message;
  },

  toJSON(message: MsgCreateRegistry): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.stakeAmount !== undefined &&
      (obj.stakeAmount = message.stakeAmount);
    message.description !== undefined &&
      (obj.description = message.description);
    message.imageUrl !== undefined && (obj.imageUrl = message.imageUrl);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateRegistry>): MsgCreateRegistry {
    const message = { ...baseMsgCreateRegistry } as MsgCreateRegistry;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.stakeAmount !== undefined && object.stakeAmount !== null) {
      message.stakeAmount = object.stakeAmount;
    } else {
      message.stakeAmount = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.imageUrl !== undefined && object.imageUrl !== null) {
      message.imageUrl = object.imageUrl;
    } else {
      message.imageUrl = "";
    }
    return message;
  },
};

const baseMsgCreateRegistryResponse: object = {};

export const MsgCreateRegistryResponse = {
  encode(
    _: MsgCreateRegistryResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateRegistryResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateRegistryResponse,
    } as MsgCreateRegistryResponse;
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

  fromJSON(_: any): MsgCreateRegistryResponse {
    const message = {
      ...baseMsgCreateRegistryResponse,
    } as MsgCreateRegistryResponse;
    return message;
  },

  toJSON(_: MsgCreateRegistryResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateRegistryResponse>
  ): MsgCreateRegistryResponse {
    const message = {
      ...baseMsgCreateRegistryResponse,
    } as MsgCreateRegistryResponse;
    return message;
  },
};

const baseMsgJoinRegistryCoOperator: object = {
  creator: "",
  registryId: "",
  stakeAmount: "",
};

export const MsgJoinRegistryCoOperator = {
  encode(
    message: MsgJoinRegistryCoOperator,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.registryId !== "") {
      writer.uint32(18).string(message.registryId);
    }
    if (message.stakeAmount !== "") {
      writer.uint32(26).string(message.stakeAmount);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgJoinRegistryCoOperator {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgJoinRegistryCoOperator,
    } as MsgJoinRegistryCoOperator;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.registryId = reader.string();
          break;
        case 3:
          message.stakeAmount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgJoinRegistryCoOperator {
    const message = {
      ...baseMsgJoinRegistryCoOperator,
    } as MsgJoinRegistryCoOperator;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.registryId !== undefined && object.registryId !== null) {
      message.registryId = String(object.registryId);
    } else {
      message.registryId = "";
    }
    if (object.stakeAmount !== undefined && object.stakeAmount !== null) {
      message.stakeAmount = String(object.stakeAmount);
    } else {
      message.stakeAmount = "";
    }
    return message;
  },

  toJSON(message: MsgJoinRegistryCoOperator): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.registryId !== undefined && (obj.registryId = message.registryId);
    message.stakeAmount !== undefined &&
      (obj.stakeAmount = message.stakeAmount);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgJoinRegistryCoOperator>
  ): MsgJoinRegistryCoOperator {
    const message = {
      ...baseMsgJoinRegistryCoOperator,
    } as MsgJoinRegistryCoOperator;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.registryId !== undefined && object.registryId !== null) {
      message.registryId = object.registryId;
    } else {
      message.registryId = "";
    }
    if (object.stakeAmount !== undefined && object.stakeAmount !== null) {
      message.stakeAmount = object.stakeAmount;
    } else {
      message.stakeAmount = "";
    }
    return message;
  },
};

const baseMsgJoinRegistryCoOperatorResponse: object = {};

export const MsgJoinRegistryCoOperatorResponse = {
  encode(
    _: MsgJoinRegistryCoOperatorResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgJoinRegistryCoOperatorResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgJoinRegistryCoOperatorResponse,
    } as MsgJoinRegistryCoOperatorResponse;
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

  fromJSON(_: any): MsgJoinRegistryCoOperatorResponse {
    const message = {
      ...baseMsgJoinRegistryCoOperatorResponse,
    } as MsgJoinRegistryCoOperatorResponse;
    return message;
  },

  toJSON(_: MsgJoinRegistryCoOperatorResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgJoinRegistryCoOperatorResponse>
  ): MsgJoinRegistryCoOperatorResponse {
    const message = {
      ...baseMsgJoinRegistryCoOperatorResponse,
    } as MsgJoinRegistryCoOperatorResponse;
    return message;
  },
};

const baseMsgJoinRegistryMember: object = {
  creator: "",
  registryId: "",
  stakeAmount: "",
};

export const MsgJoinRegistryMember = {
  encode(
    message: MsgJoinRegistryMember,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.registryId !== "") {
      writer.uint32(18).string(message.registryId);
    }
    if (message.stakeAmount !== "") {
      writer.uint32(26).string(message.stakeAmount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgJoinRegistryMember {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgJoinRegistryMember } as MsgJoinRegistryMember;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.registryId = reader.string();
          break;
        case 3:
          message.stakeAmount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgJoinRegistryMember {
    const message = { ...baseMsgJoinRegistryMember } as MsgJoinRegistryMember;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.registryId !== undefined && object.registryId !== null) {
      message.registryId = String(object.registryId);
    } else {
      message.registryId = "";
    }
    if (object.stakeAmount !== undefined && object.stakeAmount !== null) {
      message.stakeAmount = String(object.stakeAmount);
    } else {
      message.stakeAmount = "";
    }
    return message;
  },

  toJSON(message: MsgJoinRegistryMember): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.registryId !== undefined && (obj.registryId = message.registryId);
    message.stakeAmount !== undefined &&
      (obj.stakeAmount = message.stakeAmount);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgJoinRegistryMember>
  ): MsgJoinRegistryMember {
    const message = { ...baseMsgJoinRegistryMember } as MsgJoinRegistryMember;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.registryId !== undefined && object.registryId !== null) {
      message.registryId = object.registryId;
    } else {
      message.registryId = "";
    }
    if (object.stakeAmount !== undefined && object.stakeAmount !== null) {
      message.stakeAmount = object.stakeAmount;
    } else {
      message.stakeAmount = "";
    }
    return message;
  },
};

const baseMsgJoinRegistryMemberResponse: object = {};

export const MsgJoinRegistryMemberResponse = {
  encode(
    _: MsgJoinRegistryMemberResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgJoinRegistryMemberResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgJoinRegistryMemberResponse,
    } as MsgJoinRegistryMemberResponse;
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

  fromJSON(_: any): MsgJoinRegistryMemberResponse {
    const message = {
      ...baseMsgJoinRegistryMemberResponse,
    } as MsgJoinRegistryMemberResponse;
    return message;
  },

  toJSON(_: MsgJoinRegistryMemberResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgJoinRegistryMemberResponse>
  ): MsgJoinRegistryMemberResponse {
    const message = {
      ...baseMsgJoinRegistryMemberResponse,
    } as MsgJoinRegistryMemberResponse;
    return message;
  },
};

const baseMsgUnbondRegistry: object = {
  creator: "",
  registryId: "",
  unbondAmount: "",
  reason: "",
};

export const MsgUnbondRegistry = {
  encode(message: MsgUnbondRegistry, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.registryId !== "") {
      writer.uint32(18).string(message.registryId);
    }
    if (message.unbondAmount !== "") {
      writer.uint32(26).string(message.unbondAmount);
    }
    if (message.reason !== "") {
      writer.uint32(34).string(message.reason);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUnbondRegistry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUnbondRegistry } as MsgUnbondRegistry;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.registryId = reader.string();
          break;
        case 3:
          message.unbondAmount = reader.string();
          break;
        case 4:
          message.reason = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUnbondRegistry {
    const message = { ...baseMsgUnbondRegistry } as MsgUnbondRegistry;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.registryId !== undefined && object.registryId !== null) {
      message.registryId = String(object.registryId);
    } else {
      message.registryId = "";
    }
    if (object.unbondAmount !== undefined && object.unbondAmount !== null) {
      message.unbondAmount = String(object.unbondAmount);
    } else {
      message.unbondAmount = "";
    }
    if (object.reason !== undefined && object.reason !== null) {
      message.reason = String(object.reason);
    } else {
      message.reason = "";
    }
    return message;
  },

  toJSON(message: MsgUnbondRegistry): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.registryId !== undefined && (obj.registryId = message.registryId);
    message.unbondAmount !== undefined &&
      (obj.unbondAmount = message.unbondAmount);
    message.reason !== undefined && (obj.reason = message.reason);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUnbondRegistry>): MsgUnbondRegistry {
    const message = { ...baseMsgUnbondRegistry } as MsgUnbondRegistry;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.registryId !== undefined && object.registryId !== null) {
      message.registryId = object.registryId;
    } else {
      message.registryId = "";
    }
    if (object.unbondAmount !== undefined && object.unbondAmount !== null) {
      message.unbondAmount = object.unbondAmount;
    } else {
      message.unbondAmount = "";
    }
    if (object.reason !== undefined && object.reason !== null) {
      message.reason = object.reason;
    } else {
      message.reason = "";
    }
    return message;
  },
};

const baseMsgUnbondRegistryResponse: object = {};

export const MsgUnbondRegistryResponse = {
  encode(
    _: MsgUnbondRegistryResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUnbondRegistryResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUnbondRegistryResponse,
    } as MsgUnbondRegistryResponse;
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

  fromJSON(_: any): MsgUnbondRegistryResponse {
    const message = {
      ...baseMsgUnbondRegistryResponse,
    } as MsgUnbondRegistryResponse;
    return message;
  },

  toJSON(_: MsgUnbondRegistryResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUnbondRegistryResponse>
  ): MsgUnbondRegistryResponse {
    const message = {
      ...baseMsgUnbondRegistryResponse,
    } as MsgUnbondRegistryResponse;
    return message;
  },
};

const baseMsgModifyRegistry: object = {
  creator: "",
  registryId: "",
  stakeAmount: "",
  name: "",
  description: "",
  imageUrl: "",
  reason: "",
};

export const MsgModifyRegistry = {
  encode(message: MsgModifyRegistry, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.registryId !== "") {
      writer.uint32(18).string(message.registryId);
    }
    if (message.stakeAmount !== "") {
      writer.uint32(26).string(message.stakeAmount);
    }
    if (message.name !== "") {
      writer.uint32(34).string(message.name);
    }
    if (message.description !== "") {
      writer.uint32(42).string(message.description);
    }
    if (message.imageUrl !== "") {
      writer.uint32(50).string(message.imageUrl);
    }
    if (message.reason !== "") {
      writer.uint32(58).string(message.reason);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgModifyRegistry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgModifyRegistry } as MsgModifyRegistry;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.registryId = reader.string();
          break;
        case 3:
          message.stakeAmount = reader.string();
          break;
        case 4:
          message.name = reader.string();
          break;
        case 5:
          message.description = reader.string();
          break;
        case 6:
          message.imageUrl = reader.string();
          break;
        case 7:
          message.reason = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgModifyRegistry {
    const message = { ...baseMsgModifyRegistry } as MsgModifyRegistry;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.registryId !== undefined && object.registryId !== null) {
      message.registryId = String(object.registryId);
    } else {
      message.registryId = "";
    }
    if (object.stakeAmount !== undefined && object.stakeAmount !== null) {
      message.stakeAmount = String(object.stakeAmount);
    } else {
      message.stakeAmount = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.imageUrl !== undefined && object.imageUrl !== null) {
      message.imageUrl = String(object.imageUrl);
    } else {
      message.imageUrl = "";
    }
    if (object.reason !== undefined && object.reason !== null) {
      message.reason = String(object.reason);
    } else {
      message.reason = "";
    }
    return message;
  },

  toJSON(message: MsgModifyRegistry): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.registryId !== undefined && (obj.registryId = message.registryId);
    message.stakeAmount !== undefined &&
      (obj.stakeAmount = message.stakeAmount);
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined &&
      (obj.description = message.description);
    message.imageUrl !== undefined && (obj.imageUrl = message.imageUrl);
    message.reason !== undefined && (obj.reason = message.reason);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgModifyRegistry>): MsgModifyRegistry {
    const message = { ...baseMsgModifyRegistry } as MsgModifyRegistry;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.registryId !== undefined && object.registryId !== null) {
      message.registryId = object.registryId;
    } else {
      message.registryId = "";
    }
    if (object.stakeAmount !== undefined && object.stakeAmount !== null) {
      message.stakeAmount = object.stakeAmount;
    } else {
      message.stakeAmount = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.imageUrl !== undefined && object.imageUrl !== null) {
      message.imageUrl = object.imageUrl;
    } else {
      message.imageUrl = "";
    }
    if (object.reason !== undefined && object.reason !== null) {
      message.reason = object.reason;
    } else {
      message.reason = "";
    }
    return message;
  },
};

const baseMsgModifyRegistryResponse: object = {};

export const MsgModifyRegistryResponse = {
  encode(
    _: MsgModifyRegistryResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgModifyRegistryResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgModifyRegistryResponse,
    } as MsgModifyRegistryResponse;
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

  fromJSON(_: any): MsgModifyRegistryResponse {
    const message = {
      ...baseMsgModifyRegistryResponse,
    } as MsgModifyRegistryResponse;
    return message;
  },

  toJSON(_: MsgModifyRegistryResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgModifyRegistryResponse>
  ): MsgModifyRegistryResponse {
    const message = {
      ...baseMsgModifyRegistryResponse,
    } as MsgModifyRegistryResponse;
    return message;
  },
};

const baseMsgProposePrice: object = {
  creator: "",
  registryId: "",
  storeName: "",
  storeAddr: "",
  purchaseTime: "",
  prodName: "",
  price: "",
  receiptCode: "",
  reserved: "",
};

export const MsgProposePrice = {
  encode(message: MsgProposePrice, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.registryId !== "") {
      writer.uint32(18).string(message.registryId);
    }
    if (message.storeName !== "") {
      writer.uint32(26).string(message.storeName);
    }
    if (message.storeAddr !== "") {
      writer.uint32(34).string(message.storeAddr);
    }
    if (message.purchaseTime !== "") {
      writer.uint32(42).string(message.purchaseTime);
    }
    if (message.prodName !== "") {
      writer.uint32(50).string(message.prodName);
    }
    if (message.price !== "") {
      writer.uint32(58).string(message.price);
    }
    if (message.receiptCode !== "") {
      writer.uint32(66).string(message.receiptCode);
    }
    if (message.reserved !== "") {
      writer.uint32(74).string(message.reserved);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgProposePrice {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgProposePrice } as MsgProposePrice;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.registryId = reader.string();
          break;
        case 3:
          message.storeName = reader.string();
          break;
        case 4:
          message.storeAddr = reader.string();
          break;
        case 5:
          message.purchaseTime = reader.string();
          break;
        case 6:
          message.prodName = reader.string();
          break;
        case 7:
          message.price = reader.string();
          break;
        case 8:
          message.receiptCode = reader.string();
          break;
        case 9:
          message.reserved = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgProposePrice {
    const message = { ...baseMsgProposePrice } as MsgProposePrice;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.registryId !== undefined && object.registryId !== null) {
      message.registryId = String(object.registryId);
    } else {
      message.registryId = "";
    }
    if (object.storeName !== undefined && object.storeName !== null) {
      message.storeName = String(object.storeName);
    } else {
      message.storeName = "";
    }
    if (object.storeAddr !== undefined && object.storeAddr !== null) {
      message.storeAddr = String(object.storeAddr);
    } else {
      message.storeAddr = "";
    }
    if (object.purchaseTime !== undefined && object.purchaseTime !== null) {
      message.purchaseTime = String(object.purchaseTime);
    } else {
      message.purchaseTime = "";
    }
    if (object.prodName !== undefined && object.prodName !== null) {
      message.prodName = String(object.prodName);
    } else {
      message.prodName = "";
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = String(object.price);
    } else {
      message.price = "";
    }
    if (object.receiptCode !== undefined && object.receiptCode !== null) {
      message.receiptCode = String(object.receiptCode);
    } else {
      message.receiptCode = "";
    }
    if (object.reserved !== undefined && object.reserved !== null) {
      message.reserved = String(object.reserved);
    } else {
      message.reserved = "";
    }
    return message;
  },

  toJSON(message: MsgProposePrice): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.registryId !== undefined && (obj.registryId = message.registryId);
    message.storeName !== undefined && (obj.storeName = message.storeName);
    message.storeAddr !== undefined && (obj.storeAddr = message.storeAddr);
    message.purchaseTime !== undefined &&
      (obj.purchaseTime = message.purchaseTime);
    message.prodName !== undefined && (obj.prodName = message.prodName);
    message.price !== undefined && (obj.price = message.price);
    message.receiptCode !== undefined &&
      (obj.receiptCode = message.receiptCode);
    message.reserved !== undefined && (obj.reserved = message.reserved);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgProposePrice>): MsgProposePrice {
    const message = { ...baseMsgProposePrice } as MsgProposePrice;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.registryId !== undefined && object.registryId !== null) {
      message.registryId = object.registryId;
    } else {
      message.registryId = "";
    }
    if (object.storeName !== undefined && object.storeName !== null) {
      message.storeName = object.storeName;
    } else {
      message.storeName = "";
    }
    if (object.storeAddr !== undefined && object.storeAddr !== null) {
      message.storeAddr = object.storeAddr;
    } else {
      message.storeAddr = "";
    }
    if (object.purchaseTime !== undefined && object.purchaseTime !== null) {
      message.purchaseTime = object.purchaseTime;
    } else {
      message.purchaseTime = "";
    }
    if (object.prodName !== undefined && object.prodName !== null) {
      message.prodName = object.prodName;
    } else {
      message.prodName = "";
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = object.price;
    } else {
      message.price = "";
    }
    if (object.receiptCode !== undefined && object.receiptCode !== null) {
      message.receiptCode = object.receiptCode;
    } else {
      message.receiptCode = "";
    }
    if (object.reserved !== undefined && object.reserved !== null) {
      message.reserved = object.reserved;
    } else {
      message.reserved = "";
    }
    return message;
  },
};

const baseMsgProposePriceResponse: object = {};

export const MsgProposePriceResponse = {
  encode(_: MsgProposePriceResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgProposePriceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgProposePriceResponse,
    } as MsgProposePriceResponse;
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

  fromJSON(_: any): MsgProposePriceResponse {
    const message = {
      ...baseMsgProposePriceResponse,
    } as MsgProposePriceResponse;
    return message;
  },

  toJSON(_: MsgProposePriceResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgProposePriceResponse>
  ): MsgProposePriceResponse {
    const message = {
      ...baseMsgProposePriceResponse,
    } as MsgProposePriceResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateRegistry(
    request: MsgCreateRegistry
  ): Promise<MsgCreateRegistryResponse>;
  JoinRegistryCoOperator(
    request: MsgJoinRegistryCoOperator
  ): Promise<MsgJoinRegistryCoOperatorResponse>;
  JoinRegistryMember(
    request: MsgJoinRegistryMember
  ): Promise<MsgJoinRegistryMemberResponse>;
  UnbondRegistry(
    request: MsgUnbondRegistry
  ): Promise<MsgUnbondRegistryResponse>;
  ModifyRegistry(
    request: MsgModifyRegistry
  ): Promise<MsgModifyRegistryResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  ProposePrice(request: MsgProposePrice): Promise<MsgProposePriceResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateRegistry(
    request: MsgCreateRegistry
  ): Promise<MsgCreateRegistryResponse> {
    const data = MsgCreateRegistry.encode(request).finish();
    const promise = this.rpc.request(
      "pricechain.cprc.registry.Msg",
      "CreateRegistry",
      data
    );
    return promise.then((data) =>
      MsgCreateRegistryResponse.decode(new Reader(data))
    );
  }

  JoinRegistryCoOperator(
    request: MsgJoinRegistryCoOperator
  ): Promise<MsgJoinRegistryCoOperatorResponse> {
    const data = MsgJoinRegistryCoOperator.encode(request).finish();
    const promise = this.rpc.request(
      "pricechain.cprc.registry.Msg",
      "JoinRegistryCoOperator",
      data
    );
    return promise.then((data) =>
      MsgJoinRegistryCoOperatorResponse.decode(new Reader(data))
    );
  }

  JoinRegistryMember(
    request: MsgJoinRegistryMember
  ): Promise<MsgJoinRegistryMemberResponse> {
    const data = MsgJoinRegistryMember.encode(request).finish();
    const promise = this.rpc.request(
      "pricechain.cprc.registry.Msg",
      "JoinRegistryMember",
      data
    );
    return promise.then((data) =>
      MsgJoinRegistryMemberResponse.decode(new Reader(data))
    );
  }

  UnbondRegistry(
    request: MsgUnbondRegistry
  ): Promise<MsgUnbondRegistryResponse> {
    const data = MsgUnbondRegistry.encode(request).finish();
    const promise = this.rpc.request(
      "pricechain.cprc.registry.Msg",
      "UnbondRegistry",
      data
    );
    return promise.then((data) =>
      MsgUnbondRegistryResponse.decode(new Reader(data))
    );
  }

  ModifyRegistry(
    request: MsgModifyRegistry
  ): Promise<MsgModifyRegistryResponse> {
    const data = MsgModifyRegistry.encode(request).finish();
    const promise = this.rpc.request(
      "pricechain.cprc.registry.Msg",
      "ModifyRegistry",
      data
    );
    return promise.then((data) =>
      MsgModifyRegistryResponse.decode(new Reader(data))
    );
  }

  ProposePrice(request: MsgProposePrice): Promise<MsgProposePriceResponse> {
    const data = MsgProposePrice.encode(request).finish();
    const promise = this.rpc.request(
      "pricechain.cprc.registry.Msg",
      "ProposePrice",
      data
    );
    return promise.then((data) =>
      MsgProposePriceResponse.decode(new Reader(data))
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
