/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../fantasfive/params";
import { SystemInfo } from "../fantasfive/system_info";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { StoredMW } from "../fantasfive/stored_mw";
import { StoredTeam } from "../fantasfive/stored_team";

export const protobufPackage = "fantasfive.fantasfive";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetSystemInfoRequest {
  index: string;
}

export interface QueryGetSystemInfoResponse {
  systemInfo: SystemInfo | undefined;
}

export interface QueryAllSystemInfoRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllSystemInfoResponse {
  systemInfo: SystemInfo[];
  pagination: PageResponse | undefined;
}

export interface QueryGetStoredMWRequest {
  index: string;
}

export interface QueryGetStoredMWResponse {
  storedMW: StoredMW | undefined;
}

export interface QueryAllStoredMWRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllStoredMWResponse {
  storedMW: StoredMW[];
  pagination: PageResponse | undefined;
}

export interface QueryGetStoredTeamRequest {
  index: string;
}

export interface QueryGetStoredTeamResponse {
  storedTeam: StoredTeam | undefined;
}

export interface QueryAllStoredTeamRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllStoredTeamResponse {
  storedTeam: StoredTeam[];
  pagination: PageResponse | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
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

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetSystemInfoRequest: object = { index: "" };

export const QueryGetSystemInfoRequest = {
  encode(
    message: QueryGetSystemInfoRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetSystemInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetSystemInfoRequest,
    } as QueryGetSystemInfoRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetSystemInfoRequest {
    const message = {
      ...baseQueryGetSystemInfoRequest,
    } as QueryGetSystemInfoRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetSystemInfoRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetSystemInfoRequest>
  ): QueryGetSystemInfoRequest {
    const message = {
      ...baseQueryGetSystemInfoRequest,
    } as QueryGetSystemInfoRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetSystemInfoResponse: object = {};

export const QueryGetSystemInfoResponse = {
  encode(
    message: QueryGetSystemInfoResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.systemInfo !== undefined) {
      SystemInfo.encode(message.systemInfo, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetSystemInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetSystemInfoResponse,
    } as QueryGetSystemInfoResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.systemInfo = SystemInfo.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetSystemInfoResponse {
    const message = {
      ...baseQueryGetSystemInfoResponse,
    } as QueryGetSystemInfoResponse;
    if (object.systemInfo !== undefined && object.systemInfo !== null) {
      message.systemInfo = SystemInfo.fromJSON(object.systemInfo);
    } else {
      message.systemInfo = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetSystemInfoResponse): unknown {
    const obj: any = {};
    message.systemInfo !== undefined &&
      (obj.systemInfo = message.systemInfo
        ? SystemInfo.toJSON(message.systemInfo)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetSystemInfoResponse>
  ): QueryGetSystemInfoResponse {
    const message = {
      ...baseQueryGetSystemInfoResponse,
    } as QueryGetSystemInfoResponse;
    if (object.systemInfo !== undefined && object.systemInfo !== null) {
      message.systemInfo = SystemInfo.fromPartial(object.systemInfo);
    } else {
      message.systemInfo = undefined;
    }
    return message;
  },
};

const baseQueryAllSystemInfoRequest: object = {};

export const QueryAllSystemInfoRequest = {
  encode(
    message: QueryAllSystemInfoRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllSystemInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllSystemInfoRequest,
    } as QueryAllSystemInfoRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllSystemInfoRequest {
    const message = {
      ...baseQueryAllSystemInfoRequest,
    } as QueryAllSystemInfoRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllSystemInfoRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllSystemInfoRequest>
  ): QueryAllSystemInfoRequest {
    const message = {
      ...baseQueryAllSystemInfoRequest,
    } as QueryAllSystemInfoRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllSystemInfoResponse: object = {};

export const QueryAllSystemInfoResponse = {
  encode(
    message: QueryAllSystemInfoResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.systemInfo) {
      SystemInfo.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllSystemInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllSystemInfoResponse,
    } as QueryAllSystemInfoResponse;
    message.systemInfo = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.systemInfo.push(SystemInfo.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllSystemInfoResponse {
    const message = {
      ...baseQueryAllSystemInfoResponse,
    } as QueryAllSystemInfoResponse;
    message.systemInfo = [];
    if (object.systemInfo !== undefined && object.systemInfo !== null) {
      for (const e of object.systemInfo) {
        message.systemInfo.push(SystemInfo.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllSystemInfoResponse): unknown {
    const obj: any = {};
    if (message.systemInfo) {
      obj.systemInfo = message.systemInfo.map((e) =>
        e ? SystemInfo.toJSON(e) : undefined
      );
    } else {
      obj.systemInfo = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllSystemInfoResponse>
  ): QueryAllSystemInfoResponse {
    const message = {
      ...baseQueryAllSystemInfoResponse,
    } as QueryAllSystemInfoResponse;
    message.systemInfo = [];
    if (object.systemInfo !== undefined && object.systemInfo !== null) {
      for (const e of object.systemInfo) {
        message.systemInfo.push(SystemInfo.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetStoredMWRequest: object = { index: "" };

export const QueryGetStoredMWRequest = {
  encode(
    message: QueryGetStoredMWRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetStoredMWRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetStoredMWRequest,
    } as QueryGetStoredMWRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetStoredMWRequest {
    const message = {
      ...baseQueryGetStoredMWRequest,
    } as QueryGetStoredMWRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetStoredMWRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetStoredMWRequest>
  ): QueryGetStoredMWRequest {
    const message = {
      ...baseQueryGetStoredMWRequest,
    } as QueryGetStoredMWRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetStoredMWResponse: object = {};

export const QueryGetStoredMWResponse = {
  encode(
    message: QueryGetStoredMWResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.storedMW !== undefined) {
      StoredMW.encode(message.storedMW, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetStoredMWResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetStoredMWResponse,
    } as QueryGetStoredMWResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.storedMW = StoredMW.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetStoredMWResponse {
    const message = {
      ...baseQueryGetStoredMWResponse,
    } as QueryGetStoredMWResponse;
    if (object.storedMW !== undefined && object.storedMW !== null) {
      message.storedMW = StoredMW.fromJSON(object.storedMW);
    } else {
      message.storedMW = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetStoredMWResponse): unknown {
    const obj: any = {};
    message.storedMW !== undefined &&
      (obj.storedMW = message.storedMW
        ? StoredMW.toJSON(message.storedMW)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetStoredMWResponse>
  ): QueryGetStoredMWResponse {
    const message = {
      ...baseQueryGetStoredMWResponse,
    } as QueryGetStoredMWResponse;
    if (object.storedMW !== undefined && object.storedMW !== null) {
      message.storedMW = StoredMW.fromPartial(object.storedMW);
    } else {
      message.storedMW = undefined;
    }
    return message;
  },
};

const baseQueryAllStoredMWRequest: object = {};

export const QueryAllStoredMWRequest = {
  encode(
    message: QueryAllStoredMWRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllStoredMWRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllStoredMWRequest,
    } as QueryAllStoredMWRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllStoredMWRequest {
    const message = {
      ...baseQueryAllStoredMWRequest,
    } as QueryAllStoredMWRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllStoredMWRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllStoredMWRequest>
  ): QueryAllStoredMWRequest {
    const message = {
      ...baseQueryAllStoredMWRequest,
    } as QueryAllStoredMWRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllStoredMWResponse: object = {};

export const QueryAllStoredMWResponse = {
  encode(
    message: QueryAllStoredMWResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.storedMW) {
      StoredMW.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllStoredMWResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllStoredMWResponse,
    } as QueryAllStoredMWResponse;
    message.storedMW = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.storedMW.push(StoredMW.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllStoredMWResponse {
    const message = {
      ...baseQueryAllStoredMWResponse,
    } as QueryAllStoredMWResponse;
    message.storedMW = [];
    if (object.storedMW !== undefined && object.storedMW !== null) {
      for (const e of object.storedMW) {
        message.storedMW.push(StoredMW.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllStoredMWResponse): unknown {
    const obj: any = {};
    if (message.storedMW) {
      obj.storedMW = message.storedMW.map((e) =>
        e ? StoredMW.toJSON(e) : undefined
      );
    } else {
      obj.storedMW = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllStoredMWResponse>
  ): QueryAllStoredMWResponse {
    const message = {
      ...baseQueryAllStoredMWResponse,
    } as QueryAllStoredMWResponse;
    message.storedMW = [];
    if (object.storedMW !== undefined && object.storedMW !== null) {
      for (const e of object.storedMW) {
        message.storedMW.push(StoredMW.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetStoredTeamRequest: object = { index: "" };

export const QueryGetStoredTeamRequest = {
  encode(
    message: QueryGetStoredTeamRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetStoredTeamRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetStoredTeamRequest,
    } as QueryGetStoredTeamRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetStoredTeamRequest {
    const message = {
      ...baseQueryGetStoredTeamRequest,
    } as QueryGetStoredTeamRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetStoredTeamRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetStoredTeamRequest>
  ): QueryGetStoredTeamRequest {
    const message = {
      ...baseQueryGetStoredTeamRequest,
    } as QueryGetStoredTeamRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetStoredTeamResponse: object = {};

export const QueryGetStoredTeamResponse = {
  encode(
    message: QueryGetStoredTeamResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.storedTeam !== undefined) {
      StoredTeam.encode(message.storedTeam, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetStoredTeamResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetStoredTeamResponse,
    } as QueryGetStoredTeamResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.storedTeam = StoredTeam.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetStoredTeamResponse {
    const message = {
      ...baseQueryGetStoredTeamResponse,
    } as QueryGetStoredTeamResponse;
    if (object.storedTeam !== undefined && object.storedTeam !== null) {
      message.storedTeam = StoredTeam.fromJSON(object.storedTeam);
    } else {
      message.storedTeam = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetStoredTeamResponse): unknown {
    const obj: any = {};
    message.storedTeam !== undefined &&
      (obj.storedTeam = message.storedTeam
        ? StoredTeam.toJSON(message.storedTeam)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetStoredTeamResponse>
  ): QueryGetStoredTeamResponse {
    const message = {
      ...baseQueryGetStoredTeamResponse,
    } as QueryGetStoredTeamResponse;
    if (object.storedTeam !== undefined && object.storedTeam !== null) {
      message.storedTeam = StoredTeam.fromPartial(object.storedTeam);
    } else {
      message.storedTeam = undefined;
    }
    return message;
  },
};

const baseQueryAllStoredTeamRequest: object = {};

export const QueryAllStoredTeamRequest = {
  encode(
    message: QueryAllStoredTeamRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllStoredTeamRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllStoredTeamRequest,
    } as QueryAllStoredTeamRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllStoredTeamRequest {
    const message = {
      ...baseQueryAllStoredTeamRequest,
    } as QueryAllStoredTeamRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllStoredTeamRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllStoredTeamRequest>
  ): QueryAllStoredTeamRequest {
    const message = {
      ...baseQueryAllStoredTeamRequest,
    } as QueryAllStoredTeamRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllStoredTeamResponse: object = {};

export const QueryAllStoredTeamResponse = {
  encode(
    message: QueryAllStoredTeamResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.storedTeam) {
      StoredTeam.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllStoredTeamResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllStoredTeamResponse,
    } as QueryAllStoredTeamResponse;
    message.storedTeam = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.storedTeam.push(StoredTeam.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllStoredTeamResponse {
    const message = {
      ...baseQueryAllStoredTeamResponse,
    } as QueryAllStoredTeamResponse;
    message.storedTeam = [];
    if (object.storedTeam !== undefined && object.storedTeam !== null) {
      for (const e of object.storedTeam) {
        message.storedTeam.push(StoredTeam.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllStoredTeamResponse): unknown {
    const obj: any = {};
    if (message.storedTeam) {
      obj.storedTeam = message.storedTeam.map((e) =>
        e ? StoredTeam.toJSON(e) : undefined
      );
    } else {
      obj.storedTeam = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllStoredTeamResponse>
  ): QueryAllStoredTeamResponse {
    const message = {
      ...baseQueryAllStoredTeamResponse,
    } as QueryAllStoredTeamResponse;
    message.storedTeam = [];
    if (object.storedTeam !== undefined && object.storedTeam !== null) {
      for (const e of object.storedTeam) {
        message.storedTeam.push(StoredTeam.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a SystemInfo by index. */
  SystemInfo(
    request: QueryGetSystemInfoRequest
  ): Promise<QueryGetSystemInfoResponse>;
  /** Queries a list of SystemInfo items. */
  SystemInfoAll(
    request: QueryAllSystemInfoRequest
  ): Promise<QueryAllSystemInfoResponse>;
  /** Queries a StoredMW by index. */
  StoredMW(request: QueryGetStoredMWRequest): Promise<QueryGetStoredMWResponse>;
  /** Queries a list of StoredMW items. */
  StoredMWAll(
    request: QueryAllStoredMWRequest
  ): Promise<QueryAllStoredMWResponse>;
  /** Queries a StoredTeam by index. */
  StoredTeam(
    request: QueryGetStoredTeamRequest
  ): Promise<QueryGetStoredTeamResponse>;
  /** Queries a list of StoredTeam items. */
  StoredTeamAll(
    request: QueryAllStoredTeamRequest
  ): Promise<QueryAllStoredTeamResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "fantasfive.fantasfive.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  SystemInfo(
    request: QueryGetSystemInfoRequest
  ): Promise<QueryGetSystemInfoResponse> {
    const data = QueryGetSystemInfoRequest.encode(request).finish();
    const promise = this.rpc.request(
      "fantasfive.fantasfive.Query",
      "SystemInfo",
      data
    );
    return promise.then((data) =>
      QueryGetSystemInfoResponse.decode(new Reader(data))
    );
  }

  SystemInfoAll(
    request: QueryAllSystemInfoRequest
  ): Promise<QueryAllSystemInfoResponse> {
    const data = QueryAllSystemInfoRequest.encode(request).finish();
    const promise = this.rpc.request(
      "fantasfive.fantasfive.Query",
      "SystemInfoAll",
      data
    );
    return promise.then((data) =>
      QueryAllSystemInfoResponse.decode(new Reader(data))
    );
  }

  StoredMW(
    request: QueryGetStoredMWRequest
  ): Promise<QueryGetStoredMWResponse> {
    const data = QueryGetStoredMWRequest.encode(request).finish();
    const promise = this.rpc.request(
      "fantasfive.fantasfive.Query",
      "StoredMW",
      data
    );
    return promise.then((data) =>
      QueryGetStoredMWResponse.decode(new Reader(data))
    );
  }

  StoredMWAll(
    request: QueryAllStoredMWRequest
  ): Promise<QueryAllStoredMWResponse> {
    const data = QueryAllStoredMWRequest.encode(request).finish();
    const promise = this.rpc.request(
      "fantasfive.fantasfive.Query",
      "StoredMWAll",
      data
    );
    return promise.then((data) =>
      QueryAllStoredMWResponse.decode(new Reader(data))
    );
  }

  StoredTeam(
    request: QueryGetStoredTeamRequest
  ): Promise<QueryGetStoredTeamResponse> {
    const data = QueryGetStoredTeamRequest.encode(request).finish();
    const promise = this.rpc.request(
      "fantasfive.fantasfive.Query",
      "StoredTeam",
      data
    );
    return promise.then((data) =>
      QueryGetStoredTeamResponse.decode(new Reader(data))
    );
  }

  StoredTeamAll(
    request: QueryAllStoredTeamRequest
  ): Promise<QueryAllStoredTeamResponse> {
    const data = QueryAllStoredTeamRequest.encode(request).finish();
    const promise = this.rpc.request(
      "fantasfive.fantasfive.Query",
      "StoredTeamAll",
      data
    );
    return promise.then((data) =>
      QueryAllStoredTeamResponse.decode(new Reader(data))
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
