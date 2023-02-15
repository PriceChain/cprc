/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "pricechain.cprc.registry";

export interface Registry {
  id: number;
  name: string;
  stakedAmount: string;
  status: string;
  description: string;
  imageUrl: string;
  priceCount: string;
  reviewCount: string;
  timestamp: string;
  reserved: string;
  owners: string[];
}

const baseRegistry: object = {
  id: 0,
  name: "",
  stakedAmount: "",
  status: "",
  description: "",
  imageUrl: "",
  priceCount: "",
  reviewCount: "",
  timestamp: "",
  reserved: "",
  owners: "",
};

export const Registry = {
  encode(message: Registry, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.stakedAmount !== "") {
      writer.uint32(26).string(message.stakedAmount);
    }
    if (message.status !== "") {
      writer.uint32(34).string(message.status);
    }
    if (message.description !== "") {
      writer.uint32(42).string(message.description);
    }
    if (message.imageUrl !== "") {
      writer.uint32(50).string(message.imageUrl);
    }
    if (message.priceCount !== "") {
      writer.uint32(58).string(message.priceCount);
    }
    if (message.reviewCount !== "") {
      writer.uint32(66).string(message.reviewCount);
    }
    if (message.timestamp !== "") {
      writer.uint32(74).string(message.timestamp);
    }
    if (message.reserved !== "") {
      writer.uint32(82).string(message.reserved);
    }
    for (const v of message.owners) {
      writer.uint32(90).string(v!);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Registry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseRegistry } as Registry;
    message.owners = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.stakedAmount = reader.string();
          break;
        case 4:
          message.status = reader.string();
          break;
        case 5:
          message.description = reader.string();
          break;
        case 6:
          message.imageUrl = reader.string();
          break;
        case 7:
          message.priceCount = reader.string();
          break;
        case 8:
          message.reviewCount = reader.string();
          break;
        case 9:
          message.timestamp = reader.string();
          break;
        case 10:
          message.reserved = reader.string();
          break;
        case 11:
          message.owners.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Registry {
    const message = { ...baseRegistry } as Registry;
    message.owners = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.stakedAmount !== undefined && object.stakedAmount !== null) {
      message.stakedAmount = String(object.stakedAmount);
    } else {
      message.stakedAmount = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = String(object.status);
    } else {
      message.status = "";
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
    if (object.priceCount !== undefined && object.priceCount !== null) {
      message.priceCount = String(object.priceCount);
    } else {
      message.priceCount = "";
    }
    if (object.reviewCount !== undefined && object.reviewCount !== null) {
      message.reviewCount = String(object.reviewCount);
    } else {
      message.reviewCount = "";
    }
    if (object.timestamp !== undefined && object.timestamp !== null) {
      message.timestamp = String(object.timestamp);
    } else {
      message.timestamp = "";
    }
    if (object.reserved !== undefined && object.reserved !== null) {
      message.reserved = String(object.reserved);
    } else {
      message.reserved = "";
    }
    if (object.owners !== undefined && object.owners !== null) {
      for (const e of object.owners) {
        message.owners.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: Registry): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.name !== undefined && (obj.name = message.name);
    message.stakedAmount !== undefined &&
      (obj.stakedAmount = message.stakedAmount);
    message.status !== undefined && (obj.status = message.status);
    message.description !== undefined &&
      (obj.description = message.description);
    message.imageUrl !== undefined && (obj.imageUrl = message.imageUrl);
    message.priceCount !== undefined && (obj.priceCount = message.priceCount);
    message.reviewCount !== undefined &&
      (obj.reviewCount = message.reviewCount);
    message.timestamp !== undefined && (obj.timestamp = message.timestamp);
    message.reserved !== undefined && (obj.reserved = message.reserved);
    if (message.owners) {
      obj.owners = message.owners.map((e) => e);
    } else {
      obj.owners = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Registry>): Registry {
    const message = { ...baseRegistry } as Registry;
    message.owners = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.stakedAmount !== undefined && object.stakedAmount !== null) {
      message.stakedAmount = object.stakedAmount;
    } else {
      message.stakedAmount = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = "";
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
    if (object.priceCount !== undefined && object.priceCount !== null) {
      message.priceCount = object.priceCount;
    } else {
      message.priceCount = "";
    }
    if (object.reviewCount !== undefined && object.reviewCount !== null) {
      message.reviewCount = object.reviewCount;
    } else {
      message.reviewCount = "";
    }
    if (object.timestamp !== undefined && object.timestamp !== null) {
      message.timestamp = object.timestamp;
    } else {
      message.timestamp = "";
    }
    if (object.reserved !== undefined && object.reserved !== null) {
      message.reserved = object.reserved;
    } else {
      message.reserved = "";
    }
    if (object.owners !== undefined && object.owners !== null) {
      for (const e of object.owners) {
        message.owners.push(e);
      }
    }
    return message;
  },
};

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
