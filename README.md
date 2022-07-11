# Participants
Ini repository untuk service participants (service pendaftaran sekolah)

1. Lakukan copy file ".env.example" beberapa environment variable bisa diakses package godotenv melalui command line seperti di bawah ini :
```
cp .env.example .env
```
2. Bila belum ada database yang sama seperti di file ".env" . Lakukan create database (bisa dilakukan di aplikasi database management seperti DBeaver)
3. Bila belum dilakukan migration table (table belum dicreate) dan table belum ada data (table belum dilakukan seeding data). Execute command di terminal seperti di bawah ini :
```
sudo chmod +x ./pre-start.sh && ./pre-start.sh
```
4. Happy coding :)
