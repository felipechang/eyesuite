FROM alpine
RUN apk add --update g++ go leptonica tesseract-ocr tesseract-ocr-dev nodejs npm
ADD . /app/
WORKDIR /app/frontend
RUN npm install && npm run build
WORKDIR /app/
RUN go install && go build -o eyesuite
CMD ["/app/eyesuite"]