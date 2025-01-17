import { Stream } from './stream';

export function GetMultiplexedRequestSubjectName(scope: Stream, subjectName: string): string {
  return `req.${scope}-${subjectName}`;
}
