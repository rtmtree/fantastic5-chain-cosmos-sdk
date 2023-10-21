/* eslint-disable */
import { Params } from "../fantasfive/params";
import { SystemInfo } from "../fantasfive/system_info";
import { StoredMW } from "../fantasfive/stored_mw";
import { StoredTeam } from "../fantasfive/stored_team";
import { MwInfo } from "../fantasfive/mw_info";
import { TeamInfo } from "../fantasfive/team_info";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "fantasfive.fantasfive";

/** GenesisState defines the fantasfive module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  systemInfoList: SystemInfo[];
  storedMWList: StoredMW[];
  storedTeamList: StoredTeam[];
  mwInfo: MwInfo | undefined;
  /** this line is used by starport scaffolding # genesis/proto/state */
  teamInfo: TeamInfo | undefined;
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.systemInfoList) {
      SystemInfo.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.storedMWList) {
      StoredMW.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.storedTeamList) {
      StoredTeam.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    if (message.mwInfo !== undefined) {
      MwInfo.encode(message.mwInfo, writer.uint32(42).fork()).ldelim();
    }
    if (message.teamInfo !== undefined) {
      TeamInfo.encode(message.teamInfo, writer.uint32(50).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.systemInfoList = [];
    message.storedMWList = [];
    message.storedTeamList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.systemInfoList.push(
            SystemInfo.decode(reader, reader.uint32())
          );
          break;
        case 3:
          message.storedMWList.push(StoredMW.decode(reader, reader.uint32()));
          break;
        case 4:
          message.storedTeamList.push(
            StoredTeam.decode(reader, reader.uint32())
          );
          break;
        case 5:
          message.mwInfo = MwInfo.decode(reader, reader.uint32());
          break;
        case 6:
          message.teamInfo = TeamInfo.decode(reader, reader.uint32());
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
    message.systemInfoList = [];
    message.storedMWList = [];
    message.storedTeamList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.systemInfoList !== undefined && object.systemInfoList !== null) {
      for (const e of object.systemInfoList) {
        message.systemInfoList.push(SystemInfo.fromJSON(e));
      }
    }
    if (object.storedMWList !== undefined && object.storedMWList !== null) {
      for (const e of object.storedMWList) {
        message.storedMWList.push(StoredMW.fromJSON(e));
      }
    }
    if (object.storedTeamList !== undefined && object.storedTeamList !== null) {
      for (const e of object.storedTeamList) {
        message.storedTeamList.push(StoredTeam.fromJSON(e));
      }
    }
    if (object.mwInfo !== undefined && object.mwInfo !== null) {
      message.mwInfo = MwInfo.fromJSON(object.mwInfo);
    } else {
      message.mwInfo = undefined;
    }
    if (object.teamInfo !== undefined && object.teamInfo !== null) {
      message.teamInfo = TeamInfo.fromJSON(object.teamInfo);
    } else {
      message.teamInfo = undefined;
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.systemInfoList) {
      obj.systemInfoList = message.systemInfoList.map((e) =>
        e ? SystemInfo.toJSON(e) : undefined
      );
    } else {
      obj.systemInfoList = [];
    }
    if (message.storedMWList) {
      obj.storedMWList = message.storedMWList.map((e) =>
        e ? StoredMW.toJSON(e) : undefined
      );
    } else {
      obj.storedMWList = [];
    }
    if (message.storedTeamList) {
      obj.storedTeamList = message.storedTeamList.map((e) =>
        e ? StoredTeam.toJSON(e) : undefined
      );
    } else {
      obj.storedTeamList = [];
    }
    message.mwInfo !== undefined &&
      (obj.mwInfo = message.mwInfo ? MwInfo.toJSON(message.mwInfo) : undefined);
    message.teamInfo !== undefined &&
      (obj.teamInfo = message.teamInfo
        ? TeamInfo.toJSON(message.teamInfo)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.systemInfoList = [];
    message.storedMWList = [];
    message.storedTeamList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.systemInfoList !== undefined && object.systemInfoList !== null) {
      for (const e of object.systemInfoList) {
        message.systemInfoList.push(SystemInfo.fromPartial(e));
      }
    }
    if (object.storedMWList !== undefined && object.storedMWList !== null) {
      for (const e of object.storedMWList) {
        message.storedMWList.push(StoredMW.fromPartial(e));
      }
    }
    if (object.storedTeamList !== undefined && object.storedTeamList !== null) {
      for (const e of object.storedTeamList) {
        message.storedTeamList.push(StoredTeam.fromPartial(e));
      }
    }
    if (object.mwInfo !== undefined && object.mwInfo !== null) {
      message.mwInfo = MwInfo.fromPartial(object.mwInfo);
    } else {
      message.mwInfo = undefined;
    }
    if (object.teamInfo !== undefined && object.teamInfo !== null) {
      message.teamInfo = TeamInfo.fromPartial(object.teamInfo);
    } else {
      message.teamInfo = undefined;
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
