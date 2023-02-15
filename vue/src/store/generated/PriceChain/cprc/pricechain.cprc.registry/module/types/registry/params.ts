/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "pricechain.cprc.registry";

/** Params defines the parameters for the module. */
export interface Params {
  /** type of coin to mint */
  minimum_stake: string;
}

const baseParams: object = { minimum_stake: "" };

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    if (message.minimum_stake !== "") {
      writer.uint32(10).string(message.minimum_stake);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseParams } as Params;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.minimum_stake = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Params {
    const message = { ...baseParams } as Params;
    if (object.minimum_stake !== undefined && object.minimum_stake !== null) {
      message.minimum_stake = String(object.minimum_stake);
    } else {
      message.minimum_stake = "";
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.minimum_stake !== undefined &&
      (obj.minimum_stake = message.minimum_stake);
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    if (object.minimum_stake !== undefined && object.minimum_stake !== null) {
      message.minimum_stake = object.minimum_stake;
    } else {
      message.minimum_stake = "";
    }
    return message;
  },
};

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
