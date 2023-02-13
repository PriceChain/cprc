/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "pricechain.cprc.registry";

export interface RegistryOwner {
  id: number;
  registryId: string;
  stakedAmount: string;
  address: string;
  status: string;
  reserved: string;
}

const baseRegistryOwner: object = {
  id: 0,
  registryId: "",
  stakedAmount: "",
  address: "",
  status: "",
  reserved: "",
};

export const RegistryOwner = {
  encode(message: RegistryOwner, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.registryId !== "") {
      writer.uint32(18).string(message.registryId);
    }
    if (message.stakedAmount !== "") {
      writer.uint32(26).string(message.stakedAmount);
    }
    if (message.address !== "") {
      writer.uint32(34).string(message.address);
    }
    if (message.status !== "") {
      writer.uint32(42).string(message.status);
    }
    if (message.reserved !== "") {
      writer.uint32(50).string(message.reserved);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): RegistryOwner {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseRegistryOwner } as RegistryOwner;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.registryId = reader.string();
          break;
        case 3:
          message.stakedAmount = reader.string();
          break;
        case 4:
          message.address = reader.string();
          break;
        case 5:
          message.status = reader.string();
          break;
        case 6:
          message.reserved = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RegistryOwner {
    const message = { ...baseRegistryOwner } as RegistryOwner;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.registryId !== undefined && object.registryId !== null) {
      message.registryId = String(object.registryId);
    } else {
      message.registryId = "";
    }
    if (object.stakedAmount !== undefined && object.stakedAmount !== null) {
      message.stakedAmount = String(object.stakedAmount);
    } else {
      message.stakedAmount = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = String(object.status);
    } else {
      message.status = "";
    }
    if (object.reserved !== undefined && object.reserved !== null) {
      message.reserved = String(object.reserved);
    } else {
      message.reserved = "";
    }
    return message;
  },

  toJSON(message: RegistryOwner): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.registryId !== undefined && (obj.registryId = message.registryId);
    message.stakedAmount !== undefined &&
      (obj.stakedAmount = message.stakedAmount);
    message.address !== undefined && (obj.address = message.address);
    message.status !== undefined && (obj.status = message.status);
    message.reserved !== undefined && (obj.reserved = message.reserved);
    return obj;
  },

  fromPartial(object: DeepPartial<RegistryOwner>): RegistryOwner {
    const message = { ...baseRegistryOwner } as RegistryOwner;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.registryId !== undefined && object.registryId !== null) {
      message.registryId = object.registryId;
    } else {
      message.registryId = "";
    }
    if (object.stakedAmount !== undefined && object.stakedAmount !== null) {
      message.stakedAmount = object.stakedAmount;
    } else {
      message.stakedAmount = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = "";
    }
    if (object.reserved !== undefined && object.reserved !== null) {
      message.reserved = object.reserved;
    } else {
      message.reserved = "";
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
