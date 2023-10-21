/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "fantasfive.fantasfive";

export interface StoredTeam {
  index: string;
  mwId: string;
  players: string;
  captainIndex: string;
  points: string;
  rank: string;
}

const baseStoredTeam: object = {
  index: "",
  mwId: "",
  players: "",
  captainIndex: "",
  points: "",
  rank: "",
};

export const StoredTeam = {
  encode(message: StoredTeam, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.mwId !== "") {
      writer.uint32(18).string(message.mwId);
    }
    if (message.players !== "") {
      writer.uint32(26).string(message.players);
    }
    if (message.captainIndex !== "") {
      writer.uint32(34).string(message.captainIndex);
    }
    if (message.points !== "") {
      writer.uint32(42).string(message.points);
    }
    if (message.rank !== "") {
      writer.uint32(50).string(message.rank);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): StoredTeam {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseStoredTeam } as StoredTeam;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.mwId = reader.string();
          break;
        case 3:
          message.players = reader.string();
          break;
        case 4:
          message.captainIndex = reader.string();
          break;
        case 5:
          message.points = reader.string();
          break;
        case 6:
          message.rank = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StoredTeam {
    const message = { ...baseStoredTeam } as StoredTeam;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.mwId !== undefined && object.mwId !== null) {
      message.mwId = String(object.mwId);
    } else {
      message.mwId = "";
    }
    if (object.players !== undefined && object.players !== null) {
      message.players = String(object.players);
    } else {
      message.players = "";
    }
    if (object.captainIndex !== undefined && object.captainIndex !== null) {
      message.captainIndex = String(object.captainIndex);
    } else {
      message.captainIndex = "";
    }
    if (object.points !== undefined && object.points !== null) {
      message.points = String(object.points);
    } else {
      message.points = "";
    }
    if (object.rank !== undefined && object.rank !== null) {
      message.rank = String(object.rank);
    } else {
      message.rank = "";
    }
    return message;
  },

  toJSON(message: StoredTeam): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.mwId !== undefined && (obj.mwId = message.mwId);
    message.players !== undefined && (obj.players = message.players);
    message.captainIndex !== undefined &&
      (obj.captainIndex = message.captainIndex);
    message.points !== undefined && (obj.points = message.points);
    message.rank !== undefined && (obj.rank = message.rank);
    return obj;
  },

  fromPartial(object: DeepPartial<StoredTeam>): StoredTeam {
    const message = { ...baseStoredTeam } as StoredTeam;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.mwId !== undefined && object.mwId !== null) {
      message.mwId = object.mwId;
    } else {
      message.mwId = "";
    }
    if (object.players !== undefined && object.players !== null) {
      message.players = object.players;
    } else {
      message.players = "";
    }
    if (object.captainIndex !== undefined && object.captainIndex !== null) {
      message.captainIndex = object.captainIndex;
    } else {
      message.captainIndex = "";
    }
    if (object.points !== undefined && object.points !== null) {
      message.points = object.points;
    } else {
      message.points = "";
    }
    if (object.rank !== undefined && object.rank !== null) {
      message.rank = object.rank;
    } else {
      message.rank = "";
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
