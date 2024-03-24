FROM node:20 as js-builder

COPY  js /js
WORKDIR /js

RUN npm ci && npm run build

#####################################
FROM python:3.8-slim as py-base

ENV PYTHONFAULTHANDLER=1 \
    PYTHONHASHSEED=random \
    PYTHONUNBUFFERED=1

WORKDIR /app

#####################################
FROM py-base as py-builder

ENV PIP_DEFAULT_TIMEOUT=100 \
    PIP_DISABLE_PIP_VERSION_CHECK=1 \
    PIP_NO_CACHE_DIR=1 \
    POETRY_VERSION=1.3.1

RUN pip install "poetry==$POETRY_VERSION"

COPY pyproject.toml poetry.lock README.md ./
COPY api ./api
RUN poetry config virtualenvs.in-project true && \
    poetry install --only=main --no-root && \
    poetry build

#####################################
FROM py-base as final

EXPOSE 8000

COPY --from=py-builder /app/.venv ./.venv
COPY --from=py-builder /app/dist .
COPY --from=js-builder /js/dist/* /app/public

RUN ./.venv/bin/pip install *.whl
CMD ["./.venv/bin/gunicorn", "-k", "uvicorn.workers.UvicornWorker", "-b", "0.0.0.0:8000", "api.main:app"]