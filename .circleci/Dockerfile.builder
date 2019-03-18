FROM circleci/golang:1.11

# install dep
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

CMD ["/bin/sh"]