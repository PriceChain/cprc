/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "pricechain.rd_net.registry";

export interface MsgCreateRegistry {
  creator: string;
  name: string;
  stakeAmount: string;
  quorum: string;
  consensusExpiringTime: string;
}

export interface MsgCreateRegistryResponse {}

const baseMsgCreateRegistry: object = {
  creator: "",
  name: "",
  stakeAmount: "",
  quorum: "",
  consensusExpiringTime: "",
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
    if (message.quorum !== "") {
      writer.uint32(34).string(message.quorum);
    }
    if (message.consensusExpiringTime !== "") {
      writer.uint32(42).string(message.consensusExpiringTime);
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
          message.quorum = reader.string();
          break;
        case 5:
          message.consensusExpiringTime = reader.string();
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
    return message;
  },

  toJSON(message: MsgCreateRegistry): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.stakeAmount !== undefined &&
      (obj.stakeAmount = message.stakeAmount);
    message.quorum !== undefined && (obj.quorum = message.quorum);
    message.consensusExpiringTime !== undefined &&
      (obj.consensusExpiringTime = message.consensusExpiringTime);
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

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateRegistry(
    request: MsgCreateRegistry
  ): Promise<MsgCreateRegistryResponse>;
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
      "pricechain.rd_net.registry.Msg",
      "CreateRegistry",
      data
    );
    return promise.then((data) =>
      MsgCreateRegistryResponse.decode(new Reader(data))
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
