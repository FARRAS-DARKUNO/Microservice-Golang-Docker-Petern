# Menggunakan image dasar
FROM golang:1.22-alpine

# Menetapkan working directory di dalam container
WORKDIR /app

# Menyalin seluruh kode sumber ke dalam container
COPY . /app

# Menyalin go.mod dan go.sum untuk instalasi dependensi
COPY go.mod go.sum 
RUN go mod tidy

# Menyalin wait-for-postgres.sh ke dalam container
COPY wait-for-postgres.sh /wait-for-postgres.sh
RUN chmod +x /wait-for-postgres.sh

EXPOSE 80

# Menjalankan perintah untuk menunggu PostgreSQL siap dan kemudian menjalankan aplikasi
CMD ["/app/wait-for-postgres.sh", "category-db", "5434", "--", "go", "run", "main.go"]

