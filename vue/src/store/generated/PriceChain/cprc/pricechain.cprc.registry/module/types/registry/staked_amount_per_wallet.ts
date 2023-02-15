/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "pricechain.cprc.registry";

export interface StakedAmountPerWallet {
  index: string;
  wallet: string;
  stakedAmount: string;
}

const baseStakedAmountPerWallet: object = {
  index: "",
  wallet: "",
  stakedAmount: "",
};

export const StakedAmountPerWallet = {
  encode(
    message: StakedAmountPerWallet,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.wallet !== "") {
      writer.uint32(18).string(message.wallet);
    }
    if (message.stakedAmount !== "") {
      writer.uint32(26).string(message.stakedAmount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): StakedAmountPerWallet {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseStakedAmountPerWallet } as StakedAmountPerWallet;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.wallet = reader.string();
          break;
        case 3:
          message.stakedAmount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StakedAmountPerWallet {
    const message = { ...baseStakedAmountPerWallet } as StakedAmountPerWallet;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.wallet !== undefined && object.wallet !== null) {
      message.wallet = String(object.wallet);
    } else {
      message.wallet = "";
    }
    if (object.stakedAmount !== undefined && object.stakedAmount !== null) {
      message.stakedAmount = String(object.stakedAmount);
    } else {
      message.stakedAmount = "";
    }
    return message;
  },

  toJSON(message: StakedAmountPerWallet): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.wallet !== undefined && (obj.wallet = message.wallet);
    message.stakedAmount !== undefined &&
      (obj.stakedAmount = message.stakedAmount);
    return obj;
  },

  fromPartial(
    object: DeepPartial<StakedAmountPerWallet>
  ): StakedAmountPerWallet {
    const message = { ...baseStakedAmountPerWallet } as StakedAmountPerWallet;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.wallet !== undefined && object.wallet !== null) {
      message.wallet = object.wallet;
    } else {
      message.wallet = "";
    }
    if (object.stakedAmount !== undefined && object.stakedAmount !== null) {
      message.stakedAmount = object.stakedAmount;
    } else {
      message.stakedAmount = "";
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
