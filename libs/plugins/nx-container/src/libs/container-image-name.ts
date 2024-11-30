export const getContainerImageName = (projectRoot: string) => {
  return projectRoot.replace(/^(apps|libs)\//, '');
};
