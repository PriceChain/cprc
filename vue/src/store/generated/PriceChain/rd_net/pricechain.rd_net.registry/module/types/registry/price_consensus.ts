/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "pricechain.rd_net.registry";

export interface PriceConsensus {
  id: number;
  registryId: string;
  proposedPrice: string;
  proposedAt: string;
  status: string;
  yesCount: string;
  noCount: string;
  totalVoted: string;
  proposer: string;
  prodInfo: string;
  memo: string;
  reserved: string;
}

const basePriceConsensus: object = {
  id: 0,
  registryId: "",
  proposedPrice: "",
  proposedAt: "",
  status: "",
  yesCount: "",
  noCount: "",
  totalVoted: "",
  proposer: "",
  prodInfo: "",
  memo: "",
  reserved: "",
};

export const PriceConsensus = {
  encode(message: PriceConsensus, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.registryId !== "") {
      writer.uint32(18).string(message.registryId);
    }
    if (message.proposedPrice !== "") {
      writer.uint32(26).string(message.proposedPrice);
    }
    if (message.proposedAt !== "") {
      writer.uint32(34).string(message.proposedAt);
    }
    if (message.status !== "") {
      writer.uint32(42).string(message.status);
    }
    if (message.yesCount !== "") {
      writer.uint32(50).string(message.yesCount);
    }
    if (message.noCount !== "") {
      writer.uint32(58).string(message.noCount);
    }
    if (message.totalVoted !== "") {
      writer.uint32(66).string(message.totalVoted);
    }
    if (message.proposer !== "") {
      writer.uint32(74).string(message.proposer);
    }
    if (message.prodInfo !== "") {
      writer.uint32(82).string(message.prodInfo);
    }
    if (message.memo !== "") {
      writer.uint32(90).string(message.memo);
    }
    if (message.reserved !== "") {
      writer.uint32(98).string(message.reserved);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): PriceConsensus {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basePriceConsensus } as PriceConsensus;
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
          message.proposedPrice = reader.string();
          break;
        case 4:
          message.proposedAt = reader.string();
          break;
        case 5:
          message.status = reader.string();
          break;
        case 6:
          message.yesCount = reader.string();
          break;
        case 7:
          message.noCount = reader.string();
          break;
        case 8:
          message.totalVoted = reader.string();
          break;
        case 9:
          message.proposer = reader.string();
          break;
        case 10:
          message.prodInfo = reader.string();
          break;
        case 11:
          message.memo = reader.string();
          break;
        case 12:
          message.reserved = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PriceConsensus {
    const message = { ...basePriceConsensus } as PriceConsensus;
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
    if (object.proposedPrice !== undefined && object.proposedPrice !== null) {
      message.proposedPrice = String(object.proposedPrice);
    } else {
      message.proposedPrice = "";
    }
    if (object.proposedAt !== undefined && object.proposedAt !== null) {
      message.proposedAt = String(object.proposedAt);
    } else {
      message.proposedAt = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = String(object.status);
    } else {
      message.status = "";
    }
    if (object.yesCount !== undefined && object.yesCount !== null) {
      message.yesCount = String(object.yesCount);
    } else {
      message.yesCount = "";
    }
    if (object.noCount !== undefined && object.noCount !== null) {
      message.noCount = String(object.noCount);
    } else {
      message.noCount = "";
    }
    if (object.totalVoted !== undefined && object.totalVoted !== null) {
      message.totalVoted = String(object.totalVoted);
    } else {
      message.totalVoted = "";
    }
    if (object.proposer !== undefined && object.proposer !== null) {
      message.proposer = String(object.proposer);
    } else {
      message.proposer = "";
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

  toJSON(message: PriceConsensus): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.registryId !== undefined && (obj.registryId = message.registryId);
    message.proposedPrice !== undefined &&
      (obj.proposedPrice = message.proposedPrice);
    message.proposedAt !== undefined && (obj.proposedAt = message.proposedAt);
    message.status !== undefined && (obj.status = message.status);
    message.yesCount !== undefined && (obj.yesCount = message.yesCount);
    message.noCount !== undefined && (obj.noCount = message.noCount);
    message.totalVoted !== undefined && (obj.totalVoted = message.totalVoted);
    message.proposer !== undefined && (obj.proposer = message.proposer);
    message.prodInfo !== undefined && (obj.prodInfo = message.prodInfo);
    message.memo !== undefined && (obj.memo = message.memo);
    message.reserved !== undefined && (obj.reserved = message.reserved);
    return obj;
  },

  fromPartial(object: DeepPartial<PriceConsensus>): PriceConsensus {
    const message = { ...basePriceConsensus } as PriceConsensus;
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
    if (object.proposedPrice !== undefined && object.proposedPrice !== null) {
      message.proposedPrice = object.proposedPrice;
    } else {
      message.proposedPrice = "";
    }
    if (object.proposedAt !== undefined && object.proposedAt !== null) {
      message.proposedAt = object.proposedAt;
    } else {
      message.proposedAt = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = "";
    }
    if (object.yesCount !== undefined && object.yesCount !== null) {
      message.yesCount = object.yesCount;
    } else {
      message.yesCount = "";
    }
    if (object.noCount !== undefined && object.noCount !== null) {
      message.noCount = object.noCount;
    } else {
      message.noCount = "";
    }
    if (object.totalVoted !== undefined && object.totalVoted !== null) {
      message.totalVoted = object.totalVoted;
    } else {
      message.totalVoted = "";
    }
    if (object.proposer !== undefined && object.proposer !== null) {
      message.proposer = object.proposer;
    } else {
      message.proposer = "";
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
