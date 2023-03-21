/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "pricechain.cprc.registry";

export interface PriceData {
  index: string;
  registryId: string;
  creator: string;
  storeName: string;
  storeAddr: string;
  purchaseTime: string;
  prodName: string;
  price: string;
  receiptCode: string;
  reserved: string;
}

const basePriceData: object = {
  index: "",
  registryId: "",
  creator: "",
  storeName: "",
  storeAddr: "",
  purchaseTime: "",
  prodName: "",
  price: "",
  receiptCode: "",
  reserved: "",
};

export const PriceData = {
  encode(message: PriceData, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.registryId !== "") {
      writer.uint32(18).string(message.registryId);
    }
    if (message.creator !== "") {
      writer.uint32(26).string(message.creator);
    }
    if (message.storeName !== "") {
      writer.uint32(34).string(message.storeName);
    }
    if (message.storeAddr !== "") {
      writer.uint32(42).string(message.storeAddr);
    }
    if (message.purchaseTime !== "") {
      writer.uint32(50).string(message.purchaseTime);
    }
    if (message.prodName !== "") {
      writer.uint32(58).string(message.prodName);
    }
    if (message.price !== "") {
      writer.uint32(66).string(message.price);
    }
    if (message.receiptCode !== "") {
      writer.uint32(74).string(message.receiptCode);
    }
    if (message.reserved !== "") {
      writer.uint32(82).string(message.reserved);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): PriceData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basePriceData } as PriceData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.registryId = reader.string();
          break;
        case 3:
          message.creator = reader.string();
          break;
        case 4:
          message.storeName = reader.string();
          break;
        case 5:
          message.storeAddr = reader.string();
          break;
        case 6:
          message.purchaseTime = reader.string();
          break;
        case 7:
          message.prodName = reader.string();
          break;
        case 8:
          message.price = reader.string();
          break;
        case 9:
          message.receiptCode = reader.string();
          break;
        case 10:
          message.reserved = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PriceData {
    const message = { ...basePriceData } as PriceData;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.registryId !== undefined && object.registryId !== null) {
      message.registryId = String(object.registryId);
    } else {
      message.registryId = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
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

  toJSON(message: PriceData): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.registryId !== undefined && (obj.registryId = message.registryId);
    message.creator !== undefined && (obj.creator = message.creator);
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

  fromPartial(object: DeepPartial<PriceData>): PriceData {
    const message = { ...basePriceData } as PriceData;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.registryId !== undefined && object.registryId !== null) {
      message.registryId = object.registryId;
    } else {
      message.registryId = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
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
