/* eslint-disable */
export default {
  displayName: 'plugins-nx-version',
  preset: '../../../jest.preset.js',
  globals: {},
  transform: {
    '^.+\\.[tj]s$': [
      'ts-jest',
      {
        tsconfig: '<rootDir>/tsconfig.spec.json',
      },
    ],
  },
  moduleFileExtensions: ['ts', 'js', 'html'],
  coverageDirectory: '../../../coverage/libs/plugins/nx-version',
  testEnvironment: 'node',
};
