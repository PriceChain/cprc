/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Params } from "../registry/params";
import { Registry } from "../registry/registry";
import { RegistryOwner } from "../registry/registry_owner";
import { RegistryMember } from "../registry/registry_member";

export const protobufPackage = "pricechain.rd_net.registry";

/** GenesisState defines the registry module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  registryList: Registry[];
  registryCount: number;
  registryOwnerList: RegistryOwner[];
  registryOwnerCount: number;
  registryMemberList: RegistryMember[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  registryMemberCount: number;
}

const baseGenesisState: object = {
  registryCount: 0,
  registryOwnerCount: 0,
  registryMemberCount: 0,
};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.registryList) {
      Registry.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.registryCount !== 0) {
      writer.uint32(24).uint64(message.registryCount);
    }
    for (const v of message.registryOwnerList) {
      RegistryOwner.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    if (message.registryOwnerCount !== 0) {
      writer.uint32(40).uint64(message.registryOwnerCount);
    }
    for (const v of message.registryMemberList) {
      RegistryMember.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    if (message.registryMemberCount !== 0) {
      writer.uint32(56).uint64(message.registryMemberCount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.registryList = [];
    message.registryOwnerList = [];
    message.registryMemberList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.registryList.push(Registry.decode(reader, reader.uint32()));
          break;
        case 3:
          message.registryCount = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.registryOwnerList.push(
            RegistryOwner.decode(reader, reader.uint32())
          );
          break;
        case 5:
          message.registryOwnerCount = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.registryMemberList.push(
            RegistryMember.decode(reader, reader.uint32())
          );
          break;
        case 7:
          message.registryMemberCount = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.registryList = [];
    message.registryOwnerList = [];
    message.registryMemberList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.registryList !== undefined && object.registryList !== null) {
      for (const e of object.registryList) {
        message.registryList.push(Registry.fromJSON(e));
      }
    }
    if (object.registryCount !== undefined && object.registryCount !== null) {
      message.registryCount = Number(object.registryCount);
    } else {
      message.registryCount = 0;
    }
    if (
      object.registryOwnerList !== undefined &&
      object.registryOwnerList !== null
    ) {
      for (const e of object.registryOwnerList) {
        message.registryOwnerList.push(RegistryOwner.fromJSON(e));
      }
    }
    if (
      object.registryOwnerCount !== undefined &&
      object.registryOwnerCount !== null
    ) {
      message.registryOwnerCount = Number(object.registryOwnerCount);
    } else {
      message.registryOwnerCount = 0;
    }
    if (
      object.registryMemberList !== undefined &&
      object.registryMemberList !== null
    ) {
      for (const e of object.registryMemberList) {
        message.registryMemberList.push(RegistryMember.fromJSON(e));
      }
    }
    if (
      object.registryMemberCount !== undefined &&
      object.registryMemberCount !== null
    ) {
      message.registryMemberCount = Number(object.registryMemberCount);
    } else {
      message.registryMemberCount = 0;
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.registryList) {
      obj.registryList = message.registryList.map((e) =>
        e ? Registry.toJSON(e) : undefined
      );
    } else {
      obj.registryList = [];
    }
    message.registryCount !== undefined &&
      (obj.registryCount = message.registryCount);
    if (message.registryOwnerList) {
      obj.registryOwnerList = message.registryOwnerList.map((e) =>
        e ? RegistryOwner.toJSON(e) : undefined
      );
    } else {
      obj.registryOwnerList = [];
    }
    message.registryOwnerCount !== undefined &&
      (obj.registryOwnerCount = message.registryOwnerCount);
    if (message.registryMemberList) {
      obj.registryMemberList = message.registryMemberList.map((e) =>
        e ? RegistryMember.toJSON(e) : undefined
      );
    } else {
      obj.registryMemberList = [];
    }
    message.registryMemberCount !== undefined &&
      (obj.registryMemberCount = message.registryMemberCount);
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.registryList = [];
    message.registryOwnerList = [];
    message.registryMemberList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.registryList !== undefined && object.registryList !== null) {
      for (const e of object.registryList) {
        message.registryList.push(Registry.fromPartial(e));
      }
    }
    if (object.registryCount !== undefined && object.registryCount !== null) {
      message.registryCount = object.registryCount;
    } else {
      message.registryCount = 0;
    }
    if (
      object.registryOwnerList !== undefined &&
      object.registryOwnerList !== null
    ) {
      for (const e of object.registryOwnerList) {
        message.registryOwnerList.push(RegistryOwner.fromPartial(e));
      }
    }
    if (
      object.registryOwnerCount !== undefined &&
      object.registryOwnerCount !== null
    ) {
      message.registryOwnerCount = object.registryOwnerCount;
    } else {
      message.registryOwnerCount = 0;
    }
    if (
      object.registryMemberList !== undefined &&
      object.registryMemberList !== null
    ) {
      for (const e of object.registryMemberList) {
        message.registryMemberList.push(RegistryMember.fromPartial(e));
      }
    }
    if (
      object.registryMemberCount !== undefined &&
      object.registryMemberCount !== null
    ) {
      message.registryMemberCount = object.registryMemberCount;
    } else {
      message.registryMemberCount = 0;
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
