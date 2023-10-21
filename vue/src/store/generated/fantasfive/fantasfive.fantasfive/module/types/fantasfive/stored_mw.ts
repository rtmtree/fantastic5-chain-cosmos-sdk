/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "fantasfive.fantasfive";

export interface StoredMW {
  index: string;
  playerPerf: string;
}

const baseStoredMW: object = { index: "", playerPerf: "" };

export const StoredMW = {
  encode(message: StoredMW, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.playerPerf !== "") {
      writer.uint32(18).string(message.playerPerf);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): StoredMW {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseStoredMW } as StoredMW;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.playerPerf = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StoredMW {
    const message = { ...baseStoredMW } as StoredMW;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.playerPerf !== undefined && object.playerPerf !== null) {
      message.playerPerf = String(object.playerPerf);
    } else {
      message.playerPerf = "";
    }
    return message;
  },

  toJSON(message: StoredMW): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.playerPerf !== undefined && (obj.playerPerf = message.playerPerf);
    return obj;
  },

  fromPartial(object: DeepPartial<StoredMW>): StoredMW {
    const message = { ...baseStoredMW } as StoredMW;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.playerPerf !== undefined && object.playerPerf !== null) {
      message.playerPerf = object.playerPerf;
    } else {
      message.playerPerf = "";
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
