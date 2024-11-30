import * as protoParser from 'pbkit/core/parser/proto';
import { Tree } from '@nx/devkit';
import {
  Import,
  Option,
  OptionNameSegment,
  Package,
  Service,
  Statement,
  TopLevelStatement,
} from 'pbkit/core/ast';
import { underscore } from '@nx/devkit/src/utils/string-utils';
import { Dot } from 'pbkit/core/ast/lexical-elements';

export type SafeOptionsService = {
  name: string;
  grpcPort: string;
  httpPort: string;
};

export const readProto = (tree: Tree, filePath: string) => {
  const protoText = tree.read(filePath)?.toString();
  return protoParser.parse(protoText);
};

export const baseProtoDirectoryName = 'platformapis';

export const isService = (statement: TopLevelStatement): statement is Service =>
  statement.type === 'service';
export const isPackage = (statement: TopLevelStatement): statement is Package =>
  statement.type === 'package';

export const isOption = (statement: Statement): statement is Option =>
  statement.type === 'option';
export const isImport = (statement: TopLevelStatement): statement is Import =>
  statement.type === 'import';

export const isOptionSegment = (
  segment: OptionNameSegment | Dot
): segment is OptionNameSegment => segment.type === 'option-name-segment';

export const isSafeOption =
  (type: string | RegExp) =>
  ({ optionName }) => {
    const optionSegments =
      optionName.optionNameSegmentOrDots.filter(isOptionSegment);
    return optionSegments.find((f) => {
      const dotName = f.name.identOrDots.map(({ text }) => text).join('');
      if (typeof type === 'string') {
        return dotName.startsWith(type ? `platform.options.${type}` : 'platform');
      }
      return type.test(dotName);
    });
  };

export const serviceToProtoName = (service: string) => {
  return `${underscore(service.replace(/-\d+$/, ''))}.proto`;
};

export const getServices = (
  statements: Array<Statement>,
  protoText
): Array<SafeOptionsService> => {
  const [{ fullIdent }] = statements.filter(isPackage);
  const packageName = fullIdent.identOrDots.map(({ text }) => text).join('');

  const services =
    statements?.filter(isService).map((service) => {
      const {
        serviceName: { text: serviceName },
        serviceBody: { statements: serviceStatements },
      } = service;
      const options = serviceStatements.filter(isOption);
      const safeOptions = options
        .filter(isSafeOption('v2.service'))
        .map((o) => protoText.substring(o.start, o.end));
      const httpPort =
        safeOptions
          .find((o) => o.includes('http_port'))
          ?.match(/http_port:\s+(\d+)/)?.[1] ?? 8080;
      const grpcPort =
        safeOptions
          .find((o) => o.includes('grpc_port'))
          ?.match(/grpc_port:\s+(\d+)/)?.[1] ?? 50000;

      return {
        name: `${packageName}.${serviceName}`,
        grpcPort,
        httpPort,
      };
    }) ?? [];

  return services;
};

export type Version = `v${number}${'alpha' | 'beta' | ''}`;
export type ApiVisibility = 'internal' | 'poc' | 'public';
export type SafeProtoDir = 'platformapis' | 'platformapis-poc' | 'platformapis-internal';

export const apiVisibilityToSafeProtoDir: {
  [key in ApiVisibility]: SafeProtoDir;
} = {
  internal: 'platformapis-internal',
  poc: 'platformapis-poc',
  public: 'platformapis',
};

export const safeProtoDirToApiVisibility: {
  [key in SafeProtoDir]: ApiVisibility;
} = {
  'platformapis-internal': 'internal',
  'platformapis-poc': 'poc',
  safeapis: 'public',
};
