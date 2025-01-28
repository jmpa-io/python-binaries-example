# References:
# https://medium.com/dazn-tech/nuitka-python3-and-scratch-4ce209ed6dd6

FROM debian:bullseye-slim AS base

# install deps.
RUN apt-get update -y \
  && apt-get install --no-install-recommends -y \
    python3.9 \
  && rm -rf /var/cache/apt/archives
 
# setup app.
WORKDIR /var/app
COPY cmd/hello-world/run.py .

# define entrypoint, even though it's just for debugging.
ENTRYPOINT [ "python3.9", "/var/app/run.py", "--greeting"]

# ---

FROM base AS compiler

RUN apt-get update -y \
  && apt-get install --no-install-recommends -y \
    python3.9-dev \
    python3-pip \
    patchelf \
    build-essential \
    ccache \
    clang \
    libfuse-dev \
    upx \
  && python3.9 -m pip install nuitka

RUN python3.9 -m nuitka \
  --standalone \
  --nofollow-import-to=pytest \
  --python-flag=nosite,-O \
  --plugin-enable=pylint-warnings \
  --clang \
  --warn-implicit-exceptions \
  --warn-unusual-code \
  --prefer-source-code \
  run.py \
  && cd run.dist/ \
  && ldd run | grep "=> /" | awk '{print $3}' | xargs -I '{}' cp --no-clobber -v '{}' . \
  && ldd run | grep "/lib64/ld-linux-x86-64" | awk '{print $1}' | xargs -I '{}' cp --parents -v '{}' . \
  && cp --no-clobber -v /lib/x86_64-linux-gnu/libgcc_s.so.1 . \
  && mkdir -p ./lib/x86_64-linux-gnu/ \
  && cp --no-clobber -v /lib/x86_64-linux-gnu/libresolv* ./lib/x86_64-linux-gnu \
  && cp --no-clobber -v /lib/x86_64-linux-gnu/libnss_dns* ./lib/x86_64-linux-gnu \
  && upx -9 run
