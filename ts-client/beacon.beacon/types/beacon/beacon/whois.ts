/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "beacon.beacon";

export interface Whois {
  validator: string;
  latest: string;
  jitter: string;
}

function createBaseWhois(): Whois {
  return { validator: "", latest: "", jitter: "" };
}

export const Whois = {
  encode(message: Whois, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.validator !== "") {
      writer.uint32(10).string(message.validator);
    }
    if (message.latest !== "") {
      writer.uint32(18).string(message.latest);
    }
    if (message.jitter !== "") {
      writer.uint32(26).string(message.jitter);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Whois {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWhois();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.validator = reader.string();
          break;
        case 2:
          message.latest = reader.string();
          break;
        case 3:
          message.jitter = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Whois {
    return {
      validator: isSet(object.validator) ? String(object.validator) : "",
      latest: isSet(object.latest) ? String(object.latest) : "",
      jitter: isSet(object.jitter) ? String(object.jitter) : "",
    };
  },

  toJSON(message: Whois): unknown {
    const obj: any = {};
    message.validator !== undefined && (obj.validator = message.validator);
    message.latest !== undefined && (obj.latest = message.latest);
    message.jitter !== undefined && (obj.jitter = message.jitter);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Whois>, I>>(object: I): Whois {
    const message = createBaseWhois();
    message.validator = object.validator ?? "";
    message.latest = object.latest ?? "";
    message.jitter = object.jitter ?? "";
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
