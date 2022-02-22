FROM node:16.13.2-alpine
WORKDIR /root/src
COPY ./front /root/src
EXPOSE 3000
ENV NODE_ENV=production
RUN cd othello/ && yarn build
RUN cd ../
ENTRYPOINT [ "node", "server.js" ]
