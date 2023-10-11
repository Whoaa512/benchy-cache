# Build Stage
FROM rust:1.73 as builder

WORKDIR /usr/src/rust-file-cache
COPY . .

RUN cargo build --release

# Run Stage
FROM debian:buster-slim

COPY --from=builder /usr/src/rust-file-cache/target/release/rust-file-cache /usr/local/bin/rust-file-cache

EXPOSE 6942
CMD ["rust-file-cache"]
