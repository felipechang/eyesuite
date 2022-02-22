FROM alpine
RUN apk add --update g++ go \
    leptonica tesseract-ocr  \
    tesseract-ocr-dev nodejs \
    npm bash
WORKDIR /app
COPY ./ /app
COPY build.sh /app/build.sh
RUN chmod +x /app/build.sh

ENTRYPOINT ["bash","build.sh"]