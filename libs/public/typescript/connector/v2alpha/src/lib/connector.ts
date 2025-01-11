export type Connector = {
  bindings: "disconnect";
  bounds: string;
  meshSocket: string;
  procedureName: string;
  name: string;
  err: Error;
  scheme: string;
  handler: string;
  opts: string;
};

export interface ConnectorOptions {

}

export interface ConnectorInterface {

  ListenAndProcess(): void;

}

