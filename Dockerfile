FROM node:22.13.1-alpine as base

# Setup pnpm environment
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable

# Install Hugo
RUN apk add --update hugo

# Git
RUN apk add git

WORKDIR /app
COPY package.json ./
RUN pnpm install

# Expose HUGO port
EXPOSE 1010

VOLUME ./app

CMD ["npx", "nx", "serve", "docs", "--watch", "--poll", "700ms", "-p", "1010", "--bind", "0.0.0.0"]
