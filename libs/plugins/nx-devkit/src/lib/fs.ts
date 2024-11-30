import { lstat, readdir, readFile, writeFile } from 'fs/promises';
import { URL } from 'url';
import { join, parse } from 'path';
import { CheerioAPI, CheerioOptions, load } from 'cheerio';
import { getExtension } from 'node-mime-types';
import fetch from 'node-fetch';
import { directoryExists } from '@nx/plugin/testing';

export const getMatchingFileSystemPaths = async (
  directory,
  filter: (path: string) => boolean,
  filePaths = []
) => {
  if (directoryExists(directory)) {
    for (const entry of await readdir(directory)) {
      const entryPath = join(directory, entry);
      if ((await lstat(entryPath)).isDirectory()) {
        await getMatchingFileSystemPaths(entryPath, filter, filePaths);
      } else {
        if (filter(entryPath)) {
          filePaths.push(entryPath);
        }
      }
    }
  }

  return filePaths;
};

export async function readJsonFile(path: string) {
  const file = await readFile(path, 'utf-8');
  return JSON.parse(file);
}

export type JsonStringifyOptions = {
  replacer?: (this: any, key: string, value: any) => any | null;
  space?: string | number;
};

export function writeJsonFile(
  path: string,
  json: any,
  options: JsonStringifyOptions
) {
  const data = JSON.stringify(json, options.replacer, options.space) + '\n';
  return writeFile(path, data);
}

export async function updateJsonFileSystem<
  T extends object = object,
  U extends object = object
>(
  path: string,
  callback: (a: T) => U | Promise<U>,
  options: JsonStringifyOptions = { replacer: null, space: 2 }
): Promise<U> {
  const json = await readJsonFile(path);
  const updatedJson = await callback(json);
  await writeJsonFile(path, updatedJson, options);
  return updatedJson;
}

const defaultXmlOptions: CheerioOptions = {
  xml: { xmlMode: true, withStartIndices: true, decodeEntities: false },
};

export async function readXmlFileSystem(
  path: string,
  options: CheerioOptions = defaultXmlOptions
) {
  const file = await readFile(path, 'utf-8');
  return load(file, options);
}

export async function updateXmlFileSystem(
  path: string,
  callback: (a: CheerioAPI) => CheerioAPI | Promise<CheerioAPI>,
  options: CheerioOptions = defaultXmlOptions
): Promise<CheerioAPI> {
  const xml = await readXmlFileSystem(path, options);
  const updatedXml = await callback(xml);
  const data = updatedXml
    .xml()
    .replace(/\s{4}\n/gi, '')
    .replace(/"\/>/gi, '" />');
  await writeFile(path, data, 'utf-8');
  return updatedXml;
}

/**
 * Downloads a file from a URL and returns the output path.
 * See downloadBuffer for more details.
 *
 * @param url The URL to download the file from
 * @param path The directory to download the file to
 * @param fileName The optional name (without extension) of the file to place in the outputPath
 * @param fileExtension The optional extension of the file to place in the outputPath.
 * @returns The path to the downloaded file
 * @throws If the download fails
 */
export async function downloadFile(
  url: string,
  path: string,
  fileName?: string,
  fileExtension?: string
): Promise<string> {
  const [buffer, outputPath] = await downloadBuffer(
    url,
    path,
    fileName,
    fileExtension
  );

  await writeFile(outputPath, buffer);
  return outputPath;
}

/**
 * Downloads a file from a URL and returns the buffer and the output path.
 * If a fileName is not provided it will be determined from the URL, defaulting
 * to "file" if one cannot be parsed.
 * If fileExtension is not provided it will be determined from the content-type
 * header.
 *
 * @param url The URL to download the file from
 * @param path The directory to download the file to
 * @param fileName The optional name (without extension) of the file to place in the outputPath
 * @param fileExtension The optional extension of the file to place in the outputPath.
 * @returns The buffer of the downloaded file and the path to the downloaded file
 * @throws If the download fails
 */
export async function downloadBuffer(
  url: string,
  path: string,
  fileName?: string,
  fileExtension?: string
): Promise<[Buffer, string]> {
  const response = await fetch(url);
  if (!response.ok) {
    throw new Error(
      `Failed to download file. Status: ${response.status} ${response.statusText}`
    );
  }

  if (!fileExtension) {
    const contentType = response.headers.get('content-type');
    const potentialExtension = getExtension(contentType);
    if (potentialExtension instanceof Array) {
      fileExtension = potentialExtension[0];
    } else {
      fileExtension = potentialExtension;
    }
  } else if (!fileExtension.startsWith('.')) {
    fileExtension = '.' + fileExtension;
  }

  if (!fileName) {
    const urlPath = new URL(url).pathname;
    const urlFileName = parse(urlPath).name;
    if (urlFileName) {
      fileName = urlFileName;
    } else {
      fileName = 'file';
    }
  }
  const buffer = await response.buffer();
  const outputPath = join(path, fileName + fileExtension);
  return [buffer, outputPath];
}
