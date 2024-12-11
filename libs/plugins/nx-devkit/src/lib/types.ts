export type addPrefix<TKey, TPrefix extends string, TSeparator extends string> = TKey extends string
  ? `${TPrefix}${TSeparator}${TKey}`
  : never;

export type removePrefix<TPrefixedKey, TPrefix extends string, TSeparator extends string> = TPrefixedKey extends addPrefix<infer TKey, TPrefix, TSeparator>
  ? TKey
  : '';

export type prefixedValue<TObject extends object, TPrefixedKey extends string, TPrefix extends string, TSeparator extends string> = TObject extends {[K in removePrefix<TPrefixedKey, TPrefix, TSeparator>]: infer TValue}
  ? TValue
  : never;

export type PrefixKeys<TObject extends object, TPrefix extends string, TSeparator extends string = ""> = {
  [K in addPrefix<keyof TObject, TPrefix, TSeparator>]: prefixedValue<TObject, K, TPrefix, TSeparator>
}

export type KeysWithValueType<T, V> = {[K in keyof T]-?: T[K] extends V ? K : never}[keyof T];

export type Split<S extends string, D extends string> =
  string extends S ? string[] :
    S extends '' ? [] :
      S extends `${infer T}${D}${infer U}` ? [T, ...Split<U, D>] :
        [S];

export type SplitOnWordSeparator<T extends string> = Split<T, "-"|"_"|" ">;
export type UndefinedToEmptyString<T extends string> = T extends undefined ? "" : T;
export type CamelCaseStringArray<K extends string[]> = `${K[0]}${Capitalize<UndefinedToEmptyString<K[1]>>}`;
export type CamelCase<K> = K extends string ? CamelCaseStringArray<SplitOnWordSeparator<K>> : K;

export type CamelCaseKeys<T> = {
  [K in keyof T as CamelCase<K>]: T[K]
};

export type PrefixCamelCaseKeys<TObject extends object, TPrefix extends string, TSSeparator extends string = "-"> = CamelCaseKeys<PrefixKeys<TObject, TPrefix, TSSeparator>>
