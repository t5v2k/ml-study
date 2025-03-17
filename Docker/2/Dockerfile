FROM python:3.9.21-bookworm AS builder
COPY requirements.txt .
RUN pip install --user -r requirements.txt

ENV USER=appuser
ENV UID=10001
RUN adduser \
--disabled-password \
--gecos "" \
--home "/home/${USER}" \
--shell "/sbin/nologin" \
--uid "${UID}" \
"${USER}"


FROM python:3.9.21-alpine3.20
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
WORKDIR /code
COPY --chown=appuser:appuser --from=builder /root/.local/ /home/appuser/.local
COPY --chown=appuser:appuser ./code .
ENV PATH=/home/appuser/.local/bin:/home/appuser/.local/lib:$PATH
#ENV PYTHONUNBUFFERED=1
USER appuser:appuser
ENTRYPOINT [ "gunicorn", "--workers=10", "--bind=0.0.0.0:80", "--timeout=3", "app:app" ]
EXPOSE 80
