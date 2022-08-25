# OS Setup
FROM ubuntu:kinetic
RUN apt-get update
RUN apt-get upgrade -y
RUN apt-get update --fix-missing
RUN mkdir /src

# Certs
RUN apt-get install ca-certificates -y

# Network tools
RUN apt-get install net-tools -y
RUN apt-get install iputils-ping -y

# Basic tools
RUN apt-get install git -y
RUN apt-get install nano -y
RUN apt-get install p7zip-full -y
RUN apt-get install software-properties-common -y
RUN apt-get install libxml2-utils -y

# Lanuages
RUN apt-get install golang -y
RUN apt-get install python2 -y