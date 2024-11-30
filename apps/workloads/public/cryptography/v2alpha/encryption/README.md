cd proto
grpcurl \
-protoset <(buf build -o -) -plaintext \
-d '{}' \
localhost:6487 platform.cryptography.v2alpha.EncryptionService/Encrypt
