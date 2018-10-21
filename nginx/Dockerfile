FROM nginx:alpine

WORKDIR /nginx

COPY ./.nginx /nginx/conf/.nginx

CMD ["nginx", "-c", "/nginx/conf/.nginx"]