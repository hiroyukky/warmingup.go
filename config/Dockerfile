# docker build -t smarteducation/web.kns:1.0 -f config/Dockerfile . --no-cache
FROM rockylinux/rockylinux:8.4

#ENV container docker

# systemd
RUN (cd /lib/systemd/system/sysinit.target.wants/; for i in *; do [ $i == \
systemd-tmpfiles-setup.service ] || rm -f $i; done); \
rm -f /lib/systemd/system/multi-user.target.wants/*;\
rm -f /etc/systemd/system/*.wants/*;\
rm -f /lib/systemd/system/local-fs.target.wants/*; \
rm -f /lib/systemd/system/sockets.target.wants/*udev*; \
rm -f /lib/systemd/system/sockets.target.wants/*initctl*; \
rm -f /lib/systemd/system/basic.target.wants/*;\
rm -f /lib/systemd/system/anaconda.target.wants/*;

# yum install
RUN yum update -y; yum -y install epel-release; \
yum install -y \
    gcc \
    git \
    tar \
    which \
    less \
    make \
    langpacks-ja \
    python36 && \
    yum clean all;

# go install
RUN git clone https://github.com/syndbg/goenv.git ~/.goenv && \
    echo 'export GOENV_ROOT="$HOME/.goenv"' >> ~/.bashrc && \
    echo 'export PATH="$GOENV_ROOT/bin:$PATH"' >> ~/.bashrc && \
    echo 'eval "$(goenv init -)"' >> ~/.bashrc && \
    source ~/.bashrc && \
    goenv install 1.17.6 && \
    goenv global 1.17.6 && \
    go get -u github.com/kyoh86/richgo

#ENV LANG="ja_JP.UTF-8" \
#    LANGUAGE="ja_JP:ja" \
#    LC_ALL="ja_JP.UTF-8"

COPY ./ /app
WORKDIR /app
