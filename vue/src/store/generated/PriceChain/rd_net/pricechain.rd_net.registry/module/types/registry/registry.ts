/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "pricechain.rd_net.registry";

export interface Registry {
  id: number;
  name: string;
  stakedAmount: string;
  status: string;
  price: string;
  quorum: string;
  consensusExpiringTime: string;
  activeMembers: string;
  prodInfo: string;
  memo: string;
  reserved: string;
}

const baseRegistry: object = {
  id: 0,
  name: "",
  stakedAmount: "",
  status: "",
  price: "",
  quorum: "",
  consensusExpiringTime: "",
  activeMembers: "",
  prodInfo: "",
  memo: "",
  reserved: "",
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
    if (message.price !== "") {
      writer.uint32(42).string(message.price);
    }
    if (message.quorum !== "") {
      writer.uint32(50).string(message.quorum);
    }
    if (message.consensusExpiringTime !== "") {
      writer.uint32(58).string(message.consensusExpiringTime);
    }
    if (message.activeMembers !== "") {
      writer.uint32(66).string(message.activeMembers);
    }
    if (message.prodInfo !== "") {
      writer.uint32(74).string(message.prodInfo);
    }
    if (message.memo !== "") {
      writer.uint32(82).string(message.memo);
    }
    if (message.reserved !== "") {
      writer.uint32(90).string(message.reserved);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Registry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseRegistry } as Registry;
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
          message.price = reader.string();
          break;
        case 6:
          message.quorum = reader.string();
          break;
        case 7:
          message.consensusExpiringTime = reader.string();
          break;
        case 8:
          message.activeMembers = reader.string();
          break;
        case 9:
          message.prodInfo = reader.string();
          break;
        case 10:
          message.memo = reader.string();
          break;
        case 11:
          message.reserved = reader.string();
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
    if (object.price !== undefined && object.price !== null) {
      message.price = String(object.price);
    } else {
      message.price = "";
    }
    if (object.quorum !== undefined && object.quorum !== null) {
      message.quorum = String(object.quorum);
    } else {
      message.quorum = "";
    }
    if (
      object.consensusExpiringTime !== undefined &&
      object.consensusExpiringTime !== null
    ) {
      message.consensusExpiringTime = String(object.consensusExpiringTime);
    } else {
      message.consensusExpiringTime = "";
    }
    if (object.activeMembers !== undefined && object.activeMembers !== null) {
      message.activeMembers = String(object.activeMembers);
    } else {
      message.activeMembers = "";
    }
    if (object.prodInfo !== undefined && object.prodInfo !== null) {
      message.prodInfo = String(object.prodInfo);
    } else {
      message.prodInfo = "";
    }
    if (object.memo !== undefined && object.memo !== null) {
      message.memo = String(object.memo);
    } else {
      message.memo = "";
    }
    if (object.reserved !== undefined && object.reserved !== null) {
      message.reserved = String(object.reserved);
    } else {
      message.reserved = "";
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
    message.price !== undefined && (obj.price = message.price);
    message.quorum !== undefined && (obj.quorum = message.quorum);
    message.consensusExpiringTime !== undefined &&
      (obj.consensusExpiringTime = message.consensusExpiringTime);
    message.activeMembers !== undefined &&
      (obj.activeMembers = message.activeMembers);
    message.prodInfo !== undefined && (obj.prodInfo = message.prodInfo);
    message.memo !== undefined && (obj.memo = message.memo);
    message.reserved !== undefined && (obj.reserved = message.reserved);
    return obj;
  },

  fromPartial(object: DeepPartial<Registry>): Registry {
    const message = { ...baseRegistry } as Registry;
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
    if (object.price !== undefined && object.price !== null) {
      message.price = object.price;
    } else {
      message.price = "";
    }
    if (object.quorum !== undefined && object.quorum !== null) {
      message.quorum = object.quorum;
    } else {
      message.quorum = "";
    }
    if (
      object.consensusExpiringTime !== undefined &&
      object.consensusExpiringTime !== null
    ) {
      message.consensusExpiringTime = object.consensusExpiringTime;
    } else {
      message.consensusExpiringTime = "";
    }
    if (object.activeMembers !== undefined && object.activeMembers !== null) {
      message.activeMembers = object.activeMembers;
    } else {
      message.activeMembers = "";
    }
    if (object.prodInfo !== undefined && object.prodInfo !== null) {
      message.prodInfo = object.prodInfo;
    } else {
      message.prodInfo = "";
    }
    if (object.memo !== undefined && object.memo !== null) {
      message.memo = object.memo;
    } else {
      message.memo = "";
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
