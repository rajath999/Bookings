go build -o bookings.exe .\cmd\web\.
bookings.exe -dbname=bookings -dbuser=postgres -cache=false -production=false -dbpass=mysecretpassword