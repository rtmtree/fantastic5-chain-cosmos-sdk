/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "fantasfive.fantasfive";

export interface SystemInfo {
  index: string;
  nextTeamId: number;
  nextMWId: number;
}

const baseSystemInfo: object = { index: "", nextTeamId: 0, nextMWId: 0 };

export const SystemInfo = {
  encode(message: SystemInfo, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.nextTeamId !== 0) {
      writer.uint32(16).uint64(message.nextTeamId);
    }
    if (message.nextMWId !== 0) {
      writer.uint32(24).uint64(message.nextMWId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): SystemInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSystemInfo } as SystemInfo;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.nextTeamId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.nextMWId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SystemInfo {
    const message = { ...baseSystemInfo } as SystemInfo;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.nextTeamId !== undefined && object.nextTeamId !== null) {
      message.nextTeamId = Number(object.nextTeamId);
    } else {
      message.nextTeamId = 0;
    }
    if (object.nextMWId !== undefined && object.nextMWId !== null) {
      message.nextMWId = Number(object.nextMWId);
    } else {
      message.nextMWId = 0;
    }
    return message;
  },

  toJSON(message: SystemInfo): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.nextTeamId !== undefined && (obj.nextTeamId = message.nextTeamId);
    message.nextMWId !== undefined && (obj.nextMWId = message.nextMWId);
    return obj;
  },

  fromPartial(object: DeepPartial<SystemInfo>): SystemInfo {
    const message = { ...baseSystemInfo } as SystemInfo;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.nextTeamId !== undefined && object.nextTeamId !== null) {
      message.nextTeamId = object.nextTeamId;
    } else {
      message.nextTeamId = 0;
    }
    if (object.nextMWId !== undefined && object.nextMWId !== null) {
      message.nextMWId = object.nextMWId;
    } else {
      message.nextMWId = 0;
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
