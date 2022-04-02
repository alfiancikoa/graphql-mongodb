# GraphQL with MongoDB

<p align="justify"><b>MongoDB</b> (dari "humongous") adalah sistem basis data berorentasi dokumen lintas platform. Diklasifikasikan sebagai basis data "NoSQL", MongoDB menghindari struktur basis data relasional tabel berbasis tradisional yang mendukung JSON seperti dokumen dengan skema dinamis (MongoDB menyebutnya sebagai format BSON), membuat integrasi data dalam beberapa jenis aplikasi lebih mudah dan lebih cepat. Dirilis di bawah Server Side Public License, MongoDB adalah perangkat lunak bebas dan sumber terbuka.<br>
Pertama kali dikembangkan oleh perusahaan asal New York City, 10gen (sekarang MongoDB Inc.) pada bulan Oktober 2007 sebagai bagian dari platform yang direncanakan sebagai produk jasa, perusahaan bergeser ke model pembangunan sumber terbuka pada tahun 2009, dengan 10gen menawarkan dukungan komersial dan layanan lainnya.[4] Sejak itu, MongoDB telah diadopsi sebagai perangkat lunak backend oleh sejumlah situs dan layanan, termasuk Craigslist, eBay, Foursquare, SourceForge, dan The New York Times. MongoDB adalah sistem basis data NoSQL yang paling populer.<br>
(Sumber: https://id.wikipedia.org/wiki/MongoDB)
<br><br>
more detail <a href="https://github.com/mongodb/mongo-go-driver"><b>here</b></a>
</p>

### Command To Use

- Buat direktori workspace

```
$ mkdir graphql-mongodb
$ cd graphql-mongodb
$ go mod init github.com/[username]/graphql-mongodb
```

- Download library

```
# library gqlgen dari 99designs
$ go get github.com/99designs/gqlgen

# library mongo driver
$ go get go.mongodb.org/mongo-driver/mongo
```

- Terapkan skeleton/struktur folder dari 99designs

```
$ go run github.com/99designs/gqlgen init
```

more info <a href="https://github.com/99designs/gqlgen"><b>Here</b></a>
<br>
Thanks<br>
Alfian
