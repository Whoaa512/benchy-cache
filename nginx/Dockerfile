# Use the official nginx image as a base
FROM nginx:alpine

# Copy configuration file
COPY nginx.conf /etc/nginx/nginx.conf

# Copy status script
COPY status.sh /usr/local/bin/status.sh

# Provide execute permissions to the script
RUN chmod +x /usr/local/bin/status.sh

# Expose our custom port
EXPOSE 6942

CMD ["nginx", "-g", "daemon off;"]
