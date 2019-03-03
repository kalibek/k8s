FROM nginx:alpine
RUN rm -rf /etc/nginx/conf.d
COPY conf /etc/nginx
COPY index.html /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]