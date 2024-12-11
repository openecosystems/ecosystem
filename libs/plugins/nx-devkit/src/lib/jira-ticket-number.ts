export const prefixes = ['DEVOPS', 'OE', 'OECO'];

export const hasJiraTicketNumberRegEx = new RegExp(
  `(${prefixes.join('|')})-\\d+`
);

export const startsWithJiraTicketNumberRegEx = new RegExp(
  `^${hasJiraTicketNumberRegEx.source}`
);
