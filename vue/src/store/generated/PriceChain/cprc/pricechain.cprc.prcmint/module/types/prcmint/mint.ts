/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "pricechain.cprc.prcmint";

/** Minter represents the minting state. */
export interface Minter {
  /** current annual inflation rate */
  inflation: string;
  /** current annual expected provisions */
  annual_provisions: string;
}

const baseMinter: object = { inflation: "", annual_provisions: "" };

export const Minter = {
  encode(message: Minter, writer: Writer = Writer.create()): Writer {
    if (message.inflation !== "") {
      writer.uint32(10).string(message.inflation);
    }
    if (message.annual_provisions !== "") {
      writer.uint32(18).string(message.annual_provisions);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Minter {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMinter } as Minter;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.inflation = reader.string();
          break;
        case 2:
          message.annual_provisions = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Minter {
    const message = { ...baseMinter } as Minter;
    if (object.inflation !== undefined && object.inflation !== null) {
      message.inflation = String(object.inflation);
    } else {
      message.inflation = "";
    }
    if (
      object.annual_provisions !== undefined &&
      object.annual_provisions !== null
    ) {
      message.annual_provisions = String(object.annual_provisions);
    } else {
      message.annual_provisions = "";
    }
    return message;
  },

  toJSON(message: Minter): unknown {
    const obj: any = {};
    message.inflation !== undefined && (obj.inflation = message.inflation);
    message.annual_provisions !== undefined &&
      (obj.annual_provisions = message.annual_provisions);
    return obj;
  },

  fromPartial(object: DeepPartial<Minter>): Minter {
    const message = { ...baseMinter } as Minter;
    if (object.inflation !== undefined && object.inflation !== null) {
      message.inflation = object.inflation;
    } else {
      message.inflation = "";
    }
    if (
      object.annual_provisions !== undefined &&
      object.annual_provisions !== null
    ) {
      message.annual_provisions = object.annual_provisions;
    } else {
      message.annual_provisions = "";
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
