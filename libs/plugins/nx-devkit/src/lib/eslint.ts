import { sortObjectKeys } from "./sort";

export const updateOverride = (eslint, override: { "files": Array<unknown> }) => {
  const updatedEslint = {
    ...eslint,
    overrides: [
      ...eslint?.overrides?.filter(({ files }) => [...files].sort().join(",") !== [...override.files].sort().join(",")) || [],
      override
    ]
  };
  return sortObjectKeys(updatedEslint);
};
