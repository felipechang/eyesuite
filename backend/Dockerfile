# OCR Server
FROM alpine
RUN apk add --update g++ go leptonica tesseract-ocr tesseract-ocr-dev
ADD . /app/
WORKDIR /app/
RUN go install
RUN go build -o eyesuite
CMD ["/app/eyesuite"]