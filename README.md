
# LOIC![9f181b0d-49c5-46fd-a71c-d4401b5ef860](https://github.com/user-attachments/assets/b95f2174-0acf-4fba-a2a8-b278d7e1d3a9)

Low Orbit Ion Cannon rewrite in Go
how to install 
- git clone https://github.com/whitehat57/LOIC.git
- cd LOIC
- go mod init LOIC.go
- go mod tidy
- go build LOIC.go
- rm go.mod
- rm go.sum
- rm LOIC.go
- ./LOIC
# harus sudah install golang minimal versi 1.22.0

Untuk menginstal Golang versi terbaru di Termux, ikuti langkah-langkah berikut:

- Perbarui Termux: Pastikan Termux kamu up to date dengan menjalankan perintah berikut:
``` bash
pkg update && pkg upgrade
```
- Install Golang: Gunakan perintah berikut untuk menginstal Golang versi terbaru dari repositori resmi Termux:
``` bash
pkg install golang
```
- Cek Versi Golang: Setelah proses instalasi selesai, pastikan Golang terinstal dengan benar dan cek versi yang diinstal:
```bash
go version
```
- Atur Path (Opsional jika diperlukan): Jika kamu memerlukan path khusus untuk Go workspace atau binari, kamu dapat menambahkannya ke .bashrc atau .zshrc kamu:
```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```
- Setelah menambahkan, jalankan perintah berikut untuk menerapkan perubahan:
```bash
source ~/.bashrc
```
