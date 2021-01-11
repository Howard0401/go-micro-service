FROM alpine
ADD user /user
ENTRYPOINT [ "/user" ]

# FROM golang:latest
# ENV GOPROXY=https://proxy.golang.org
# WORKDIR /
# ADD . /
# RUN cd /user && go build  

# EXPOSE 8000
# ENTRYPOINT [ "./user" ]

# Damn, Windows couldn't use powershell,  you should use CMD
# reference: https://blog.csdn.net/hx7013/article/details/91489642