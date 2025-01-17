nx g @nx/node:application apps/connectors/public/advertisement/v1/decision \
--name=apps-connectors-public-advertisement-v1-decision \
--bundler=esbuild \
--framework=none \
--linter=eslint \
--unitTestRunner=jest \
--e2eTestRunner=none


nx g @nx/node:library \
--directory=libs/partner/typescript/nats/v2 \
--name=libs-partner-typescript-nats-v2 \
--importPath=@openecosystems/natsv2 \
--buildable=true \
--compiler=tsc \
--linter=eslint \
--publishable=true \
--unitTestRunner=jest

nx g @nx/node:library \
--directory=libs/public/typescript/connector/v2alpha \
--name=libs-public-typescript-connector-v2alpha \
--importPath=@openecosystems/connectorv2alpha \
--buildable=true \
--compiler=tsc \
--linter=eslint \
--publishable=true \
--unitTestRunner=jest

nx g @nx/node:library \
--directory=libs/public/typescript/sdk/v2alpha \
--name=libs-public-typescript-sdk-v2alpha \
--importPath=@openecosystems/sdkv2alpha \
--buildable=true \
--compiler=tsc \
--linter=eslint \
--publishable=true \
--unitTestRunner=jest

nx g @nx/node:library \
--directory=libs/public/typescript/model \
--name=libs-public-typescript-model \
--importPath=@openecosystems/model-public \
--buildable=true \
--compiler=tsc \
--linter=eslint \
--publishable=true \
--unitTestRunner=jest

nx g @nx/node:library \
--directory=libs/private/typescript/model \
--name=libs-private-typescript-model \
--importPath=@openecosystems/model-private \
--buildable=true \
--compiler=tsc \
--linter=eslint \
--publishable=true \
--unitTestRunner=jest

nx g @nx/node:library \
--directory=libs/partner/typescript/model \
--name=libs-partner-typescript-model \
--importPath=@openecosystems/model-partner \
--buildable=true \
--compiler=tsc \
--linter=eslint \
--publishable=true \
--unitTestRunner=jest

nx g @nx/node:library \
--directory=libs/poc/typescript/model \
--name=libs-poc-typescript-model \
--importPath=@openecosystems/model-poc \
--buildable=true \
--compiler=tsc \
--linter=eslint \
--publishable=true \
--unitTestRunner=jest
