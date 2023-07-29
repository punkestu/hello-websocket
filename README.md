# WEBSOCKET
## Review
Websocket merupakan salah satu metode koneksi dimana setelah client terkonek dengan service maka koneksi tidak akan langsung diputus setelah
kegiatan selesai. Oleh karena itu websocket biasa digunakan untuk melakukan transmisi dan monitoring data secara terus menerus misal chat

## Alur Kerja
Pada dasarnya websocket sama seperti koneksi web yang lainnya. Namun yang membedakan adalah koneksinya akan terus berjalan walaupun kegiatan
atau request sudah selesai. Sistemnya di sini akan menjadi mirip dengan message broker dimana client akan melakukan publish dan subscribe pada topic 
tertentu. Sehingga alurnya akan menjadi:
1. Client mengirim message ke server dengan topic yang telah ditentukan
2. Server memonitor message pada topic yang telah ditentukan
3. Server menerima message dan memprosesnya jika perlu
4. Server melakukan broadcast message melalui topic yang dimonitor oleh client
5. Client memonitor message pada topic tertentu
6. Client menerima message melalui topic yang dimonitor

## Alur Pembuatan
1. --- Server ---
2. Buka koneksi websocket dan monitor sebuah topic
3. Terima message pada topic yang dimonitor
4. Proses message jika perlu
5. Broadcast message melalui sebuah topic
6. --- Client ---
7. Buat Koneksi dengan server dan monitor sebuah topic
8. Kirim message dengan sebuah topic
9. Dapatkan message dari topic yang dimonitor ketika ada message baru
