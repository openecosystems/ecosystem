import { Option, OptionNameSegment, Statement } from 'pbkit/core/ast';
import { Dot } from 'pbkit/core/ast/lexical-elements';
import { Rpc } from 'pbkit/core/ast/top-level-definitions';
export type SafeOptionsRpcService = {
  name: string;
  type: string;
  request: string;
  response: string;
};
export const baseProtoDirectoryName = 'platform';

export const isRPC = (statement: Statement): statement is Rpc =>
  statement.type === 'rpc';

export const isOption = (statement: Statement): statement is Option =>
  statement.type === 'option';

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
export const getRpcServices = (
  statements: Array<Statement>,
  protoText
): Array<SafeOptionsRpcService> => {
  const services =
    statements?.filter(isRPC).map((service) => {
      const {
        rpcName: { text: serviceName },
        //semiOrRpcBody: { statements: serviceStatements },
      } = service;
      if ('statements' in service.semiOrRpcBody) {
        const options = service.semiOrRpcBody.statements.filter(isOption);
        const safeOptions = options
          .filter(isSafeOption('v2.cqrs'))
          .map((o) => protoText.substring(o.start, o.end));
        const type =
          safeOptions
            .find((o) => o.includes('type'))
            ?.match(/type:\s+(\D+)(?=})/, 'g')?.[1] ?? '';

        const request = service.reqType.messageType.identOrDots[0].text;
        const response = service.resType.messageType.identOrDots[0].text;
        return {
          name: `${serviceName}`,
          type: type.trim(),
          request: request,
          response: response,
        };
      }
    }) ?? [];

  return services;
};
