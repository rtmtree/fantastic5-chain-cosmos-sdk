/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "fantasfive.fantasfive";

export interface MsgCreateTeam {
  creator: string;
  player0: string;
  player1: string;
  player2: string;
  player3: string;
  player4: string;
}

export interface MsgCreateTeamResponse {
  teamId: string;
}

const baseMsgCreateTeam: object = {
  creator: "",
  player0: "",
  player1: "",
  player2: "",
  player3: "",
  player4: "",
};

export const MsgCreateTeam = {
  encode(message: MsgCreateTeam, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.player0 !== "") {
      writer.uint32(18).string(message.player0);
    }
    if (message.player1 !== "") {
      writer.uint32(26).string(message.player1);
    }
    if (message.player2 !== "") {
      writer.uint32(34).string(message.player2);
    }
    if (message.player3 !== "") {
      writer.uint32(42).string(message.player3);
    }
    if (message.player4 !== "") {
      writer.uint32(50).string(message.player4);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateTeam {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateTeam } as MsgCreateTeam;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.player0 = reader.string();
          break;
        case 3:
          message.player1 = reader.string();
          break;
        case 4:
          message.player2 = reader.string();
          break;
        case 5:
          message.player3 = reader.string();
          break;
        case 6:
          message.player4 = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateTeam {
    const message = { ...baseMsgCreateTeam } as MsgCreateTeam;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.player0 !== undefined && object.player0 !== null) {
      message.player0 = String(object.player0);
    } else {
      message.player0 = "";
    }
    if (object.player1 !== undefined && object.player1 !== null) {
      message.player1 = String(object.player1);
    } else {
      message.player1 = "";
    }
    if (object.player2 !== undefined && object.player2 !== null) {
      message.player2 = String(object.player2);
    } else {
      message.player2 = "";
    }
    if (object.player3 !== undefined && object.player3 !== null) {
      message.player3 = String(object.player3);
    } else {
      message.player3 = "";
    }
    if (object.player4 !== undefined && object.player4 !== null) {
      message.player4 = String(object.player4);
    } else {
      message.player4 = "";
    }
    return message;
  },

  toJSON(message: MsgCreateTeam): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.player0 !== undefined && (obj.player0 = message.player0);
    message.player1 !== undefined && (obj.player1 = message.player1);
    message.player2 !== undefined && (obj.player2 = message.player2);
    message.player3 !== undefined && (obj.player3 = message.player3);
    message.player4 !== undefined && (obj.player4 = message.player4);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateTeam>): MsgCreateTeam {
    const message = { ...baseMsgCreateTeam } as MsgCreateTeam;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.player0 !== undefined && object.player0 !== null) {
      message.player0 = object.player0;
    } else {
      message.player0 = "";
    }
    if (object.player1 !== undefined && object.player1 !== null) {
      message.player1 = object.player1;
    } else {
      message.player1 = "";
    }
    if (object.player2 !== undefined && object.player2 !== null) {
      message.player2 = object.player2;
    } else {
      message.player2 = "";
    }
    if (object.player3 !== undefined && object.player3 !== null) {
      message.player3 = object.player3;
    } else {
      message.player3 = "";
    }
    if (object.player4 !== undefined && object.player4 !== null) {
      message.player4 = object.player4;
    } else {
      message.player4 = "";
    }
    return message;
  },
};

const baseMsgCreateTeamResponse: object = { teamId: "" };

export const MsgCreateTeamResponse = {
  encode(
    message: MsgCreateTeamResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.teamId !== "") {
      writer.uint32(10).string(message.teamId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateTeamResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateTeamResponse } as MsgCreateTeamResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.teamId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateTeamResponse {
    const message = { ...baseMsgCreateTeamResponse } as MsgCreateTeamResponse;
    if (object.teamId !== undefined && object.teamId !== null) {
      message.teamId = String(object.teamId);
    } else {
      message.teamId = "";
    }
    return message;
  },

  toJSON(message: MsgCreateTeamResponse): unknown {
    const obj: any = {};
    message.teamId !== undefined && (obj.teamId = message.teamId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateTeamResponse>
  ): MsgCreateTeamResponse {
    const message = { ...baseMsgCreateTeamResponse } as MsgCreateTeamResponse;
    if (object.teamId !== undefined && object.teamId !== null) {
      message.teamId = object.teamId;
    } else {
      message.teamId = "";
    }
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateTeam(request: MsgCreateTeam): Promise<MsgCreateTeamResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateTeam(request: MsgCreateTeam): Promise<MsgCreateTeamResponse> {
    const data = MsgCreateTeam.encode(request).finish();
    const promise = this.rpc.request(
      "fantasfive.fantasfive.Msg",
      "CreateTeam",
      data
    );
    return promise.then((data) =>
      MsgCreateTeamResponse.decode(new Reader(data))
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
