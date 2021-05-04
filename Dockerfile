FROM ubuntu:latest

COPY ./pluto /
COPY listed_company.csv /
ENTRYPOINT [ "/pluto" ] 
