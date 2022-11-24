# use this file to update the app version from git. Make this file executeble

git pull

soda migrate

go build -o bookings cmd/web/*.go

sudo supervisorctl stop bookings
sudo supervisorctl start bookings