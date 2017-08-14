// package: 
// file: chat.proto

import * as chat_pb from "./chat_pb";
export class Chat {
  static serviceName = "Chat";
}
export namespace Chat {
  export class Join {
    static readonly methodName = "Join";
    static readonly service = Chat;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = chat_pb.JoinRequest;
    static readonly responseType = chat_pb.JoinResponse;
  }
  export class Say {
    static readonly methodName = "Say";
    static readonly service = Chat;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = chat_pb.SayRequest;
    static readonly responseType = chat_pb.SayResponse;
  }
}
