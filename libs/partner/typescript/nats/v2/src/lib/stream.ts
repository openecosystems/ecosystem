// Interface for Stream
export interface Stream {
  StreamPrefix(): string;
}

// InboundStream Class
export class InboundStream implements Stream {
  topicWildcard: string;
  streamName: string;

  constructor() {
    this.topicWildcard = "";
    this.streamName = "";
  }

  StreamPrefix(): string {
    return InboundStream.StreamPrefix();
  }

  static StreamPrefix(): string {
    return "inbound";
  }

}

// InternalStream Class
export class InternalStream implements Stream {
  topicWildcard: string;
  streamName: string;

  constructor() {
    this.topicWildcard = "";
    this.streamName = "";
  }

  StreamPrefix(): string {
    return InternalStream.StreamPrefix();
  }

  static StreamPrefix(): string {
    return "internal";
  }
}

// OutboundStream Class
export class OutboundStream implements Stream {
  topicWildcard: string;
  streamName: string;

  constructor() {
    this.topicWildcard = "";
    this.streamName = "";
  }

  StreamPrefix(): string {
    return OutboundStream.StreamPrefix();
  }

  static StreamPrefix(): string {
    return "outbound";
  }
}
