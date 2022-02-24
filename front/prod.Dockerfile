FROM node:16.13.2-alpine
WORKDIR /root/src
COPY ./front /root/src
COPY ./proto /root/src/proto
EXPOSE 3000
ENV NODE_ENV=production
RUN yarn
RUN cd othello/ && yarn && yarn build
RUN cd ../
ENTRYPOINT [ "node", "server.js" ]
