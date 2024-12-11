export const notCharsRegex = /[^A-Za-z]/gi;

export const sortNotChars = (
  [a]: [string, string],
  [b]: [string, string],
  reverse = false
) => {
  const ax = a.replace(notCharsRegex, '');
  const bx = b.replace(notCharsRegex, '');
  return reverse ? bx.localeCompare(ax) : ax.localeCompare(bx);
};

export const sortObjectKeys = <T extends object>(obj: T, groups = []) => {
  const scriptEntries = Object.entries(obj);
  const matchingSorted = groups.flatMap((g) => {
    return scriptEntries
      .filter(([key]) => key.startsWith(g))
      .sort(sortNotChars);
  });

  const matchingKeys = matchingSorted.flatMap(([key]) => key);
  const nonMatchingSorted = scriptEntries
    .filter(([key]) => !matchingKeys.includes(key))
    .sort(sortNotChars);

  const allSorted = [...matchingSorted, ...nonMatchingSorted];
  return Object.fromEntries(allSorted) as T;
};
