certbot certonly --noninteractive --agree-tos --cert-name stockswebcert -d www.srcport.net --register-unsafely-without-email --webroot -w /var/www/html/
a2enmod ssl
a2ensite default-ssl
mv /var/www/html/apache/default-ssl.conf /etc/apache2/sites-enabled/default-ssl.conf
service apache2 reload