# Pull base image.
FROM quay.io/eris/base
MAINTAINER Levvel <hello@levvel.io>

# configure install
ENV NAME         blockchain-digital-notary
ENV REPO 	getlevvel/blockchain-digital-notary
ENV ALIAS 	$REPO
ENV BRANCH       master
ENV BINARY_PATH  ./cmd/$NAME
ENV CLONE_PATH   $GOPATH/src/github.com/$ALIAS
ENV INSTALL_PATH $INSTALL_BASE/$NAME
ENV INSTALL_BASE /usr/local/bin

# install
WORKDIR $CLONE_PATH
RUN git clone -q https://github.com/$REPO $CLONE_PATH && \
  go get github.com/eris-ltd/eris-cli/cmd/eris && \
  go build && go install

# cleanup install
#RUN rm -rf $GOPATH/src/* && \
#  unset NAME && \
#  unset INSTALL_BASE && \
#  unset REPO && \
#  unset CLONE_PATH && \
#  unset BINARY_PATH && \
#  unset INSTALL_PATH && \
#  unset BRANCH

# start script
COPY start.sh $INSTALL_BASE/start

# set user
USER $USER
WORKDIR $ERIS

# boot
VOLUME $ERIS
EXPOSE 11113
CMD ["start"]

